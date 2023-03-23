package configs

import (
	"database-migration/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter(controller controllers.CategoryController) *gin.Engine {
	controller.SetupEngine().POST("/api/categories", controller.Create)
	controller.SetupEngine().GET("/api/categories/:id", controller.FindById)
	controller.SetupEngine().GET("/api/categories", controller.FindAll)

	return controller.SetupEngine()
}
