package model

type Location struct {
	Id int
	Name string
	Slug string
	Locations *Location
}
