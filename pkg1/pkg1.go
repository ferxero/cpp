package pkg1

import (
	. "fmt"
)

// SayHello name, lname are values by method.
func SayHello(name, lname string) string {
	return Sprintf("Hi there, %v %v !", name, lname)
}