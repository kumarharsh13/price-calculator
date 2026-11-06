package conversion

import (
	"errors"
	"strconv"
)

func StringToFloat(strings []string) ([]float64, error) {
	floats := []float64{} // local — each goroutine gets its own
	for _, stringVal := range strings {
		floatVal, err := strconv.ParseFloat(stringVal, 64)
		if err != nil {
			return nil, errors.New("failed to convert string to float")
		}
		floats = append(floats, floatVal)
	}
	return floats, nil
}
