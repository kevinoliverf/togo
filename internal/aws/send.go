package aws

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	eventhandler "github.com/kozloz/togo/pkg/event_handler"
)

func (c *Client) Push(ctx context.Context, event eventhandler.Event) error {

	log.Printf("%v\n", event)
	messageBody := base64.StdEncoding.EncodeToString([]byte(event.GetBody()))
	result, err := c.client.SendMessage(&sqs.SendMessageInput{
		MessageGroupId:         aws.String(event.GetType()),
		MessageDeduplicationId: aws.String(event.GetID()),
		MessageBody:            aws.String(messageBody),
		QueueUrl:               &c.queueURL,
	})

	if err != nil {
		fmt.Println("Error", err)
		return err
	}

	fmt.Println("Success", *result.MessageId)
	return nil
}
