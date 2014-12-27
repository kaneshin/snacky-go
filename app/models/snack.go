package models

import (
	"github.com/kaneshin/snacky-go/app"
	"github.com/kaneshin/snacky-go/app/entities"
)

type Snack struct {
	*Model
}

func NewSnack() *Snack {
	instance := &Snack{}
	instance.InitModel()
	return instance
}

func (self *Snack) Create(id int32, name string) error {
	snack := entities.NewSnack(id, name)
	return app.Dbd.Dbm.Insert(snack.Scheme())
}
