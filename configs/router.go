package configs

import (
	"database-migration/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter(controller controllers.CategoryController, film controllers.FilmController) *gin.Engine {
	controller.SetupEngine().POST("/api/categories", controller.Create)
	controller.SetupEngine().GET("/api/categories/:id", controller.FindById)
	controller.SetupEngine().GET("/api/categories", controller.FindAll)

	controller.SetupEngine().POST("/api/film", film.FindById)
	controller.SetupEngine().POST("/api/test", film.GetDataCategoryFIlm)

	return controller.SetupEngine()
}
