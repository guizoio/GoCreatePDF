package infra

import (
	"CreateFilePDF/src/configs/database"
	"CreateFilePDF/src/generator"
	"CreateFilePDF/src/infra/adapters/gorm/repository"
	"CreateFilePDF/src/user_case/create_file"
	"gorm.io/gorm"
	"os"
)

type ContainerDI struct {
	DB            *gorm.DB
	CreateHandler create_file.CreateHandler
	CreatePDF     generator.CreatePDF
}

func NewContainerDI() *ContainerDI {

	container := &ContainerDI{}

	config := database.Config{
		Hostname: os.Getenv("DB_HOST_LOCAL"),
		Port:     os.Getenv("DB_PORT"),
		UserName: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	}
	container.DB = database.InitGorm(&config)

	container.build()
	return container
}

func (c *ContainerDI) build() {

	repositoryCreate := repository.NewCreateRepository(c.DB)
	c.CreatePDF = generator.NewCreatePDF(
		c.CreatePDF.HeaderPDF,
		c.CreatePDF.People,
		c.CreatePDF.Company,
		repositoryCreate,
	)
	c.CreateHandler = create_file.NewCreateHandler(c.CreatePDF)
}
func (c *ContainerDI) ShutDown() {}
