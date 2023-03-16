/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/r-scheele/utils"
	"github.com/spf13/cobra"
)

var commandMap map[string]func(string, string) error

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert a configuration file from one format to another",
	Run: func(cmd *cobra.Command, args []string) {

		path, _ := cmd.Flags().GetString("path")
		path_ext := strings.Split(path, ".")[1]

		output_file, _ := cmd.Flags().GetString("output")
		output_file_ext := strings.Split(output_file, ".")[1]

		command := fmt.Sprintf("%s2%s", path_ext, output_file_ext)

		err := commandMap[command](path, output_file)
		if err != nil {
			fmt.Println(err)
		}

	},
}

func init() {

	commandMap = make(map[string]func(string, string) error)
	commandMap["json2env"] = utils.Json2Env
	commandMap["env2json"] = utils.Env2JSON

	convertCmd.Flags().StringP("path", "p", "", "Path of the file to convert")
	convertCmd.MarkFlagRequired("path")
	convertCmd.Flags().StringP("output", "o", "", "Output file format")
	convertCmd.MarkFlagRequired("output")
	rootCmd.AddCommand(convertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// convertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
