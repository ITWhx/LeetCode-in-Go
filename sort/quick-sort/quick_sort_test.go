/**
*
*@author 吴昊轩
*@create 2019-08-2319:28
 */
package quick_sort

import (
	"fmt"
	"testing"
)

type question struct {
	para para
	ans  ans
}

type para struct {
	one []int
}
type ans struct {
	one []int
}

func TestName(t *testing.T) {

	questions := []question{
		question{
			para{[]int{1, 3, 4, 5, 6, 2, 3}},
			ans{[]int{1, 2, 3, 3, 4, 5, 6}},
		},
	}

	for _, v := range questions {
		p := v.para
		fmt.Printf("%v\n", p)
		quick_sort(p.one, 0, len(p.one)-1)
		fmt.Printf("%v\n", p)
	}
}
