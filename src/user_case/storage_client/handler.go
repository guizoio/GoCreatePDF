package storage_client

import "github.com/gofiber/fiber/v2"

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
	ref.Service.ListBuckets()
	return c.Status(fiber.StatusOK).JSON("")
}
