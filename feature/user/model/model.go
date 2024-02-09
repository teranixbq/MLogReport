package model

type Users struct {
	Nim         string `gorm:"primaryKey"`
	Name        string
	Password    string
	Class       string
	Program     string
	Total_Score float64
}

