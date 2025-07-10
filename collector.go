package main

type Collector struct {
	Name string
	Item string

	IsOn        bool
	IsFueled    bool
	TickCounter int

	MaterialCost map[string]int

	EnergyPerBatch int
	ItemsPerBatch  int
	TicksPerBatch  int
}
