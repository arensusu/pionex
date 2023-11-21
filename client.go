package pionex

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Client struct {
	ApiKey    string
	ApiSecret string
	Endpoint  string
}

func NewClient(key, secret string) *Client {
	return &Client{
		ApiKey:    key,
		ApiSecret: secret,
		Endpoint:  "https://api.pionex.com",
	}
}

func (c *Client) Sign(method, api, timestamp string, param url.Values, body string) string {
	param.Add("timestamp", timestamp)
	paramsString := "?" + param.Encode()

	hash := hmac.New(sha256.New, []byte(c.ApiSecret))
	hash.Write([]byte(method + api + paramsString + body))
	sign := hash.Sum(nil)
	return hex.EncodeToString(sign)
}

func (c *Client) HttpGet(api string, param url.Values) ([]byte, error) {
	u, err := url.Parse(c.Endpoint)
	if err != nil {
		return nil, err
	}

	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	signature := c.Sign(http.MethodGet, api, timestamp, param, "")

	u.Path = api
	u.RawQuery = param.Encode()

	request, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("PIONEX-KEY", c.ApiKey)
	request.Header.Add("PIONEX-SIGNATURE", signature)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("%d: %s", response.StatusCode, responseBody)
	}

	return responseBody, nil
}

func (c *Client) HttpPost(api string, body any) ([]byte, error) {
	method := "POST"
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)

	bodyJson, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	signature := c.Sign(method, api, timestamp, url.Values{}, string(bodyJson))

	uri := c.Endpoint + api + "?timestamp=" + timestamp

	reader := bytes.NewReader(bodyJson)
	request, err := http.NewRequest(method, uri, reader)
	if err != nil {
		return nil, err
	}

	request.Header.Add("PIONEX-KEY", c.ApiKey)
	request.Header.Add("PIONEX-SIGNATURE", signature)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("%d: %s", response.StatusCode, responseBody)
	}

	return responseBody, nil
}

func (c *Client) HttpDelete(api string, body any) ([]byte, error) {
	method := "DELETE"
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)

	bodyJson, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	signature := c.Sign(method, api, timestamp, url.Values{}, string(bodyJson))

	uri := c.Endpoint + api + "?timestamp=" + timestamp

	reader := bytes.NewReader(bodyJson)
	request, err := http.NewRequest(method, uri, reader)
	if err != nil {
		return nil, err
	}

	request.Header.Add("PIONEX-KEY", c.ApiKey)
	request.Header.Add("PIONEX-SIGNATURE", signature)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("%d: %s", response.StatusCode, responseBody)
	}

	return responseBody, nil
}
