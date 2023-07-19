package RestApiProject

func SelSort(ar []int) {
	for i := 0; i < len(ar)-1; i++ {
		minElement := i
		for j := i + 1; j < len(ar); j++ {
			if ar[minElement] > ar[j] {
				minElement = j
			}
		}
		temp := ar[minElement]
		ar[minElement] = ar[i]
		ar[i] = temp

	}
}
