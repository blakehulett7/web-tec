package main

import "github.com/google/uuid"

type State struct {
	Id    uuid.UUID
	E     int
	IData ISet `gorm:"serializer:json"`

	FuelSystem     []FuelEvent
	MutationSystem []MutationEvent
}

type FuelEvent struct {
	StateId        uuid.UUID
	EntityId       uuid.UUID
	Cost           ISet `gorm:"serializer:json"`
	TicksRemaining int
	EperTick       int
}

type MutationEvent struct {
	StateId  uuid.UUID
	EntityId uuid.UUID
	Remove   ISet `gorm:"serializer:json"`
	Add      ISet `gorm:"serializer:json"`
}
