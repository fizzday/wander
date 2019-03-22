package coder

import (
	"fmt"
	"testing"
)

func TestBinarySearch_LessThanOfLast(t *testing.T) {
	bs := &BinarySearch{}

	var arr = []int{1,3,3,9,12,34,34}

	res := bs.LessThanOfLast(arr, len(arr), 9)

	if res==2 {
		t.Log(fmt.Sprintf("PASS: 查找结果为: %v", res))
	} else {
		t.Error("FAIL: 查找结果为: ", res)
	}
}

func TestBinarySearch_MoreThanOfFirst(t *testing.T) {
	bs := &BinarySearch{}

	var arr = []int{1,3,3,9,12,34,34}

	res := bs.MoreThanOfFirst(arr, len(arr), 3)

	if res==3 {
		t.Log(fmt.Sprintf("PASS: 查找结果为: %v", res))
	} else {
		t.Error("FAIL: 查找结果为: ", res)
	}
}

func TestBinarySearch_EqualLast(t *testing.T) {
	bs := &BinarySearch{}

	var arr = []int{1,3,3,9,12,34,34}

	res := bs.EqualLast(arr, len(arr), 12)

	if res==4 {
		t.Log(fmt.Sprintf("PASS: 查找结果为: %v", res))
	} else {
		t.Error("FAIL: 查找结果为: ", res)
	}
}

func TestBinarySearch_EqualFirst(t *testing.T) {
	bs := &BinarySearch{}

	var arr = []int{1,3,3,9,12,12,34}

	res := bs.EqualFirst(arr, len(arr), 12)

	if res==4 {
		t.Log(fmt.Sprintf("PASS: 查找结果为: %v", res))
	} else {
		t.Error("FAIL: 查找结果为: ", res)
	}
}

func TestBinarySearch_Normal(t *testing.T) {
	bs := &BinarySearch{}

	var arr = []int{1,3,9,12,34}

	res := bs.Normal(arr, 12)

	if res==3 {
		t.Log(fmt.Sprintf("PASS: 查找结果为: %v", res))
	} else {
		t.Error("FAIL: 查找结果为: ", res)
	}
}

func TestBinarySearch_Recursion(t *testing.T) {
	bs := &BinarySearch{}

	var arr = []int{1,3,9,12,34}

	res := bs.Recursion(arr,0,len(arr)-1, 12)

	if res==3 {
		t.Log(fmt.Sprintf("PASS: 查找结果为: %v", res))
	} else {
		t.Error("FAIL: 查找结果为: ", res)
	}
}
