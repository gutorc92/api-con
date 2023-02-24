package bootstrap

import (
	"context"
	"fmt"

	"github.com/gutorc92/go-backend-clean-architecture/database"
	"github.com/gutorc92/go-backend-clean-architecture/mongo"
)

const (
	MONGO = "mongo"
)

type Application struct {
	Env      *Env
	Database database.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	fmt.Printf("db kind %s\n", app.Env.DBKind)
	if app.Env.DBKind == MONGO {
		app.Database = mongo.NewClient(context.TODO(), app.Env.DBHost, app.Env.DBPort, app.Env.DBUser, app.Env.DBPass)
	}
	return *app
}

func (app *Application) CloseDBConnection() {
	app.Database.Disconnect(context.TODO())
}
