package problem0011

func maxArea(height []int) int {
	// 从两端开始寻找，至少保证了宽度是最大值
	i, j := 0, len(height)-1
	max := 0

	for i < j {
		a, b := height[i], height[j]
		h := min(a, b)

		area := h * (j - i)
		if max < area {
			max = area
		}

		// 朝着area具有变大的可能性方向变化。
		if a < b {
			i++
		} else {
			j--
		}
	}

	return max
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func maxArea1(height []int) int {

	max, l, r := 0, 0, len(height)-1
	tmp := 0
	for l < r {

		if height[l] < height[r] {
			tmp = height[l] * (r - l)
			l++
		} else {
			tmp = height[r] * (r - l)
			r--
		}

		if tmp > max {
			max = tmp
		}

	}

	return max
}
