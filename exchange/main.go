package main

import (
	"./utils"
	"fmt"
	"net/http"
)

func main() {
	rest := utils.Rest{}
	rest.Method = "GET"
	rest.Url = "http://www.baidu.com"
	rest.Payload = map[string]string{"a": "111"}
	rest.Header = http.Header{"Content-Type":"application/json"}
	result := utils.HttpRestApi(rest)
	fmt.Println(result)
}
