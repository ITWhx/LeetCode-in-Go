package problem0022

func generateParenthesis(n int) []string {
	res := make([]string, 0, n*n)
	bytes := make([]byte, n*2)
	dfs(n, n, 0, bytes, &res)
	return res
}

func dfs(left, right, idx int, bytes []byte, res *[]string) {
	// 所有符号都添加完毕
	if left == 0 && right == 0 {
		*res = append(*res, string(bytes))
		return
	}

	// "(" 不用担心匹配问题，
	// 只要 left > 0 就可以直接添加
	if left > 0 {
		bytes[idx] = '('
		dfs(left-1, right, idx+1, bytes, res)
	}

	// 想要添加 ")" 时
	// 需要 left < right，
	// 即在 bytes[:idx] 至少有一个 "(" 可以与 这个 ")" 匹配
	if right > 0 && left < right {
		bytes[idx] = ')'
		dfs(left, right-1, idx+1, bytes, res)
	}
}

func generateParenthesis2(n int) []string {
	res := make([]string, n*n)
	item := ""
	generate(item, n, n, &res)
	return res
}

func generate(item string, left, right int, res *[]string) {
	//左括号和右括号满足上述条件的前提下都为n个，添加这个答案
	if left == 0 && right == 0 {
		*res = append(*res, item)
		return
	}
	//左括号的个数小于n 才能继续放左括号
	if left > 0 {
		generate(item+"(", left-1, right, res)
	}
	//左括号个数必须大于右括号的放置个数 才能继续放右括号
	if left < right {
		generate(item+")", left, right+1, res)
	}
}
