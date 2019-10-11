package problem0049

import "strconv"

func groupAnagrams(ss []string) [][]string {
	tmp := make(map[int][]string, len(ss)/2)
	for _, s := range ss {
		c := encode(s)
		tmp[c] = append(tmp[c], s)
	}

	res := make([][]string, 0, len(tmp))
	for _, v := range tmp {
		res = append(res, v)
	}

	return res
}

// prime 与 A～Z 对应，英文中出现概率越大的字母，选用的质数越小
var prime = []int{5, 71, 37, 29, 2, 53, 59, 19, 11, 83, 79, 31, 43, 13, 7, 67, 97, 23, 17, 3, 41, 73, 47, 89, 61, 101}

func encode(s string) int {
	res := 1
	for i := range s {
		res *= prime[s[i]-'a']
	}
	return res
}

func groupAnagrams2(strs []string) [][]string {
	var res [][]string
	m := make(map[string][]string)
	for _, str := range strs {
		key := getKey(str)
		m[key] = append(m[key], str)
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

// 用一个大小为26的int数组来统计每个单词中字符出现的次数，
// 然后将int数组转为一个唯一的字符串，跟字符串数组进行映射，
func getKey(s string) string {
	freq := make([]int, 26)
	for _, c := range s {
		freq[c-'a']++
	}
	key := ""
	for _, n := range freq {
		key += strconv.Itoa(n) + "/"
	}
	return key
}
