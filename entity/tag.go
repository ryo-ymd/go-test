package entity

// Tag is a tag model property
type Tag struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Key         string `json:"key"`
	ColorCode string `json:"colorCode"`
	ContentType string `json:"contentType"`
	Description  string `json:"description"`
}
