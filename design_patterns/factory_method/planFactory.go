package main

import "fmt"

func getPlan(planType string) (iPlan, error) {
	if planType == "domestic" {
		return newDomesticPlan(), nil
	}
	if planType == "commercial" {
		return newCommercialPlan(), nil
	}

	return nil, fmt.Errorf("Invalid plan selected")
}
