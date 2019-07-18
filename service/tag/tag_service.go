package tag

import (
	"github.com/ryo-ymd/go-test/db"
	"github.com/ryo-ymd/go-test/entity"
)

type Service struct{}

type Tag entity.Tag

// TagServiceはTagのDB操作を行う
func (s Service) GetAll() ([]Tag, error) {
	database := db.GetDB()

	var t []Tag

	if err := database.Find(&t).Error; err != nil {
		return nil, err
	}

	return t, nil
}

func (s Service) FindById(id uint) (Tag, error) {

	database := db.GetDB()
	var t Tag

	if err := database.Where("id = ?", id).First(&t).Error; err != nil {
		return t, err
	}

	return t, nil
}

func (s Service) GetTree(id uint) ([]Tag, error) {
	database := db.GetDB()
	database.LogMode(true)
	var t []Tag

	if err := database.Debug().
		Select("tags.*").
		Joins("inner join tag_hierarchies on tags.id = tag_hierarchies.ancestor_id").
		Where("tag_hierarchies.descendant_id = ?", id).
		Order("tag_hierarchies.generations asc").Find(&t).Error; err != nil {
		return t, err
	}
	return t, nil
}
