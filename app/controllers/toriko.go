package controllers

import (
	t "github.com/kaneshin/snacky-go/app/models/toriko"
	"github.com/revel/revel"
)

type Toriko struct {
	Controller
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
