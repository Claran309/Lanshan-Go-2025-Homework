package lv2_func_test

import (
	"Lesson_1/Lanshan-lesson2/lv2/lv2_func"
	"math"
	"testing"
)

func TestInToPost(t *testing.T) {
	result := lv2_func.Run("114*(5.20+13.14)-514/191+98.10")
	result = math.Round(result*10000) / 10000 // math.Round - 四舍五入函数
	expect := float64(2186.1689)
	if result != expect {
		t.Errorf("InToPost expect %f, got %f", expect, result)
	}
}
