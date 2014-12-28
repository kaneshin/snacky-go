package entities

import (
	"time"

	schm "github.com/kaneshin/snacky-go/app/entities/schemata"
)

type Snack struct {
	*Base
	schm.Snack
}

func NewSnack(id int32, name string) *Snack {
	snack := Snack{
		Base: NewBase(),
	}
	snack.InitBase()
	snack.Id = id
	snack.Name = name
	snack.CreatedAt = time.Now()
	snack.UpdatedAt = time.Now()
	return &snack
}

func (self *Snack) Scheme() *schm.Snack {
	return &self.Snack
}
