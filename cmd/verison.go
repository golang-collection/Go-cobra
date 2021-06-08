package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

/**
* @Author: super
* @Date: 2021-06-08 13:26
* @Description:
**/

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of myCobra",
	Long:  `All software has versions. This is myCobra's`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("myCobra v0.9 -- HEAD")
	},
}
