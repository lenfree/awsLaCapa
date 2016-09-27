package models

import (
        "github.com/astaxie/beego"
        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/awserr"
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go/service/ec2"
)

/* This is a hack.
At the moment, I couldn't figure out a way to use type alias to generate
object model. Hence, use a model to describe ec2.DescribeDhcpOptions and
not being used anywhere else except for Swagger API documentation.
*/

type DHCPConfiguration struct {
        Key string       `json:"Key"`
        Values []struct {
            Value string `json:"Value"`
        } `json:"Values"`
}

type DHCPOption struct {
        DhcpConfigurations []DHCPConfiguration `json:"DhcpOption"`
}

type DHCPOptionSet struct {
        DhcpOptions []DHCPOption `json:"DhcpOptions"`
        DhcpOptionsID string     `json:"DhcpOptionsId"`
        Tags interface{}         `json:"Tags"`
}

func DHCPOptionsList() (*ec2.DescribeDhcpOptionsOutput, error) {
        svc := ec2.New(session.New(), &aws.Config{
                Region: aws.String("ap-southeast-2"),
        })

        params := &ec2.DescribeDhcpOptionsInput{}

        resp, err := svc.DescribeDhcpOptions(params)
beego.Info(resp)
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
