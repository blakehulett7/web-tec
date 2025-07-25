package main

import "github.com/google/uuid"

type State struct {
	Id         uuid.UUID
	Energy     int
	Items      Items       `gorm:"embedded"`
	Collectors []Collector `gorm:"foreignKey:UserId"`
}
