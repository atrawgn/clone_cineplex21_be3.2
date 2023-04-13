package migration

import (
	"fmt"
	"github.com/atrawiguna/golang-restapi-gorm/database"
	"github.com/atrawiguna/golang-restapi-gorm/model/entity"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{}, &entity.Film{}, &entity.Theater{}, &entity.Comment{}, &entity.TheaterList{}, &entity.Book{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Databases migrated successfully")
}
