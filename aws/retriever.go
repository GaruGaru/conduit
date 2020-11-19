package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"strconv"
)

type Retriever struct {
	sqsWrapper SQSWrapper
}

func NewRetriever(SQSWrapper SQSWrapper) *Retriever {
	return &Retriever{
		sqsWrapper: SQSWrapper,
	}
}

func (be *Retriever) Retrieve(queue string) ([]*sqs.Message, error) {

	result, err := be.sqsWrapper.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            &queue,
		MaxNumberOfMessages: aws.Int64(10),
		VisibilityTimeout:   aws.Int64(20),
		WaitTimeSeconds:     aws.Int64(0),
	})

	if err != nil {
		return []*sqs.Message{}, err
	}

	return result.Messages, err
}

func (be *Retriever) GetApproximateNumberOfMessages(queue string) (int64, error) {
	result, err := be.sqsWrapper.GetQueueAttributes(&sqs.GetQueueAttributesInput{
		AttributeNames: aws.StringSlice([]string{"ApproximateNumberOfMessages"}),
		QueueUrl:       &queue,
	})

	if err != nil {
		return 0, err
	}

	if result.Attributes["ApproximateNumberOfMessages"] == nil {
		return 0, fmt.Errorf("cannot retrieve ApproximateNumberOfMessages of queue: %s", queue)
	}

	count, err := strconv.ParseInt(*result.Attributes["ApproximateNumberOfMessages"], 10, 64)

	if err != nil {
		return 0, nil
	}

	return count, err
}
