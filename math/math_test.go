package math

import (
	"testing"
)

// TestIsOddNumber 测试奇数判断是否正确
func TestIsOddNumber(t *testing.T) {
	// 先来个正常的
	if IsOddNumber(99) != true {
		t.Log("99 是奇数.")
		t.Fail()
	}

	// 再来个边界值
	if IsOddNumber(1) != true {
		t.Log("1 是奇数.")
		t.Fail()
	}

	// 来个异常的
	if IsOddNumber(0) != false {
		t.Log("0 不是奇数.")
		t.Fail()
	}
}

// TestIsEvenNumber 测试偶数判断是否正确
func TestIsEvenNumber(t *testing.T) {
	// 正常
	if IsEvenNumber(100) != true {
		t.Log("100 是偶数.")
		t.Fail()
	}

	// 正常负数
	if IsEvenNumber(-2) != true {
		t.Log("-2 是偶数.")
		t.Fail()
	}

	// 边界值
	if IsEvenNumber(0) != true {
		t.Log("0 是偶数.")
		t.Fail()
	}

	// 异常
	if IsEvenNumber(1) != false {
		t.Log("1 不是偶数.")
		t.Fail()
	}

	// 异常负数
	if IsEvenNumber(-1) {
		t.Log("-1 不是偶数.")
		t.Fail()
	}
}
