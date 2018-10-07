package helper

import (
	"net/http"
)

type HttpClient struct {
	client *http.Client
}

func setHeader(rq *http.Request, header map[string]interface{}) *http.Request {

	return rq
}
func setBody(rq *http.Request, body map[string]interface{}) *http.Request {

	return rq
}

func (httpClient *HttpClient) request(url string, method string, header map[string]interface{}, body map[string]interface{}) *http.Response {

	req, err := http.NewRequest(method, url, nil)
	if err == nil {
		panic("HttpClient: request exception")
	}
	if len(header) > 0 {
		req = setHeader(req, header)
	}
	if len(body) > 0 {
		req = setBody(req, header)
	}

	resp, err := httpClient.client.Do(req)
	return resp
}
