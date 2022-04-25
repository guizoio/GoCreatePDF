package infra

import (
	"CreateFilePDF/src/configs/database"
	"CreateFilePDF/src/user_case/create_file"
	"gorm.io/gorm"
	"os"
)

type ContainerDI struct {
	DB            *gorm.DB
	CreateHandler create_file.CreateHandler
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
	c.CreateHandler = create_file.NewCreateHandler()
}
func (c *ContainerDI) ShutDown() {}
