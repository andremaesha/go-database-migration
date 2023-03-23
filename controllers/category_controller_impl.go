package controllers

import (
	"database-migration/models/web"
	"database-migration/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryControllerImpl struct {
	engine          *gin.Engine
	CategoryService service.CategoryService
}

func NewCategoryControllerImpl(engine *gin.Engine, categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{engine: engine, CategoryService: categoryService}
}

func (controller *CategoryControllerImpl) Create(c *gin.Context) {
	request := web.CategoryCreateRequest{}

	err := c.Bind(&request)
	if err != nil {
		panic(err)
	}

	categoryResponse := controller.CategoryService.Create(request)
	c.JSON(http.StatusOK, categoryResponse)
}

func (controller *CategoryControllerImpl) Update(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *CategoryControllerImpl) Delete(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *CategoryControllerImpl) FindById(c *gin.Context) {
	id := c.Param("id")

	response := controller.CategoryService.FindById(id)

	c.JSON(200, response)
}

func (controller *CategoryControllerImpl) FindAll(c *gin.Context) {
	responses := controller.CategoryService.FindAll()

	c.JSON(http.StatusOK, responses)
}

func (controller *CategoryControllerImpl) SetupEngine() *gin.Engine {
	return controller.engine
}
