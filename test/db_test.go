package test

import (
	"database-migration/configs"
	"database-migration/models/domain"
	"database-migration/models/web"
	"database-migration/repository"
	"database-migration/service"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"testing"
)

func TestConnectToDB(t *testing.T) {
	configs.NewDB()
}

func TestAddFirst(t *testing.T) {
	category := domain.Category{
		Name: "test golang hihi",
	}

	db := configs.NewDB()
	err := db.Table("first").Debug().Create(&category).Clauses(clause.Returning{}).Error
	if err != nil {
		panic(err)
	}

	fmt.Println(category)
}

func TestAddTxFirst(t *testing.T) {
	category := domain.Category{Name: "andre maesha"}

	db := configs.NewDB()
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := db.Table("first").Debug().Select("name").Create(&category).Clauses(clause.Returning{}).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(category)
}

func TestQueryFirst(t *testing.T) {
	category := domain.Category{Name: "andre"}
	film := domain.Film{}

	db := configs.NewDB()
	//datas := make(map[string]interface{})
	//datasFilm := make(map[string]interface{})

	//err := db.Table("first").Select("id,name").Find(datas).Scan(&category).Error
	//if err != nil {
	//	panic(err)
	//}

	err := db.Table("film").Select("film_id,title,description,release_year,fulltext").First(&film).Error
	if err != nil {
		panic(err)
	}

	fmt.Println(category)
	fmt.Println(film)
}

func TestQueryAllFIrst(t *testing.T) {
	var datas []domain.Category

	err := configs.NewDB().Table("first").Select("id,name").Find(&datas).Error
	if err != nil {
		panic(err)
	}

	fmt.Println(datas)
}

func TestRepositoryServiceFirst(t *testing.T) {
	request := web.CategoryCreateRequest{Name: "tot 1"}

	db := configs.NewDB()
	categoryRepository := repository.NewCategoryRepositoryImpl()
	categoryService := service.NewCategoryServiceImpl(db, categoryRepository)
	response := categoryService.Create(request)

	fmt.Println(response)
}
