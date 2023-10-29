package workers

import (
	"github.com/hibiken/asynq"
)

type TaskDistributor interface {
	DistributeTaskSendVerifyEmail(payload *PayloadSendVerifyEmail, opts ...asynq.Option) error
}
