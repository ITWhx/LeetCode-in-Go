package problem0027

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
// one 代表第一个参数
type para struct {
	one []int
	two int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one []int
}

func Test_Problem0027(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[]int{3, 2, 2, 3}, 3},
			ans{[]int{2, 2}},
		},

		question{
			para{[]int{3, 1, 5, 7, 2, 2, 3}, 3},
			ans{[]int{2, 1, 5, 7, 2}},
		},

		question{
			para{[]int{1, 5, 7, 2, 2}, 3},
			ans{[]int{1, 5, 7, 2, 2}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, p.one[:removeElement(p.one, p.two)], "输入:%v", p)
	}
}

func TestName(t *testing.T) {
	arrInt32 := [...]uint32{5, 4, 2, 1, 3, 17, 13}

	var arrMax uint32 = 20
	bit := NewBitmap(arrMax)

	for _, v := range arrInt32 {
		bit.Set(v)
	}

	fmt.Println("排序后:")
	for i := uint32(0); i < arrMax; i++ {
		if k := bit.Test(i); k == 1 {
			fmt.Printf("%d ", i)
		}
	}
}

const (
	BitSize = 8 //一个字节8位
)

type Bitmap struct {
	BitArray  []byte
	ArraySize uint32
}

func NewBitmap(max uint32) *Bitmap {
	var r uint32
	switch {
	case max <= BitSize:
		r = 1
	default:
		r = max / BitSize
		if max%BitSize != 0 {
			r += 1
		}
	}

	fmt.Println("数组大小:", r)
	return &Bitmap{BitArray: make([]byte, r), ArraySize: r}
}

func (bitmap *Bitmap) Set(i uint32) {
	idx, pos := bitmap.calc(i)
	bitmap.BitArray[idx] |= 1 << pos
	fmt.Println("set()  value=", i, " idx=", idx, " pos=", pos, ByteToBinaryString(bitmap.BitArray[idx]))
}

func (bitmap *Bitmap) Test(i uint32) byte {
	idx, pos := bitmap.calc(i)
	return bitmap.BitArray[idx] >> pos & 1
}

func (bitmap *Bitmap) Clear(i uint32) {
	idx, pos := bitmap.calc(i)
	bitmap.BitArray[idx] &^= 1 << pos
}

func (bitmap *Bitmap) calc(i uint32) (idx, pos uint32) {

	idx = i >> 3 //相当于i / 8,即字节位置
	if idx >= bitmap.ArraySize {
		panic("数组越界.")
		return
	}
	pos = i % BitSize //位位置
	return
}

//ByteToBinaryString函数来源:
// Go语言版byte变量的二进制字符串表示
// http://www.sharejs.com/codes/go/4357
func ByteToBinaryString(data byte) (str string) {
	var a byte
	for i := 0; i < 8; i++ {
		a = data
		data <<= 1
		data >>= 1

		switch a {
		case data:
			str += "0"
		default:
			str += "1"
		}

		data <<= 1
	}
	return str
}
