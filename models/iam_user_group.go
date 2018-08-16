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
object model. Hence, use a model to describe iam.GetGroup and not being
used anywhere else except for Swagger API documentation.Furthermore,
there's a known limitation with non built in data types such as time.Time
*/
type IAMUserGroupOutput struct {
	IAMUserGroups []*iam.GetGroupOutput `json:"IAM_user_groups"`
}

type IAMUserGroup struct {
	Group       IAMGroup    `json:"Group"`
	IsTruncated bool        `json:"IsTruncated"`
	Marker      interface{} `json:"Marker"`
	Users       []IAMUser   `json:"Users"`
}

type IAMUserGroups struct {
	IAMUserGroups []IAMUserGroup `json:"IAM_user_groups"`
}

func IAMUserGroupList() (*IAMUserGroupOutput, error) {
	groups, err := IAMGroupList()
	if err != nil {
		beego.Error(err)
	}
	resp, err := get_user_from_group(groups)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func get_user_from_group(groups *iam.ListGroupsOutput) (*IAMUserGroupOutput, error) {
	svc := iam.New(session.New(), &aws.Config{
		Region: aws.String(beego.AppConfig.String("awsRegion")),
	})
	data := new(IAMUserGroupOutput)
	for _, g := range groups.Groups {

		params := &iam.GetGroupInput{
			GroupName: aws.String(*g.GroupName),
		}

		resp, err := svc.GetGroup(params)

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
		data.IAMUserGroups = append(data.IAMUserGroups, resp)
	}
	return data, nil
}
