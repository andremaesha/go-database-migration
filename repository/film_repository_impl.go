package repository

import (
	"database-migration/models/domain"
	"gorm.io/gorm"
)

type FilmRepositoryImpl struct {
}

func NewFilmRepositoryImpl() FilmRepository {
	return &FilmRepositoryImpl{}
}

func (reposiory *FilmRepositoryImpl) GetById(db *gorm.DB, film domain.Film) domain.Film {
	err := db.Debug().Table("film").First(&film).Error
	if err != nil {
		panic(err)
	}

	return film
}
