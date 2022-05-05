package minio_client

import (
	"context"
	"github.com/minio/minio-go/v7"
	"os"
)

type FaceClientMinio interface {
	CheckLife() string
	ListBuckets() ([]minio.BucketInfo, error)
	UploadObject(bucketName, fileName string) error
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
	//for _, bucket := range buckets {
	//	fmt.Println(bucket.Name + " - " + bucket.CreationDate.String())
	//}
	return buckets, nil
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
