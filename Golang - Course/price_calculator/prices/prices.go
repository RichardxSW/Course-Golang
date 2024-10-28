package prices

import (
	// "bufio"
	"fmt"
	// "os"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager        iomanager.IOManager `json:"-"`
	TaxRate          float64             `json:"tax_rate"`
	InputPrices      []float64           `json:"input_prices"`
	TaxIncludedPrice map[string]string   `json:"tax_included_price"`
}

func (job *TaxIncludedPriceJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		return err
	}
	job.InputPrices = prices
	return nil
}

func (job TaxIncludedPriceJob) Calculate(doneChan chan bool, errorChan chan error) {
	err := job.LoadData()

	if err != nil {
		// return err
		errorChan <- err
        return
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	job.TaxIncludedPrice = result
	// fmt.Println(result)
	job.IOManager.WriteResult(job)
	doneChan <- true // signal completion to main goroutine
}

func New(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}
