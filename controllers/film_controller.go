package controllers

import "github.com/gin-gonic/gin"

type FilmController interface {
	FindById(c *gin.Context)
	GetDataCategoryFIlm(c *gin.Context)
}
