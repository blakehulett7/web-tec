package main

import "github.com/google/uuid"

type State struct {
	InstanceId uuid.UUID
	Energy     int
	Items      map[string]int
	Collectors []Collector
}

func init_state(id uuid.UUID) State {
	return State{
		InstanceId: id,
		Energy:     0,
		Items:      map[string]int{},
		Collectors: []Collector{},
	}
}
