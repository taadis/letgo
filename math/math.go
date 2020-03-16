// math 是对标准包 math 的一些补充, 以便使用.
package math

// 高效判断的方法，右移一位，再左移一位，如果与原来相等，则是偶数，否则是奇数.
// 如何判断上面的方法是否高效?

//  IsOddNumber 函数用于判断某个数是否是奇数.
func IsOddNumber(number int) bool {
	if number%2 == 1 {
		return true
	}
	return false
}

// IsEvenNumber 函数用于判断某个数是否是偶数.
func IsEvenNumber(number int) bool {
	if number%2 == 0 {
		return true
	}
	return false
}
