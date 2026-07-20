package syncx

import "fmt"

type AlternatePrinter struct {
	total int

	ch chan int
}

func NewAlternatePrinter(total int) *AlternatePrinter {
	return &AlternatePrinter{
		total: total,
		ch:    make(chan int),
	}
}

func (p *AlternatePrinter) produce() {
	for i := 0; i <= p.total; i++ {
		p.ch <- i
	}

	close(p.ch)
}

func (p *AlternatePrinter) Run() {
	fmt.Printf("total is %d\n", p.total)

	go p.produce()

	for num := range p.ch {
		fmt.Printf("%d\n", num)
	}

	fmt.Printf("all done.\n")
}
