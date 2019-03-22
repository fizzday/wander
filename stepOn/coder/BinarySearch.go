package coder

// IBinarySearch
// 二分法查找及其变种

type IBinarySearch interface {
	// 参数释义
	// arr 给定有序数组
	// start 初始序号
	// end 结尾序号
	// length 数组长度
	// searchValue 待查找的值

	// 二分查找 - 递归查找
	Recursion(arr []int, start, end, searchValue int) int
	// 二分查找 - 常规非递归查找
	Normal(arr []int, searchValue int) int
	// 二分查找变种 - 查找与所查找的数相等的第一个数
	EqualFirst(arr []int, length, searchValue int) int
	// 二分查找变种 - 查找与所查找的数相等的第一个数
	EqualLast(arr []int, length, searchValue int) int
	// 二分查找变种 - 查找第一个大于key的值，也就是说返回大于key的最左边元素的下标。
	MoreThanOfFirst(arr []int, length, searchValue int) int
	// 二分查找变种 - 查找最后一个小于key的元素，也就是说返回小于key的最右边的元素下标。
	LessThanOfLast(arr []int, length, searchValue int) int
}

type BinarySearch struct {}

func (b *BinarySearch) LessThanOfLast(arr []int, length, searchValue int) int {
	var left = 0
	var right = length - 1
	if arr[left] >= searchValue {
		return -1
	}
	for left <= right {
		// 中间的key
		midKey := left + (right-left)/2
		if arr[midKey] >= searchValue {
			right = midKey - 1
		} else {
			left = midKey + 1
		}
	}

	return right
}

func (b *BinarySearch) MoreThanOfFirst(arr []int, length, searchValue int) int {
	var left = 0
	var right = length - 1
	if arr[right] <= searchValue {
		return -1
	}
	for left <= right {
		// 中间的key
		midKey := left + (right-left)/2
		if arr[midKey] <= searchValue {
			left = midKey + 1
		} else {
			right = midKey - 1
		}
	}

	return left
}

func (b *BinarySearch) EqualLast(arr []int, length, searchValue int) int {
	var left = 0
	var right = length - 1
	for left <= right {
		// 中间的key
		midKey := left + (right-left)/2
		if arr[midKey] > searchValue {
			right = midKey - 1
		} else {
			left = midKey + 1
		}
	}
	if right > 0 && arr[right] == searchValue {
		return right
	}
	return -1
}

func (b *BinarySearch) EqualFirst(arr []int, length, searchValue int) int {
	var left = 0
	var right = length - 1
	for left <= right {
		// 中间的key
		midKey := left + (right-left)/2
		if arr[midKey] >= searchValue {
			right = midKey - 1
		} else {
			left = midKey + 1
		}
	}
	// 如果左边key小于总长度,且左边key对应的值等于查找值, 则取出即为第一个相等的值
	if left < length && arr[left] == searchValue {
		return left
	}

	return -1
}

func (b *BinarySearch) Normal(arr []int, searchValue int) int {
	var left = 0
	var right = len(arr) - 1
	for {
		if left <= right {
			// 中间的key
			midKey := left + (right-left)/2
			if arr[midKey] == searchValue {
				return midKey
			} else if arr[midKey] > searchValue {
				right = midKey - 1
			} else {
				left = midKey + 1
			}
		} else {
			return -1
		}

	}
}

func (b *BinarySearch) Recursion(arr []int, start, end, searchValue int) int {
	if start <= end {
		// 中间的key
		midKey := start + (end-start)/2
		// 如果中间值等于查找值,则返回,若大于查找值,则对左边继续进行查找, 否则对右边继续查找
		if arr[midKey] == searchValue {
			return midKey
		} else if arr[midKey] > searchValue {
			return b.Recursion(arr, start, midKey-1, searchValue)
		} else {
			return b.Recursion(arr, midKey+1, end, searchValue)
		}
	}

	return -1
}
