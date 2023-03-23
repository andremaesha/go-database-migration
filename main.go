package main

import (
	"database-migration/configs"
	"database-migration/controllers"
	"database-migration/repository"
	"database-migration/service"
	"github.com/gin-gonic/gin"
)

func main() {
	db := configs.NewDB()
	engine := gin.Default()
	categoryRepository := repository.NewCategoryRepositoryImpl()
	categoryService := service.NewCategoryServiceImpl(db, categoryRepository)
	categoryController := controllers.NewCategoryControllerImpl(engine, categoryService)

	router := configs.NewRouter(categoryController)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
