package main

import (
	"encoding/json"
	"os"
	"reflect"
)

type S struct {
	Name string
}

type D struct {
	Pants bool
}

func main() {
	a := Combine(&S{"Bob"}, &D{true})
	json.NewEncoder(os.Stderr).Encode(a)
}

func Combine(v ...interface{}) interface{} {
	f := make([]reflect.StructField, len(v))
	for i, u := range v {
		f[i].Type = reflect.TypeOf(u)
		f[i].Anonymous = true
	}

	r := reflect.New(reflect.StructOf(f)).Elem()
	for i, u := range v {
		r.Field(i).Set(reflect.ValueOf(u))
	}
	return r.Addr().Interface()
}
