package controllers

import (
	t "github.com/kaneshin/snacky-go/app/models/toriko"
	"github.com/revel/revel"
)

type Toriko struct {
	Controller
}

func (c Toriko) UpdateSnacksByKind(kind int, limit int, offset int) revel.Result {
	toriko := t.NewToriko()
	err := toriko.UpdateSnacksByKind(kind, limit, offset)
	res := c.ResultJson(err)
	return c.RenderJson(res)
}

func (c Toriko) UpdateSnacksByYear(year int, limit int, offset int) revel.Result {
	toriko := t.NewToriko()
	err := toriko.UpdateSnacksByYear(year, limit, offset)
	res := c.ResultJson(err)
	return c.RenderJson(res)
}

func (c Toriko) GetSnacksByKind(kind int, limit int, offset int) revel.Result {
	toriko := t.NewToriko()
	items, err := toriko.GetSnacksByKind(kind, limit, offset)
	res := c.ResultJson(err)
	res["items"] = items
	return c.RenderJson(res)
}

func (c Toriko) GetSnacksByYear(year int, limit int, offset int) revel.Result {
	toriko := t.NewToriko()
	items, err := toriko.GetSnacksByYear(year, limit, offset)
	res := c.ResultJson(err)
	res["items"] = items
	return c.RenderJson(res)
}

func (c Toriko) UpdateKind() revel.Result {
	toriko := t.NewToriko()
	err := toriko.UpdateKindMaster()
	res := c.ResultJson(err)
	return c.RenderJson(res)
}

func (c Toriko) UpdateArea() revel.Result {
	toriko := t.NewToriko()
	err := toriko.UpdateAreaMaster()
	res := c.ResultJson(err)
	return c.RenderJson(res)
}

func (c Toriko) UpdateMaker() revel.Result {
	toriko := t.NewToriko()
	err := toriko.UpdateMakerMaster()
	res := c.ResultJson(err)
	return c.RenderJson(res)
}
