package aws

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	eventhandler "github.com/kozloz/togo/pkg/event_handler"
)

func (c *SQSClient) Fetch(ctx context.Context) ([]eventhandler.Event, error) {
	// Receive messages from the SQS queue
	result, err := c.client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		MessageSystemAttributeNames: []types.MessageSystemAttributeName{
			types.MessageSystemAttributeNameSentTimestamp,
			types.MessageSystemAttributeNameMessageGroupId,
		},
		MessageAttributeNames: []string{
			"All",
		},
		QueueUrl:            &c.queueURL,
		MaxNumberOfMessages: 10,
		VisibilityTimeout:   60, // 60 seconds
		WaitTimeSeconds:     0,
	})
	if err != nil {
		log.Printf("Error receiving message with error %v", err)
		return nil, err
	}
	if len(result.Messages) == 0 {
		log.Printf("Received no messages")
		return nil, nil
	}

	var messageReceiptHandles []types.DeleteMessageBatchRequestEntry
	events := []eventhandler.Event{}
	for _, message := range result.Messages {
		log.Printf("Message ID: %s\n", *message.MessageId)
		log.Printf("    Body: %s\n", *message.Body)
		for _, attr := range message.MessageAttributes {
			log.Printf("    Attributes: %v\n", attr)
		}

		event := SQSEvent{
			ID:      *message.MessageId,
			GroupID: message.Attributes["MessageGroupId"],
			Body:    *message.Body,
		}

		// Add each event to the list of events to be returned
		events = append(events, &event)
		// Add the message to the list of messages to be deleted
		messageReceiptHandles = append(messageReceiptHandles, types.DeleteMessageBatchRequestEntry{
			Id:            message.MessageId,
			ReceiptHandle: message.ReceiptHandle,
		})

	}

	// Delete the messages from the queue so to not receive them again
	_, err = c.client.DeleteMessageBatch(ctx, &sqs.DeleteMessageBatchInput{
		QueueUrl: &c.queueURL,
		Entries:  messageReceiptHandles,
	})
	if err != nil {
		log.Printf("Error deleting message %v", err)
		return nil, err
	}

	for _, event := range events {
		log.Printf("Received messages: %+v", event)
	}
	return events, nil
}
