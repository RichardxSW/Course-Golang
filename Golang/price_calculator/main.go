package main

// import "fmt"
import (
	"fmt"

	// "example.com/price-calculator/cmdmanager"
	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100)) 
		// cmd := cmdmanager.New()
		priceJob := prices.New(fm, taxRate)
		go priceJob.Calculate(doneChans[index], errorChans[index]) 
		// if err != nil {
		// 	fmt.Println("Couldn't calculate the job")
		// 	fmt.Println(err)
		// }
	}

	for i := range taxRates{
		select {
		case err := <-errorChans[i]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[i]:
			fmt.Println("Job completed")   
		}
	} 
}
