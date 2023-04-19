package main

import (
	"fmt"

	"github.com/thedevsaddam/gojsonq/v2"
)

func main() {
	user1 := `{
		"name": "tian",
		"married": true,
		"address": {
		  "city": "beijing",
		  "country": "China"
		}
	  }`

	user1json := gojsonq.New().FromString(user1)
	name1 := user1json.Find("name").(string)

	user1json.Reset()
	city1 := user1json.Find("address.city")
	fmt.Printf("The name is %s and the city is %v", name1, city1)
}
