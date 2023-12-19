package internal

import "github.com/spf13/cobra"

const (
	ProjectId      = "projectid"
	SubscriptionId = "subscriptionid"
	TopicId        = "topicid"
)

type Config struct {
	ProjectId      string `json:"project_id"`
	SubscriptionId string `json:"subscription_id,omitempty"`
	TopicId        string `json:"topic_id,omitempty"`
}

func InitConfig(cmd *cobra.Command) (*Config, error) {
	projectId, err := cmd.Flags().GetString(ProjectId)
	if err != nil {
		return nil, err
	}

	subscriptionId, err := cmd.Flags().GetString(SubscriptionId)
	if err != nil {
		return nil, err
	}

	topicId, err := cmd.Flags().GetString(TopicId)
	if err != nil {
		return nil, err
	}

	return &Config{
		ProjectId:      projectId,
		SubscriptionId: subscriptionId,
		TopicId:        topicId,
	}, nil
}
