/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
<<<<<<< HEAD
=======

>>>>>>> service2
*/
package cmd

import (
<<<<<<< HEAD
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
=======
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"net/http"
	"github.com/spf13/cobra"
 )

// command3Cmd represents the command3 command
var command3Cmd = &cobra.Command{
	Use:   "echopost [name1] [name2]",
	Short: "Posts some data as json to postman-echo API and displays the echoed response",
	Long: `Posts some data as json to the postman-echo API and displays the response that is echoed back.`,
	Run: func(cmd *cobra.Command, args []string) {
		postBody, _ := json.Marshal(map[string]string{
			"name1": os.Args[2],
			"name2": os.Args[3],
		 })
		 responseBody := bytes.NewBuffer(postBody)
	  	
		 //Leverage Go's HTTP Post function to make request
		 resp, err := http.Post("https://postman-echo.com/post", "application/json", responseBody)
	  	
		 //Handle Error
		 if err != nil {
			log.Fatalf("An Error Occured %v", err)
		 }
		 defer resp.Body.Close()
	  	
		 //Read the response body
		 body, err := ioutil.ReadAll(resp.Body)
		 if err != nil {
			log.Fatalln(err)
		 }
		 sb := string(body)
		 log.Printf(sb)
>>>>>>> service2
	},
}

func init() {
	rootCmd.AddCommand(command3Cmd)
<<<<<<< HEAD
=======

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// command3Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// command3Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
>>>>>>> service2
}
