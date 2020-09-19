package algorithm

// 插入排序
func InsertSort(values []int) []int {
	for i := 1; i < len(values); i++ {
		temp := values[i] // 把当前元素拎出来, 方便后续顺序比对以及插队
		j := i
		for j > 0 && temp < values[j-1] {
			values[j] = values[j-1] // 每个大的元素往后移
			j--
		}
		values[j] = temp
	}
	return values
}
