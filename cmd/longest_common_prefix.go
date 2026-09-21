package cmd

import (
	"fmt"
	"leet-code/internal/longestcommonprefix"

	"github.com/spf13/cobra"
)

var longestCommonPrefixCmd = &cobra.Command{
	Use:   "longest-common-prefix",
	Short: "Solve the Longest Common Prefix problem",

	RunE: func(cmd *cobra.Command, args []string) error {

		strs := []string{"flower", "flow", "flight"}

		result := longestcommonprefix.LongestCommonPrefix(strs)
		fmt.Println("Result:", result)

		return nil
	},
}
