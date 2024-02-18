// Package main is the entry point of this Go program.
package main

import (
// Importing the "fmt" package for formatted I/O.
   "fmt"
)

func (e email) cost() int {
   // If the recipient is subscribed, the cost is the length of the email body multiplied by 2.
   if e.isSubscribed == true {
   	return len(e.body) * 2
   // Otherwise, the cost is the length of the email body multiplied by 5.
   } else {
   	return len(e.body) * 5
   }
}
// The method (e email) format() string creates a formatted string representing the email, including its body and the subscription status of the recipient.
func (e email) format() string {
   // If the recipient is subscribed, the formatted string is "'email body' | Subscribed".
   if e.isSubscribed == true {
   	msg := fmt.Sprintf("'%s' | Subscribed", e.body)
   	return msg
   // Otherwise, the formatted string is "'email body' | Not Subscribed".
   } else {
   	msg := fmt.Sprintf("'%s' | Not Subscribed", e.body)
   	return msg
   }
}

// The type expense is an interface with a single method, cost(), which returns a float64 value.
type expense interface {
   cost() float64
}

// The type formatter is an interface with a single method, format(), which returns a string value.
type formatter interface {
   format() string
}

// The type email is a struct containing two fields: isSubscribed (a boolean indicating the subscription status of the recipient) and body (a string representing the email body).
type email struct {
   isSubscribed bool
   body         string
}