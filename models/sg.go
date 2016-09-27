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
object model. Hence, use a model to describe ec2.DescribeSecurityGroups and
not being used anywhere else except for Swagger API documentation.
*/
type SecurityGroup struct {
        Description string `json:"Description"`
        GroupID string `json:"GroupId"`
        GroupName string `json:"GroupName"`
        SecurityGroupIPPermission []SecurityGroupIPPermission `json:"IPPermissions"`
        SecurityGroupIPPermissionEgress []SecurityGroupIPPermissionEgress `json:"IPPermissionsEgress"`
        OwnerID string `json:"OwnerId"`
        Tags interface{} `json:"Tags"`
        VpcID string `json:"VpcId"`
}

type SecurityGroupIPPermission struct {
        FromPort interface{} `json:"FromPort"`
        IPProtocol string `json:"IpProtocol"`
        IPRanges interface{} `json:"IpRanges"`
        PrefixListIds interface{} `json:"PrefixListIds"`
        ToPort interface{} `json:"ToPort"`
        UserIDGroupPairs []SecurityGroupUserIDGroupPairs `json:"UserIDGroupPairs"`
}

type SecurityGroupIPPermissionEgress struct {
        FromPort interface{} `json:"FromPort"`
        IPProtocol string `json:"IpProtocol"`
        IPRanges []SecurityGroupIPRanges `json:"IpRanges"`
        PrefixListIds interface{} `json:"PrefixListIds"`
        ToPort interface{} `json:"ToPort"`
        UserIDGroupPairs interface{} `json:"UserIdGroupPairs"`
}

type SecurityGroupIPRanges struct {
        CidrIP string `json:"CidrIp"`
}

type SecurityGroupUserIDGroupPairs struct {
        UserIDGroupPairs []struct {
            GroupID string `json:"GroupId"`
            GroupName interface{} `json:"GroupName"`
            PeeringStatus interface{} `json:"PeeringStatus"`
            UserID string `json:"UserId"`
            VpcID interface{} `json:"VpcId"`
            VpcPeeringConnectionID interface{} `json:"VpcPeeringConnectionId"`
        } `json:"UserIdGroupPairs"`
}

type SecurityGroups struct {
        SecurityGroups []SecurityGroup `json:"security_groups"`
}

func SecurityGroupsList() (*ec2.DescribeSecurityGroupsOutput, error) {
        svc := ec2.New(session.New(), &aws.Config{
                Region: aws.String("ap-southeast-2"),
        })

        params := &ec2.DescribeSecurityGroupsInput{}
        resp, err := svc.DescribeSecurityGroups(params)
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
