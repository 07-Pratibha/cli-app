/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"./config"

	"github.com/spf13/cobra"
)

var command3Cmd = &cobra.Command{
	Use:   "command3",
	Short: "Call API for command3",
	Run: func(cmd *cobra.Command, args []string) {
		// Load API configuration from config.yaml
		configFile := "cmd/config.yaml"
		apiConfig, err := config.LoadConfig(configFile)
		if err != nil {
			fmt.Printf("Failed to load config file: %v\n", err)
			return
		}

		// Retrieve the API configuration for command2
		apiInfo, ok := apiConfig.Apis["command2"]
		if !ok {
			fmt.Println("API configuration for command2 not found")
			return
		}

		// Make the API call using the retrieved configuration
		fmt.Printf("Calling command2 API: %s %s\n", apiInfo.Method, apiInfo.URL)
		// Add your API call logic here

		// Example: Making a GET request using the net/http package
		resp, err := http.Get(apiInfo.URL)
		if err != nil {
			fmt.Printf("Failed to make API call: %v\n", err)
			return
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Failed to read response body: %v\n", err)
			return
		}

		// Store the API response in a variable or process it further
		apiResponse := string(body)
		fmt.Printf("API response: %s\n", apiResponse)
	},
}

func init() {
	rootCmd.AddCommand(command2Cmd)
}
