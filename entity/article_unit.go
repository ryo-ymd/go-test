package entity

// ArticleUnit is a article unit model property
type ArticleUnit struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	ArticleID    uint   `json:"articleId"`
	UnitableType string `json:"unitableType"`
	UnitableID   uint   `json:"unitableId"`
	Position     rune   `json:"position"`
	CreatedAt    string `json:"createdAt"`
}
