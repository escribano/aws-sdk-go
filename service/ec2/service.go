package ec2

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/protocol/ec2query"
	"github.com/awslabs/aws-sdk-go/internal/signer/v4"
)

// EC2 is a client for Amazon EC2.
type EC2 struct {
	*aws.Service
}

// New returns a new EC2 client.
func New(config *aws.Config) *EC2 {
	if config == nil {
		config = &aws.Config{}
	}

	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config),
		ServiceName: "ec2",
		APIVersion:  "2014-10-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(ec2query.Build)
	service.Handlers.Unmarshal.PushBack(ec2query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(ec2query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(ec2query.UnmarshalError)

	return &EC2{service}
}
