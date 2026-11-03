package conversion

import (
	"errors"
	"strconv"
)

var floats []float64

func StringToFloat(strings []string) ([]float64, error) {
	for _, stringVal := range strings {
		floatVal, err := strconv.ParseFloat(stringVal, 64)
		if err != nil {
			return nil, errors.New("Failed to convert string to float..!!")
		}

		floats = append(floats, floatVal)
	}
	return floats, nil
}
