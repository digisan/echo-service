package main

import (
	"fmt"

	. "github.com/digisan/echo-service/client"
	lk "github.com/digisan/logkit"
)

func main() {
	resp, err := Fetch("GET", "https://127.0.0.1:1323/api/testmsg", nil, nil)
	lk.FailOnErr("%v", err)
	fmt.Println("https://127.0.0.1:1323/api/testmsg", string(resp))
}
