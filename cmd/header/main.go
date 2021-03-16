package header 

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/spf13/cobra"
)

var (
	_author string
	_mail string
)


var CMD = &cobra.Command{
	Use:   "header",
	Short: "add header for all files",
	Long:  `add header for all files`,
	Run:   _main,
}




var (
	// only files with these extensions will be considered
	extensions = []string{".go"}

	// paths with any of these prefixes will be skipped
	skipPrefixes = []string{
		// boring stuff
		"vendor/", 
		"tests/testdata/", 
	}

	// this regexp must match the entire license comment at the
	// beginning of each file.
	headerCommentRE = regexp.MustCompile(`^//\s*(Copyright|Author:).*?\n(?://.*?\n)*\n*`)
)


var headerTpl = template.Must(template.New("").Parse(`
// Copyright {{.Author}}©{{.Year}} All rights reserved.
// Author: {{.Author}} {{.Mail}}
// license that can be found in the LICENSE file

`[1:]))

type info struct {
	file string
	Year int64
	Author string
	Mail string
}

func init() {
	CMD.PersistentFlags().StringVarP(&_author, "author", "a", "CZ", " name of author ")
	CMD.PersistentFlags().StringVarP(&_mail, "mail", "m", "cz.theng@gmail.com", " email of author ")
}

func _main(cmd *cobra.Command, args []string) {
	processFiles()
}

func processFiles() {
	var (
		files = getFiles()
		filec = make(chan string)
		infoc = make(chan *info, 20)
		wg    sync.WaitGroup
	)

	go func() {
		for _, f := range files {
			filec <- f
		}
		close(filec)
	}()
	for i := runtime.NumCPU(); i >= 0; i-- {
		// getting file info is slow and needs to be parallel.
		// it traverses git history for each file.
		wg.Add(1)
		go getInfo(filec, infoc, &wg)
	}
	go func() {
		wg.Wait()
		close(infoc)
	}()
	appendHeaders(infoc)
}

func skipFile(path string) bool {
	if strings.Contains(path, "/testdata/") {
		return true
	}
	for _, p := range skipPrefixes {
		if strings.HasPrefix(path, p) {
			return true
		}
	}
	return false
}

func getFiles() []string {
	cmd := exec.Command("git", "ls-tree", "-r", "--name-only", "HEAD")
	var files []string
	err := doLines(cmd, func(line string) {
		if skipFile(line) {
			return
		}
		ext := filepath.Ext(line)
		for _, wantExt := range extensions {
			if ext == wantExt {
				goto keep
			}
		}
		return
	keep:
		files = append(files, line)
	})
	if err != nil {
		log.Fatal("error getting files:", err)
	}
	return files
}

func getInfo(files <-chan string, out chan<- *info, wg *sync.WaitGroup) {
	for file := range files {
		stat, err := os.Lstat(file)
		if err != nil {
			fmt.Printf("ERROR %s: %v\n", file, err)
			continue
		}
		if !stat.Mode().IsRegular() {
			continue
		}
		if isGenerated(file) {
			continue
		}
		info, err := fileInfo(file)
		if err != nil {
			fmt.Printf("ERROR %s: %v\n", file, err)
			continue
		}
		out <- info
	}
	wg.Done()
}

func isGenerated(file string) bool {
	fd, err := os.Open(file)
	if err != nil {
		return false
	}
	defer fd.Close()
	buf := make([]byte, 2048)
	n, _ := fd.Read(buf)
	buf = buf[:n]
	for _, l := range bytes.Split(buf, []byte("\n")) {
		if bytes.HasPrefix(l, []byte("// Code generated")) {
			return true
		}
	}
	return false
}

// fileInfo finds the lowest year in which the given file was committed.
func fileInfo(file string) (*info, error) {
	info := &info{
		file: file, 
		Year: int64(time.Now().Year()),
		Mail: _mail,
		Author: _author,
	}
	cmd := exec.Command("git", "log", "--follow", "--find-renames=80", "--find-copies=80", "--pretty=format:%ai", "--", file)
	err := doLines(cmd, func(line string) {
		y, err := strconv.ParseInt(line[:4], 10, 64)
		if err != nil {
			fmt.Printf("cannot parse year: %q", line[:4])
		}
		if y < info.Year {
			info.Year = y
		}
	})
	return info, err
}

func appendHeaders(infos <-chan *info) {
	for i := range infos {
		appendHeader(i)
	}
}

func appendHeader(info *info) {
	fi, err := os.Stat(info.file)
	if os.IsNotExist(err) {
		fmt.Println("skipping (does not exist)", info.file)
		return
	}
	if err != nil {
		log.Fatalf("error stat'ing %s: %v\n", info.file, err)
	}
	content, err := ioutil.ReadFile(info.file)
	if err != nil {
		log.Fatalf("error reading %s: %v\n", info.file, err)
	}
	// Construct new file content.
	buf := new(bytes.Buffer)
	headerTpl.Execute(buf, info)
	if m := headerCommentRE.FindIndex(content); m != nil && m[0] == 0 {
		buf.Write(content[:m[0]])
		buf.Write(content[m[1]:])
	} else {
		buf.Write(content)
	}
	// Write it to the file.
	if bytes.Equal(content, buf.Bytes()) {
		fmt.Println("skipping (no changes)", info.file)
		return
	}
	if err := ioutil.WriteFile(info.file, buf.Bytes(), fi.Mode()); err != nil {
		log.Fatalf("error writing %s: %v", info.file, err)
	}
}

func doLines(cmd *exec.Cmd, f func(string)) error {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	s := bufio.NewScanner(stdout)
	for s.Scan() {
		f(s.Text())
	}
	if s.Err() != nil {
		return s.Err()
	}
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("%v (for %s)", err, strings.Join(cmd.Args, " "))
	}
	return nil
}
