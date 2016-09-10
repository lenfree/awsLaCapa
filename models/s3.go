package models

import (
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/lenfree/awsRestWrapper/connect"
        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/service/s3"
        "github.com/aws/aws-sdk-go/aws/awserr"
        "github.com/astaxie/beego"
)

/* This is a hack.
At the moment, I couldn't figure out a way to use type alias to generate
object model. Hence, use a model to describe s3.ListBucketsOutput and
not being used anywhere else except for Swagger API documentation.
*/
type Bucket struct {
        CreationDate string `json:"CreationDate"`
        Name         string `json:"Name"`
}

type Owner struct {
        DisplayName string `json:"DisplayName"`
        ID          string `json:"ID"`
}

type Buckets struct {
        Buckets []Bucket `json:"buckets"`
        Owner   Owner    `json:"owner"`
}

func S3List() (*s3.ListBucketsOutput, error){
        // TODO: Make region be smart and dynamic
        svc := s3.New(session.New(), &aws.Config{Region: aws.String("ap-southeast-2")})
        resp, err := connect.S3connect(svc)
        if err != nil {
          if awsErr, ok := err.(awserr.Error); ok {
           beego.Error(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
           if reqErr, ok := err.(awserr.RequestFailure); ok {
               beego.Error(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
            }
          } else {
              beego.Debug(err.Error())
          }
          return nil, err
        }
        return resp, nil
}
