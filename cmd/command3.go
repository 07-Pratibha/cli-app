/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

var command3Cmd = &cobra.Command{
	Use:   "command3",
	Short: "Make an API call and display the response",
	Long:  `Make an API call to a publicly available API and display the response in the CLI.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Make the API call
		response, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
		if err != nil {
			fmt.Println("Failed to make API call:", err)
			return
		}
		defer response.Body.Close()

		// Process the API response
		// For example, read the response body and display it
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Failed to read API response:", err)
			return
		}

		fmt.Println("API response:")
		fmt.Println(string(body))
	},
}

func init() {
	rootCmd.AddCommand(command3Cmd)
}
