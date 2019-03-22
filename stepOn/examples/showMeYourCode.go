package main

import (
	"fmt"
	"math"
)

func main() {
	// 1.
	var str = []string{"a", "b", "c"}
	fmt.Println("第1题:")
	getRankPermutationAndCombination(str, 0, len(str))

	//// 2
	//fmt.Println("\n第2题:")
	//fmt.Println(getSecondMaxNumber([]int{1, 3, 9, 22, 3, 5, 99}, 7))
	//
	//// 3
	//fmt.Println("\n第3题:")
	//fmt.Println(checkValInArray([][]int{
	//	{1, 3, 5, 13, 26},
	//	{2, 5, 7, 16, 33},
	//	{6, 9, 12, 23, 55},
	//}, 12))
	//
	//// 4
	//fmt.Println(getBiggestSumFromMatrix([][]int{
	//	{1, 3, 5, 13, 26},
	//	{2, 5, 7, 16, 33},
	//	{6, 9, 12, 23, 55},
	//}, 3,5))
}

// 1. 输入一个字符串, 求它的所有排列组合.
// 全排列：还是abc三个字符，全排列即字符不能重复。最后 3*2 =6种结果
// 可以利用1中的方法，只要判断3个字符是否相等，都不相等的才是需要的全排列里的一个。这样的时间复杂度为n^n,而全排列的种类为n！所以需要设计一种n！的算法。
// 也可以利用递归，第一个字符串一共有n种选择，剩下的变成一个n-1规模的递归问题。而第一个字符的n种选择，都是字符串里面的。因此可以使用第一个字符与1-n的位置上进行交换，得到n中情况，然后递归处理n-1的规模，只是处理完之后需要在换回来，变成原来字符的样子。
func getRankPermutationAndCombination(data []string, start, dataLength int) {
	if dataLength == start+1 {
		fmt.Println(data)
	} else {
		for i := start; i < dataLength; i++ {
			// 让当前首位依次为后面的每一个数
			data[start], data[i] = data[i], data[start]
			// 递归后面的情况
			getRankPermutationAndCombination(data, start+1, dataLength)
		}
	}
}
// 2.
func getSecondMaxNumber(arr []int, n int) []int {
	var i, max, second_max int
	max = arr[0]
	for i = 1; i < n; i++ {
		if arr[i] > max {
			second_max = max
			max = arr[i]
		} else {
			if arr[i] > second_max {
				second_max = arr[i]
			}
		}
	}
	return []int{max, second_max}
}
// 3.
func checkValInArray(data [][]int, searchValue int) bool {
	var row = 0
	var col = len(data[0])-1
	for row<=len(data)-1 && col>=0 {
		if data[row][col] == searchValue {
			return true
		} else if data[row][col] < searchValue {
			row++
		} else {
			col--
		}
	}

	return false
}
// 4.
func getBiggestSumFromMatrix(data [][]int, row, col int) [][]int{
	var x,y,i,j,maxSum int
	for i=0;i<row-1;i++ {
		for j=0;j<col-1;j++ {
			tmpSum := data[i][j]+data[i][j+1]+data[i+1][j]+data[i+1][j+1]
			if tmpSum>maxSum {
				maxSum = tmpSum
				x=i
				y=j
			}
		}
	}
	return [][]int{
		{data[x][y],data[x][y+1]},
		{data[x+1][y],data[x+1][y+1]},
	}
}
// 5.
func snake(row,col int)  {
	var start=1
	var x,y int
	var res [][]int
	var circles = (row+1)/2

	for i:=1;i<=circles;i++ {

	}
}
