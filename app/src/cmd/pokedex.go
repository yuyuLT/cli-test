/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

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
		isQuiz, _ := cmd.Flags().GetBool("quiz")
		text := getText(pokemon)

		if isQuiz {
			rand.Seed(time.Now().UnixNano())
			random := rand.Intn(898) + 1
			dummyText := getText(strconv.Itoa(random))

			choices := rand.Intn(2)

			if choices == 0 {
				fmt.Println(text)
			} else {
				fmt.Println(dummyText)
			}

			var answer string
			fmt.Print("Is this the Pokémon you chose? (y/n) > ")
			fmt.Scan(&answer)

			fmt.Println()
			if (answer == "y" && choices == 0) || (answer == "n" && choices == 1) {
				fmt.Println("Right!!")
			} else {
				fmt.Println("Wrong!!")
			}
			fmt.Println()

		} else {
			if text == "" {
				fmt.Println("データがありません")
			} else {
				fmt.Println(text)
			}
		}

	},
}

func getText(pokemon string) string {
	p, _ := pokeapi.PokemonSpecies(pokemon)
	flavor := p.FlavorTextEntries

	if flavor != nil {
		for i := 0; i < len(flavor); i++ {
			if flavor[i].Language.Name == "ja" {
				return flavor[i].FlavorText
			}
		}
	}

	return ""
}

func init() {
	rootCmd.AddCommand(pokedexCmd)
	pokedexCmd.Flags().BoolP("quiz", "q", false, "quiz mode")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pokedexCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pokedexCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
