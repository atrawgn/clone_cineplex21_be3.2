package controller

import (
	"github.com/atrawiguna/golang-restapi-gorm/database"
	"github.com/atrawiguna/golang-restapi-gorm/model/entity"
	"github.com/atrawiguna/golang-restapi-gorm/model/request"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

// User Section

func TheaterControllerGetByKota(ctx *fiber.Ctx) error {
	kota := ctx.Params("kota") // get the value of the 'kota' parameter from the request URL
	userInfo := ctx.Locals("userInfo")
	log.Println("user info data :: ", userInfo)

	var theaters []entity.Theater
	if err := database.DB.Where("kota = ?", kota).Find(&theaters).Error; err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).SendString("Internal Server Error")
	}

	return ctx.JSON(theaters)
}

/*func TheaterControllerGetDetails(ctx *fiber.Ctx) error {
	theaterid := ctx.QueryInt("theaterid")
	var film []entity.TheaterList
	err := database.DB.Raw(`
		SELECT t.id, t.kota, t.theater, t.phone
		FROM theaters
		INNER JOIN theater_lists l ON l.theater_id = theater.id
		WHERE theaters.id = ?`, theaterid).Scan(&theater)

	if err.Error != nil {
		log.Println(err.Error)
	}

	return ctx.JSON(fiber.Map{
		"message": "successfully",
		"data":    theater,
	})
}*/

func TheaterControllerGetDetails(ctx *fiber.Ctx) error {
	theaterId := ctx.QueryInt("theaterid")

	var theater entity.Theater

	err := database.DB.First(&theater, "id = ?", theaterId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data tidak ditemukan",
		})
	}

	var film []entity.TheaterId
	err = database.DB.Raw(`
		SELECT f.id, f.judul, l.theater_id AS theater_id, f.jenis_film, f. produser, f.sutradara, f.penulis, f.produksi, f.casts, f.sinopsis, f.like
		FROM films f
		INNER JOIN theater_lists l ON l.film_id = f.id
		WHERE l.theater_id = ?`, theaterId).Scan(&film).Error

	var theaterdetails entity.TheaterDetails
	theaterdetails.ID = theater.ID
	theaterdetails.Kota = theater.Kota
	theaterdetails.Theater = theater.Theater
	theaterdetails.Phone = theater.Phone
	theaterdetails.Film = film

	return ctx.JSON(fiber.Map{
		"theater": theater,
		"film":    film,
		"details": theaterdetails,
		"message": "success",
	})
}

func TheaterControllerCreate(ctx *fiber.Ctx) error {
	theater := new(request.TheaterCreateRequest)
	if err := ctx.BodyParser(theater); err != nil {
		return err
	}

	// VALIDASI REQUEST
	validate := validator.New()
	errValidate := validate.Struct(theater)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Gagal",
			"error":   errValidate.Error(),
		})
	}

	newTheater := entity.Theater{
		Kota:    theater.Kota,
		Theater: theater.Theater,
		Phone:   theater.Phone,
	}

	errCreateTheater := database.DB.Create(&newTheater).Error
	if errCreateTheater != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Tidak berhasil menyimpan data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Berhasil",
		"data":    newTheater,
	})
}

func TheaterControllerGetById(ctx *fiber.Ctx) error {
	theaterId := ctx.Params("id")

	var theater entity.Theater
	err := database.DB.First(&theater, "id = ?", theaterId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data tidak ditemukan",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Sukses",
		"data":    theater,
	})
}

func TheaterControllerUpdate(ctx *fiber.Ctx) error {
	theaterRequest := new(request.TheaterUpdateRequest)
	if err := ctx.BodyParser(theaterRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var theater entity.Theater

	theaterId := ctx.Params("id")
	// CHECK AVAILABLE USER
	err := database.DB.First(&theater, "id = ?", theaterId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data tidak valid",
		})
	}

	// UPDATE USER DATA
	if theaterRequest.Theater != "" {
		theater.Theater = theaterRequest.Theater
	}

	errUpdate := database.DB.Save(&theater).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Sukses",
		"data":    theater,
	})
}

func TheaterControllerDelete(ctx *fiber.Ctx) error {
	theaterId := ctx.Params("id")
	var theater entity.Theater

	// CHECK AVAILABLE USER
	err := database.DB.Debug().First(&theater, "id=?", theaterId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "theater tidak ditemukan",
		})
	}

	errDelete := database.DB.Debug().Delete(&theater).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "theater telah dihapus",
	})
}

func TheaterControllerCreateList(ctx *fiber.Ctx) error {
	Theater := new(request.TheaterListCreateRequest)

	if err := ctx.BodyParser(Theater); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(Theater)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "gagal",
			"error":   errValidate.Error(),
		})
	}

	newTheaterList := entity.TheaterList{
		TheaterID: Theater.TheaterId,
		FilmID:    Theater.FilmId,
	}

	errCreateTheater := database.DB.Create(&newTheaterList).Error
	if errCreateTheater != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "gagal menyimpan data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newTheaterList,
	})
}
