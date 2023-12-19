package cmd

import (
	monitoring "cloud.google.com/go/monitoring/apiv3/v2"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Vinayaks439/gcp-monitoring-go/internal"
	"github.com/Vinayaks439/gcp-monitoring-go/pkg"
	"github.com/spf13/cobra"
	"log"
	"os"
	"time"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gcpmonitoring",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		c, err := monitoring.NewQueryClient(ctx)
		if err != nil {
			log.Fatal(err)
		}
		defer c.Close()
		config, err := internal.InitConfig(cmd)
		if err != nil {
			os.Exit(1)
		}
		monitor := &pkg.Monitor{c}
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				go func() {
					result, err := monitor.NoUndeliveredMessages(ctx, config)
					if err != nil {
						log.Fatal(err)
					}
					body, err := json.Marshal(result)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println(string(body))
				}()
				go func() {
					result, err := monitor.OldestUnackedMessageAge(ctx, config)
					if err != nil {
						log.Fatal(err)
					}
					body, err := json.Marshal(result)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println(string(body))
				}()
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().String(internal.ProjectId, "", "Project ID")
	rootCmd.Flags().String(internal.SubscriptionId, "", "Subscription ID")
	rootCmd.Flags().String(internal.TopicId, "", "Topic ID")

	rootCmd.MarkFlagRequired(internal.ProjectId)
}
