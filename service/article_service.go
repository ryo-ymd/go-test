package article

import (
	"github.com/ryo-ymd/go-test/db"
	"github.com/ryo-ymd/go-test/entity"
)

type Service struct {}

type Article entity.Article

// ArticleServiceはArticleのDB操作を行う
func (s Service) GetAll() ([]Article, error) {
	db := db.GetDB()
	var a []Article

	if err := db.Preload("ArticleUnits").Find(&a).Error; err != nil {
		return nil, err
	}

	return a, nil
}
