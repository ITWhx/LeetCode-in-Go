/**
*
*@author 吴昊轩
*@create 2019-08-2319:06
 */
package quick_sort

func quick_sort(s []int, left, right int) {
	if left < right {
		par := getPar(s, left, right)
		quick_sort(s, left, par-1)
		quick_sort(s, par+1, right)
	}
}

func getPar(s []int, left, right int) int {
	fast := left
	last := right

	leftLen := 0
	rightLen := 0
	l1 := left
	r2 := right
	tmp := s[left]
	for left < right {
		for right > left && s[right] >= tmp {
			if s[right] == tmp {
				s[right], s[r2] = s[r2], s[right]
				rightLen++
				r2--
			}
			right--
		}
		s[left] = s[right]
		for right > left && s[left] <= tmp {
			if s[left] == tmp {
				s[l1], s[left] = s[left], s[l1]
				leftLen++
				l1++
			}
			left++
		}
		s[right] = s[left]
	}
	s[left] = tmp

	i := left - 1
	j := fast
	for j < l1 && s[i] != tmp {
		s[j], s[i] = s[i], s[j]
		j++
		i--
	}

	i2 := left + 1
	j2 := last

	for j2 > r2 && s[i2] != tmp {
		s[i2], s[j2] = s[j2], s[i2]
		i2++
		j2--
	}

	return left
}

func setPartion(s []int, l, r int) {
	mid := (l + r) / 2

	if s[mid] > s[r] {
		s[mid], s[r] = s[r], s[mid]
	}
	if s[l] > s[mid] {
		s[l], s[mid] = s[mid], s[l]
	}

	if s[mid] < s[r] {
		s[mid], s[l] = s[l], s[mid]
	}
}
