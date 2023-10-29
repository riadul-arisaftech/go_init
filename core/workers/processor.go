package workers

import (
	"context"
	"github.com/go_sample/core/config"
	"github.com/go_sample/core/mail"
	"github.com/go_sample/core/utils"
	"github.com/go_sample/database/repository"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

const (
	QueueCritical = "critical"
	QueueDefault  = "default"
)

type RedisTaskProcessor struct {
	server *asynq.Server
	store  repository.Store
	mailer mail.EmailSender
	config *config.Configuration
}

func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt, store repository.Store, mailer mail.EmailSender, config *config.Configuration) TaskProcessor {
	logger := utils.NewLogger()
	redis.SetLogger(logger)

	server := asynq.NewServer(
		redisOpt,
		asynq.Config{
			Queues: map[string]int{
				QueueCritical: 10,
				QueueDefault:  5,
			},
			ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
				log.Error().Err(err).Str("type", task.Type()).
					Bytes("payload", task.Payload()).Msg("process task failed")
			}),
			Logger: logger,
		},
	)
	return &RedisTaskProcessor{
		server: server,
		store:  store,
		mailer: mailer,
		config: config,
	}
}

func (processor *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()
	mux.HandleFunc(TaskSendVerifyEmail, processor.ProcessTaskSendVerifyEmail)
	return processor.server.Start(mux)
}

func RunTaskProcessor(config *config.Configuration, redisOpt asynq.RedisClientOpt, store repository.Store) {
	mailer := mail.NewGmailSender(config.Email.SenderName, config.Email.SenderAddress, config.Email.SenderPassword)
	taskProcessor := NewRedisTaskProcessor(redisOpt, store, mailer, config)
	log.Info().Msg("start task processor")
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start task processor")
	}
}
