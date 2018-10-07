package application

import (
	"application/components"
	"application/config"
	"sync"
)

const (
	VERSION string = "1.0.0"
)

type App struct {
	appName    string
	Logger     *components.Logger
	config     map[string]interface{}
	env        string
	wg         *sync.WaitGroup
	components map[string]interface{}
}

func NewApp(name string, configPath string) *App {

	app := &App{
		appName: name,
		Logger:  components.GetLogger(),
		config:  config.GetConfig(configPath),
	}
	value, ok := app.config["env"]
	if !ok {
		app.env = "dev"
	} else {
		app.env = value.(string)
	}

	return app
}

func (a *App) GetEnv() string {
	return a.env
}
func (a *App) GetAppName() string {
	return a.appName
}
func (a *App) GetConfig() map[string]interface{} {
	return a.config
}
