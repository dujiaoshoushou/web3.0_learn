package main

/*
*
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
示例 1：

输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：

输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。

提示：

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] 如果非空，则仅由小写英文字母组成
*/
func main() {

	var strs = []string{"fliower", "fliow", "flight"}
	println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}
	minLen := len(strs[0])
	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
		}
	}
	for i := 0; i < minLen; i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != c {
				return strs[0][:i]
			}
		}
	}
	return strs[0][:minLen]
}
