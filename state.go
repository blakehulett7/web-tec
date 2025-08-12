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
	IsFueled       bool
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

func (app *State) Tick() {
	for _, event := range app.FuelSystem {
		if !event.IsFueled {
			is_paid := app.IData.Pay(event.Cost)
			if is_paid {
				event.IsFueled = true
			}
		}

		app.E += event.EperTick
		event.EperTick -= 1
	}

	for _, event := range app.MutationSystem {
		app.IData.Pay(event.Remove)
		app.IData.Add(event.Add)
	}
}
