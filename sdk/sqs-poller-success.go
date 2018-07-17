package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"

	"fmt"
)

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
		fmt.Println(message)

		deleteMessage(message, service, &queueUrl)
	}
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
