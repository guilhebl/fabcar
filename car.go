package main

import (
	"github.com/google/uuid"
	"time"
)

// Define the car structure.  Structure tags are used by encoding/json library
type Car struct {
	Id      string    `json:"id"`
	Make    string    `json:"make"`
	Model   string    `json:"model"`
	Colour  string    `json:"colour"`
	Owner   string    `json:"owner"`
	Created time.Time `json:"created"`
	Parts   []CarPart `json:"parts"`
}

// A car part is part of a car and
type CarPart struct {
	Id       string    `json:"id"`
	ParentId string    `json:"parentId"`
	Model    string    `json:"model"`
	Created  time.Time `json:"created"`
}

func NewCar(make, model, color, owner string, parts ...string) *Car {
	id, _ := uuid.NewRandom()
	idStr := id.String()

	p := make([]CarPart, 0)
	for i := 0; i < len(parts); i++ {
		part := NewCarPart(parts[i], idStr)
		p = append(p, *part)
	}

	return &Car{
		Id:      idStr,
		Make:    make,
		Model:   model,
		Colour:  color,
		Owner:   owner,
		Created: time.Now(),
		Parts:   p,
	}
}

func NewCarPart(model, parentId string) *CarPart {
	id, _ := uuid.NewRandom()

	return &CarPart{
		Id:       id.String(),
		Model:    model,
		ParentId: parentId,
		Created:  time.Now(),
	}
}
