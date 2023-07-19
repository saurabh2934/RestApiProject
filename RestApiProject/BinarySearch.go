package RestApiProject

import "strconv"

func BinarySearch(ar []int, find int) string {
	str := ""
	si := 0
	ei := len(ar) - 1

	for si <= ei {
		mid := (si + ei) / 2
		if ar[mid] == find {
			fs := strconv.Itoa(find)
			itos := strconv.Itoa(mid)
			str = str + "Element " + fs + " find at index " + itos
			return str
		} else if ar[mid] < find {
			si = mid + 1
		} else {
			ei = mid - 1
		}
	}
	return "Element not found"
}
