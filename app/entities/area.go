package entities

import (
	"time"

	schm "github.com/kaneshin/snacky-go/app/entities/schemata"
)

type Area struct {
	*Base
	schm.Area
}

func NewArea(id int32, name string) *Area {
	area := Area{
		Base: NewBase(),
	}
	area.InitBase()
	area.Id = id
	area.Name = name
	area.CreatedAt = time.Now().Unix()
	area.UpdatedAt = time.Now().Unix()
	return &area
}
