package controllers

import "github.com/revel/revel"

type Controller struct {
	*revel.Controller
}

func (c Controller) ResultJson(err error) map[string]interface{} {
	res := map[string]interface{}{}
	if err != nil {
		res["status"] = "ng"
		res["error"] = err
	} else {
		res["status"] = "ok"
	}
	return res
}
