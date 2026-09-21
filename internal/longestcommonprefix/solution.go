package longestcommonprefix

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Iterate through the characters of the first string
	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]

		// Check this character against the same position in all other strings
		for j := 1; j < len(strs); j++ {
			// If the current string is shorter than index 'i' or characters don't match
			if i == len(strs[j]) || strs[j][i] != char {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}
