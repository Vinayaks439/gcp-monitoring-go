package pkg

import (
	monitoring "cloud.google.com/go/monitoring/apiv3/v2"
	"context"
	"github.com/Vinayaks439/gcp-monitoring-go/internal"
)

type PubSub interface {
	NoUndeliveredMessages(ctx context.Context, config *internal.Config) (*internal.Result, error)
	OldestUnackedMessageAge(ctx context.Context, config *internal.Config) (*internal.Result, error)
}

type Monitor struct {
	*monitoring.QueryClient
}

var _ PubSub = (*Monitor)(nil)
