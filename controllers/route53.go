package controllers

import (
        "github.com/astaxie/beego"
        "github.com/lenfree/awsLaCapa/models"
)

type Route53Controller struct {
        beego.Controller
}

// @Title GetAll
// @Description Return all Route53 Hosted Zones
// @Success 200 {object} models.Route53HostedZones
// @router / [get]
func (c *Route53Controller) GetAll() {
        hostedZones, err := models.Route53HostedZoneList()
        if err != nil {
                c.Data["json"] = err.Error()
        } else {
                c.Data["json"] = hostedZones
        }
        c.ServeJSON()
}

// @Title Get
// @Description Return all Route53 Hosted Zone Resource Record Set
// @Success 200 {object} models.Route53HostedZoneRRSet
// @router /hostedzone/:id [get]
func (c *Route53Controller) GetHostedZoneRRSet() {
        hostedZoneID := c.GetString(":id")
        hostedZoneRRSet, err := models.Route53HostedZoneRRSetByID("/hostedzone/" + hostedZoneID)
        if err != nil {
                c.Data["json"] = err.Error()
        } else {
                c.Data["json"] = hostedZoneRRSet
        }
        c.ServeJSON()
}
