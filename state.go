package main

import "github.com/google/uuid"

type State struct {
	Id              uuid.UUID
	E               int
	IData           []IMat `gorm:"foreignKey:UserId"`
	FuelSystem      []FuelEvent
	IMutationSystem []MutationEvent
}

func NewState() State {
	id := uuid.New()

	IList := []IMat{}
	for _, i := range IMap {
		IList = append(IList, IMat{
			UserId: id,
			Name:   i,
			Count:  0,
		})
	}

	return State{
		Id:    id,
		E:     0,
		IData: IList,
	}
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
