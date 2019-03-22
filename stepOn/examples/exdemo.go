package main

import "fmt"

func main()  {
	findBigestNumber()
}

func findBigestNumber(arr []int, count int) []int {
	var tmp []int
	for item:=range arr {
		// 第1~count个直接放入tmp

		// 第 count+1 个开始与tmp中的数据对比, 比tmp中的数据小则移除, 比tmp中的任意一个大, 则移除tmp中最小的数据
	}
	return arr
}