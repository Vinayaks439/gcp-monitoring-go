package pkg

import (
	"cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	"context"
	"errors"
	"fmt"
	"github.com/Vinayaks439/gcp-monitoring-go/internal"
	"google.golang.org/api/iterator"
	"log"
	"time"
)

func (m Monitor) NoUndeliveredMessages(ctx context.Context, config *internal.Config) (*internal.Result, error) {
	result := &internal.Result{
		Metric:         "pubsub.googleapis.com/subscription/num_undelivered_messages",
		TimeSeriesData: nil,
		MetricData:     nil,
	}
	var metric internal.Metric
	var label internal.Label
	req := &monitoringpb.QueryTimeSeriesRequest{
		Name: fmt.Sprintf("projects/%s", config.ProjectId), // optional
		Query: fmt.Sprintf(`fetch pubsub_subscription
		| metric 'pubsub.googleapis.com/subscription/num_undelivered_messages'
		| group_by 1m,
			[value_num_undelivered_messages_mean:
			   mean(value.num_undelivered_messages)]
		| within 5m`),
	}
	it := m.QueryTimeSeries(ctx, req)
	for {
		timeseries, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		for _, points := range timeseries.PointData {
			for _, values := range points.Values {
				metric.Value = values.GetDoubleValue()
			}
			metric.StartTime = time.Unix(points.TimeInterval.StartTime.Seconds, 0)
			metric.EndTime = time.Unix(points.TimeInterval.EndTime.Seconds, 0)
			result.MetricData = append(result.MetricData, metric)
		}
		for _, labels := range timeseries.LabelValues {
			label.ProjectId = labels.GetStringValue()
			label.SubscriptionId = labels.GetStringValue()
		}
		result.LabelData = label
		result.TimeSeriesData = timeseries
		return result, nil
	}
	return result, nil
}

func (m Monitor) OldestUnackedMessageAge(ctx context.Context, config *internal.Config) (*internal.Result, error) {
	result := &internal.Result{
		Metric:         "pubsub.googleapis.com/subscription/oldest_unacked_message_age",
		TimeSeriesData: nil,
		MetricData:     nil,
	}
	var metric internal.Metric
	var label internal.Label
	req := &monitoringpb.QueryTimeSeriesRequest{
		Name: fmt.Sprintf("projects/%s", config.ProjectId), // optional
		Query: fmt.Sprintf(`fetch pubsub_subscription
		| metric 'pubsub.googleapis.com/subscription/oldest_unacked_message_age'
		| group_by 1m,
			[value_oldest_unacked_message_age_max:
			   max(value.oldest_unacked_message_age)]
		| within 5m`),
	}
	it := m.QueryTimeSeries(ctx, req)
	for {
		timeseries, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		for _, points := range timeseries.PointData {
			for _, values := range points.Values {
				metric.Value = values.GetInt64Value()
			}
			metric.StartTime = time.Unix(points.TimeInterval.StartTime.Seconds, 0)
			metric.EndTime = time.Unix(points.TimeInterval.EndTime.Seconds, 0)
			result.MetricData = append(result.MetricData, metric)
		}
		for _, labels := range timeseries.LabelValues {
			label.ProjectId = labels.GetStringValue()
			label.SubscriptionId = labels.GetStringValue()
		}
		result.LabelData = label
		result.TimeSeriesData = timeseries
		return result, nil
	}
	return result, nil
}
