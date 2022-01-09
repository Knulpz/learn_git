package main

import "math"

func findShortestSubArray(nums []int) int {
	//使用map
	mapArray := make(map[int][]int, 0) //num => []pos
	//遍历数组
	for pos, v := range nums {
		mapArray[v] = append(mapArray[v], pos)
	}
	var target, max int
	min := math.MaxInt64
	for num, array := range mapArray {
		//筛选出现次数最多且长度最短的情况
		count := len(array)
		len := array[len(array)-1] - array[0] + 1
		//同时更新max和min的值
		if count > max {
			max = count
			min = len
			target = num
		} else if count == max && len < min {
			min = len
			target = num
		}
	}

	//返回target所指的长度
	arr := mapArray[target]
	return arr[len(arr)-1] - arr[0] + 1
}
