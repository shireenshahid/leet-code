package cmd

import (
	"fmt"
	"leet-code/internal/containerwithmostwater"

	"github.com/spf13/cobra"
)

var ContainerWithMostWaterCmd = &cobra.Command{
	Use:   "container-with-most-water [solution]",
	Short: "Solve the Container With Most Water problem",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {

		height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
		switch args[0] {
		case "brute-force":
			result := containerwithmostwater.BruteForce(height)
			fmt.Println("Result:", result)

		case "two-pointer":
			result := containerwithmostwater.TwoPointer(height)
			fmt.Println("Result:", result)

		default:
			return fmt.Errorf("unknown solution: %s", args[0])
		}

		return nil
	},
}
