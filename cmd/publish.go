package cmd

import (
	"fmt"
	"github.com/GaruGaru/conduit/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/spf13/cobra"
)

var (
	destinationQueue string
	times            int
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "publish message to a queue",
	Run: func(cmd *cobra.Command, args []string) {
		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))

		sqsClient := sqs.New(sess)

		publisher := aws.NewPublisher(sqsClient)

		for _, msg := range args {
			fmt.Println(msg)
			for i := 0; i < times; i++ {
				err := publisher.Publish(msg, destinationQueue)

				if err != nil {
					panic(err)

				}
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(publishCmd)
	publishCmd.PersistentFlags().StringVarP(&destinationQueue, "destination", "d", "", "sqs destination queue")
	publishCmd.PersistentFlags().IntVarP(&times, "times", "t", 1, "number of message publish")

	cobra.MarkFlagRequired(publishCmd.PersistentFlags(), "destination")

}
