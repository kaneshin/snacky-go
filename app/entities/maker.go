package entities

import (
	"time"

	schm "github.com/kaneshin/snacky-go/app/entities/schemata"
)

type Maker struct {
	*Base
	schm.Maker
}

func NewMaker(id int32, name string) *Maker {
	maker := Maker{
		Base: NewBase(),
	}
	maker.InitBase()
	maker.Id = id
	maker.Name = name
	maker.CreatedAt = time.Now()
	maker.UpdatedAt = time.Now()
	return &maker
}

func (e *Maker) Scheme() *schm.Maker {
	return &e.Maker
}
