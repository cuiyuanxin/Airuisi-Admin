package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/cuiyuanxin/airuisi-admin/cmd/service"

	"github.com/cuiyuanxin/airuisi-admin/cmd/version"

	"github.com/cuiyuanxin/airuisi-admin/global"

	"github.com/spf13/cobra"
	"github.com/zztroot/color"
)

func tip() {
	usageStr := `欢迎使用 ` + color.Coat(fmt.Sprintf("%s %s", global.AppName, global.Version), color.Green) + ` 可以使用 ` + color.Coat("-h", color.Red) + ` 查看命令`
	usageStr1 := `也可以参考xxx官方文档的相关内容`
	fmt.Printf("%s\n", usageStr)
	fmt.Printf("%s\n", usageStr1)
}

var rootCmd = &cobra.Command{
	Use:          global.AppName,
	Short:        global.AppName,
	SilenceUsage: true,
	Long:         global.AppName,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New(color.Coat("requires at least one arg", color.Red))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func init() {
	rootCmd.AddCommand(service.Cmd)
	rootCmd.AddCommand(version.Cmd)
}

// Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
