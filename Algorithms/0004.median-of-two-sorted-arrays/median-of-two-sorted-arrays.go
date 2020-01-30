package problem0004

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := combine(nums1, nums2)
	return medianOf(nums)
}

func combine(mis, njs []int) []int {
	lenMis, i := len(mis), 0
	lenNjs, j := len(njs), 0
	res := make([]int, lenMis+lenNjs)

	for k := 0; k < lenMis+lenNjs; k++ {
		if i == lenMis ||
			(i < lenMis && j < lenNjs && mis[i] > njs[j]) {
			res[k] = njs[j]
			j++
			continue
		}

		if j == lenNjs ||
			(i < lenMis && j < lenNjs && mis[i] <= njs[j]) {
			res[k] = mis[i]
			i++
		}
	}

	return res
}

func medianOf(nums []int) float64 {
	l := len(nums)

	if l == 0 {
		panic("切片的长度为0，无法求解中位数。")
	}

	if l%2 == 0 {
		return float64(nums[l/2]+nums[l/2-1]) / 2.0
	}

	return float64(nums[l/2])
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	left := (len1 + len2 + 1) / 2
	right := (len1 + len2 + 2) / 2

	return float64(findMedian(nums1, 0, len1-1, nums2, 0, len2-1, left)+findMedian(nums1, 0, len1-1, nums2, 0, len2-1, right)) * 0.5
}

func findMedian(nums1 []int, start1, end1 int, nums2 []int, start2, end2, k int) int {
	len1 := end1 - start1 + 1
	len2 := end2 - start2 + 1

	if len1 == 0 {
		return nums2[start2+k-1]
	}
	if len2 == 0 {
		return nums1[start1+k-1]
	}
	if k == 1 {
		return min(nums1[start1], nums2[start2])
	}

	i := start1 + min(len1, k/2) - 1
	j := start2 + min(len2, k/2) - 1

	if nums1[i] > nums2[j] {
		return findMedian(nums1, start1, end1, nums2, j+1, end2, k-(j-start2+1))
	}
	return findMedian(nums1, i+1, end1, nums2, start2, end2, k-(i-start1+1))

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
