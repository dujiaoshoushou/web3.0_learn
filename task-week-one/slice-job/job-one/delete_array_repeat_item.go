package main

import (
	"fmt"
)

/*
*给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。

考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：

更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
返回 k 。
判题标准:

系统会用下面的代码来测试你的题解:

int[] nums = [...]; // 输入数组
int[] expectedNums = [...]; // 长度正确的期望答案

int k = removeDuplicates(nums); // 调用

assert k == expectedNums.length;

	for (int i = 0; i < k; i++) {
	    assert nums[i] == expectedNums[i];
	}

如果所有断言都通过，那么您的题解将被 通过。
示例 1：

输入：nums = [1,1,2]
输出：2, nums = [1,2,_]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
示例 2：

输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]
解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。

提示：

1 <= nums.length <= 3 * 104
-104 <= nums[i] <= 104
nums 已按 非严格递增 排列
*/
func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 4, 4}
	num, slice := removeDuplicates3(&nums)
	//num := removeDuplicates4(nums)
	fmt.Println(num, slice)
	//fmt.Println(num)
}

func removeDuplicates(nums []int) (int, []int) {
	var slice []int
lab:
	for _, value := range nums {
		if slice == nil {
			slice = append(slice, value)
		} else {
			for _, v := range slice {
				if value == v {
					continue lab
				}
			}
			slice = append(slice, value)
		}
	}
	return len(slice), slice
}

func removeDuplicates2(nums []int) (int, []int) {
	var slice []int

	for _, value := range nums {
		if slice == nil {
			slice = append(slice, value)
		} else {
			flag := true
			for _, v := range slice {
				if value == v {
					flag = false
					break
				}
			}
			if flag {
				slice = append(slice, value)
			}
		}
	}

	return len(slice), slice
}

func removeDuplicates3(nums *[]int) (int, []int) {

	slow_index := 1
	for fast_index := 1; fast_index < len(*nums); fast_index++ {
		if (*nums)[fast_index] != (*nums)[fast_index-1] {
			(*nums)[slow_index] = (*nums)[fast_index]
			slow_index++
		}
	}
	return slow_index, (*nums)
}

func removeDuplicates4(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 1 // 慢指针初始位置（第一个元素必然保留）
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast] // 写入新唯一元素
			slow++
		}
	}
	return slow
}
