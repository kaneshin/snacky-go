package controllers

import (
	m "github.com/kaneshin/snacky-go/app/models"
	"github.com/revel/revel"
)

type App struct {
	Controller
}

func (c App) Index() revel.Result {
	snackModel := m.NewSnack()
	snackModel.Create(11, "UmaiZO!")
	// m := m.NewToriko()
	// fmt.Println(m.GetSnackByKind(1, 3, 0))
	// fmt.Println(m.GetSnackByYear(2014, 3, 0))
	return c.Render()
}
