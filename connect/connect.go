package connect

import (
  "github.com/aws/aws-sdk-go/service/ec2"
  "github.com/aws/aws-sdk-go/service/s3"
  "fmt"
)


var connectString = "Attempting to connect..."

func EC2connect(svc *ec2.EC2) (*ec2.DescribeInstancesOutput, error) {
  fmt.Printf("%s", connectString)
  resp, err := svc.DescribeInstances(nil)
  if err != nil {
    return nil, err
  }
  return resp, nil
}

func S3connect(svc *s3.S3) (*s3.ListBucketsOutput, error) {
  fmt.Printf("%s", connectString)
  var params *s3.ListBucketsInput
  resp, err := svc.ListBuckets(params)
  if err != nil {
    return nil, err
  }
  return resp, nil
}

