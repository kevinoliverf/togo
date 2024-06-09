package aws

import (
	"context"
	"encoding/base64"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	eventhandler "github.com/kozloz/togo/pkg/event_handler"
)

func (c *SQSClient) Push(ctx context.Context, event eventhandler.Event) error {

	log.Printf("Pushing event %v\n", event)

	// Encode the event body to base64
	messageBody := base64.StdEncoding.EncodeToString([]byte(event.GetBody()))

	// Send the message to the SQS queue
	result, err := c.client.SendMessage(ctx, &sqs.SendMessageInput{
		MessageGroupId:         aws.String(event.GetType()),
		MessageDeduplicationId: aws.String(event.GetID()),
		MessageBody:            aws.String(messageBody),
		QueueUrl:               &c.queueURL,
	})
	if err != nil {
		log.Printf("Error sending message %v", err)
		return err
	}

	log.Printf("Successfully pushed with message id %v", *result.MessageId)
	return nil
}
