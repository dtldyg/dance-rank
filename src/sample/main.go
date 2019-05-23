package main

import (
	"net/http"
	"fmt"
)

func main() {
	hc := http.Client{}
	resp, err := hc.Get("https://api.bilibili.com/x/web-interface/newlist?rid={rid}&pn={pn}&ps={ps}")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
}
