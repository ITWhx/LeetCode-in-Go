package problem0190

func reverseBits(n uint32) uint32 {
	n = (n >> 16) | (n << 16)
	n = ((n & 0xff00ff00) >> 8) | ((n & 0x00ff00ff) << 8)
	n = ((n & 0xf0f0f0f0) >> 4) | ((n & 0x0f0f0f0f) << 4)
	n = ((n & 0xcccccccc) >> 2) | ((n & 0x33333333) << 2)
	n = ((n & 0xaaaaaaaa) >> 1) | ((n & 0x55555555) << 1)
	return n
}

func reverseBits2(n uint32) uint32 {
	res := uint32(0)

	for i := uint32(0); i < 32; i++ {
		tmp := n >> i
		tmp = tmp & 1
		tmp = tmp<<32 - i
		res = res | tmp
	}

	return res
}

// for 8 bit binary number abcdefgh, the process is as follow:
// abcdefgh -> efghabcd -> ghefcdab -> hgfedcba
