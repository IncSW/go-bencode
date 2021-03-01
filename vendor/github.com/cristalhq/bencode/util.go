package bencode

import "sort"

func sortStrings(ss []string) {
	if len(ss) <= strSliceLen {
		// for i := 1; i < len(ss); i++ {
		// 	for j := i; j > 0 && ss[j] < ss[j-1]; j-- {
		// 		ss[j], ss[j-1] = ss[j-1], ss[j]
		// 	}
		// }
		// below is the code above, but (almost) without bound checks

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
