package algorithm

// 冒泡排序
func BubbleSort(values []int) []int {
	for i, l := 0, len(values); i < l; i++ {
		for j := 0; j < l-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				continue
			}
		}
	}
	return values
}
