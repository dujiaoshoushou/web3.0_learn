package main

import "strconv"

/*
给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间
*/
func main() {
	var arrs = []int{1, 1, 2, 2, 3, 4, 4, 5, 5}
	var maps = make(map[string]int)
	for i := 0; i < len(arrs); i++ {
		value, flag := maps[strconv.Itoa(arrs[i])]
		if flag == true {
			maps[strconv.Itoa(arrs[i])] = (value + 1)
		} else {
			maps[strconv.Itoa(arrs[i])] = 1
		}
	}
	for key, value := range maps {
		if value == 1 {
			println(key)
		}
	}
}
