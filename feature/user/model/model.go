package model

type Users struct {
	Nim         string
	Name        string
	Password    string
	Class       string
	Program     string
	Total_Score float64
}

// func (U *Users) BeforeCreate(tx *gorm.DB) (err error) {
// 	newUuid := uuid.New()
// 	U.Id = newUuid.String()

// 	return nil
// }
