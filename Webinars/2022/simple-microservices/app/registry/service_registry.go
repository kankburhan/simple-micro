package registry

import (
	rp "simple-micro/app/repository"
	"simple-micro/app/service"

	"github.com/go-kit/kit/log"
	"gorm.io/gorm"
)

func NewTodoService(db *gorm.DB, logger log.Logger) service.TodoService {
	return service.NewtodoServiceImpl(
		rp.NewBaseRepository(db),
		rp.NewTodoRepository(rp.NewBaseRepository(db)),
	)
}
