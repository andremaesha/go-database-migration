package service

import (
	"database-migration/models/domain"
	"database-migration/models/web"
	"database-migration/repository"
	"gorm.io/gorm"
)

type FilmServiceImpl struct {
	DB         *gorm.DB
	repository repository.FilmRepository
}

func NewFilmServiceImpl(DB *gorm.DB, repository repository.FilmRepository) FilmService {
	return &FilmServiceImpl{DB: DB, repository: repository}
}

func (service *FilmServiceImpl) GetDataById(id uint64) web.FilmResponse {
	film := domain.Film{FilmId: id}

	data := service.repository.GetById(service.DB, film)

	return web.FilmResponse{
		Message:     "success",
		Code:        200,
		IdFilm:      int(data.FilmId),
		Title:       data.Title,
		Description: data.Description,
		ReleaseYear: data.ReleaseYear,
		FullText:    data.Fulltext,
	}
}
