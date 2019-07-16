package article

import (
	"github.com/ryo-ymd/go-test/db"
	"github.com/ryo-ymd/go-test/entity"
)

type Service struct {
}

type Article entity.Article

func (s Service) GetAll() ([]Article, error) {
	db := db.GetDB()
	var a []Article

	if err := db.Find(&a).Error; err != nil {
		return nil, err
	}

	return a, nil
}
