package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gutorc92/go-backend-clean-architecture/api/middleware"
	"github.com/gutorc92/go-backend-clean-architecture/bootstrap"
	"github.com/gutorc92/go-backend-clean-architecture/database"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db database.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewProfileRouter(env, timeout, db, protectedRouter)
	NewTaskRouter(env, timeout, db, protectedRouter)
}
