package main

import (
	"fmt"

	"github.com/kumarharsh13/price-calculator/fileManager"
	"github.com/kumarharsh13/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	results := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		fm := fileManager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		pricesJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		pricesJob.Process()
	}
	fmt.Println(results)
}
