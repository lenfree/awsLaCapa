package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/lenfree/awsLaCapa/controllers:IAMController"] = append(beego.GlobalControllerRouter["github.com/lenfree/awsLaCapa/controllers:IAMController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
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

	beego.GlobalControllerRouter["github.com/lenfree/awsLaCapa/controllers:VPCController"] = append(beego.GlobalControllerRouter["github.com/lenfree/awsLaCapa/controllers:VPCController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
