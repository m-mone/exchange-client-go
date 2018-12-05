package utils

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

type Rest struct {
	Method  string
	Url     string
	Payload map[string]string
	Header  http.Header
}

func HttpRestApi(httpRest Rest) string {
	bytesData, err := json.Marshal(httpRest.Payload)

	payload := bytes.NewReader(bytesData)
	req, err := http.NewRequest(httpRest.Method, httpRest.Url, payload)
	if err != nil {

	}

	req.Header = httpRest.Header
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		defer resp.Body.Close()
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	//str := (*string)(unsafe.Pointer(&respBytes))
	//fmt.Println(*str)
	return string(body)
}
