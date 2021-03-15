package app

import (
	"github.com/spf13/cobra"
)

var (
	_initFlag bool
)

var AppCMD = &cobra.Command{
	Use:   "app",
	Short: "create an app templment",
	Long:  `create an app templment`,
	Run:   _main,
}

func init() {
	AppCMD.PersistentFlags().BoolVarP(&_initFlag, "init", "i", false, "create and init an app templement")

}

func _main(cmd *cobra.Command, args []string) {
	if _initFlag {
		createAndInitDir()
		return
	}
}

func createAndInitDir() {

}
