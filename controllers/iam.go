package controllers

import (
    "github.com/astaxie/beego"
    "github.com/lenfree/awsRestWrapper/models"
)

type IAMController struct {
    beego.Controller
}

// @Title GetAll
// @Description get all S3 Buckets
// @Success 200 {object} models.Bucket
// @router / [get]
func (c *IAMController) GetAll() {
    IAMUsers, err := models.IAMUserList()
    if err != nil {
        c.Data["json"] = err.Error()
    } else {
        c.Data["json"] = IAMUsers
    }
    c.ServeJSON()
}
