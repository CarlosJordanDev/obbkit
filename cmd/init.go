/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	project string
	name    string
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Lists available starterkits on onboardbase.",
	Long:    `Lists avaialable starterkits on onboardbase.`,
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Enter project type: ")
		fmt.Scan(&project)
		fmt.Printf("Enter project name: ")
		fmt.Scan(&name)
		input := name
		switch project {
		case "reactjs":
			fmt.Println("You reactjs project has been created with title name of ->>", input)
		case "vuejs":
			fmt.Println("You vuejs project has been created with title name of ->>", input)
		case "flask":
			fmt.Println("You flask project has been created with title name of ->>", input)
		default:
			const colorRed = "\033[0;31m"
			const colorNone = "\033[0m"
			fmt.Fprintln(os.Stdout, "None:", colorRed, "please enter a valid project type.")
			fmt.Fprintln(os.Stdout, "Help:", colorNone, "Use the obbkit list command to see available project types.")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.
	//rootCmd.Flags().StringVarP(&project, "project", "p", "YOUR PROJECT", "project type (required)")
	//rootCmd.MarkFlagRequired("project")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
