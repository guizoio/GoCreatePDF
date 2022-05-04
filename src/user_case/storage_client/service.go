package storage_client

import (
	"CreateFilePDF/src/infra/adapters/minio_client"
)

type FaceServiceStorage interface {
	CheckLife() string
	ListBuckets()
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

func (s *ServiceStorage) ListBuckets() {
	s.client.ListBuckets()
}

func (s *ServiceStorage) UploadFile(bucketName, fileName string) error {
	err := s.client.UploadObject(bucketName, fileName)
	if err != nil {
		return err
	}
	return nil
}
