package leetcode

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	} else if rowIndex == 1 {
		return []int{1, 1}
	}
	res, head, tail, last, curr, level := make([]int, rowIndex+2), 0, 2, 0, 0, 1
	res[head], res[head+1] = 1, 1
	for level < rowIndex {
		// 出队
		curr = res[head]
		head = (head + 1) % (rowIndex + 2)
		// 入队
		res[tail] = curr + last
		tail = (tail + 1) % (rowIndex + 2)
		// 判断是否在行尾，last不为0且curr为1
		if last != 0 && curr == 1 {
			res[tail] = 1
			tail = (tail + 1) % (rowIndex + 2)
			last = 0
			level++
		} else {
			last = curr
		}
	}
	if head == 0 {
		return res[head:tail]
	}
	return append(res[head:], res[0:tail]...)
}
