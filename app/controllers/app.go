package controllers

import (
	"fmt"

	m "github.com/kaneshin/snacky-go/app/models"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	snackModel := m.NewSnack()
	fmt.Println(snackModel.Create(10, "Umai!"))
	return c.Render()
}
