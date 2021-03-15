package aws

import (
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/satori/go.uuid"
)

type Publisher struct {
	Sqs *sqs.SQS
	id  string
}

func NewPublisher(SQS *sqs.SQS) *Publisher {
	return &Publisher{
		Sqs: SQS,
		id:  uuid.NewV4().String(),
	}
}

func (p *Publisher) Publish(messageBody string, queue string) error {

	request := &sqs.SendMessageInput{
		QueueUrl:    &queue,
		MessageBody: &messageBody,
	}

	_, err := p.Sqs.SendMessage(request)

	if err != nil {
		return err
	}

	return nil
}

func (p *Publisher) Redeliver(messages []*sqs.Message, queue string) error {

	batch := make([]*sqs.SendMessageBatchRequestEntry, len(messages))
	for i := 0; i < len(messages); i++ {
		batch[i] = &sqs.SendMessageBatchRequestEntry{
			Id:          messages[i].MessageId,
			MessageBody: messages[i].Body,
		}
	}

	_, err := p.Sqs.SendMessageBatch(
		&sqs.SendMessageBatchInput{
			QueueUrl: &queue,
			Entries:  batch,
		},
	)

	if err != nil {
		return err
	}

	//if len(out.Failed) != 0 {
	//	return fmt.Errorf("failed publish of %d / %d messages", len(out.Failed), len(batch))
	//}

	return nil
}
