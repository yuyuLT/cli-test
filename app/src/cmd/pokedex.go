/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/mtslzr/pokeapi-go"
	"github.com/spf13/cobra"
)

var pokedexCmd = &cobra.Command{
	Use:   "pokedex",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		language := "ja"
		isEn, _ := cmd.Flags().GetBool("english")
		if isEn {
			language = "en"
		}

		pokemon := args[0]
		isQuiz, _ := cmd.Flags().GetBool("quiz")
		text := getText(pokemon, language)

		if isQuiz {
			rand.Seed(time.Now().UnixNano())
			random := rand.Intn(898) + 1
			dummyText := getText(strconv.Itoa(random), language)

			choices := rand.Intn(2)

			if choices == 0 {
				fmt.Println(text)
			} else {
				fmt.Println(dummyText)
			}

			var answer string

			fmt.Println()
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
				name := getName(pokemon, language)
				fmt.Println(name, text)
			}
		}

	},
}

func getText(pokemon string, language string) string {
	p, _ := pokeapi.PokemonSpecies(pokemon)
	flavor := p.FlavorTextEntries

	if flavor != nil {
		for i := 0; i < len(flavor); i++ {
			if flavor[i].Language.Name == language {
				return strings.ReplaceAll(flavor[i].FlavorText, "\n", " ")
			}
		}
	}

	return ""
}

func getName(pokemon, language string) string {
	p, _ := pokeapi.PokemonSpecies(pokemon)
	names := p.Names

	if names != nil {
		for i := 0; i < len(names); i++ {
			if names[i].Language.Name == language {
				return strings.ReplaceAll(names[i].Name, "\n", " ")
			}
		}
	}

	return ""
}

func init() {
	rootCmd.AddCommand(pokedexCmd)
	pokedexCmd.Flags().BoolP("quiz", "q", false, "quiz mode")
	pokedexCmd.Flags().BoolP("english", "e", false, "english mode")
}
