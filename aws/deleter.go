package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Deleter struct {
	Sqs   *sqs.SQS
}

func NewDeleter(  SQS *sqs.SQS) *Deleter {
	return &Deleter{
		Sqs:   SQS,
	}
}

func (d *Deleter) Delete(messages []*sqs.Message, queue string) error {
	deleteBatch := make([]*sqs.DeleteMessageBatchRequestEntry, len(messages))
	for i := 0; i < len(messages); i++ {
		deleteBatch[i] = &sqs.DeleteMessageBatchRequestEntry{
			Id:            messages[i].MessageId,
			ReceiptHandle: messages[i].ReceiptHandle,
		}
	}

	deleteOut, err := d.Sqs.DeleteMessageBatch(&sqs.DeleteMessageBatchInput{
		QueueUrl: &queue,
		Entries:  deleteBatch,
	})

	if err != nil {
		panic(err)
	}

	if len(deleteOut.Failed) != 0 {
		 return fmt.Errorf("failed delete %d / %d messages", len(deleteOut.Failed), len(messages))
	}

	return err
}
