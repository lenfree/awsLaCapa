package models

import (
        "github.com/astaxie/beego"
        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/awserr"
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go/service/iam"
)

/* This is a hack.
At the moment, I couldn't figure out a way to use type alias to generate
object model. Hence, use a model to describe iam.ListGroupsOutput and
not being used anywhere else except for Swagger API documentation. Futhermore,
there's a known limitation with non built in data types such as time.Time
*/
type IAMGroup struct {
        Arn        string    `json:"Arn"`
        CreateDate string    `json:"CreateDate"`
        GroupID    string    `json:"GroupId"`
        GroupName  string    `json:"GroupName"`
        Path       string    `json:"Path"`
}

type IAMGroups struct {
        Groups      []IAMGroup  `json:"IAM_groups"`
        IsTruncated bool        `json:"IsTruncated"`
        Marker      interface{} `json:"Marker"`
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
