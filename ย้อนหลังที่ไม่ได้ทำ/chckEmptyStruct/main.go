package main

import (
	"fmt"
	"reflect"
	"time"
)

type Session struct {
	playerId  string
	beehive   string
	timestamp time.Time
}

func (s Session) IsEmpty() bool {
	return reflect.DeepEqual(s, Session{})
}

func main() {
	x := Session{}
	if x.IsEmpty() {
		fmt.Print("is empty")
	}
}
