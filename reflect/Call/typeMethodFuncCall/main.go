package main

import (
	"fmt"
	"reflect"
)

type dog struct {
	name string
}

func (a *dog) Speak() {
	fmt.Println("speaking!")
}
func (a *dog) Walk() {
	fmt.Println("walking!")
}

func (a *dog) GetName() (string, error) {
	fmt.Println("Getting name!")
	return a.name, nil
}
func main() {
	var tudou = &dog{name: "tudou"}
	// 获取reflect.Type
	objectType := reflect.TypeOf(tudou)
	// 批量执行方法
	for i := 0; i < objectType.NumMethod(); i++ {
		fmt.Printf("Now method: %v is being executed...\n", objectType.Method(i).Name)
		// Call的第一个入参对应receiver，即方法的接受者本身
		ret := objectType.Method(i).Func.Call([]reflect.Value{reflect.ValueOf(tudou)})
		fmt.Printf("The return of method: %s is %v\n\n", objectType.Method(i).Name, ret)
	}
}
