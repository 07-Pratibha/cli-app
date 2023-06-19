/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/spf13/cobra"
)

// Struct for storing the data from the dad joke json
type Response struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
   }

// command2Cmd represents the command2 command
var command2Cmd = &cobra.Command{
	Use:   "dadjoke",
	Short: "Gets a random dad joke from an online web app",
	Long: `Obtains a random dad joke from an online web app that hosts such jokes and stores it in a json. 
		That json's fields are then displayed to the user.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Calling the dadjoke.com API...")
		client := &http.Client{}
		req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
		if err != nil {
		fmt.Print(err.Error())
		}
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
		fmt.Print(err.Error())
		}
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
		fmt.Print(err.Error())
		}
		var responseObject Response
		json.Unmarshal(bodyBytes, &responseObject)
		// fmt.Printf("API Response as struct %+v\n", responseObject)
		fmt.Printf("Only the joke retrieved from the response .json stored in a struct: \n\n%s\n", responseObject.Joke)
	},
}

func init() {
	rootCmd.AddCommand(command2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// command2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// command2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
