package storage_client

import (
	"CreateFilePDF/src/infra/adapters/minio_client"
	"github.com/minio/minio-go/v7"
)

type FaceServiceStorage interface {
	CheckLife() string
	ListBuckets() ([]minio.BucketInfo, error)
	UploadFile(bucketName, fileName string) error
}

type ServiceStorage struct {
	client minio_client.FaceClientMinio
}

func NewServiceStorage(client minio_client.FaceClientMinio) *ServiceStorage {
	return &ServiceStorage{client}
}

func (s *ServiceStorage) CheckLife() string {
	return s.client.CheckLife()
}

func (s *ServiceStorage) ListBuckets() ([]minio.BucketInfo, error) {
	return s.client.ListBuckets()
}

func (s *ServiceStorage) UploadFile(bucketName, fileName string) error {
	err := s.client.UploadObject(bucketName, fileName)
	if err != nil {
		return err
	}
	return nil
}
