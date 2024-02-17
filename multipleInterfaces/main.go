package main

import (
	"fmt"
)

func (e email) cost() int {
	// ?
	if e.isSubscribed == true {
		return len(e.body) * 2
	} else {
		return len(e.body) * 5
	}
}

func (e email) format() string {
	// ?
	if e.isSubscribed == true {
		msg := fmt.Sprintf("'%s' | Subscribed", e.body)
		return msg
	} else {
		msg := fmt.Sprintf("'%s' | Not Subscribed", e.body)
		return msg
	}
}

type expense interface {
	cost() float64
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
