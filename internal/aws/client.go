package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	eventhandler "github.com/kozloz/togo/pkg/event_handler"
)

type SQSClient struct {
	eventhandler.EventSource
	queueURL string
	client   *sqs.Client
}

func NewClient(queueURL string) (*SQSClient, error) {
	// Create a SQS service client.
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithClientLogMode(aws.LogRetries|aws.LogRequest),
		config.WithRegion("ap-southeast-2"))
	if err != nil {
		return nil, err
	}

	return &SQSClient{
		client:   sqs.NewFromConfig(cfg),
		queueURL: queueURL,
	}, nil
}
