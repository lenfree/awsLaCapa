package models

import (
        "github.com/astaxie/beego"
        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/awserr"
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go/service/s3"
        "github.com/bitly/go-simplejson"
        "io/ioutil"
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

func S3List() (*s3.ListBucketsOutput, error) {
        // TODO: Make region be smart and dynamic
        svc := s3.New(session.New(), &aws.Config{Region: aws.String("ap-southeast-2")})

        params := &s3.ListBucketsInput{}

        resp, err := svc.ListBuckets(params)

        if err != nil {
                if awsErr, ok := err.(awserr.Error); ok {
                        beego.Error(awsErr.Code(),
                                awsErr.Message(),
                                awsErr.OrigErr(),
                        )
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

func S3GetObjectByKey(bucket string, objectKey string) (*simplejson.Json, error) {
        svc := s3.New(session.New(), &aws.Config{Region: aws.String("ap-southeast-2")})
        params := &s3.GetObjectInput{
                Bucket:              aws.String(bucket),
                Key:                 aws.String(objectKey),
                ResponseContentType: aws.String("application/json"),
        }
        resp, err := svc.GetObject(params)
        if err != nil {
                beego.Debug(err)
                return nil, err
        }
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                beego.Debug(err)
                return nil, err
        }
        data, err := simplejson.NewJson(body)
        if err != nil {
                beego.Debug(err)
                return nil, err
        }
        return data, nil
}

func S3ListObjects(bucket string) (*s3.ListObjectsV2Output, error) {
        svc := s3.New(session.New(), &aws.Config{Region: aws.String("ap-southeast-2")})
        params := &s3.ListObjectsV2Input{
                Bucket: aws.String(bucket),
        }
        resp, err := svc.ListObjectsV2(params)
        if err != nil {
                beego.Debug(err)
                return nil, err
        }
        return resp, nil
}
