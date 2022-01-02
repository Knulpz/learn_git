package main

//Leetcode 66 加一
func plusOne(digits []int) []int {
	//1.暴力解法不可行，int会溢出
	//2.总结规律得：逆序遍历数组，得到第一个不为9得位，使其+1，并将之后的9全部置0；对于全为9的特殊情况，返回一个新值
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			for j := len(digits) - 1; j > i; j-- {
				//此处不需要判断是否为9，因为能运行到此处就说明了这一点
				digits[j] = 0
			}
			return digits
		}
	}
	//前面不return说明全是9
	digits = make([]int, len(digits)+1)
	digits[0] = 1
	return digits
}
