package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"

	"fmt"
)

func main() {
	session := session.Must(session.NewSession(&aws.Config{Region: aws.String("us-east-1")}))
	service := sns.New(session)

	topicArn := "<YOUR_TOPIC_ARN>"

	input := &sns.PublishInput{
		Message:  aws.String("Hello"),
		TopicArn: aws.String(topicArn),
	}

	result, err := service.Publish(input)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(result)
}
