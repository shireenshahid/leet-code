package cmd

import (
	"fmt"
	"leet-code/internal/twosum"

	"github.com/spf13/cobra"
)

var twoSumCmd = &cobra.Command{
	Use:   "two-sum [solution]",
	Short: "Solve the Two Sum problem",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {

		nums := []int{2, 7, 11, 15}
		target := 9

		switch args[0] {
		case "brute-force":
			result := twosum.BruteForce(nums, target)
			fmt.Println("Result:", result)

		case "hash-map":
			result := twosum.HashMap(nums, target)
			fmt.Println("Result:", result)

		default:
			return fmt.Errorf("unknown solution: %s", args[0])
		}

		return nil
	},
}
