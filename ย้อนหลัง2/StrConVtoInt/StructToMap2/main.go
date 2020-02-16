package main

import (
	"fmt"
	"reflect"
)

func ToMap(in interface{}, tag string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts structs; got %T", v)
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {

		fi := typ.Field(i)
		if tagv := fi.Tag.Get(tag); tagv != "" {
			out[tagv] = v.Field(i).Interface()
		}
	}
	return out, nil
}

type Foo struct {
	A int    `m:"a"`
	B string `m:"b"`
	C string
}

func main() {
	f := Foo{A: 1, B: "hello", C: "world"}
	fmt.Printf("%[1]T, %+[1]v\n", f)

	g, err := ToMap(f, "m")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%[1]T, %+[1]v\n", g)
}
