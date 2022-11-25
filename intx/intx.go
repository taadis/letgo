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
