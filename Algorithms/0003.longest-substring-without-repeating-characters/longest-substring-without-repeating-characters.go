package problem0003

func lengthOfLongestSubstring(s string) int {
	// location[s[i]] == j 表示：
	// s中第i个字符串，上次出现在s的j位置，所以，在s[j+1:i]中没有s[i]
	// location[s[i]] == -1 表示： s[i] 在s中第一次出现
	location := [256]int{} // 只有256长是因为，假定输入的字符串只有ASCII字符
	for i := range location {
		location[i] = -1 // 先设置所有的字符都没有见过
	}

	maxLen, left := 0, 0

	for i := 0; i < len(s); i++ {
		// 说明s[i]已经在s[left:i+1]中重复了
		// 并且s[i]上次出现的位置在location[s[i]]
		if location[s[i]] >= left {
			left = location[s[i]] + 1 // 在s[left:i+1]中去除s[i]字符及其之前的部分
		} else if i+1-left > maxLen {
			// fmt.Println(s[left:i+1])
			maxLen = i + 1 - left
		}
		location[s[i]] = i
	}

	return maxLen
}

func lengthOfLongestSubstring2(s string) int {
	start, maxlen := 0, 0
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if idx, ok := m[s[i]]; ok {
			//注意这里的判断条件！！！idx>=start   idx表示出现重复字符，如果重复字符在start之前，就没必要更新start了
			if idx >= start {
				start = idx + 1
			}
		}
		m[s[i]] = i
		maxlen = max(maxlen, i-start+1)
	}
	return maxlen
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
