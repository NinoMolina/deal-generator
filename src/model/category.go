package model

type Category struct {
	Id           int
	Uuid         string
	ParentId     int
	Name         string
	BusinessUnit string
	Parent       *Category
}
