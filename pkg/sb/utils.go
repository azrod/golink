package sb

import (
	"fmt"
	"log"
)

// convertToSliceStruct converts a []models.Model to a []T.
func convertToSliceStruct[T any](inters []interface{}) ([]T, error) {
	result := []T{}
	for _, inter := range inters {
		x, err := convertToStruct[T](inter)
		if err != nil {
			return nil, err
		}
		result = append(result, x)
	}
	return result, nil
}

// convertToStruct converts a interface{} to a T.
func convertToStruct[T any](i interface{}) (T, error) {
	log.Default().Printf("convertToStruct: %T", i)
	value, ok := i.(T)
	if !ok {
		return value, fmt.Errorf("failed to convert interface{} to %T", value)
	}

	return value, nil
}
