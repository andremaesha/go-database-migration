package service

import (
	"database-migration/models/domain"
	"database-migration/models/web"
	"database-migration/repository"
	"gorm.io/gorm"
)

type ServiceImpl struct {
	DB                 *gorm.DB
	repositoryFilm     repository.FilmRepository
	repositoryCategory repository.CategoryRepository
}

func NewServiceImpl(DB *gorm.DB, repositoryFilm repository.FilmRepository, repositoryCategory repository.CategoryRepository) Service {
	return &ServiceImpl{DB: DB, repositoryFilm: repositoryFilm, repositoryCategory: repositoryCategory}
}

func (service *ServiceImpl) GetCategoryAndFilmById(categoryId string, filmId uint64) web.ResponseCategoryFilm {
	categoryData, err := service.repositoryCategory.FindById(service.DB, categoryId)
	if err != nil {
		panic(err)
	}
	filmData := service.repositoryFilm.GetById(service.DB, domain.Film{FilmId: filmId})

	return web.ResponseCategoryFilm{
		CategoryResponse: web.CategoryResponse{
			Id:   categoryData.Id,
			Name: categoryData.Name,
		},
		FilmResponse: web.FilmResponse{
			Message:     "success",
			Code:        200,
			IdFilm:      int(filmData.FilmId),
			Title:       filmData.Title,
			Description: filmData.Description,
			ReleaseYear: filmData.ReleaseYear,
			FullText:    filmData.Fulltext,
		},
	}
}
