package main

import (
	"fmt"

	"github.com/prashantbhat84/price-calculator/prices"
)

func main() {
	fmt.Println("Price Calculator Starts")

	var taxRates []float64 = []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}

}
