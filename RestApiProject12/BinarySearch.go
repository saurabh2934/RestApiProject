package RestApiProject

import (
	"errors"
	"strconv"
)

func BinarySearch(ar []int, find int) error {
	str := ""
	si := 0
	ei := len(ar) - 1

	for si <= ei {
		mid := (si + ei) / 2
		if ar[mid] == find {
			fs := strconv.Itoa(find)
			itos := strconv.Itoa(mid)
			str = str + "Element " + fs + " find at index " + itos
			return errors.New(str)
		} else if ar[mid] < find {
			si = mid + 1
		} else {
			ei = mid - 1
		}
	}
	return errors.New("Element Not! Found")
}
