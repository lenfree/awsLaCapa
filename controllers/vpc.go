package controllers

import (
    "github.com/astaxie/beego"
    "github.com/lenfree/awsRestWrapper/models"
)

type VPCController struct {
    beego.Controller
}

// @Title GetAll
// @Description Return all VPCs
// @Success 200 {object} models.VPCs
// @router / [get]
func (c *VPCController) GetAll() {
    VPCs, err := models.VPCList()
    if err != nil {
        c.Data["json"] = err.Error()
    } else {
        c.Data["json"] = VPCs
    }
    c.ServeJSON()
}
