package main

import (
	"fmt"

	"github.com/cz-theng/czkit-go"
	"github.com/cz-theng/czkit-go/cmd/app"
	"github.com/cz-theng/czkit-go/cmd/header"
	"github.com/spf13/cobra"
)

var (
	_version bool
)

var rootCMD = &cobra.Command{
	Use:   "czkit",
	Short: "czkit",
	Long:  `czkit`,
	Run:   _main,
}

func init() {
	rootCMD.PersistentFlags().BoolVarP(&_version, "version", "v", false, "print version of czkit")

	rootCMD.AddCommand(app.CMD)
	rootCMD.AddCommand(header.CMD)
}

func main() {
	rootCMD.Execute()
}

func dumpVersion() {
	fmt.Printf("%s\n", czkit.Version())
}

func _main(cmd *cobra.Command, args []string) {
	if _version {
		dumpVersion()
		return
	}

}
