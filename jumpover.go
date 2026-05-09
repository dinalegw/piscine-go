package piscine

func JumpOver(str string) string {
	// If empty or has no 3rd character
	if str == "" || len(str) < 3 {
		return "\n"
	}

	result := ""

	// Take every 3rd character → indexes: 2, 5, 8, ...
	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}

	// Output must end with newline
	return result + "\n"
}
