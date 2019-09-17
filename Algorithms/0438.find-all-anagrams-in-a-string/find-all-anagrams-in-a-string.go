package problem0438

func findAnagrams(s string, p string) []int {
	//var res = []int{}
	//
	//var target, window [26]int
	//for i := 0; i < len(p); i++ {
	//	target[p[i]-'a']++
	//}
	//
	//check := func(i int) {
	//	if window == target {
	//		res = append(res, i)
	//	}
	//}
	//
	//for i := 0; i < len(s); i++ {
	//	window[s[i]-'a']++
	//	if i == len(p)-1 {
	//		check(0)
	//	} else if len(p) <= i {
	//		window[s[i-len(p)]-'a']--
	//		check(i - len(p) + 1)
	//	}
	//}
	//
	//return res
	var res []int
	var target, window [26]int

	for i := 0; i < len(p); i++ {
		target[p[i]-'a']++
	}

	left, right := 0, 0
	//当窗口右边出界 则退出
	for right < len(s) {
		//获取当前最右边的字符的hash值
		idx := s[right] - 'a'
		window[idx]++
		right++
		//当前字符的hash值 大于 目标数组（溢出） 1.当前字符不在目标数组中 2.当前字符有重复数组 左边界右移缩小窗口
		for target[idx] < window[idx] {
			window[s[left]-'a']--
			left++
		}
		//当出现窗口大小等于目标数组的大小的时候 匹配成功+1
		if right-left == len(p) {
			res = append(res, left)
		}
	}
	return res
}
