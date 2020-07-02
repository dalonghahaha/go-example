package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	subCmd.PersistentFlags().StringP("name", "n", "", "名称")
	sub2Cmd.PersistentFlags().StringP("address", "a", "beijing", "地址")
	subCmd.AddCommand(sub2Cmd)
	subCmd.AddCommand(sub3Cmd)
	rootCmd.AddCommand(subCmd)
}

var subCmd = &cobra.Command{
	Use: "sub",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("this is sub cmd")
		//获取flag值
		fmt.Println("name:", cmd.Flag("name").Value.String())
	},
}

var sub2Cmd = &cobra.Command{
	Use: "sub2",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("this is sub2 cmd")
		//获取flag值
		fmt.Println("name:", cmd.Flag("name").Value.String())
		fmt.Println("address:", cmd.Flag("address").Value.String())
	},
}

var sub3Cmd = &cobra.Command{
	Use: "sub3",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("this is sub2 cmd")
		//获取flag值
		fmt.Println("name:", cmd.Flag("name").Value.String())
	},
}
