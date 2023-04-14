package entity

type Theater struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Kota    string `json:"kota"`
	Theater string `json:"theater"`
	Phone   string `json:"phone"`
}

type TheaterDetails struct {
	ID      uint        `json:"id" gorm:"primaryKey"`
	Kota    string      `json:"kota"`
	Theater string      `json:"theater"`
	Phone   string      `json:"phone"`
	Film    []TheaterId `json:"film"`
}
