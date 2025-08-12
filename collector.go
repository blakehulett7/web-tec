package main

import "github.com/google/uuid"

func NewCM(app *State) {
	entity_id := uuid.New()
	add_set := ISet{}
	add_set[C] = 1

	mutation_event := MutationEvent{
		StateId:      app.Id,
		EntityId:     entity_id,
		ECostperTick: 1,
		Add:          add_set,
	}

	app.MutationSystem = append(app.MutationSystem, mutation_event)
}

func NewCBurner(app *State) {
	entity_id := uuid.New()
	cost := ISet{}
	cost[C] = 1

	fuel_event := FuelEvent{
		StateId:        app.Id,
		EntityId:       entity_id,
		IsFueled:       false,
		Cost:           cost,
		TicksRemaining: 10,
		EperTick:       1,
	}
	app.FuelSystem = append(app.FuelSystem)
}
