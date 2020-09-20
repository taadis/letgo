package algorithm

// SelectionSort
func SelectionSort(values []int) []int{
	for i, n := 0, len(values) ;i < n; i++{
		min := i
		for j := i + 1; j < n; j ++{
			if values[min] > values[j] {
				min = j
			}
		}
		values[i], values[min]= values[min], values[i]
	}
	return values
}
