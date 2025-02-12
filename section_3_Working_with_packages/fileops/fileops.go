package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Reading file
func GetFloatFromFile(fileName string, defaultValue float64) (float64, error) {
	// if we don't use some argument we use _

	data, err := os.ReadFile(fileName)
	if err != nil {
		return defaultValue, errors.New("failed to find file")
	}
	valueConvertToText := string(data)
	value, err := strconv.ParseFloat(valueConvertToText, 64)
	if err != nil {
		return defaultValue, errors.New("failed to parse stored value")
	}
	return value, nil
}

// writing data into txt
func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}
