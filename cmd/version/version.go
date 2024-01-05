package version

import (
	"fmt"

	"github.com/cuiyuanxin/airuisi-admin/global"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "version",
	Short:   "Get version info",
	Example: global.AppName + "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(global.Version)
	},
}
