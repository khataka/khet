package main

import (
    "fmt"
    "reflect"
)

type bill struct {
    N1 int
    N2 string
    n3 string
}

func main() {
    a := bill{4, "dhfthf", "fdgdf"}

    v := reflect.ValueOf(a)

    values := make(map[string]interface{}, v.NumField())

    for i := 0; i < v.NumField(); i++ {
        if v.Field(i).CanInterface() {
            values[v.Type().Field(i).Name] = v.Field(i).Interface()
        } else {
            fmt.Printf("sorry you have a unexported field (lower case) value you are trying to sneak past. I will not allow it: %v\n", v.Type().Field(i).Name)
        }
    }

    fmt.Println(values)

    passObject(&values)
}

func passObject(v1 *map[string]interface{}) {
    fmt.Println("yoyo")
}