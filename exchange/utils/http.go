package utils

import (
	"sync"
	"net/url"
	"net/http"
	"strings"
	"encoding/json"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
)

const GetMethod = "GET"
const PostMethod = "POST"
const DeleteMethod = "DELETE"
const PutMethod = "PUT"
const SendTypeForm = "form"
const SendTypeJson = "json"

type HttpSend struct {
	Url      string
	SendType string
	Header   map[string]string
	Body     map[string]string
	sync.RWMutex
}

func NewHttpSend(url string) *HttpSend {
	return &HttpSend{
		Url:      url,
		SendType: SendTypeJson,
	}
}

func (h *HttpSend) SetBody(body map[string]string) {
	h.Lock()
	defer h.Unlock()
	h.Body = body
}

func (h *HttpSend) SetHeader(header map[string]string) {
	h.Lock()
	defer h.Unlock()
	h.Header = header
}

func (h *HttpSend) SetSendType(sendType string) {
	h.Lock()
	defer h.Unlock()
	h.SendType = sendType
}

func GetUrlBuild(Url string, data map[string]string) string {
	u, _ := url.Parse(Url)
	q := u.Query()
	for k, v := range data {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func (h *HttpSend) send(method string) ([]byte, error) {
	var (
		req      *http.Request
		resp     *http.Response
		client   http.Client
		sendData string
		err      error
	)
	if len(h.Body) > 0 {
		if strings.ToLower(h.SendType) == SendTypeJson {
			sendBody, jsonErr := json.Marshal(h.Body)
			if jsonErr != nil {
				return nil, jsonErr
			}
			sendData = string(sendBody)
		} else {
			sendBody := http.Request{}
			sendBody.ParseForm()
			for k, v := range h.Body {
				sendBody.Form.Add(k, v)
			}
			sendData = sendBody.Form.Encode()
		}
	}

	//忽略https证书
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	req, err = http.NewRequest(method, h.Url, strings.NewReader(sendData))

	defer req.Body.Close()

	//设置默认header
	if len(h.Header) == 0 {

		if strings.ToLower(h.SendType) == SendTypeJson {
			h.Header = map[string]string{
				"Content-Type": "application/json;charset=utf-8",
			}
		} else {
			h.Header = map[string]string{
				"Content-Type": "application/x-www-form-urlencoded",
			}
		}

	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36")

	// 添加请求头
	for k, v := range h.Header {
		if strings.ToLower(k) == "host" {
			req.Host = v
		} else {
			req.Header.Add(k, v)
		}
	}

	// 发送请求
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("error http code :%d", resp.StatusCode))
	}

	return ioutil.ReadAll(resp.Body)
}
func (h *HttpSend) Get() ([]byte, error) {
	return h.send(GetMethod)
}

func (h *HttpSend) Post() ([]byte, error) {
	return h.send(PostMethod)
}

func (h *HttpSend) Delete() ([]byte, error) {
	return h.send(DeleteMethod)
}

func (h *HttpSend) Put() ([]byte, error) {
	return h.send(PutMethod)
}
