package application

type Imoudule interface {
	Register(app *App)

	//UnRegister only run before process exit
	UnRegister(app *App)
	Name() string
}
