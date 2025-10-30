package Lv1

func BucketSort(n []int) map[int]int {
	m := make(map[int]int)
	for i := 0; i < len(n); i++ {
		m[n[i]]++
	}
	return m
}

//func main() {
//	var length int
//	fmt.Printf("请输入数组长度：")
//	fmt.Scanf("%d", &length)
//	arr := make([]int, length)
//	fmt.Printf("\n请输入数组：")
//	for i := 0; i < length; i++ {
//		fmt.Scanf("%d", &arr[i])
//	}
//	fmt.Print("\n")
//	bucketMap := BucketSort(arr)
//	for k, v := range bucketMap {
//		fmt.Printf("元素`%d`在数组中出现了%d次\n", k, v)
//	}
//}
