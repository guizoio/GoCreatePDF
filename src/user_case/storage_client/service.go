package storage_client

import (
	"CreateFilePDF/src/entity"
	"CreateFilePDF/src/infra/adapters/minio_client"
)

type FaceServiceStorage interface {
	CheckLife() string
	ListBuckets() ([]*entity.BucketInfo, error)
	ListObjects(bucket string) ([]*entity.ObjectIndo, error)
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

func (s *ServiceStorage) ListBuckets() ([]*entity.BucketInfo, error) {
	return s.client.ListBuckets()
}

func (s *ServiceStorage) ListObjects(bucket string) ([]*entity.ObjectIndo, error) {
	result, err := s.client.ListBucketObjects(bucket)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ServiceStorage) UploadFile(bucketName, fileName string) error {
	err := s.client.UploadObject(bucketName, fileName)
	if err != nil {
		return err
	}
	return nil
}
