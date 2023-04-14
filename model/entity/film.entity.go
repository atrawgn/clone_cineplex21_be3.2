package entity

type Film struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Cover     string `json:"cover"`
	Judul     string `json:"judul"`
	JenisFilm string `json:"jenis_film"`
	Produser  string `json:"produser"`
	Sutradara string `json:"sutradara"`
	Penulis   string `json:"penulis"`
	Produksi  string `json:"produksi"`
	Casts     string `json:"casts"`
	Sinopsis  string `json:"sinopsis"`
	Like      uint   `json:"like"`
}

type TheaterId struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	TheaterId uint   `json:"theater_id"`
	Cover     string `json:"cover"`
	Judul     string `json:"judul"`
	JenisFilm string `json:"jenis_film"`
	Produser  string `json:"produser"`
	Sutradara string `json:"sutradara"`
	Penulis   string `json:"penulis"`
	Produksi  string `json:"produksi"`
	Casts     string `json:"casts"`
	Sinopsis  string `json:"sinopsis"`
	Like      uint   `json:"like"`
}
