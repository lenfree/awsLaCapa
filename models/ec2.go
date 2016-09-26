package models

import (
        "github.com/astaxie/beego"
        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/awserr"
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go/service/ec2"
        "time"
)

/* This is a hack.
At the moment, I couldn't figure out a way to use type alias to generate
object model. Hence, use a model to describe ec2.DescribeInstancesOutput and
not being used anywhere else except for Swagger API documentation.
*/

type EC2Instance struct {
        AmiLaunchIndex      int    `json:"AmiLaunchIndex"`
        Architecture        string `json:"Architecture"`
        BlockDeviceMappings []struct {
                DeviceName string `json:"DeviceName"`
                Ebs        struct {
                        AttachTime          time.Time `json:"AttachTime"`
                        DeleteOnTermination bool      `json:"DeleteOnTermination"`
                        Status              string    `json:"Status"`
                        VolumeID            string    `json:"VolumeId"`
                } `json:"Ebs"`
        } `json:"BlockDeviceMappings"`
        ClientToken        string      `json:"ClientToken"`
        EbsOptimized       bool        `json:"EbsOptimized"`
        EnaSupport         interface{} `json:"EnaSupport"`
        Hypervisor         string      `json:"Hypervisor"`
        IamInstanceProfile struct {
                Arn string `json:"Arn"`
                ID  string `json:"Id"`
        } `json:"IamInstanceProfile"`
        ImageID           string      `json:"ImageId"`
        InstanceID        string      `json:"InstanceId"`
        InstanceLifecycle interface{} `json:"InstanceLifecycle"`
        InstanceType      string      `json:"InstanceType"`
        KernelID          interface{} `json:"KernelId"`
        KeyName           string      `json:"KeyName"`
        LaunchTime        time.Time   `json:"LaunchTime"`
        Monitoring        struct {
                State string `json:"State"`
        } `json:"Monitoring"`
        NetworkInterfaces []struct {
                Association interface{} `json:"Association"`
                Attachment  struct {
                        AttachTime          time.Time `json:"AttachTime"`
                        AttachmentID        string    `json:"AttachmentId"`
                        DeleteOnTermination bool      `json:"DeleteOnTermination"`
                        DeviceIndex         int       `json:"DeviceIndex"`
                        Status              string    `json:"Status"`
                } `json:"Attachment"`
                Description string `json:"Description"`
                Groups      []struct {
                        GroupID   string `json:"GroupId"`
                        GroupName string `json:"GroupName"`
                } `json:"Groups"`
                MacAddress         string `json:"MacAddress"`
                NetworkInterfaceID string `json:"NetworkInterfaceId"`
                OwnerID            string `json:"OwnerId"`
                PrivateDNSName     string `json:"PrivateDnsName"`
                PrivateIPAddress   string `json:"PrivateIpAddress"`
                PrivateIPAddresses []struct {
                        Association      interface{} `json:"Association"`
                        Primary          bool        `json:"Primary"`
                        PrivateDNSName   string      `json:"PrivateDnsName"`
                        PrivateIPAddress string      `json:"PrivateIpAddress"`
                } `json:"PrivateIpAddresses"`
                SourceDestCheck bool   `json:"SourceDestCheck"`
                Status          string `json:"Status"`
                SubnetID        string `json:"SubnetId"`
                VpcID           string `json:"VpcId"`
        } `json:"NetworkInterfaces"`
        Placement struct {
                Affinity         interface{} `json:"Affinity"`
                AvailabilityZone string      `json:"AvailabilityZone"`
                GroupName        string      `json:"GroupName"`
                HostID           interface{} `json:"HostId"`
                Tenancy          string      `json:"Tenancy"`
        } `json:"Placement"`
        Platform         interface{} `json:"Platform"`
        PrivateDNSName   string      `json:"PrivateDnsName"`
        PrivateIPAddress string      `json:"PrivateIpAddress"`
        ProductCodes     []struct {
                ProductCodeID   string `json:"ProductCodeId"`
                ProductCodeType string `json:"ProductCodeType"`
        } `json:"ProductCodes"`
        PublicDNSName   string      `json:"PublicDnsName"`
        PublicIPAddress interface{} `json:"PublicIpAddress"`
        RamdiskID       interface{} `json:"RamdiskId"`
        RootDeviceName  string      `json:"RootDeviceName"`
        RootDeviceType  string      `json:"RootDeviceType"`
        SecurityGroups  []struct {
                GroupID   string `json:"GroupId"`
                GroupName string `json:"GroupName"`
        } `json:"SecurityGroups"`
        SourceDestCheck       bool        `json:"SourceDestCheck"`
        SpotInstanceRequestID interface{} `json:"SpotInstanceRequestId"`
        SriovNetSupport       interface{} `json:"SriovNetSupport"`
        State                 struct {
                Code int    `json:"Code"`
                Name string `json:"Name"`
        } `json:"State"`
        StateReason           interface{} `json:"StateReason"`
        StateTransitionReason string      `json:"StateTransitionReason"`
        SubnetID              string      `json:"SubnetId"`
        Tags                  []struct {
                Key   string `json:"Key"`
                Value string `json:"Value"`
        } `json:"Tags"`
        VirtualizationType string `json:"VirtualizationType"`
        VpcID              string `json:"VpcId"`
}

type EC2Instances struct {
        ECInstance    []EC2Instance `json:"Instances"`
        Groups        interface{}   `json:"Groups"`
        OwnerID       string        `json:"OwnerId"`
        RequesterID   string        `json:"RequesterId"`
        ReservationID string        `json:"ReservationId"`
}

type EC2Reservations struct {
        EC2Reservations []EC2Instances `json:"Reservations"`
}

func EC2InstanceList() (*ec2.DescribeInstancesOutput, error) {
        svc := ec2.New(session.New(), &aws.Config{
                Region: aws.String("ap-southeast-2"),
        })

        params := &ec2.DescribeInstancesInput{}

        resp, err := svc.DescribeInstances(params)

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
