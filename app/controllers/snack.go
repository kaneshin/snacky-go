package controllers

import "github.com/revel/revel"

type Snack struct {
	Controller
}

func (c Snack) Get() revel.Result {
	return c.Render()
}

func (c Snack) GetById() revel.Result {
	return c.Render()
}

func (c Snack) List() revel.Result {
	return c.Render()
}
