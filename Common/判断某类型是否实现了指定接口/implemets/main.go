package main

import (
	"fmt"
	"reflect"
)

type dog struct {
}
type animal interface {
	speak()
}

func (a *dog) speak() {
	fmt.Println("wang wang wang!")
}

func main() {
	typ_dog := reflect.TypeOf(&dog{})
	typ_animal := reflect.TypeOf((*animal)(nil)).Elem()
	fmt.Println(typ_dog.Implements(typ_animal))
}
