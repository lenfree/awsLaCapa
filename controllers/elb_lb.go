package controllers

import (
        "github.com/astaxie/beego"
        "github.com/lenfree/awsLaCapa/models"
)

type ELBController struct {
        beego.Controller
}

// @Title Get v1 ELBs
// @Description Return all ELBv1 LBs
// @Success 200 {object} models.ELBLoadBalancers
// @router / [get]
func (c *ELBController) GetELBs() {
        elbs, err := models.ELBLoadBalancerList()
        if err != nil {
                c.Data["json"] = err.Error()
        } else {
                c.Data["json"] = elbs
        }
        c.ServeJSON()
}
