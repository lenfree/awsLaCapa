package controllers

import (
        "github.com/astaxie/beego"
        "github.com/lenfree/awsLaCapa/models"
)

type DHCPController struct {
        beego.Controller
}

// @Title GetAll
// @Description Return all DHCP Options
// @Success 200 {object} models.DHCPOptionSet
// @router / [get]
func (c *DHCPController) GetAll() {
        DHCPs, err := models.DHCPOptionsList()
        if err != nil {
                c.Data["json"] = err.Error()
        } else {
                c.Data["json"] = DHCPs
        }
        c.ServeJSON()
}
