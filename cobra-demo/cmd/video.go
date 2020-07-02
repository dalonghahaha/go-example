package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"cobra-demo/services"
)

func init() {
	videoCmd.PersistentFlags().StringP("id", "i", "127.0.0.1", "video ID")
	rootCmd.AddCommand(videoCmd)
}

var videoCmd = &cobra.Command{
	Use: "video",
	Run: func(cmd *cobra.Command, args []string) {
		id := cmd.Flag("id").Value.String()
		_id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			fmt.Println("参数id格式不正确")
			return
		}
		info := services.NewVideoService().GetVideo(_id)
		if info == nil {
			fmt.Println("视频不存在")
			return
		}
		fmt.Println("img:", info.VideoImg)
		fmt.Println("url:", info.VideoURL)
	},
}
