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

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db database.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
}
