package main

type Imat int

const (
	Wd Imat = iota
	Stn
	Fe
	C
	Cu
	Sn
	Ni
)

type ISet [7]int
