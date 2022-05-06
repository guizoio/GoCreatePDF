package minio_client

import (
	"CreateFilePDF/src/entity"
	"context"
	"github.com/minio/minio-go/v7"
	"io"
	"os"
)

type FaceClientMinio interface {
	CheckLife() string
	ListBuckets() ([]*entity.BucketInfo, error)
	UploadObject(bucketName, fileName string) error
	ListBucketObjects(bucket string) ([]*entity.ObjectInfo, error)
	DownloadObject(bucket, fileName string) error
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

func (c ClientMinio) ListBuckets() ([]*entity.BucketInfo, error) {
	buckets, err := c.minioClient.ListBuckets(context.Background())
	if err != nil {
		return nil, err
	}
	var list []*entity.BucketInfo
	for _, bucket := range buckets {
		data := entity.BucketInfo{
			Name:         bucket.Name,
			CreationDate: bucket.CreationDate,
		}
		list = append(list, data.ToDomain())
	}
	return list, nil
}

func (c ClientMinio) ListBucketObjects(bucket string) ([]*entity.ObjectInfo, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	objectCh := c.minioClient.ListObjects(ctx, bucket, minio.ListObjectsOptions{
		Recursive: true,
	})

	var list []*entity.ObjectInfo
	for object := range objectCh {
		if object.Err != nil {
			return nil, object.Err
		}
		objects := entity.ObjectInfo{
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
	_, err = c.minioClient.PutObject(
		context.Background(),
		bucketName,
		fileName,
		file,
		fileStat.Size(),
		minio.PutObjectOptions{ContentType: "application/octet-stream"},
	)

	if err != nil {
		return err
	}
	return nil
}

func (c ClientMinio) DownloadObject(bucket, fileName string) error {
	object, err := c.minioClient.GetObject(context.Background(), bucket, fileName, minio.GetObjectOptions{})
	if err != nil {
		return err
	}
	localFile, err := os.Create("./" + fileName)
	if err != nil {
		return err
	}
	if _, err = io.Copy(localFile, object); err != nil {
		return err
	}
	return nil
}
