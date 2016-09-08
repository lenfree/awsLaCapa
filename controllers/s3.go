package controllers

import (
    "github.com/astaxie/beego"
    "github.com/lenfree/awsRestWrapper/models"
)

// Category Object
type S3Controller struct {
    beego.Controller
}

// @Title GetAll
// @Description get all S3 Buckets
// @Success 200 {object} models.Bucket
// @router / [get]
func (c *S3Controller) GetAll() {
    buckets, err := models.S3List()
    if err != nil {
        c.Data["json"] = err.Error()
    } else {
        c.Data["json"] = buckets
    }
    c.ServeJSON()
}
