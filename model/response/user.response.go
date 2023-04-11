package response

type UserResponse struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"-" gorm:"column:password"`
}
