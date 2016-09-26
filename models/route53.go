package models

import (
        "github.com/astaxie/beego"
        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go/service/route53"
        "time"
)

/* At the moment, I couldn't figure out a way to use type alias to generate
object model. Hence, use a model to describe route53.ListHostedZonesOUtput and not being
used anywhere else except for Swagger API documentation
*/
type Route53HostedZone struct {
        CallerReference time.Time `json:"CallerReference"`
        Config          struct {
                Comment     string `json:"Comment"`
                PrivateZone bool   `json:"PrivateZone"`
        } `json:"Config"`
        ID                     string `json:"Id"`
        Name                   string `json:"Name"`
        ResourceRecordSetCount int    `json:"ResourceRecordSetCount"`
}

type Route53HostedZones struct {
        Route53HostedZones []Route53HostedZone `json:"Route53_hosted_zones"`
        IsTruncated        bool                `json:"IsTruncated"`
        Marker             interface{}         `json:"Marker"`
        MaxItems           string              `json:"MaxItems"`
        NextMarker         interface{}         `json:"NextMarker"`
}

type Route53HostedZoneRRSet struct {
        IsTruncated          bool        `json:"IsTruncated"`
        MaxItems             string      `json:"MaxItems"`
        NextRecordIdentifier interface{} `json:"NextRecordIdentifier"`
        NextRecordName       interface{} `json:"NextRecordName"`
        NextRecordType       interface{} `json:"NextRecordType"`
        ResourceRecordSets   []struct {
                AliasTarget     interface{} `json:"AliasTarget"`
                Failover        interface{} `json:"Failover"`
                GeoLocation     interface{} `json:"GeoLocation"`
                HealthCheckID   interface{} `json:"HealthCheckId"`
                Name            string      `json:"Name"`
                Region          interface{} `json:"Region"`
                ResourceRecords []struct {
                        Value string `json:"Value"`
                } `json:"ResourceRecords"`
                SetIdentifier           interface{} `json:"SetIdentifier"`
                TTL                     int         `json:"TTL"`
                TrafficPolicyInstanceID interface{} `json:"TrafficPolicyInstanceId"`
                Type                    string      `json:"Type"`
                Weight                  interface{} `json:"Weight"`
        } `json:"ResourceRecordSets"`
}

func Route53HostedZoneList() (*route53.ListHostedZonesOutput, error) {
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

func Route53HostedZoneRRSetByID(id string) (*route53.ListResourceRecordSetsOutput, error) {
        svc := route53.New(session.New(), &aws.Config{
                Region: aws.String("ap-southeast-2"),
        })

        params := &route53.ListResourceRecordSetsInput{
                HostedZoneId: aws.String(id),
        }

        resp, err := svc.ListResourceRecordSets(params)

        if err != nil {
                beego.Error(err)
                return nil, err
        }
        return resp, nil
}
