package storage

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func InitStorage(config *Config) *minio.Client {
	minioClient, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
		Secure: config.UseSSL,
	})
	if err != nil {
		panic("failed to connect storage: " + err.Error())
	}

	return minioClient
}
