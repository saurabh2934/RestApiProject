package RestApiProject

import "strconv"

func LinearSearching(ar []int, find int) string {
	str := ""
	i := 0
	for ; i < len(ar)-1; i++ {
		if find == ar[i] {
			fs := strconv.Itoa(find)
			itos := strconv.Itoa(i)
			str = str + "Element " + fs + " is available on index " + itos
			return str
		}
	}
	return "Element not Found"
}
