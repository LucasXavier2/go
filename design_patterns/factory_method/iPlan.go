package main

type iPlan interface {
	getRate() float32
	calculateBill(units int)
}
