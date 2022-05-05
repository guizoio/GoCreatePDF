package minio_client

import (
	"CreateFilePDF/src/entity"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"os"
)

type FaceClientMinio interface {
	CheckLife() string
	ListBuckets() ([]minio.BucketInfo, error)
	UploadObject(bucketName, fileName string) error
	ListBucketObjects(bucket string) ([]*entity.ObjectIndo, error)
}

type ClientMinio struct {
	minioClient *minio.Client
}

func NewClientMinio(minioClient *minio.Client) *ClientMinio {
	return &ClientMinio{minioClient}
}

func (c ClientMinio) CheckLife() string {
	return c.minioClient.EndpointURL().String()
}

func (c ClientMinio) ListBuckets() ([]minio.BucketInfo, error) {
	buckets, err := c.minioClient.ListBuckets(context.Background())
	if err != nil {
		return nil, err
	}
	return buckets, nil
}

func (c ClientMinio) ListBucketObjects(bucket string) ([]*entity.ObjectIndo, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	objectCh := c.minioClient.ListObjects(ctx, bucket, minio.ListObjectsOptions{
		Recursive: true,
	})

	var list []*entity.ObjectIndo

	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return nil, object.Err
		}
		objects := entity.ObjectIndo{
			Name: object.Key,
			Date: object.LastModified,
			Size: object.Size,
		}
		list = append(list, objects.ToDomain())
	}
	return list, nil
}

func (c ClientMinio) UploadObject(bucketName, fileName string) error {
	file, errOpenFile := os.Open(fileName)
	if errOpenFile != nil {
		return errOpenFile
	}
	defer file.Close()
	fileStat, err := file.Stat()
	if err != nil {
		return err
	}
	_, err = c.minioClient.PutObject(context.Background(), bucketName, fileName, file, fileStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return err
	}
	return nil
}
