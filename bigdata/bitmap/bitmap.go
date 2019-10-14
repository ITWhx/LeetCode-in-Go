/**
*
*@author 吴昊轩
*@create 2019-08-2918:42
 */
package bitmap

type bitMap struct {
	bitArray  []byte
	arraySize uint32
}

func NewBitMap(size uint32) *bitMap {

	r := size / 8
	i := size % 8
	if i > 0 {
		r++
	}
	return &bitMap{make([]byte, r), r}
}

func (bm *bitMap) set(num uint32) {
	idx, pos := bm.calc(num)

	bm.bitArray[idx] |= 1 << pos
}

func (bm *bitMap) calc(num uint32) (idx, pos uint32) {
	idx = num >> 8
	if idx >= bm.arraySize {
		panic("数组越界")
		return
	}

	pos = num % 8
	return idx, pos
}
