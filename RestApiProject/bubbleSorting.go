package RestApiProject

func BubbleSorting(ar []int) {
	for i := 0; i < len(ar)-1; i++ {
		for j := 0; j < len(ar)-i-1; j++ {
			if ar[j] > ar[j+1] {
				temp := ar[j]
				ar[j] = ar[j+1]
				ar[j+1] = temp
			}
		}
	}
}
