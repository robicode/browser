package util

// Squeeze removes duplicates from s
func Squeeze(s string) string {
	result := make([]rune, 0)
	var previous rune
	for _, rune := range s {
		if rune != previous {
			result = append(result, rune)
		}
		previous = rune
	}
	return string(result)
}
