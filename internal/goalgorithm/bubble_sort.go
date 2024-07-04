package goalgorithm

func BubbleSort(unSorted []int) []int {
	for i := 0; i < len(unSorted); i++ {
		for j := i + 1; j < len(unSorted); j++ {
			if unSorted[i] > unSorted[j] {
				unSorted[i], unSorted[j] = unSorted[j], unSorted[i]
			}
		}
	}
	return unSorted
}
