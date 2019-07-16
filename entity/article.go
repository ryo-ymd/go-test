package entity

// Article is a article model property
type Article struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	ArticleUnits []ArticleUnit
}
