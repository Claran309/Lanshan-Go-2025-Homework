package lv1_test

import (
	"Lesson_1/Lanshan-lesson2/lv1"
	"reflect"
	"testing"
)

func TestBucketSort(t *testing.T) {
	n := []int{1, 1, 4, 5, 1, 4, 1, 9, 1, 9, 8, 1, 0}
	result := Lv1.BucketSort(n)
	expect := map[int]int{
		0: 1,
		1: 6,
		4: 2,
		5: 1,
		8: 1,
		9: 2,
	}
	if !reflect.DeepEqual(result, expect) { //由于不知道怎么比较两个map是否相等，我去浅学了以下reflet库
		t.Errorf("BucketSort failed, expect:%v, got:%v", expect, result)
	}
}
