package main

import "github.com/google/uuid"

type IMat struct {
	UserId uuid.UUID
	Name   string
	Count  int
}
