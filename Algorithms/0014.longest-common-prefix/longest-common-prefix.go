package problem0014

func longestCommonPrefix(strs []string) string {
	short := shortest(strs)
	for i, r := range short {
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != byte(r) {
				return strs[j][:i]
			}
		}
	}

	return short
}

func shortest(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	res := strs[0]
	for _, s := range strs {
		if len(res) > len(s) {
			res = s
		}
	}

	return res
}

func longestCommonPrefix2(strs []string) string {
	short := shortest2(strs)

	for i, v := range short {
		for _, c := range strs {
			if c[i] != byte(v) {
				return short[:i]
			}
		}
	}

	return short

}

func shortest2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	short := strs[0]

	for _, v := range strs {
		if len(v) < len(short) {
			short = v
		}
	}
	return short
}
