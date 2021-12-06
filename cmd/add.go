package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

/**
* @Author: super
* @Date: 2021-06-08 13:55
* @Description:
**/

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new file",
	Long:  "Help you create the stand file",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("need input the file name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var (
			err      error
			fileName string
			file     *os.File
		)

		fileName = "./folder/" + args[0]
		_ = os.Mkdir("./folder", 644)
		_, err = os.Stat(fileName)
		if os.IsNotExist(err) {
			file, err = os.Create(fileName)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer file.Close()
			fmt.Println("Create file: ", fileName)
		fmt.Println("hello")
		}
	},
}
