package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "db.vjpgdoelfnmgwmpdokkn.supabase.co"
	port     = 5432
	user     = "postgres"
	password = "SuretyBond2023!"
	dbname   = "Cineplex_Team_3"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Tidak dapat terkoneksi ke database")
	}
	fmt.Println("Terkoneksi ke Database")

}
