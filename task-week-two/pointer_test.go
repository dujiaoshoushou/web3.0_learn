package main

import (
	"fmt"
	"testing"
)

func TestEditNun(t *testing.T) {
	var num int = 10
	EditNun(&num)
	fmt.Println(num)
}

func TestSliceNum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	SliceNum(&nums)
	fmt.Println(nums)
}

func TestSliceNum2(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	slice := []*int{&nums[0], &nums[1], &nums[2], &nums[3], &nums[4]}
	SliceNum2(slice)
	fmt.Println(nums)
}
