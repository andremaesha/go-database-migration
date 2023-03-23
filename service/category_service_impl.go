package service

import (
	"database-migration/models/domain"
	"database-migration/models/web"
	"database-migration/repository"
	"gorm.io/gorm"
)

type CategoryServiceImpl struct {
	DB         *gorm.DB
	repository repository.CategoryRepository
}

func NewCategoryServiceImpl(DB *gorm.DB, repository repository.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{DB: DB, repository: repository}
}

func (service *CategoryServiceImpl) Create(request web.CategoryCreateRequest) web.CategoryResponse {
	category := domain.Category{
		Name: request.Name,
	}

	save := service.repository.Save(service.DB, category)

	return web.CategoryResponse{
		Id:   save.Id,
		Name: save.Name,
	}
}

func (service *CategoryServiceImpl) Update(request web.CategoryUpdateRequest) web.CategoryResponse {
	//TODO implement me
	panic("implement me")
}

func (service *CategoryServiceImpl) Delete(categoryId int) {
	//TODO implement me
	panic("implement me")
}

func (service *CategoryServiceImpl) FindById(categoryId string) web.CategoryResponse {
	response := web.CategoryResponse{}

	data, err := service.repository.FindById(service.DB, categoryId)
	if err != nil {
		panic(err)
	}

	response.Id = data.Id
	response.Name = data.Name

	return response
}

func (service *CategoryServiceImpl) FindAll() []web.CategoryResponse {
	var response []web.CategoryResponse

	datas := service.repository.FindAll(service.DB)

	for _, item := range datas {
		response = append(response, web.CategoryResponse{
			Id:   item.Id,
			Name: item.Name,
		})
	}

	return response
}
