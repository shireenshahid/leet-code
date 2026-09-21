package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "leet-code",
	Short: "my personal leetcode solution",
	Long:  `Leet is a CLI application that provides solutions to various coding problems and challenges. It allows users to explore different algorithms, data structures, and problem-solving techniques through an interactive command-line interface.`,

	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Welcome to Leet CLI! Use 'leet-code [command] --help' to see available commands.")
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.AddCommand(twoSumCmd)
	rootCmd.AddCommand(ContainerWithMostWaterCmd)
	rootCmd.AddCommand(removeDuplicatesCmd)
	rootCmd.AddCommand(longestCommonPrefixCmd)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
