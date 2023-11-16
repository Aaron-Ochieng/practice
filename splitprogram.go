package practice

func SplitProgram(str string, sep string) []string {
	// Initialize an empty string array
	var result []string

	// Initialize the start index and end index
	start := 0
	end := 0

	// Iterate over the string
	for i := 0; i < len(str); i++ {
		// Check if the current character is the separator
		if str[i:i+len(sep)] == sep {
			// Append the substring to the result array
			result = append(result, str[start:end])

			// Update the start index
			start = end + len(sep)

			// Update the end index
			end = start
		} else {
			// Increment the end index
			end++
		}
	}

	// Append the last substring to the result array
	result = append(result, str[start:end])

	return result
}
