package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"fmt"
)

func main() {
	session := session.Must(session.NewSession(&aws.Config{Region: aws.String("us-east-1")}))
	service := s3.New(session)

	input := &s3.CreateBucketInput{
		Bucket: aws.String("<YOUR_BUCKET_NAME>"),
	}
	result, err := service.CreateBucket(input)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
