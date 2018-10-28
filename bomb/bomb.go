package bomb

import (
	"SmsBomb/helper"
	"application"
	"fmt"
)

type Sender struct {
	httpClient  *helper.HttpClient
	app         *application.App
	concurrency int
}

func (s *Sender) send() {
	for i := 0; i < s.concurrency; i++ {
		s.app.Wg.Add(1)
		go s.sendToItisw()
	}
	s.app.Wg.Wait()
}

func (s *Sender) sendToItisw() {
	defer s.app.Wg.Add(-1)
	for {
		res := s.httpClient.Request("https://itisw.com", "GET")
		fmt.Printf(res)
	}
}

func NewSenderModule(app *application.App) *Sender {
	sd := &Sender{
		app:         app,
		concurrency: 10,
		httpClient:  helper.NewHttpClient(true, "tcp://127.0.0.1:1080"),
	}
	return sd
}

func (s *Sender) Run() {
	s.send()
}

func (s *Sender) resetHttpClient(useProxy bool) {
	s.httpClient = helper.NewHttpClient(true, "tcp://127.0.0.1:1080")
}
