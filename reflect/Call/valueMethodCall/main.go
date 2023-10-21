package main

import (
	"fmt"
	"reflect"
)

type dog struct {
	name string
}

func (a *dog) GetName() (string, error) {
	fmt.Println("Getting name!")
	return a.name, nil
}
func (a *dog) Speak() {
	fmt.Println("speaking!")
}

// func (a *dog) SetName(name string) error {
// 	fmt.Println("Setting name!")
// 	a.name = name
// 	return nil
// }

func main() {
	var tudou = &dog{name: "tudou"}
	// 获取reflect.Value
	objectValue := reflect.ValueOf(tudou)
	// 根据方法名获取method，执行Call
	objectValue.MethodByName("Speak").Call(nil)
	// objectValue.MethodByName("SetName").Call([]reflect.Value{reflect.ValueOf("newName")})
	objectValue.MethodByName("GetName").Call(nil)

	// 批量执行
	objectType := reflect.TypeOf(tudou)
	for i := 0; i < objectValue.NumMethod(); i++ {
		fmt.Printf("the method name is %s\n", objectType.Method(i).Name)
		objectValue.Method(i).Call(nil)
	}
}
