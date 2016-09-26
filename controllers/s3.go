package controllers

import (
        "github.com/astaxie/beego"
        "github.com/lenfree/awsLaCapa/models"
)

// Category Object
type S3Controller struct {
        beego.Controller
}

// @Title GetAll
// @Description Return all S3 Buckets
// @Success 200 {object} models.Buckets
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

// @Title GetByBucketName
// @Description Return S3 Bucket Directories/Prefix
// @Success 200
// @router /:bucket [get]
func (c *S3Controller) GetByPrefix() {
        bucket := c.GetString(":bucket")
        object, err := models.S3ListObjects(bucket)
        if err != nil {
                c.Data["json"] = err.Error()
        } else {
                c.Data["json"] = object
        }
        c.ServeJSON()
}

// @Title GetByObjectKey
// @Description Return S3 Bucket Object
// @Success 200
// @router /:bucket/:directory/:objkey [get]
func (c *S3Controller) GetObjectByKey() {
        bucket := c.GetString(":bucket")
        directory := c.GetString(":directory")
        objectKey := c.GetString(":objkey")
        object, err := models.S3GetObjectByKey(bucket, directory+"/"+objectKey)
        if err != nil {
                c.Data["json"] = err.Error()
        } else {
                c.Data["json"] = object
        }
        c.ServeJSON()
}
