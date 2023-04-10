package aws

import (
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

func NewSession() client.ConfigProvider {
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}
	return sess
}

func NewDynamoClient() *dynamo.DB {
	return dynamo.New(NewSession())
}
