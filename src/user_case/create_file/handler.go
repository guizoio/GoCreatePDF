package create_file

import "github.com/gofiber/fiber/v2"

type CreateHandler struct {
}

func NewCreateHandler() CreateHandler {
	return CreateHandler{}
}

func (ref *CreateHandler) Get(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON("success")
}
