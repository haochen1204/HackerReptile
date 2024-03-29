package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// 版本信息
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of HackerReptile",
	Long:  `All software has versions. This is HackerReptile's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("HackerReptile Static Site Generator v0.1 -- HEAD")
	},
}

// init
//
//	@Description:
func init() {
	rootCmd.AddCommand(versionCmd)
}
