package piscine

func ShoppingListSort(slice []string) []string {
	// Simple selection sort based on string length
	for i := 0; i < len(slice)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(slice); j++ {
			if len(slice[j]) < len(slice[minIndex]) {
				minIndex = j
			}
		}
		// Swap
		slice[i], slice[minIndex] = slice[minIndex], slice[i]
	}
	return slice
}
