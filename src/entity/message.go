package entity

import (
	"gorm.io/datatypes"
	"time"
)

type Test struct {
	ID      string
	Name    string
	Content datatypes.JSON
}

type HeaderPDF struct {
	FilePDF int64
	FileIMG string
}

type PublishMessageKafkaPeople struct {
	ID       string
	Date     time.Time
	FIleName string
	People   People
}

type People struct {
	Name      string
	CPF       string
	RG        string
	BirthDate string
	Address   Address
	Contact   Contact
}

type Company struct {
	Name              string
	CNPJ              string
	StateRegistration string
	OpeningDate       string
	Site              string
	Address           Address
	Contact           Contact
}

type Contact struct {
	Email     string
	Cell      string
	Telephone string
}

type Address struct {
	CodePostal string
	Address    string
	Number     string
	District   string
	City       string
	State      string
}
