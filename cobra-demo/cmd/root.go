package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"julive/components/db"
	"julive/components/logger"
)

var rootCmd = &cobra.Command{
	Use: "cobra-demo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("welcome to use cobra demo, use -h or --help to see more detail")
	},
}

var registers = [...]func() error{
	logger.Register,
	db.Register,
}

func Init() (err error) {
	//依次初始化组件,只要其中一个失败直接退出
	for _, register := range registers {
		err = register()
		if err != nil {
			return
		}
	}
	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
