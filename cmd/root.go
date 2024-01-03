/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"HackerReptile/console"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var Url string
var UrlList string
var Cookie string

var header = `
 _   _            _            ______           _   _ _      
| | | |          | |           | ___ \         | | (_) |     
| |_| | __ _  ___| | _____ _ __| |_/ /___ _ __ | |_ _| | ___ 
|  _  |/ _` + "`" + ` |/ __| |/ / _ \ '__|    // _ \ '_ \| __| | |/ _ \
| | | | (_| | (__|   <  __/ |  | |\ \  __/ |_) | |_| | |  __/
\_| |_/\__,_|\___|_|\_\___|_|  \_| \_\___| .__/ \__|_|_|\___|
                                         | |                 
                                         |_|

一个专用于渗透测试的爬虫
目的在于根据需求，在渗透测试、攻防演练中快速爬取需要的信息
`

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "HackerReptile",
	Short: "一个专用于渗透测试的爬虫",
	Long:  header,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(header + "\n")
		console.Console()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.HackerReptile.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().StringVarP(&Url, "url", "u", "", "爬取的url地址")
	rootCmd.PersistentFlags().StringVarP(&UrlList, "list", "l", "", "爬取的url地址文件")
	rootCmd.PersistentFlags().StringVarP(&Cookie, "cookie", "c", "", "设置cookie的值")

	// 绑定到viper
	viper.BindPFlag("url", rootCmd.PersistentFlags().Lookup("url"))
	viper.BindPFlag("list", rootCmd.PersistentFlags().Lookup("list"))
	viper.BindPFlag("cookie", rootCmd.PersistentFlags().Lookup("cookie"))
}
