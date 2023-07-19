package RestApiProject

func InsertionSoting(ar []int) {
	var i = 0
	var j, temp int
	for i = 1; i < len(ar); i++ {
		temp = ar[i]
		j = i - 1
		for j >= 0 && temp <= ar[j] {
			ar[j+1] = ar[j]
			j = j - 1
		}
		ar[j+1] = temp
	}

}
