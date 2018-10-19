package helper

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type HttpClient struct {
}

var DefaultClient *HttpClient

var client *http.Client

func (hc *HttpClient) Request(url string, method string) string {

	req, _ := http.NewRequest(method, url, nil)
	req.Header.Add("user-agent", `Mozilla/5.0 (Macintosh; SmsBomb AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.67 Safari/537.36`)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "error"
	}
	defer func() {
		resp.Body.Close()
	}()
	return "1"

}

func (*HttpClient) bodyFormat(body io.ReadCloser) string {

	// 拿到数据
	bytes, err := ioutil.ReadAll(body)

	if err != nil {
		panic(err)
	}
	// 这里要格式化再输出，因为 ReadAll 返回的是字节切片
	return fmt.Sprintf("%s\n", bytes)
}

func init() {

	proxyUrl, err := url.Parse("tcp://183.163.142.218:4265")
	if err != nil {
		panic(err)
	}
	client = &http.Client{

		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
		Timeout: time.Duration(60 * time.Second),
	}
}
