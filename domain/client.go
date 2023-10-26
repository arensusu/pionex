package domain

import "net/url"

type Client interface {
	Sign(method string, api string, timestamp string, param url.Values, body string) string
	HttpGet(api string, param url.Values) ([]byte, error)
	HttpPost(api string, body any) ([]byte, error)
}

type HttpResponse struct {
	Result    bool   `json:"result"`
	Timestamp int64  `json:"timestamp"`
	Code      string `json:"code"`
	Message   string `json:"message"`
}
