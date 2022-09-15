/*
Copyright © 2022 Mike Pianka <mikepianka@protonmail.com>
*/
package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func isDirectory(path string) bool {
	info, err := os.Stat(path)

	if os.IsNotExist(err) {
		log.Fatal("input path does not exist; cannot continue")
	} else {
		if !info.Mode().IsDir() {
			log.Fatal("input path is not a directory; cannot continue")
		}
	}

	return true
}

func listSubdirectories(path string) ([]string, error) {
	contents, err := os.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	var dirs []string

	for _, item := range contents {
		if item.IsDir() {
			dirs = append(dirs, item.Name())
		}
	}

	if len(dirs) > 0 {
		return dirs, nil
	} else {
		return dirs, errors.New("no subdirectories found in input path")
	}
}

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find subdirectory sizes in a directory.",
	Long: `Search a directory to find the size
of each top level subdirectory.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		isDirectory(args[0])

		subdirs, err := listSubdirectories(args[0])

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(subdirs)
	},
}

func init() {
	rootCmd.AddCommand(findCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	findCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
