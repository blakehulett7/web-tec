package main

import "github.com/google/uuid"

type Collector struct {
	UserId uuid.UUID

	Name string
	Item string

	IsOn        bool
	IsFueled    bool
	TickCounter int

	MaterialCost IMat `gorm:"embedded"`

	EnergyPerBatch int
	ItemsPerBatch  int
	TicksPerBatch  int
}
