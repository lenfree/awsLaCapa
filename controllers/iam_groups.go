package controllers

import (
        "github.com/astaxie/beego"
        "github.com/lenfree/awsLaCapa/models"
)

type IAMGroupsController struct {
        beego.Controller
}

// @Title GetAll
// @Description Return all IAM groups
// @Success 200 {object} models.IAMGroups
// @router /groups [get]
func (c *IAMGroupsController) GetAllIAMGroups() {
        IAMGroups, err := models.IAMGroupList()
        if err != nil {
                c.Data["json"] = err.Error()
        } else {
                c.Data["json"] = IAMGroups
        }
        c.ServeJSON()
}
