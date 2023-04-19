package main

import (
	"fmt"

	"github.com/bitly/go-simplejson"
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

	user1json, err := simplejson.NewJson([]byte(user1))
	if err != nil {
		fmt.Println(err)
	}
	name1, _ := user1json.Get("name").String()
	city1, _ := user1json.Get("address").Get("city").String()
	fmt.Printf("The name is %s and the city is %s", name1, city1)
}
