package entity

import (
	"time"
)

type BucketInfo struct {
	Name         string    `json:"name"`
	CreationDate time.Time `json:"creationDate"`
}

func (ref *BucketInfo) ToDomain() *BucketInfo {
	return &BucketInfo{
		Name:         ref.Name,
		CreationDate: ref.CreationDate,
	}
}

type ObjectIndo struct {
	Name string
	Date time.Time
	Size int64
}

func (ref *ObjectIndo) ToDomain() *ObjectIndo {
	return &ObjectIndo{
		Name: ref.Name,
		Date: ref.Date,
		Size: ref.Size,
	}
}