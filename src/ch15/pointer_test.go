package ch15

import (
	"fmt"
	"testing"
)

type Dog struct {
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func TestName(t *testing.T) {
	dog := new(Dog)
	dog.SetName("dogName")
	fmt.Printf("the dog name is %s\n", dog.name)
}