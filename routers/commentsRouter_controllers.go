package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/lenfree/awsRestWrapper/controllers:IAMController"] = append(beego.GlobalControllerRouter["github.com/lenfree/awsRestWrapper/controllers:IAMController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lenfree/awsRestWrapper/controllers:S3Controller"] = append(beego.GlobalControllerRouter["github.com/lenfree/awsRestWrapper/controllers:S3Controller"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
