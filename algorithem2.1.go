package main

import (
	"fmt"
	"strconv"
)

//首先这道题显然是用map来处理的
func subdomainVisits(cpdomains []string) []string {
	res := make([]string, 0)
	mapDomainCnt := make(map[string]int)
	//遍历所有的域名
	for _, k := range cpdomains {
		bytes := []byte(k)
		//每个域名的次数
		byteNum := make([]byte, 0)
		//每个域名一个栈
		stack := make([]byte, 0)
		stack = append(stack, ' ') //在域名前放个空格用于判断结束
		for _, v := range bytes {
			//数字
			if v >= '0' && v <= '9' {
				byteNum = append(byteNum, v)
				continue
			}
			//空格
			if v == ' ' {
				continue
			}
			//域名的部分先存到栈里面，先进后出
			if v >= 'a' && v <= 'z' || v == '.' {
				stack = append(stack, v)
				continue
			}
		}
		//处理域名，从后开始遍历
		for i := len(stack) - 1; i >= 0; i-- {
			//遇到.则将后面的域名放到map中
			if stack[i] == '.' || stack[i] == ' ' {
				fmt.Println("test", string(stack[i+1:len(stack)]), FlushByteSLiceToInt(byteNum))
				mapDomainCnt[string(stack[i+1:len(stack)])] += FlushByteSLiceToInt(byteNum)
			}
		}
	}

	for k, v := range mapDomainCnt {
		singleCp := make([]byte, 0)
		singleCp = append(singleCp, FlushIntToByteSLice(v)...) //次数
		singleCp = append(singleCp, ' ')                       //空格
		singleCp = append(singleCp, k...)                      //域名
		res = append(res, string(singleCp))
	}

	return res
}

func FlushByteSLiceToInt(byte_slice []byte) int {
	res, _ := strconv.Atoi(string(byte_slice))
	return res
}

func FlushIntToByteSLice(value int) []byte {
	res := []byte(strconv.Itoa(value))
	return res
}
