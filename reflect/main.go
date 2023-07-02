package main

import (
	"fmt"
	"reflect"
)

type dog struct {
}

func (a *dog) Speak() error {
	fmt.Println("wang wang wang!")
	return nil
}
func (a *dog) Walk() error {
	fmt.Println("walk walk walk!")
	return fmt.Errorf("can not walk")
}

func main() {
	var tudou = &dog{}
	objectType := reflect.TypeOf(tudou)
	for i := 0; i < objectType.NumMethod(); i++ {
		ret := objectType.Method(i).Func.Call([]reflect.Value{reflect.ValueOf(tudou)})
		// fmt.Println(ret)
		for _, v := range ret {
			fmt.Println(v)
		}
	}
}
