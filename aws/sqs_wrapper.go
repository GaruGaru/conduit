package aws

import "github.com/aws/aws-sdk-go/service/sqs"

//please AWS use interfaces to simplify test writing.
type SQSWrapper interface {
	ReceiveMessage(input *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error)
	GetQueueAttributes(input *sqs.GetQueueAttributesInput) (*sqs.GetQueueAttributesOutput, error)
}

type SQSWrapperImpl struct {
	sqs *sqs.SQS
}

func NewSQSWrapperImpl(sqs *sqs.SQS) *SQSWrapperImpl {
	return &SQSWrapperImpl{sqs: sqs}
}

func (s SQSWrapperImpl) ReceiveMessage(input *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
	return s.sqs.ReceiveMessage(input)
}

func (s SQSWrapperImpl) GetQueueAttributes(input *sqs.GetQueueAttributesInput) (*sqs.GetQueueAttributesOutput, error) {
	return s.sqs.GetQueueAttributes(input)
}
