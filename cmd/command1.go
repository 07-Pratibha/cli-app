/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
	"github.com/spf13/cobra"
)

// A Response struct to map the Entire Response
type ResponseP struct {
    Name    string    `json:"name"`
    Pokemon []Pokemon `json:"pokemon_entries"`
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
    EntryNo int            `json:"entry_number"`
    Species PokemonSpecies `json:"pokemon_species"`
}

// A struct to map our Pokemon's Species which includes it's name
type PokemonSpecies struct {
    Name string `json:"name"`
}

// command1Cmd represents the command1 command
var command1Cmd = &cobra.Command{
	Use:   "pokedex",
	Short: "Displays a list of Pokemons from the Kanto region",
	Long: `Obtains the Kanto pokedex from a web service through an API call and displays the list of pokemon from it.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Calling the pokeapi.co API...")
		response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		var responseObject ResponseP
		json.Unmarshal(responseData, &responseObject)

		fmt.Println(responseObject.Name)
		fmt.Println(len(responseObject.Pokemon), "pokemon in this region.")

		for i := 0; i < len(responseObject.Pokemon); i++ {
			fmt.Println("#", i+1, ": ",responseObject.Pokemon[i].Species.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(command1Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// command1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// command1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
