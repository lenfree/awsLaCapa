package models

import (
        "github.com/astaxie/beego"
        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/awserr"
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go/service/iam"
        "time"
)

/* This is a hack.
At the moment, I couldn't figure out a way to use type alias to generate
object model. Hence, use a model to describe iam.ListUsersOutput and
not being used anywhere else except for Swagger API documentation.
*/
type IAMGroup struct {
        Arn        string    `json:"Arn"`
        CreateDate time.Time `json:"CreateDate"`
        GroupID    string    `json:"GroupId"`
        GroupName  string    `json:"GroupName"`
        Path       string    `json:"Path"`
}

type IAMGroups struct {
        IAMGroups   []IAMGroup  `json:"Groups"`
        IsTruncated bool        `json:"IsTruncated"`
        Marker      interface{} `json:"Marker"`
}

func IAMUserList() (*iam.ListUsersOutput, error) {
        svc := iam.New(session.New(), &aws.Config{
                Region: aws.String("ap-southeast-2"),
        })

        params := &iam.ListUsersInput{}

        resp, err := svc.ListUsers(params)

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

func IAMGroupList() (*iam.ListGroupsOutput, error) {
        svc := iam.New(session.New(), &aws.Config{
                Region: aws.String("ap-southeast-2"),
        })

        params := &iam.ListGroupsInput{}

        resp, err := svc.ListGroups(params)

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
                        return nil, err
                }
        }
        return resp, nil
}
