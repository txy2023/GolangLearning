package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name    string
	Married bool
	Address struct {
		City    string
		Country string
	}
}

func main() {
	user1 := `{
		"name": "tian",
		"married": false,
		"address": {
		  "city": "beijing",
		  "country": "China"
		}
	  }`
	user1Struct := &user{}
	json.Unmarshal([]byte(user1), user1Struct)
	fmt.Printf("解码后的结果为：%v", *user1Struct)
}
