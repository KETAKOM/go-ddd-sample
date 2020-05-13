package main

import "fmt"

type FullName struct {
	firstName, lastName string
}

func (f *FullName) Equals(name FullName) bool {
	if f.firstName == name.firstName &&
		f.lastName == name.lastName {
		return true
	}
	return false
}

func NewName(f, l string) FullName {
	return FullName{firstName: f, lastName: l}
}

func main() {
	yamada := NewName("taro", "yamada")
	suzuki := NewName("taro", "yamada")

	if t := yamada.Equals(suzuki); !t {
		fmt.Println("not equals")
		return
	}
	fmt.Println("equals")
}
