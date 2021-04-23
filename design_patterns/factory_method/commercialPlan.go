package main

type CommercialPlan struct {
	plan
}

func (cp *CommercialPlan) getRate() float32 {
	return cp.rate
}

func newCommercialPlan() iPlan {
	return &CommercialPlan{
		plan: plan{
			rate: 1.68,
		},
	}
}
