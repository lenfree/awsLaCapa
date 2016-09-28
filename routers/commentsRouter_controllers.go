package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/lenfree/awsLaCapa/controllers:DHCPController"] = append(beego.GlobalControllerRouter["github.com/lenfree/awsLaCapa/controllers:DHCPController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lenfree/awsLaCapa/controllers:EC2Controller"] = append(beego.GlobalControllerRouter["github.com/lenfree/awsLaCapa/controllers:EC2Controller"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lenfree/awsLaCapa/controllers:IAMGroupsController"] = append(beego.GlobalControllerRouter["github.com/lenfree/awsLaCapa/controllers:IAMGroupsController"],
		beego.ControllerComments{
			Method: "GetAllIAMGroups",
			Router: `/groups`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lenfree/awsLaCapa/controllers:IAMUsersController"] = append(beego.GlobalControllerRouter["github.com/lenfree/awsLaCapa/controllers:IAMUsersController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/users`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lenfree/awsLaCapa/controllers:Route53Controller"] = append(beego.GlobalControllerRouter["github.com/lenfree/awsLaCapa/controllers:Route53Controller"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lenfree/awsLaCapa/controllers:Route53Controller"] = append(beego.GlobalControllerRouter["github.com/lenfree/awsLaCapa/controllers:Route53Controller"],
		beego.ControllerComments{
			Method: "GetHostedZoneRRSet",
			Router: `/hostedzone/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lenfree/awsLaCapa/controllers:S3Controller"] = append(beego.GlobalControllerRouter["github.com/lenfree/awsLaCapa/controllers:S3Controller"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lenfree/awsLaCapa/controllers:S3Controller"] = append(beego.GlobalControllerRouter["github.com/lenfree/awsLaCapa/controllers:S3Controller"],
		beego.ControllerComments{
			Method: "GetByPrefix",
			Router: `/:bucket`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lenfree/awsLaCapa/controllers:S3Controller"] = append(beego.GlobalControllerRouter["github.com/lenfree/awsLaCapa/controllers:S3Controller"],
		beego.ControllerComments{
			Method: "GetObjectByKey",
			Router: `/:bucket/:directory/:objkey`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lenfree/awsLaCapa/controllers:SecurityGroupController"] = append(beego.GlobalControllerRouter["github.com/lenfree/awsLaCapa/controllers:SecurityGroupController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lenfree/awsLaCapa/controllers:VPCController"] = append(beego.GlobalControllerRouter["github.com/lenfree/awsLaCapa/controllers:VPCController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lenfree/awsLaCapa/controllers:VPCController"] = append(beego.GlobalControllerRouter["github.com/lenfree/awsLaCapa/controllers:VPCController"],
		beego.ControllerComments{
			Method: "GetAllVPCPeering",
			Router: `/vpc_peering_connections`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
