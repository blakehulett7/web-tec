package main

import "github.com/google/uuid"

type State struct {
	Id    uuid.UUID
	E     int
	IData ISet

	FuelSystem      []FuelEvent
	IMutationSystem []MutationEvent
}

type FuelEvent struct {
	Id       uuid.UUID
	Cost     ISet
	Duration int
	EperT    int
}

type MutationEvent struct {
	Id     uuid.UUID
	Remove ISet
	Add    ISet
}
