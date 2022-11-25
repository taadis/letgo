package algorithm

// SelectionSort 选择排序
// O(N^2/2):嵌套循环
//选择排序的步数大约只有冒泡排序的一般
// 也就是说选择排序比冒泡排序快一倍左右
func SelectionSort(values []int) []int {
	for i, n := 0, len(values); i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if values[min] > values[j] {
				min = j
			}
		}
		values[i], values[min] = values[min], values[i]
	}
	return values
}
