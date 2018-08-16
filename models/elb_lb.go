package models

import (
	"github.com/astaxie/beego"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elb"
)

/* This is a hack.
At the moment, I couldn't figure out a way to use type alias to generate
object model. Hence, use a model to describe iam.ListServerCertificatesOutput and
not being used anywhere else except for Swagger API documentation.Furthermore,
there's a known limitation with non built in data types such as time.Time
*/

func ELBLoadBalancerList() (*elb.DescribeLoadBalancersOutput, error) {
	svc := elb.New(session.New(), &aws.Config{
		Region: aws.String(beego.AppConfig.String("awsRegion")),
	})

	params := &elb.DescribeLoadBalancersInput{}

	resp, err := svc.DescribeLoadBalancers(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			beego.Error(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				beego.Error(
					reqErr.Code(),
					reqErr.Message(),
					reqErr.StatusCode(),
					reqErr.RequestID(),
				)
			}
		} else {
			beego.Debug(err.Error())
		}
		return nil, err
	}
	return resp, nil
}
