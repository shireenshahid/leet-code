package cmd

import (
	"fmt"
	"leet-code/internal/removeduplicates"

	"github.com/spf13/cobra"
)

var removeDuplicatesCmd = &cobra.Command{
	Use:   "remove-duplicates [solution]",
	Short: "Solve the Remove Duplicates from Sorted Array problem",

	RunE: func(cmd *cobra.Command, args []string) error {

		nums := []int{1, 1, 2, 2, 3, 3}

		result := removeduplicates.RemoveDuplicates(nums)
		fmt.Println("Result:", result)

		return nil
	},
}
