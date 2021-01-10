package leetcode

import "strconv"

func summaryRanges(nums []int) []string {
	l := len(nums)
	if l == 0 {
		return []string{}
	}

	start, end, res := 0, 1, make([]string, 0, l)
	for start < l && end < l {
		if nums[end]-nums[end-1] != 1 {
			// 获取范围字符串
			res = append(res, getRangeStr(nums, start, end-1))
			start = end
		}
		end++
	}
	// 获取最后一个范围字符串
	res = append(res, getRangeStr(nums, start, end-1))

	return res
}

func getRangeStr(nums []int, start, end int) string {
	str := ""
	if start == end {
		str = strconv.Itoa(nums[start])
	} else {
		str = strconv.Itoa(nums[start]) + "->" + strconv.Itoa(nums[end])
	}
	return str
}
