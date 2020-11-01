package search

// LinearSearch Algorithm
func LinearSearch(input []int, target int) bool {
	for _, key := range input {
		if key == target {
			return true
		}
	}
	return false
}
