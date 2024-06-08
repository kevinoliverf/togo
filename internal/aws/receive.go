package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	eventhandler "github.com/kozloz/togo/pkg/event_handler"
)

func (c *Client) Fetch(ctx context.Context) ([]eventhandler.Event, error) {
	result, err := c.client.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
			aws.String(sqs.MessageSystemAttributeNameMessageGroupId),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            &c.queueURL,
		MaxNumberOfMessages: aws.Int64(10),
		VisibilityTimeout:   aws.Int64(60), // 60 seconds
		WaitTimeSeconds:     aws.Int64(0),
	})
	if err != nil {
		fmt.Println("Error", err)
		return nil, err
	}
	if len(result.Messages) == 0 {
		fmt.Println("Received no messages")
		return nil, nil
	}
	_, err = c.client.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      &c.queueURL,
		ReceiptHandle: result.Messages[0].ReceiptHandle,
	})

	if err != nil {
		fmt.Println("Delete Error", err)
		return nil, err
	}
	fmt.Printf("Success: %+v\n", result.Messages)

	events := []eventhandler.Event{}
	for _, message := range result.Messages {
		fmt.Printf("Message ID: %s\n", *message.MessageId)
		fmt.Printf("    Body: %s\n", *message.Body)
		for _, attr := range message.MessageAttributes {
			fmt.Printf("    Attributes: %s\n", attr)
		}

		event := SQSEvent{
			ID:      *message.MessageId,
			GroupID: *message.Attributes["MessageGroupId"],
			Body:    *message.Body,
		}
		events = append(events, &event)
	}

	for _, event := range events {
		log.Printf("Received messages: %+v", event)
	}
	return events, nil
}
