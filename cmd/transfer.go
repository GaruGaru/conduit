package cmd

import (
	"bufio"
	"fmt"
	"github.com/GaruGaru/conduit/transfer"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/spf13/cobra"
	"os"
)

var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "moves all the messages from a queue to another",
	Run:   runTransfer,
}

var (
	transferSourceQueue        string
	transferDestinationQueue   string
	transferConcurrency        int
	transferBatchSize          int
	transferDeleteAfterPublish bool
)

func init() {

	transferCmd.PersistentFlags().StringVarP(&transferSourceQueue, "source", "s", "", "sqs source queue")
	transferCmd.PersistentFlags().StringVarP(&transferDestinationQueue, "destination", "d", "", "sqs destination queue")
	transferCmd.PersistentFlags().IntVarP(&transferConcurrency, "concurrency", "c", 1, "transfer number of workers")
	transferCmd.PersistentFlags().IntVarP(&transferBatchSize, "batch", "b", 10, "transfer messages batch size (max 10)")
	transferCmd.PersistentFlags().BoolVarP(&transferDeleteAfterPublish, "delete", "e", true, "delete message after publish")

	cobra.MarkFlagRequired(transferCmd.PersistentFlags(), "source")
	cobra.MarkFlagRequired(transferCmd.PersistentFlags(), "destination")

	rootCmd.AddCommand(transferCmd)
}

func runTransfer(cmd *cobra.Command, args []string) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	sqsClient := sqs.New(sess)

	transferJob := transfer.New(sqsClient, transferSourceQueue, transferDestinationQueue, transferDeleteAfterPublish, transferConcurrency, transferBatchSize)

	transferJob.RunAsync(func() {
		fmt.Println("completed")
		os.Exit(0)
	}, func(e error) {
		panic(e)
	})

	buf := bufio.NewReader(os.Stdin)

	fmt.Println("transfer running, press enter to interrupt.")
	_, err := buf.ReadBytes('\n')

	if err != nil {
		panic(err)
	}

	transferJob.Interrupt()

	fmt.Println("transfer interrupted")
}
