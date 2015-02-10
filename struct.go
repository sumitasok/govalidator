package validator

import (
	// "bytes"
	// "fmt"
	"reflect"
	"strings"
)

type Element interface {
}

func Parse(element Element) bool {
	typeOf := reflect.TypeOf(element)
	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)

		// val := reflect.ValueOf(element).Field(i)
		validatorKeys := field.Tag.Get("validator")
		strs := strings.Split(validatorKeys, ",")
		for _, str := range strs {
			switch str {
			case "required":
				// fmt.Println("got a required")
				// if bytes == nil {
				// 	return false
				// }
			}
		}
		// on the field value
		// for switch case type of validation

	}
	return true
}
