package problem0387

import "fmt"

func firstUniqChar(s string) int {
	rec := make([]int, 26)
	for _, b := range s {
		i := b - 'a'
		fmt.Println(i)
		rec[b-'a']++
	}

	for i, b := range s {
		if rec[b-'a'] == 1 {
			return i
		}
	}

	return -1
}
