package model


type Application struct {
	Model
	Name string `gorm:"not null" json:"name"`
	Type string `gorm:"type:text;not null" json:"type"`
	Description string `gorm:"type:text;not null"  json:"description"`
	AgeGroup AgeGroup  `json:"age-group"`
	Purpose string  `gorm:"type:text;" json:"purpose"`
	BaseURL string`json:"baseURL"`
	Company string `json:"company"`

}

type AgeGroup struct {
	Min int `json:"min"`
	Max int `json:"max"`
}