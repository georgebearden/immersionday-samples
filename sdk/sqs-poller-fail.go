package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"

	"encoding/json"
	"fmt"
)

type CustomMessage struct {
	Name string
}

func main() {
	session := session.Must(session.NewSession(&aws.Config{Region: aws.String("us-east-1")}))
	service := sqs.New(session)

	queueUrl := "<YOUR_QUEUE_URL>"

	result, err := service.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl: &queueUrl,
	})

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	numMessages := len(result.Messages)
	fmt.Printf("Received %d messages.\n", numMessages)

	for _, message := range result.Messages {

		customMessage, err := decodeMessage(message)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(customMessage)
		deleteMessage(message, service, &queueUrl)
	}
}

func decodeMessage(message *sqs.Message) (*CustomMessage, error) {
	var customMessage CustomMessage

	if err := json.Unmarshal([]byte(*message.Body), &customMessage); err != nil {
		return nil, err
	}

	return &customMessage, nil
}

func deleteMessage(message *sqs.Message, service *sqs.SQS, queueUrl *string) {
	resultDelete, err := service.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      queueUrl,
		ReceiptHandle: message.ReceiptHandle,
	})

	if err != nil {
		fmt.Println("Delete Error", err)
		return
	}

	fmt.Println("Message Deleted", resultDelete)
}
