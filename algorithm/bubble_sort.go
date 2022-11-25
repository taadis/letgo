package algorithm

// BubbleSort 冒泡排序
// O(N^2):循环嵌套,随着元素个数增长,步数却呈指数增长,比较低效
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
