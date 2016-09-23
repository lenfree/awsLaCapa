package models

import (
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go/aws"
        "github.com/astaxie/beego"
        "github.com/aws/aws-sdk-go/service/route53"
        "time"
)

/* At the moment, I couldn't figure out a way to use type alias to generate
 object model. Hence, use a model to describe ec2.DescribeVpcs and not being
 used anywhere else except for Swagger API documentation
 */
type Route53HostedZone struct {
        CallerReference time.Time `json:"CallerReference"`
        Config struct {
            Comment string `json:"Comment"`
            PrivateZone bool `json:"PrivateZone"`
        } `json:"Config"`
        ID string `json:"Id"`
        Name string `json:"Name"`
        ResourceRecordSetCount int `json:"ResourceRecordSetCount"`
}

type Route53HostedZones struct {
        Route53HostedZones            []Route53HostedZone `json:route53_hosted_zones`
        IsTruncated            bool                `json:"IsTruncated"`
        Marker                 interface{}         `json:"Marker"`
        MaxItems               string              `json:"MaxItems"`
        NextMarker             interface{}         `json:"NextMarker"`
}

func Route53HostedZoneList() (*route53.ListHostedZonesOutput , error) {
        svc := route53.New(session.New(), &aws.Config{ 
                Region: aws.String("ap-southeast-2"),
        })

        params := &route53.ListHostedZonesInput{}

        resp, err := svc.ListHostedZones(params)

        if err != nil {
                beego.Error(err)
                return nil, err
        }
        return resp, nil
}
