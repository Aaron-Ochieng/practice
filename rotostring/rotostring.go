package practice

func RoToString(str string) string {
	s := str + " "
	words := []string{}
	word := ""
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && s[i+1] != ' ' {
			word += string(rune(s[i]))
		} else if s[i] != ' ' && s[i+1] == ' ' {
			word += string(rune(s[i]))
			words = append(words, word)
			word = ""
		}
	}
	arr := words[1:]
	arr = append(arr, words[0])

	returnString := ""

	for i := 0; i < len(arr); i++ {
		returnString += arr[i]
		if i < len(arr)-1 {
			returnString += " "
		}
	}
	return returnString
}
