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

	tmp := s[left]
	for left < right {
		for right > left && s[right] >= tmp {
			right--
		}
		s[left] = s[right]
		for right > left && s[left] <= tmp {
			left++
		}
		s[right] = s[left]
	}
	s[left] = tmp
	return left
}
