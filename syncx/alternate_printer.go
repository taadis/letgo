package syncx

import (
	"fmt"
	"time"
)

type AlternatePrinter struct {
	total int

	ch  chan int
	ch1 chan struct{}
	ch2 chan struct{}
}

func NewAlternatePrinter(total int) *AlternatePrinter {
	return &AlternatePrinter{
		total: total,
		ch:    make(chan int),
		ch1:   make(chan struct{}),
		ch2:   make(chan struct{}),
	}
}

func (p *AlternatePrinter) produce() {
	for i := 0; i <= p.total; i++ {
		p.ch <- i
	}

	close(p.ch)
}

func (p *AlternatePrinter) worker1() {
	for {
		<-p.ch1

		num, ok := <-p.ch
		if !ok {
			close(p.ch2)
			return
		}

		fmt.Printf("%d\n", num)

		p.ch2 <- struct{}{}
	}
}

func (p *AlternatePrinter) worker2() {
	for {
		<-p.ch2

		num, ok := <-p.ch
		if !ok {
			close(p.ch1)
			return
		}

		fmt.Printf("%d\n", num)

		p.ch1 <- struct{}{}
	}
}

func (p *AlternatePrinter) Run() {
	fmt.Printf("total is %d\n", p.total)

	go p.produce()
	go p.worker1()
	go p.worker2()

	p.ch1 <- struct{}{}

	time.Sleep(1 * time.Second)

	fmt.Printf("all done.\n")
}
