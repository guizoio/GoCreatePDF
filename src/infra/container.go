package infra

import (
	"CreateFilePDF/src/configs/database"
	"gorm.io/gorm"
	"os"
)

type ContainerDI struct {
	DB *gorm.DB
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

	return container
}

func (c *ContainerDI) build()    {}
func (c *ContainerDI) ShutDown() {}
