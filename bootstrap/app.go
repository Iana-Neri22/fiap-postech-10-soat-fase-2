package bootstrap

import "lanchonete/mongo"

type Application struct {
	Env   *Env
	Mongo mongo.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = NewDatabase(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseConnection(app.Mongo)
}