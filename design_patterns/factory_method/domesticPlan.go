package main

type DomesticPlan struct {
	plan
}

func (dp *DomesticPlan) getRate() float32 {
	return dp.rate
}

func newDomesticPlan() iPlan {
	return &DomesticPlan{
		plan: plan{
			rate: 1.32,
		},
	}
}
