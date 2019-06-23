package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Retriever struct {
	Sqs *sqs.SQS
}

func NewRetriever(SQS *sqs.SQS) *Retriever {
	return &Retriever{
		Sqs: SQS,
	}
}

func (be *Retriever) Retrieve(queue string) ([]*sqs.Message, error) {
	result, err := be.Sqs.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            &queue,
		MaxNumberOfMessages: aws.Int64(10),
		VisibilityTimeout:   aws.Int64(20),
		WaitTimeSeconds:     aws.Int64(0),
	})

	return result.Messages, err
}
