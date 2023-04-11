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

	/*hashedPassword, err := utils.HashingPassword(user.Password)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	newUser.Password = hashedPassword*/

	errCreateUser := database.DB.Create(&newTheater).Error
	if errCreateUser != nil {
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
	userRequest := new(request.UserUpdateRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var user entity.User

	userId := ctx.Params("id")
	// CHECK AVAILABLE USER
	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data tidak valid",
		})
	}

	// UPDATE USER DATA
	if userRequest.Nama != "" {
		user.Nama = userRequest.Nama
	}

	errUpdate := database.DB.Save(&user).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Sukses",
		"data":    user,
	})
}

func TheaterControllerDelete(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	var user entity.User

	// CHECK AVAILABLE USER
	err := database.DB.Debug().First(&user, "id=?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user tidak ditemukan",
		})
	}

	errDelete := database.DB.Debug().Delete(&user).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "user telah dihapus",
	})
}

/*func TheaterControllerUpdateEmail(ctx *fiber.Ctx) error {
	userRequest := new(request.UserEmailRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var user entity.User

	userId := ctx.Params("id")
	// CHECK AVAILABLE USER
	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data tidak valid",
		})
	}

	// UPDATE USER DATA
	user.Email = userRequest.Email

	errUpdate := database.DB.Save(&user).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Sukses",
		"data":    user,
	})
}*/
