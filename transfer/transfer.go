package transfer

import (
	"github.com/GaruGaru/conduit/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"strconv"
	"sync"
)

type Job struct {
	Sqs              *sqs.SQS
	SourceQueue      string
	DestinationQueue string
	Delete           bool
	Concurrency      int
	retriever        aws.Retriever
	deleter          aws.Deleter
	publisher        aws.Publisher
	terminationCh    chan struct{}
	errorsCh         chan error
}

func New(sqs *sqs.SQS, sourceQueue string, destinationQueue string, delete bool, concurrency int) *Job {

	if concurrency <= 0 {
		panic("invalid concurrency " + strconv.Itoa(concurrency))
	}

	return &Job{
		Sqs:              sqs,
		SourceQueue:      sourceQueue,
		DestinationQueue: destinationQueue,
		Delete:           delete,
		Concurrency:      concurrency,
		retriever:        *aws.NewRetriever(sqs),
		deleter:          *aws.NewDeleter(sqs),
		publisher:        *aws.NewPublisher(sqs),
		terminationCh:    make(chan struct{}),
		errorsCh:         make(chan error, concurrency),
	}
}

func (t *Job) workerFn(wg *sync.WaitGroup) {
	defer wg.Done()
	for ; ; {
		select {
		case <-t.terminationCh:
			return
		default:

			messages, err := t.retriever.Retrieve(t.SourceQueue)

			if err != nil {
				t.errorsCh <- err
				return
			}

			if len(messages) == 0 {
				return
			}

			err = t.publisher.Redeliver(messages, t.DestinationQueue)

			if err != nil {
				t.errorsCh <- err
				return
			}

			if t.Delete {
				err = t.deleter.Delete(messages, t.SourceQueue)

				if err != nil {
					t.errorsCh <- err
					return
				}
			}

		}

	}
}

func (t *Job) Interrupt() {
	t.terminationCh <- struct{}{}
}

func (t *Job) RunAsync(onCompleteFn func(), onErrorFn func(error)) {
	go func() {

		err := t.Run()

		if err != nil {
			onErrorFn(err)
		} else {
			onCompleteFn()
		}

	}()
}

func (t *Job) Run() error {

	var wg sync.WaitGroup
	wg.Add(t.Concurrency)

	for i := 0; i < t.Concurrency; i++ {
		go t.workerFn(&wg)
	}

	wg.Wait()

	close(t.terminationCh)
	close(t.errorsCh)

	for err := range t.errorsCh {
		return err
	}

	return nil
}
