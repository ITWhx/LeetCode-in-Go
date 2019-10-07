/**
*
*@author 吴昊轩
*@create 2019-10-0315:42
 */
package heap

type Interface interface {
	Less(i, j int) bool
	Swap(i, j int)
	Len() int
	Push(x interface{})
	Pop() interface{}
}

func Init(x Interface) {
	n := x.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down(x, i, n)
	}
}

func Push(h Interface, x interface{}) {
	n := h.Len()
	h.Push(x)
	up(h, n)
}

func Pop(h Interface) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

func Remove(h Interface, i int) interface{} {
	n := h.Len() - 1
	if n != i {
		h.Swap(i, n)
		Fix(h, i)
	}
	return h.Pop()
}
func Fix(x Interface, i int) {
	if !down(x, i, x.Len()) {
		up(x, i)
	}
}
func up(h Interface, i int) {

	for {
		j := (i - 1) / 2
		if j == i || h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
}
func down(h Interface, i0, n int) bool {
	i := i0

	for {
		j1 := 2*i + 1
		//left <0 after int overflow
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2
		}

		if h.Less(i, j) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	//如果i>i0表示满足堆结构，未向下调整
	return i > i0
}
