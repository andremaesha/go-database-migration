package repository

import (
	"database-migration/models/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepositoryImpl() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (c *CategoryRepositoryImpl) Save(db *gorm.DB, category domain.Category) domain.Category {
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := db.Table("first").Debug().Select("name").Create(&category).Clauses(clause.Returning{}).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	return category
}

func (c *CategoryRepositoryImpl) Update(db *gorm.DB, category domain.Category) domain.Category {
	//TODO implement me
	panic("implement me")
}

func (c *CategoryRepositoryImpl) Delete(db *gorm.DB, category domain.Category) {
	//TODO implement me
	panic("implement me")
}

func (c *CategoryRepositoryImpl) FindById(db *gorm.DB, categoryId string) (domain.Category, error) {
	data := domain.Category{}

	err := db.Table("first").Select("id,name").Where("id = ?", categoryId).Scan(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (c *CategoryRepositoryImpl) FindAll(db *gorm.DB) []domain.Category {
	var datas []domain.Category

	err := db.Table("first").Select("id,name").Find(&datas).Error
	if err != nil {
		panic(err)
	}

	return datas
}
