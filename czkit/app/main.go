package app

import (
	"log"
	"os"

	"github.com/cz-theng/czkit-go/czkit/app/asset"
	"github.com/spf13/cobra"
)

//go:generate go-bindata -o=app/asset/asset.go -pkg=asset test/...

var (
	_initFlag bool
)

var CMD = &cobra.Command{
	Use:   "app",
	Short: "create an app templment",
	Long:  `create an app templment`,
	Run:   _main,
}

func init() {
	CMD.PersistentFlags().BoolVarP(&_initFlag, "init", "i", false, "create and init an app templement")

}

func _main(cmd *cobra.Command, args []string) {
	if _initFlag {
		createAndInitDir()
		return
	}
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func createAndInitDir() {

	if !exists("./cmd") {
		if err := os.Mkdir("cmd", os.ModePerm); err != nil {
			log.Fatal("create directory cmd error")
		}
	}

	if err := asset.RestoreAssets("./cmd", "test"); err != nil {
		log.Fatalf("expand asset error:%s", err.Error())
	}
}
