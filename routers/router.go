// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
        "github.com/lenfree/awsLaCapa/controllers"

        "github.com/astaxie/beego"
)

func init() {
        ns := beego.NewNamespace("/v1",
                beego.NSNamespace("/s3",
                        beego.NSInclude(
                                &controllers.S3Controller{},
                        ),
                ),
                beego.NSNamespace("/iam",
                        beego.NSInclude(
                                &controllers.IAMController{},
                        ),
                ),
                beego.NSNamespace("/vpc",
                        beego.NSInclude(
                                &controllers.VPCController{},
                        ),
                ),
                beego.NSNamespace("/route53",
                        beego.NSInclude(
                                &controllers.Route53Controller{},
                        ),
                ),
                beego.NSNamespace("/ec2",
                        beego.NSInclude(
                                &controllers.EC2Controller{},
                        ),
                ),
        )
        beego.AddNamespace(ns)
}
