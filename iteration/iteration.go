package iteration

// Repeat characters for the specified number of times
func Repeat(character string, number int) string {
	var repeated string
	for i := 0; i < number; i++ {
		repeated += character
	}
	return repeated
}