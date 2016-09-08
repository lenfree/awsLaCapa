package models

import (
        "encoding/json"
        "github.com/lenfree/awsRestWrapper/connect"
        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/service/s3"
        "github.com/aws/aws-sdk-go/aws/awserr"
        "github.com/aws/aws-sdk-go/aws/awsutil"
        "time"
        "strings"
        "github.com/astaxie/beego"
)

type Bucket struct {
        Name         string     `json:"bucket_name"`
        CreationDate *time.Time `json:"bucket_creation_date"`
}

func S3List() ([]Bucket, error){
        // TODO: Make region be smart and dynamic
        svc := s3.New(&aws.Config{Region: "ap-southeast-2"})
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
        buckets := s3Buckets(resp)
        for _, bucket := range buckets {
                data, _ := json.Marshal(Bucket{
                        CreationDate: bucket.CreationDate,
                        Name: bucket.Name,
                })
                beego.Info(string(data))
        }
        return buckets, nil
}

func s3Buckets(resp *s3.ListBucketsOutput) ([]Bucket) {
        var buckets []Bucket
        for _, bucket := range resp.Buckets {
          bucket := Bucket{
            Name: strings.Replace(awsutil.StringValue(bucket.Name), "\"", "", -1),
            CreationDate: bucket.CreationDate,
          }
          buckets = append(buckets, bucket)
        }
        return buckets
}
