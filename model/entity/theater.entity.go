package entity

type Theater struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Kota    string `json:"kota"`
	Theater string `json:"theater"`
	Phone   string `json:"phone"`
}
