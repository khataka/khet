package main

import (
	"fmt"
	"reflect"
)

type AnonymousType reflect.Value

func (a AnonymousType) IsA(typeToAssert interface{}) bool {
	return typeToAssert == reflect.Value(a).Kind()
}
func ToAnonymousType(obj interface{}) AnonymousType {
	return AnonymousType(reflect.ValueOf(obj))
}

func main() {

	var f float64 = 3.4
	var s string = "hello"
	var struc struct{}

	anonFloat := ToAnonymousType(f)
	anonString := ToAnonymousType(s)
	anonStruct := ToAnonymousType(struc)

	fmt.Println("Anon Float64 is a Float64 %t", anonFloat.IsA(reflect.Float64))
	fmt.Println("Anon Float64 is a Float32 %t", anonFloat.IsA(reflect.Float32))
	fmt.Println("Anon String is a String %t", anonString.IsA(reflect.String))
	fmt.Println("Anon String is a Float32 %t", anonString.IsA(reflect.Float32))
	fmt.Println("Anon Struct is a Struct %t", anonStruct.IsA(reflect.Struct))
	fmt.Println("Anon Struct is a Slice %t", anonStruct.IsA(reflect.Slice))

	anon := ToAnonymousType(f)
	if anon.IsA(reflect.String) {
		fmt.Println("Its A String!")
	} else if anon.IsA(reflect.Float32) {
		fmt.Println("Its A Float32!")
	} else if anon.IsA(reflect.Float64) {
		fmt.Println("Its A Float64!")
	} else {
		fmt.Println("Failed")
	}

}
