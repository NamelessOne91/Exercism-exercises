package flatten

import (
	"reflect"
)

func Flatten(nested interface{}) []interface{} {
	flat := []interface{}{}

	t := reflect.TypeOf(nested).Kind()
	if t == reflect.Slice || t == reflect.Array {
		for _, el := range nested.([]interface{}) {
			if el != nil {
				flat = append(flat, Flatten(el)...)
			}
		}
	} else if nested != nil {
		flat = append(flat, nested)
	}

	return flat
}
