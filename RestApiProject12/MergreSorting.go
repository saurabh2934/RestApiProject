package RestApiProject

func merge(s1, s2, arr []int) {
	i := 0
	k := 0
	j := 0
	for i < len(s1) && j < len(s2) {
		if s1[i] < s2[j] {
			arr[k] = s1[i]
			i++
			k++
		} else {
			arr[k] = s2[j]
			j++
			k++
		}
	}
	if i < len(s1) {
		for i < len(s1) {
			arr[k] = s1[i]
			i++
			k++
		}
	}
	if j < len(s2) {
		for j < len(s2) {
			arr[k] = s2[j]
			j++
			k++
		}
	}
}
func MergerSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	n := len(arr) / 2
	//n1 := len(arr) - n
	var a []int
	var b []int
	// let's copy the element in both split array
	for i := 0; i < n; i++ {
		a = append(a, arr[i])
	}
	for i := n; i < len(arr); i++ {
		b = append(b, arr[i])
	}
	MergerSort(a)
	MergerSort(b)
	merge(a, b, arr)
	return arr
}
