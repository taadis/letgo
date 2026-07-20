package syncx

import "fmt"

type AlternatePrinter struct {
	total int
}

func NewAlternatePrinter(total int) *AlternatePrinter {
	return &AlternatePrinter{}
}

func (p *AlternatePrinter) Run() {
	fmt.Printf("total is %d", p.total)
}
