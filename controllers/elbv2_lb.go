package controllers

import (
        "github.com/astaxie/beego"
        "github.com/lenfree/awsLaCapa/models"
)

type ELBv2Controller struct {
        beego.Controller
}

// @Title Get v2 ELBs
// @Description Return all ELBv2 LBs
// @Success 200 {object} models.ELBv2LoadBalancers
// @router / [get]
func (c *ELBv2Controller) GetELBs() {
        elbs, err := models.ELBv2LoadBalancerList()
        if err != nil {
                c.Data["json"] = err.Error()
        } else {
                c.Data["json"] = elbs
        }
        c.ServeJSON()
}
