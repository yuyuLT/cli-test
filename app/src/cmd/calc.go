/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"
	"errors"
	"github.com/spf13/cobra"
)

// calcCmd represents the calc command
var calcCmd = &cobra.Command{
	Use:   "calc",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 3 {
			return errors.New("requires one sign and two numbers.")
		}

		if args[0] != "add" && args[0] != "sub" {
			return errors.New("sign is 'add' or 'sub'.")
		}

		_, err := strconv.Atoi(args[1])
		if err != nil{
			fmt.Println(args[1])
			return errors.New("This argument could not be converted.")
		}

		_, err = strconv.Atoi(args[2])
		if err != nil{
			fmt.Println(args[2])
			return errors.New("This argument could not be converted.")
		}
		
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {

		result := 0
		num1, _ := strconv.Atoi(args[1])
		num2, _ := strconv.Atoi(args[2])
		if args[0] == "add"{
			result = num1 + num2
		}else if args[0] == "sub"{
			result = num1 - num2
		}

		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(calcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// calcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// calcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
