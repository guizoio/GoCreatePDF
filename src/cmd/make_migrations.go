package cmd

import (
	"CreateFilePDF/src/infra/adapters/gorm/model"
	"gorm.io/gorm"
)

type DatabaseMakeMigrations struct {
	database *gorm.DB
}

func NewDatabaseMakeMigrations(database *gorm.DB) *DatabaseMakeMigrations {
	return &DatabaseMakeMigrations{database}
}

func (d *DatabaseMakeMigrations) MakeMigrations() {
	err := d.database.Migrator().AutoMigrate(&model.Create{})
	if err != nil {
		panic("Error Migrating Create Table")
	}
}
