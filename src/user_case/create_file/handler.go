package create_file

import (
	"CreateFilePDF/src/entity"
	"CreateFilePDF/src/generator"
	"CreateFilePDF/src/user_case/storage_client"
	"github.com/gofiber/fiber/v2"
	"os"
	"strings"
	"time"
)

type CreateHandler struct {
	CreatePDF      generator.CreatePDF
	ServiceStorage storage_client.FaceServiceStorage
	BucketStorage  string
}

func NewCreateHandler(CreatePDF generator.CreatePDF, ServiceStorage storage_client.FaceServiceStorage, BucketStorage string) CreateHandler {
	return CreateHandler{CreatePDF, ServiceStorage, BucketStorage}
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

	fileName := rideName(buff.Name)
	ref.CreatePDF.CreatePDF(fileName)

	defer os.Remove("./" + fileName)

	ref.ServiceStorage.UploadFile(ref.BucketStorage, fileName)
	return c.Status(fiber.StatusOK).JSON(fileName)
}

func rideName(name string) string {
	year := time.Now().Format("2006")
	month := time.Now().Format("01")
	day := time.Now().Format("02")
	hour := time.Now().Format("15")
	minute := time.Now().Format("04")
	second := time.Now().Format("05")
	return year + "_" + month + "_" + day + "_" + hour + "_" + minute + "_" + second + "_" +
		strings.Replace(name, " ", "_", -1) + ".pdf"
}
