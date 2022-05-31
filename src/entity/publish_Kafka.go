package entity

import "time"

type PublishMessageKafkaPeople struct {
	ID       string
	Date     time.Time
	FIleName string
	People   People
}
