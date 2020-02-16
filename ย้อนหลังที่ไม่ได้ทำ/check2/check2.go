package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	type MyFloat float64
	type Vertex struct {
		X, Y float64
	}


	type EmptyInterface interface{}

	type Abser interface {
	Abs() float64}
	
}



func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) Abs() float64 {
	return math.Abs(float64(f))

}

	var ia, ib Abser
	ia = Vertex{1, 2}
	ib = MyFloat(1)
	fmt.Println(reflect.TypeOf(ia))
	fmt.Println(reflect.TypeOf(ia).Kind())
	fmt.Println(reflect.TypeOf(ib))
	fmt.Println(reflect.TypeOf(ib).Kind())


	if reflect.TypeOf(ia) != reflect.TypeOf(ib) {
		fmt.Println("Not equal typeOf")
	}
	if reflect.TypeOf(ia).Kind() != reflect.TypeOf(ib).Kind() {
		fmt.Println("Not equal kind")
	}

	ib = Vertex{3, 4}
	if reflect.TypeOf(ia) == reflect.TypeOf(ib) {
		fmt.Println("Equal typeOf")
	}
	if reflect.TypeOf(ia).Kind() == reflect.TypeOf(ib).Kind() {
		fmt.Println("Equal kind")
	}

