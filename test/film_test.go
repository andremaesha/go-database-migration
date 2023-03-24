package test

import (
	"database-migration/configs"
	"database-migration/models/domain"
	"fmt"
	"testing"
)

func TestFIlmGetById(t *testing.T) {
	film := domain.Film{
		FilmId: 133,
	}

	db := configs.NewDB()
	err := db.Debug().Table("film").First(&film).Error
	if err != nil {
		panic(err)
	}

	fmt.Println(film)
}
