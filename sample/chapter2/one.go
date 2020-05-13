package main

import (
	"fmt"
)

type FullName struct {
	firstName, lastName string
}

func NewName(f, l string) FullName {
	return FullName{firstName: f, lastName: l}
}

func (f *FullName) GetLastName() string {
	return f.lastName
}

func main() {
	fullName := NewName("shunki", "kotake")
	lastname := fullName.GetLastName()
	fmt.Println(lastname)
}
