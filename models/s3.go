package models

import (
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/lenfree/awsRestWrapper/connect"
        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/service/s3"
        "github.com/aws/aws-sdk-go/aws/awserr"
        "github.com/astaxie/beego"
)

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
