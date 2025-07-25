package main

import "github.com/google/uuid"

type State struct {
	Id         uuid.UUID
	Energy     int
	Items      map[string]int
	Collectors []Collector
}

func NewState(id uuid.UUID) State {
	return State{
		Id:         id,
		Energy:     0,
		Items:      map[string]int{},
		Collectors: []Collector{},
	}
}
