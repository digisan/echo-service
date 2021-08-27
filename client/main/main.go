package main

import (
	"fmt"

	. "github.com/digisan/echo-service/client"
)

func main() {
	resp, err := Fetch("GET", "http://127.0.0.1:1545/api/test", nil, nil)
	if err == nil {
		fmt.Println(string(resp))
	}
}
