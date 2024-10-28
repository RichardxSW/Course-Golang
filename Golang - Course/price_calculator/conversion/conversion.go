package conversion

import (
	"fmt"
	// "os"
	"strconv"
	"errors"
)

func StringsToFloats(strings []string) ([]float64, error) {
	var floats []float64
	for _, stringVal := range strings {
		floatVal, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
			return nil, errors.New(fmt.Sprintf("Failed to convert string '%s' to float", stringVal))
		}

		floats = append(floats, floatVal)
	}
	return floats, nil
}