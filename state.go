package main

import "github.com/google/uuid"

type State struct {
	Id         uuid.UUID
	E          int
	IData      []IMat      `gorm:"foreignKey:UserId"`
	Collectors []Collector `gorm:"foreignKey:UserId"`
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
		Id:         id,
		E:          0,
		IData:      IList,
		Collectors: []Collector{},
	}
}
