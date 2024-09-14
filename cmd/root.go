/*
Copyright © 2024 Mihail Andritchi <mihail.andirtchi@gmail.com>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "json_parser",
	Short: "A tiny json parser tool",
	Long: `Another coding challenge to learn the go programming language:
				This one should simply parse strings or files into json format.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		if len(filePath) > 0 {
			data, err := os.ReadFile(filePath)
			if err != nil {
				log.Fatalf("Error reading file: %s", err)
			}

			var jsonData interface{}
			err = json.Unmarshal(data, &jsonData)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			} else {
				fmt.Println("Valid JSON")
				os.Exit(0)
			}

		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of json parser",
	Long:  "JSON Parser version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("json_parser v0.1")
	},
}
