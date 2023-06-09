package request

type FilmCreateRequest struct {
	Judul     string `json:"judul" validate:"required"`
	JenisFilm string `json:"jenis_film"`
	Produser  string `json:"produser" validate:"required"`
	Sutradara string `json:"sutradara" validate:"required"`
	Penulis   string `json:"penulis" validate:"required"`
	Produksi  string `json:"produksi" validate:"required"`
	Casts     string `json:"casts" validate:"required"`
	Sinopsis  string `json:"sinopsis" validate:"required"`
	Cover     string `json:"cover"`
}

type FilmUpdateRequest struct {
	Judul     string `json:"judul"`
	JenisFilm string `json:"jenis_film"`
	Produser  string `json:"produser"`
	Sutradara string `json:"sutradara"`
	Penulis   string `json:"penulis"`
	Produksi  string `json:"produksi"`
	Casts     string `json:"casts"`
	Sinopsis  string `json:"sinopsis"`
}

type FilmLikeUpdateRequest struct {
	Like uint `json:"like"`
}

// USER SECTION
type UserCreateRequest struct {
	Nama     string `json:"nama" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserUpdateRequest struct {
	Nama string `json:"nama"`
}

type UserEmailRequest struct {
	Email string `json:"email" validate:"required"`
}

//THEATER SECTION

type TheaterCreateRequest struct {
	Kota    string `json:"kota"`
	Theater string `json:"theater"`
	Phone   string `json:"phone"`
}
type TheaterUpdateRequest struct {
	Kota    string `json:"kota"`
	Theater string `json:"theater"`
	Phone   string `json:"phone"`
}
type TheaterListCreateRequest struct {
	TheaterId uint `json:"theater_id"`
	FilmId    uint `json:"film_id"`
}

//BOOK SECTION

type BookCreateRequest struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Title    string `json:"title"`
	Synopsis string `json:"sinopsis"`
	Content  string `json:"content"`
	Author   string `json:"author"`
}

type BookUpdateRequest struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Title    string `json:"title"`
	Synopsis string `json:"sinopsis"`
	Content  string `json:"content"`
	Author   string `json:"author"`
}

//COMMENT

type CommentCreateRequest struct {
	FilmId  uint   `json:"film_id" validate:"required"`
	Comment string `json:"comment" validate:"required"`
}
