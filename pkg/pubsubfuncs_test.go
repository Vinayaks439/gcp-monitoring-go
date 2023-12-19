package pkg

import (
	monitoring "cloud.google.com/go/monitoring/apiv3/v2"
	"context"
	"github.com/Vinayaks439/gcp-monitoring-go/internal"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestNoUndeliveredMessages(t *testing.T) {
	ctx := context.Background()
	c, err := monitoring.NewQueryClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	config := &internal.Config{
		ProjectId: "myproject",
	}
	monitor := &Monitor{c}
	result, err := monitor.NoUndeliveredMessages(ctx, config)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, "pubsub.googleapis.com/subscription/num_undelivered_messages", result.Metric)
	require.NotNil(t, result.TimeSeriesData)
	require.NotNil(t, result.MetricData)
	require.NotNil(t, result.LabelData)
}
func TestOldestUnackedMessageAge(t *testing.T) {
	ctx := context.Background()
	c, err := monitoring.NewQueryClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	config := &internal.Config{
		ProjectId: "myproject",
	}
	monitor := &Monitor{c}
	result, err := monitor.OldestUnackedMessageAge(ctx, config)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, "pubsub.googleapis.com/subscription/oldest_unacked_message_age", result.Metric)
	require.NotNil(t, result.TimeSeriesData)
	require.NotNil(t, result.MetricData)
	require.NotNil(t, result.LabelData)
}
