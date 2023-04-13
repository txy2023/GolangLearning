package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	body := strings.NewReader(`{"name":"tian"}`)
	// body := strings.NewReader("{\"name\":\"tian\"}")
	req, _ := http.NewRequest("POST", "http://127.0.0.1:8080/register", body)
	res, _ := http.DefaultClient.Do(req)
	ret, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(ret))
}
