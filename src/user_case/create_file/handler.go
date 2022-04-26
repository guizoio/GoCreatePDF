package create_file

import (
	"CreateFilePDF/src/entity"
	"CreateFilePDF/src/generator"
	"github.com/gofiber/fiber/v2"
	"os"
)

type CreateHandler struct {
	CreatePDF generator.CreatePDF
}

func NewCreateHandler(CreatePDF generator.CreatePDF) CreateHandler {
	return CreateHandler{CreatePDF}
}

func (ref *CreateHandler) Check(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON("success")
}

func (ref *CreateHandler) CreateFilePDF(c *fiber.Ctx) error {
	var request RequestCreatePDF
	if result := c.BodyParser(&request); result != nil {
		return c.Status(fiber.StatusBadRequest).JSON("ERROR: " + result.Error())
	}
	buff := request.ToDomain()
	var headerPDF entity.HeadlerPDF
	var companyPDF entity.Company

	headerPDF.FilePDF = request.FilePDF
	headerPDF.FileIMG = request.FileIMG

	ref.CreatePDF.HeaderPDF = headerPDF
	ref.CreatePDF.People = buff
	ref.CreatePDF.Company = companyPDF

	ref.CreatePDF.CreatePDF()

	//return c.Status(fiber.StatusOK).JSON(buff)
	defer os.Remove("./Registration.pdf")
	return c.Download("./Registration.pdf", "Registration.pdf")
}
