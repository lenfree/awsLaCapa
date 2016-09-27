package controllers

import (
        "github.com/astaxie/beego"
        "github.com/lenfree/awsLaCapa/models"
)

type SecurityGroupController struct {
        beego.Controller
}

// @Title GetAll
// @Description Return all Security Group Set
// @Success 200 {object} models.SecurityGroups
// @router / [get]
func (c *SecurityGroupController) GetAll() {
        sgs, err := models.SecurityGroupsList()
        if err != nil {
                c.Data["json"] = err.Error()
        } else {
                c.Data["json"] = sgs
        }
        c.ServeJSON()
}
