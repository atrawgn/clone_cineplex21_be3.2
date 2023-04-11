package entity

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
