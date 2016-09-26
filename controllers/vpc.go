package controllers

import (
        "github.com/astaxie/beego"
        "github.com/lenfree/awsLaCapa/models"
)

type VPCController struct {
        beego.Controller
}

// @Title GetAll
// @Description Return all VPCs
// @Success 200 {object} models.VPCs
// @router /vpcs [get]
func (c *VPCController) GetAll() {
        VPCs, err := models.VPCList()
        if err != nil {
                c.Data["json"] = err.Error()
        } else {
                c.Data["json"] = VPCs
        }
        c.ServeJSON()
}

// @Title GetAllVPCPeering
// @Description Return all VPC Peering Connections
// @Success 200 {object} models.VPCPeeringConnections
// @router /vpc_peering_connections [get]
func (c *VPCController) GetAllVPCPeering() {
        VPCPeeringConnections, err := models.VPCPeeringList()
        if err != nil {
                c.Data["json"] = err.Error()
        } else {
                c.Data["json"] = VPCPeeringConnections
        }
        c.ServeJSON()
}
