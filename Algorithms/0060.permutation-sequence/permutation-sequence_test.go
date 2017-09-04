package Problem0060

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

// para 是参数
type para struct {
	n int
	k int
}

// ans 是答案
type ans struct {
	one string
}

func Test_Problem0060(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				4,
				4,
			},
			ans{
				"1342",
			},
		},

		question{
			para{
				3,
				3,
			},
			ans{
				"213",
			},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, getPermutation(p.n, p.k), "输入:%v", p)
	}
}
