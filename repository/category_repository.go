package repository

import (
	"database-migration/models/domain"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Save(db *gorm.DB, category domain.Category) domain.Category
	Update(db *gorm.DB, category domain.Category) domain.Category
	Delete(db *gorm.DB, category domain.Category)
	FindById(db *gorm.DB, categoryId string) (domain.Category, error)
	FindAll(db *gorm.DB) []domain.Category
}
