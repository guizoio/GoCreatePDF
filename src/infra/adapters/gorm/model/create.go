package model

import "gorm.io/datatypes"

type Create struct {
	ID       string         `gorm:"column:id"`
	Name     string         `gorm:"column:name"`
	Content  datatypes.JSON `gorm:"column:content"`
	FileName string         `gorm:"column:file_name"`
}

func (c *Create) TableName() string {
	return "create"
}
