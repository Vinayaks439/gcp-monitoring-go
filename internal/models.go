package internal

import (
	"cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	"time"
)

type Result struct {
	Metric         string                       `json:"metric"`
	TimeSeriesData *monitoringpb.TimeSeriesData `json:"time_series_data"`
	MetricData     []Metric                     `json:"metric_data"`
	LabelData      Label                        `json:"label_data"`
}

type Metric struct {
	Value     interface{} `json:"value"`
	StartTime time.Time   `json:"start_time"`
	EndTime   time.Time   `json:"end_time"`
}

type Label struct {
	ProjectId      string `json:"project_id"`
	SubscriptionId string `json:"subscription_id"`
}
