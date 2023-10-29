package workers

import (
	"github.com/hibiken/asynq"
)

type RedisTaskDistributor struct {
	Client *asynq.Client
}

func NewRedisTaskDistributor(redisOpt asynq.RedisClientOpt) TaskDistributor {
	client := asynq.NewClient(redisOpt)
	return &RedisTaskDistributor{
		Client: client,
	}
}
