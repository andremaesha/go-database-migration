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

	filmRepository := repository.NewFilmRepositoryImpl()
	filmService := service.NewFilmServiceImpl(db, filmRepository)

	services := service.NewServiceImpl(db, filmRepository, categoryRepository)

	filmController := controllers.NewFilmControllerImpl(engine, filmService, services)

	router := configs.NewRouter(categoryController, filmController)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
