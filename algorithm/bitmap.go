package algorithm

type Bitmap struct {
	words []uint64
}

func (m *Bitmap) getWordAndBit(num int) (int, uint) {
	word, bit := num/64, uint(num%64)
	return word, bit
}

func (m *Bitmap) Has(num int) bool {
	word, bit := m.getWordAndBit(num)
	return word < len(m.words) && m.words[word]&(1<<bit) != 0
}

func (m *Bitmap) Add(num int) {
	word, bit := m.getWordAndBit(num)
	for word >= len(m.words) {
		m.words = append(m.words, 0)
	}

	if m.words[word]&(1<<bit) == 0 {
		m.words[word] |= 1 << bit
	}
}
