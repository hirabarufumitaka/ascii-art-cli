/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"

	"github.com/hirabarufumitaka/ascii-art-cli/aa"
	"github.com/spf13/cobra"
)

// birdCmd represents the bird command
var birdCmd = &cobra.Command{
	Use:   "bird",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := aa.Aa.Open("bird.txt")
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		// ファイルを読み込んで出力
		buf := new(bytes.Buffer)
		buf.ReadFrom(file)

		fmt.Print(buf.String())
	},
}

func init() {
	rootCmd.AddCommand(birdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// birdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// birdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
