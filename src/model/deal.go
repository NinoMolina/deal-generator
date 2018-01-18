package model

import "DealGenerator/src/util"

type Deal struct {
	ID       int
	Title    string
	Image    string
	Price    int
	Category string
	Location string
	Merchant string
}

func (p Deal) ToString() string {
	return util.ToJson(p)
}
