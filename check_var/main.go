package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var value = 23
	var value = "Bimasakti Multi Sinergi"
	//var value = true

	var reflectValue = reflect.ValueOf(value)

	fmt.Println("tipe  variabel :", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel :", reflectValue)
	} else if reflectValue.Kind() == reflect.String {
		fmt.Println("nilai variabel :", reflectValue)
	} else if reflectValue.Kind() == reflect.Bool {
		fmt.Println("nilai variabel :", reflectValue)
	}
}
