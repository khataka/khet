package main

import (
	"fmt"
	"time"
)

type Session struct {
	playerId  string
	beehive   string
	timestamp time.Time
}

func (s Session) Equal(o Session) bool {
	if s.playerId != o.playerId {
		return false
	}
	if s.beehive != o.beehive {
		return false
	}
	if s.timestamp != o.timestamp {
		return false
	}
	return true
}

func (s Session) IsEmpty() bool {
	return s.Equal(Session{})
}

func main() {
	x := Session{}
	if x.IsEmpty() {
		fmt.Print("is empty")
	}
}
