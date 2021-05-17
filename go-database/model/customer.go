package model

import "time"

type Customer struct {
	Id        int32
	Name      string
	Email     string
	Balance   int64
	Rating    float64
	BirthDate time.Time
	Married   bool
	CreatedAt time.Time
}
