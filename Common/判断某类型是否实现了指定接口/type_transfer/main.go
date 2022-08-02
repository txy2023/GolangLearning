package main

import "fmt"

type dog struct {
}
type animal interface {
	speak()
}

func (a *dog) speak() {
	fmt.Println("wang wang wang!")
}

func test(a animal) {
	if _, ok := a.(*dog); ok {
		fmt.Println("success")
	}
}

func main() {
	var a animal
	// a = &dog{}
	// fmt.Println(a)
	if _, ok := a.(*dog); !ok {
		fmt.Println(ok)
	}
	fmt.Println(a)
	test(&dog{})
}

// func main() {
// var _ animal = &dog{}
// var _ animal = (*dog)(nil)
// }
