/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

// command4Cmd represents the command4 command
var command4Cmd = &cobra.Command{
	Use:   "command4",
	Short: "Make a POST API call and display the response",
	Long:  `Make a POST API call to a publicly available API and display the response in the CLI.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Define the request payload
		payload := struct {
			Title  string `json:"title"`
			Body   string `json:"body"`
			UserID int    `json:"userId"`
		}{
			Title:  "Create a CLI",
			Body:   "Such a fun problem statement",
			UserID: 4,
		}

		// Convert the payload to JSON
		payloadJSON, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Failed to marshal JSON payload:", err)
			return
		}

		// Make the POST API call
		response, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(payloadJSON))
		if err != nil {
			fmt.Println("Failed to make POST API call:", err)
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
	rootCmd.AddCommand(command4Cmd)
}
