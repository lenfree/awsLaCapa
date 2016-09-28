package controllers

import (
        "github.com/astaxie/beego"
        "github.com/lenfree/awsLaCapa/models"
)

type IAMUserGroupsController struct {
        beego.Controller
}

// @Title GetAll User Groups
// @Description Return all IAM users in specified IAM User Groups
// @Success 200 {object} models.IAMUserGroups
// @router /user_groups [get]
func (c *IAMUserGroupsController) GetAll() {
        IAMUserGroups, err := models.IAMUserGroupList()
        if err != nil {
                c.Data["json"] = err.Error()
        } else {
                c.Data["json"] = IAMUserGroups
        }
        c.ServeJSON()
}
