package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// user结构体，用于保存request Body中json解码后的内容
type user struct {
	Name string
}

func registe(w http.ResponseWriter, r *http.Request) {
	user1 := &user{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, user1)
	fmt.Fprintf(w, "The name you submit is %s", user1.Name)
}

func main() {
	http.HandleFunc("/register", registe)
	http.ListenAndServe(":8080", nil)
}
