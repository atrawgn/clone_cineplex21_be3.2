package controller

import (
	"github.com/atrawiguna/golang-restapi-gorm/database"
	"github.com/atrawiguna/golang-restapi-gorm/model/entity"
	"github.com/atrawiguna/golang-restapi-gorm/model/request"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
)

func BookControllerGet(ctx *fiber.Ctx) error {
	var book []entity.Book
	err := database.DB.Find(&book)

	if err.Error != nil {
		log.Println(err.Error)
	}
	return ctx.JSON(book)
}

func BookControllerCreate(ctx *fiber.Ctx) error {
	book := new(request.BookCreateRequest)
	if err := ctx.BodyParser(book); err != nil {
		return err
	}

	// VALIDASI REQUEST
	validate := validator.New()
	errValidate := validate.Struct(book)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Gagal",
			"error":   errValidate.Error(),
		})
	}

	newBook := entity.Book{
		Title:    book.Title,
		Synopsis: book.Synopsis,
		Content:  book.Content,
		Author:   book.Author,
	}

	errCreateBook := database.DB.Create(&newBook).Error
	if errCreateBook != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Tidak berhasil menyimpan data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Berhasil",
		"data":    newBook,
	})
}

func BookControllerGetById(ctx *fiber.Ctx) error {
	bookId := ctx.Params("id")

	var book entity.Book
	err := database.DB.First(&book, "id = ?", bookId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data tidak ditemukan",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Sukses",
		"data":    book,
	})
}

func BookControllerUpdate(ctx *fiber.Ctx) error {
	bookRequest := new(request.BookUpdateRequest)
	if err := ctx.BodyParser(bookRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var book entity.Book

	bookId := ctx.Params("id")
	// CHECK AVAILABLE USER
	err := database.DB.First(&book, "id = ?", bookId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data tidak valid",
		})
	}

	// UPDATE USER DATA
	if bookRequest.Title != "" {
		book.Title = bookRequest.Title
	}
	errUpdate := database.DB.Save(&book).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	book.Title = bookRequest.Title

	return ctx.JSON(fiber.Map{
		"message": "Sukses",
		"data":    book,
	})
}

func BookControllerDelete(ctx *fiber.Ctx) error {
	bookId := ctx.Params("id")
	var book entity.Book

	// CHECK AVAILABLE USER
	err := database.DB.Debug().First(&book, "id=?", bookId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "buku tidak ditemukan",
		})
	}

	errDelete := database.DB.Debug().Delete(&book).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "buku telah dihapus",
	})
}
