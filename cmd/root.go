package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

/**
* @Author: super
* @Date: 2021-06-08 13:24
* @Description:
**/

var rootCmd = &cobra.Command{
	Use:   "myCobra",
	Short: "MyCobra is a test case",
	Long:  `Have fun`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
