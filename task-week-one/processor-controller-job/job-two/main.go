package main

/*
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
例如，121 是回文，而 123 不是。
示例 1：
输入：x = 121
输出：true

示例 2：
输入：x = -121
输出：false
解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。

示例 3：
输入：x = 10
输出：false
解释：从右向左读, 为 01 。因此它不是一个回文数。
*/
func main() {
	var x_num int = 121
	//var x_num1 int = -121
	println(isPalindrome(x_num))
}

func isPalindrome(num int) bool {
	if num < 0 || (num%10 == 0 && num != 0) {
		return false
	}

	var i int = 0
	for i < num {
		i = i*10 + num%10
		num /= 10
	}
	if i == num || num == i/10 {
		return true
	} else {
		return false
	}
}
