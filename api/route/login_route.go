package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gutorc92/go-backend-clean-architecture/api/controller"
	"github.com/gutorc92/go-backend-clean-architecture/bootstrap"
	"github.com/gutorc92/go-backend-clean-architecture/database"
	"github.com/gutorc92/go-backend-clean-architecture/domain"
	"github.com/gutorc92/go-backend-clean-architecture/repository"
	"github.com/gutorc92/go-backend-clean-architecture/usecase"
)

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db database.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}
