package service

import (
	"database-migration/models/web"
)

type CategoryService interface {
	Create(request web.CategoryCreateRequest) web.CategoryResponse
	Update(request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(categoryId int)
	FindById(categoryId string) web.CategoryResponse
	FindAll() []web.CategoryResponse
}
