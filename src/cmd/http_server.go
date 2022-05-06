package cmd

import (
	"CreateFilePDF/src/infra"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"time"
)

func StartHttp(ctx context.Context, containerDI *infra.ContainerDI) {
	app := fiber.New(fiber.Config{
		StrictRouting: true,
	})

	go func() {
		for {
			select {
			case <-ctx.Done():
				if err := app.Shutdown(); err != nil {
					panic(err)
				}
				return
			default:
				time.Sleep(1 * time.Second)
			}
		}
	}()

	app.Use(cors.New(cors.Config{
		AllowHeaders: "*",
	}))

	app.Get("/check", containerDI.CreateHandler.Check)
	app.Post("/create", containerDI.CreateHandler.CreateFilePDF)

	app.Get("/storage/check", containerDI.StorageClient.Check)
	app.Get("/storage/list/buckets", containerDI.StorageClient.ListBuckets)
	app.Get("/storage/list/:bucket", containerDI.StorageClient.ListObjects)
	app.Get("/storage/download/:bucket/:object", containerDI.StorageClient.Download)

	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
