package main

import (
	"time"

	"github.com/gin-gonic/gin"
	route "github.com/gutorc92/go-backend-clean-architecture/api/route"
	"github.com/gutorc92/go-backend-clean-architecture/bootstrap"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	db := app.Database.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, db, gin)

	gin.Run(env.ServerAddress)
}
