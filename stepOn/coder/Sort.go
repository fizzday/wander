package coder

// 排序
type ISort interface{
	// BubbleSort 冒泡排序 - 先从数组中找到最大值(或最小值)并放到数组最左端(或最右端)，
	// 然后在剩下的数字中找到次大值(或次小值)，以此类推，直到数组有序排列。算法的时间复杂度为O(n^2)。
	// 最佳情况：T(n) = O(n)   最差情况：T(n) = O(n2)   平均情况：T(n) = O(n2)
	BubbleSort(arr []int) []int

	// SelectionSort 选择排序 - 首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置，然后，
	// 再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。以此类推，直到所有元素均排序完毕
	// 最佳情况：T(n) = O(n2)  最差情况：T(n) = O(n2)  平均情况：T(n) = O(n2)
	SelectionSort(arr []int) []int

	// InsertSort插入排序 - 通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。
	// 插入排序在实现上，通常采用in-place排序（即只需用到O(1)的额外空间的排序），因而在从后向前扫描过程中，
	// 需要反复把已排序元素逐步向后挪位，为最新元素提供插入空间。
	// 最佳情况：T(n) = O(n)   最坏情况：T(n) = O(n2)   平均情况：T(n) = O(n2)
	InsertSort(arr []int) []int

	// QuickSort 快速排序 - 通过一趟排序将待排记录分隔成独立的两部分，其中一部分记录的关键字
	// 均比另一部分的关键字小，则可分别对这两部分记录继续进行排序，以达到整个序列有序。
	// 最佳情况：T(n) = O(nlogn)   最差情况：T(n) = O(n2)   平均情况：T(n) = O(nlogn)　
	QuickSort(arr []int,start,end int) []int
}

type Sort struct{}

func (s *Sort) BubbleSort(arr []int) []int {
	if len(arr) > 0 {
		var i, j int

		for i = 0; i < len(arr); i++ {
			for j = 0; j < len(arr)-1-i; j++ {
				if arr[j+1] < arr[j] {
					arr[j],arr[j+1] = arr[j+1],arr[j]
				}
			}
		}
	}

	return arr
}

func (s *Sort) SelectionSort(arr []int) []int {
	if len(arr) > 0 {
		var i, j int

		for i = 0; i < len(arr); i++ {
			var minIndex = i
			for j = i; j < len(arr); j++ {
				if arr[j] < arr[minIndex] {
					minIndex = j
				}
			}
			arr[minIndex],arr[i] = arr[i],arr[minIndex]
		}
	}

	return arr
}

func (s *Sort) InsertSort(arr []int) []int {
	if len(arr) > 0 {
		var i int
		var current int

		for i = 0; i < len(arr)-1; i++ {
			var preIndex = i
			current = arr[i+1]
			for preIndex >= 0 && current < arr[preIndex] {
				arr[preIndex+1] = arr[preIndex]
				preIndex--
			}
			arr[preIndex+1] = current
		}
	}

	return arr
}

func (s *Sort) QuickSort(arr []int, startIndex, endIndex int) []int {
	if startIndex < endIndex {
		var i = startIndex
		var j = endIndex
		var baseValue = arr[i]
		for i < j {
			// 从右向左找比基准数小的数
			for i < j && arr[j] >= baseValue {
				j--
			}
			if i < j {
				arr[i] = arr[j]
				i++
			}
			// 从左向右找比基准数大的数
			for i < j && arr[i] < baseValue {
				i++
			}
			if i < j {
				arr[j] = arr[i]
				j--
			}
		}
		// 把基准数放到i的位置
		arr[i] = baseValue
		// 递归
		s.QuickSort(arr, startIndex, i-1);
		s.QuickSort(arr, i+1, endIndex);
	}
	return arr
}
