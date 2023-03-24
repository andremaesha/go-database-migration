package service

import "database-migration/models/web"

type Service interface {
	GetCategoryAndFilmById(categoryId string, filmId uint64) web.ResponseCategoryFilm
}
