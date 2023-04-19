package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	user1 := `{
		"name": "tian",
		"married": false,
		"address": {
		  "city": "beijing",
		  "country": "China"
		}
	  }`
	// 指定配置的类型为json
	viper.SetConfigType("json")
	// 读取数据
	if err := viper.ReadConfig(strings.NewReader(user1)); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("数据的所有键值: %v\n", viper.AllKeys())
	fmt.Printf("解析后的数据：%v\n", viper.AllSettings())
	fmt.Printf("the type of \"married\" is %s\n", reflect.TypeOf(viper.Get("married")))
	fmt.Printf("The name is %s and the country is %s\n", viper.Get("name"), viper.Get("address.country"))
}
