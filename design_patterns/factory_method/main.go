package main

import "fmt"

func main() {
	myDomesticPlan, _ := getPlan("domestic")
	myCommercialPlan, _ := getPlan("commercial")

	fmt.Printf("Domestic plan rate is: %f\n", myDomesticPlan.getRate())
	fmt.Printf("Commercial plan rate is: %f\n", myCommercialPlan.getRate())

	myDomesticPlan.calculateBill(5)
	myCommercialPlan.calculateBill(5)
}
