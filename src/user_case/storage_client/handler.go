package storage_client

import (
	"github.com/gofiber/fiber/v2"
	"os"
)

type StorageClient struct {
	Service FaceServiceStorage
}

func NewStorageClient(Service FaceServiceStorage) StorageClient {
	return StorageClient{Service}
}

func (ref *StorageClient) Check(c *fiber.Ctx) error {
	result := ref.Service.CheckLife()
	return c.Status(fiber.StatusOK).JSON(result)
}

func (ref *StorageClient) ListBuckets(c *fiber.Ctx) error {
	result, err := ref.Service.ListBuckets()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("ERROR: " + err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (ref *StorageClient) ListObjects(c *fiber.Ctx) error {
	bucket := c.Params("bucket")
	result, err := ref.Service.ListObjects(bucket)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("ERROR: " + err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (ref *StorageClient) Download(c *fiber.Ctx) error {
	bucket := c.Params("bucket")
	object := c.Params("object")
	defer os.Remove("./" + object)
	err := ref.Service.DownloadFile(bucket, object)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("ERROR: " + err.Error())
	}
	return c.Download("./"+object, object)
}
