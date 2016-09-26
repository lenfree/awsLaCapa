package controllers

import (
        "github.com/astaxie/beego"
        "github.com/lenfree/awsLaCapa/models"
)

type EC2Controller struct {
        beego.Controller
}

// @Title GetAll
// @Description Return all EC2 instances
// @Success 200 {object} models.EC2Reservations
// @router / [get]
func (c *EC2Controller) GetAll() {
        EC2Instances, err := models.EC2InstanceList()
        if err != nil {
                c.Data["json"] = err.Error()
        } else {
                c.Data["json"] = EC2Instances
        }
        c.ServeJSON()
}
