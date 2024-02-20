package main

import (
	"fmt"
	"reflect"
)

func PopulateStruct(data map[string]interface{}, persn interface{}) {
	// Get the reflection value of the persn interface and navigate to its underlying struct
	resultValue := reflect.ValueOf(persn).Elem()

	for key, value := range data {

		personField := resultValue.FieldByName(key)

		if personField.IsValid() {
			// if the personfield is valid and is a struct itself , recursively call the populate struct
			if personField.Kind() == reflect.Struct {

				if nestedMap, ok := value.(map[string]interface{}); ok {
					// Creating a new instance of the nested struct type
					nestedStruct := reflect.New(personField.Type()).Interface()

					PopulateStruct(nestedMap, nestedStruct)

					personField.Set(reflect.ValueOf(nestedStruct).Elem())
				}
			} else {
				// setting the value directly, if it is a non struct field
				personField.Set(reflect.ValueOf(value))
			}
		}
	}

}

type Person struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	City  string
	State string
}

func main() {
	data := map[string]interface{}{
		"Name":        "Chetan Thakral",
		"Age":         21,
		"pincode":     110018,
		"RandomField": "random",
		"Address": map[string]interface{}{
			"City":  "New Delhi",
			"State": "Delhi",
		},
	}

	var personPtr *Person = &Person{}

	PopulateStruct(data, personPtr)

	fmt.Printf("%+v\n", *personPtr)
}
