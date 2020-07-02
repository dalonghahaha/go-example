package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	testCmd.PersistentFlags().StringP("host", "H", "127.0.0.1", "参数数目")
	rootCmd.AddCommand(testCmd)
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "命令描述文字",
	Run: func(cmd *cobra.Command, args []string) {
		//参数
		fmt.Println("参数:", args)
		//获取flag值
		fmt.Println("host:", cmd.Flag("host").Value.String())
	},
}
