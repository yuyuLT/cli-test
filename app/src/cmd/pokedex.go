/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"github.com/mtslzr/pokeapi-go"
	"github.com/spf13/cobra"
)

// pokedexCmd represents the pokedex command
var pokedexCmd = &cobra.Command{
	Use:   "pokedex",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		
		pokemon := args[0] 

		p, _ := pokeapi.PokemonSpecies(pokemon)
		flavor := p.FlavorTextEntries

		if(flavor != nil){
			for i := 0; i < len(flavor); i++ {
				if flavor[i].Language.Name == "ja" {
					fmt.Println(flavor[i].FlavorText)
					os.Exit(0)
				}
			}
		}

		fmt.Println("データがありません")

		
	
	

	},
}

func init() {
	rootCmd.AddCommand(pokedexCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pokedexCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pokedexCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
