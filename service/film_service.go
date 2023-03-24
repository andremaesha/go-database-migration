package service

import "database-migration/models/web"

type FilmService interface {
	GetDataById(id uint64) web.FilmResponse
}
