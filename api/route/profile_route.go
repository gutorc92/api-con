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

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db database.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}
