package models

import (
        "fmt"
        "github.com/astaxie/beego"
        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/awserr"
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go/service/iam"
        "github.com/lenfree/awsRestWrapper/connect"
        "time"
)

type IAMUser struct {
        Arn                     *string `json:"arn"`
        CreateDate              *time.Time `"json:creation_date" "timestampFormat:"iso8601"`
        PasswordLastUsed        *time.Time `"json:password_last_used" "timestampFormat:"iso8601"`
        Path                    *string `json:"path"`
        UserId                  *string `json:"user_id"`
        UserName                *string `json:"user_name"`
        UserPolicyList          []*iam.PolicyDetail `json:"user_policy_list"`
        AttachedManagedPolicies []*iam.AttachedPolicy `json:"attached_managed_policies"`
}

func IAMUserList() (*iam.ListUsersOutput, error) {
        svc := iam.New(session.New(), &aws.Config{Region: aws.String("ap-southeast-2")})
        resp, err := connect.IAMconnect(svc)
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

       fmt.Println(resp)
       return resp, nil
}
