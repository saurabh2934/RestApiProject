package RestApiProject

func partition(arr []int, low, high int) int {
	pivot := arr[high]

	i := (low - 1)

	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}
	temp := arr[i+1]
	arr[i+1] = arr[high]
	arr[high] = temp
	return (i + 1)
}
func QuickFunction(arr []int, low, high int) []int {
	if low < high {
		pi := partition(arr, low, high)
		QuickFunction(arr, low, pi-1)
		QuickFunction(arr, pi+1, high)
	}
	return arr
}
