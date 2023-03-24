package repository

import (
	"database-migration/models/domain"
	"gorm.io/gorm"
)

type FilmRepository interface {
	GetById(db *gorm.DB, film domain.Film) domain.Film
}
