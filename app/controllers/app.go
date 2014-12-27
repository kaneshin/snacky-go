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
	// snackModel := m.NewSnack()
	// snackModel.Create(10, "Umai!")
	m := m.NewToriko()
	fmt.Println(m.GetKindMaster())
	fmt.Println(m.GetAreaMaster())
	fmt.Println(m.GetMakerMaster())
	fmt.Println(m.GetSnackByKind(1, 3, 0))
	fmt.Println(m.GetSnackByYear(2014, 3, 0))
	return c.Render()
}
