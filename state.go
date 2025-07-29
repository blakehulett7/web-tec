package main

import "github.com/google/uuid"

type State struct {
	Id         uuid.UUID
	E          int
	IMap       Items       `gorm:"embedded"`
	Collectors []Collector `gorm:"foreignKey:UserId"`
}
