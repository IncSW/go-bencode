package internal

import "sort"

const strSliceLen = 20

func SortStrings(ss []string) {
	if len(ss) <= strSliceLen {
		for i := 1; i < len(ss); i++ {
			for j := i; j > 0; j-- {
				if ss[j] >= ss[j-1] {
					break
				}
				ss[j], ss[j-1] = ss[j-1], ss[j]
			}
		}
	} else {
		sort.Strings(ss)
	}
}
