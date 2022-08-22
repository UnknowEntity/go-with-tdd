package iteration

// Repeat take the character input and return a string of them repeat count times
func Repeat(character string, count int) string {
	var result string
	for i := 0; i < count; i++ {
		result += character
	}

	return result
}
