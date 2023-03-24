package controllers

import (
	"database-migration/models/web"
	"database-migration/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type FilmControllerImpl struct {
	engine      *gin.Engine
	FilmService service.FilmService
	Service     service.Service
}

func NewFilmControllerImpl(engine *gin.Engine, filmService service.FilmService, service service.Service) FilmController {
	return &FilmControllerImpl{engine: engine, FilmService: filmService, Service: service}
}

func (controller *FilmControllerImpl) GetDataCategoryFIlm(c *gin.Context) {
	request := web.RequestCategoryFilm{}

	err := c.Bind(&request)
	if err != nil {
		panic(err)
	}
	fmt.Println(request)

	data := controller.Service.GetCategoryAndFilmById(request.CategoryId, request.FilmId)

	fmt.Println(data)

	c.JSON(200, data)
}

func (controller *FilmControllerImpl) FindById(c *gin.Context) {
	request := web.FilmRequest{}
	err := c.Bind(&request)
	if err != nil {
		panic(err)
	}

	response := controller.FilmService.GetDataById(request.Id)

	c.JSON(200, response)
}
