package main

import "fmt"

type plan struct {
	rate float32
}

func (p *plan) getRate() float32 {
	return -1.0
}

func (p *plan) calculateBill(months int) {
	fmt.Printf("%.2f\n", p.rate*float32(months))
}
