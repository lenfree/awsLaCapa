package controllers

import (
        "github.com/astaxie/beego"
        "github.com/lenfree/awsLaCapa/models"
)

type IAMServerCertsController struct {
        beego.Controller
}

// @Title GetServerCerts
// @Description Return all IAM Server Certificates
// @Success 200 {object} models.IAMServerCertificateMetadataList
// @router /server_certificates [get]
func (c *IAMUsersController) GetServerCerts() {
        IAMServerCerts, err := models.IAMServerCertificteList()
        if err != nil {
                c.Data["json"] = err.Error()
        } else {
                c.Data["json"] = IAMServerCerts
        }
        c.ServeJSON()
}
