package infra

import (
	"CreateFilePDF/src/configs/database"
	"CreateFilePDF/src/configs/storage"
	"CreateFilePDF/src/generator"
	"CreateFilePDF/src/infra/adapters/gorm/repository"
	kafkaService "CreateFilePDF/src/infra/adapters/kafka"
	"CreateFilePDF/src/infra/adapters/minio_client"
	"CreateFilePDF/src/user_case/create_file"
	"CreateFilePDF/src/user_case/storage_client"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
	"os"
	"strconv"
)

type ContainerDI struct {
	DB             *gorm.DB
	CreateHandler  create_file.CreateHandler
	CreatePDF      generator.CreatePDF
	Storage        *minio.Client
	StorageConnect minio_client.ClientMinio
	StorageClient  storage_client.StorageClient
	PublishMessage generator.PublishMessage
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

	useSSL, err := strconv.ParseBool(os.Getenv("USE_SSL"))
	if err != nil {
		panic("Error convert useSSL to bool: " + err.Error())
	}
	configMinio := storage.Config{
		Endpoint:        os.Getenv("ENDPOINT"),
		AccessKeyID:     os.Getenv("ACCESS_KEY_ID"),
		SecretAccessKey: os.Getenv("SECRET_ACCESS_KET"),
		UseSSL:          useSSL,
	}
	container.Storage = storage.InitStorage(&configMinio)

	configKafka := kafkaService.NewKafkaWriter(os.Getenv("KAFKA_SERVER"), false)
	container.PublishMessage = kafkaService.NewKafkaClient(configKafka)

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
		c.PublishMessage,
	)

	clientMinio := minio_client.NewClientMinio(c.Storage)
	serviceMinio := storage_client.NewServiceStorage(clientMinio)
	c.StorageClient = storage_client.NewStorageClient(serviceMinio)

	c.CreateHandler = create_file.NewCreateHandler(c.CreatePDF, serviceMinio, os.Getenv("BUCKET"))
}
func (c *ContainerDI) ShutDown() {}
