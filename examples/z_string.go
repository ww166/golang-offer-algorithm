package examples

func ZConvert(s string, numRows int) string {
	sLen := len(s)
	step := numRows*2 - 2
	var ret string

	if 1 == numRows {
		return s
	}

	// https://pic.leetcode-cn.com/d610b140dd0789204efe699672dc72a83e7b826da0165bbf083d24fc97ecdea7-image.png
	// 0                       2*n-2
	// 1                 2*n-3 2*n-1
	// 2          2*n-4        2*n
	// .            .           .
	// .         .              .
	// .      .                 .
	// n-2 n                   3*n-4
	// n-1                     3*n-3

	// 第一行或者最后一行的 间距为 step
	// 其他行间距: step(eg numRows * 2 - 2) - 2 * iRow(以0起始，行数序号), 2 *  iRow 交互循环
	for x := 0; x < numRows; x++ {
		idx := x
		offset := x * 2

		for idx < sLen {
			ret += string(s[idx])
			offset = step - offset
			if 0 == x || numRows-1 == x {
				idx += step
			} else {
				idx += offset
			}
		}
	}

	return ret
}
