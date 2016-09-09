package connect

import (
        "github.com/aws/aws-sdk-go/service/ec2"
        "github.com/aws/aws-sdk-go/service/s3"
        "github.com/aws/aws-sdk-go/service/iam"
        "github.com/astaxie/beego"
)

var connectString = "Attempting to connect..."

func EC2connect(svc *ec2.EC2) (*ec2.DescribeInstancesOutput, error) {
  beego.Info("%s", connectString)
  resp, err := svc.DescribeInstances(nil)
  if err != nil {
    return nil, err
  }
  return resp, nil
}

func S3connect(svc *s3.S3) (*s3.ListBucketsOutput, error) {
  beego.Info("%s", connectString)
  var params *s3.ListBucketsInput
  resp, err := svc.ListBuckets(params)
  if err != nil {
    return nil, err
  }
  return resp, nil
}

func IAMconnect(svc *iam.IAM) (*iam.ListUsersOutput, error) {
  beego.Info("%s", connectString)
  var params *iam.ListUsersInput
  resp, err := svc.ListUsers(params)
  if err != nil {
    return nil, err
  }
  return resp, nil
}

