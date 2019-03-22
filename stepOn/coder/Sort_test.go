package coder

import (
	"fmt"
	"testing"
)

var sort = &Sort{}
var prearr = []int{3, 8, 1, 22, 5}

func TestSort_BubbleSort(t *testing.T) {
	res := sort.BubbleSort(prearr)
	if res[0] != 1 || res[2] != 5 || res[4] != 22 {
		t.Error(fmt.Sprintf("Fail: %v", res))
	} else {
		t.Log(fmt.Sprintf("Pass: %v", res))
	}
}

func TestSort_SelectionSort(t *testing.T) {
	res := sort.SelectionSort(prearr)
	if res[0] != 1 || res[2] != 5 || res[4] != 22 {
		t.Error(fmt.Sprintf("Fail: %v", res))
	} else {
		t.Log(fmt.Sprintf("Pass: %v", res))
	}
}

func TestSort_InsertSort(t *testing.T) {
	res := sort.InsertSort(prearr)
	if res[0] != 1 || res[2] != 5 || res[4] != 22 {
		t.Error(fmt.Sprintf("Fail: %v", res))
	} else {
		t.Log(fmt.Sprintf("Pass: %v", res))
	}
}

func TestSort_QuickSort(t *testing.T) {
	res := sort.QuickSort(prearr, 0, 4)
	if res[0] != 1 || res[2] != 5 || res[4] != 22 {
		t.Error(fmt.Sprintf("Fail: %v", res))
	} else {
		t.Log(fmt.Sprintf("Pass: %v", res))
	}
}
