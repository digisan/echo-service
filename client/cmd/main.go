package main

import (
	"fmt"

	. "github.com/digisan/echo-service/client"
	lk "github.com/digisan/logkit"
)

const (
	addr = "https://192.168.31.157:1323"
)

func main() {
	SetFetchCert("./cert/public.pem")
	resp, err := Fetch("GET", addr+"/api/testmsg", nil, nil)
	lk.FailOnErr("%v", err)
	fmt.Println(addr+"/api/testmsg", string(resp))
}
