package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) outputData() {
	fmt.Println(m)
}

func main() {
	//kalau pakai make bisa diisi data scr lgsg tanpa append, kalau pakai make slice, slot yang kosong gabisa di isi, hrus pake append
	userNames := make([]string, 3)
	// userNames := []string{}

	userNames[0] = "John"
	userNames[2] = "Jane"
	userNames = append(userNames, "Mark")

	fmt.Println(userNames)

	// courseRatings := make(map[string]float64, 3)
	courseRatings := make(floatMap, 3)
	courseRatings["Python"] = 4.5
	courseRatings["JavaScript"] = 4.8
	courseRatings["Golang"] = 4.2

	courseRatings.outputData()

	for index, value := range userNames {
		fmt.Println("Index: ", index)
		fmt.Println("Value: ", value)
	}

	for key, value := range courseRatings {
		fmt.Println("Key: ", key)
		fmt.Println("Value: ", value)
	}
}