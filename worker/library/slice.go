package library

// SliceFind searchs an item in the slice
func SliceFind(slice []string, val string) int {
	for i, item := range slice {
		if item == val {
			return i
		}
	}
	return -1
}
