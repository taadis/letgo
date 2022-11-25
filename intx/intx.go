package intx

// HasDuplicateValue0 判断数组中是否有重复值
// O(N^2):嵌套循环,步数指数级增长,性能较差
func HasDuplicateValue0(values []int) bool {
	for i := 0; i < len(values); i++ {
		for j := 0; j < len(values); j++ {
			if i != j && values[i] == values[j] {
				return true
			}
		}
	}
	return false
}

// HasDuplicateValue1 判断数组中是否有重复值
// O(N):更优实现,只有一层循环,但多了一个变量存储,空间占用会变大
func HasDuplicateValue1(values []int) bool {
	existingNumbers := make(map[int]int, len(values))
	for i := 0; i < len(values); i++ {
		if existingNumbers[values[i]] == 0 {
			existingNumbers[values[i]] = 1
		} else {
			return true
		}
	}
	return false
}
