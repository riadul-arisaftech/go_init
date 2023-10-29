package workers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const TaskSendVerifyEmail = "task:send_verify_email"

type PayloadSendVerifyEmail struct {
	Email string `json:"email"`
}

func (distributor *RedisTaskDistributor) DistributeTaskSendVerifyEmail(
	payload *PayloadSendVerifyEmail,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %w", err)
	}

	task := asynq.NewTask(TaskSendVerifyEmail, jsonPayload, opts...)
	info, err := distributor.Client.EnqueueContext(context.Background(), task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %w", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("enqueued task")
	return nil
}

func (processor *RedisTaskProcessor) ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error {
	var payload PayloadSendVerifyEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	fmt.Println("payload: %w", payload.Email)

	// begin transaction
	//tx := processor.store.CreateUser()
	//var user entities.User
	//if err := tx.Where("email = ?", payload.Email).Find(&user).Error; err != nil {
	//	tx.Rollback()
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		return fmt.Errorf("user does not exist: %w", asynq.SkipRetry)
	//	}
	//	return fmt.Errorf("failed to get user: %w", err)
	//}
	//
	//// to send email from here
	//otp := utils.GenerateOTP()
	//verificationData := entities.UserVerification{
	//	Email:     user.Email,
	//	OTP:       otp,
	//	Type:      entities.MailConfirmation,
	//	ExpiresAt: time.Now().Add(time.Minute * time.Duration(10)),
	//}
	//
	//if err := tx.Create(&verificationData).Error; err != nil {
	//	tx.Rollback()
	//	return fmt.Errorf("failed to create verify email: %w", err)
	//}
	//tx.Commit()
	// end transaction

	//subject := "Welcome to CV Sorting Project"
	//// verifyUrl := fmt.Sprintf("http://localhost:8080/v1/verify_email?email_id=%d&secret_code=%s",
	//// 	verifyEmail.ID, verifyEmail.SecretCode)
	//content := fmt.Sprintf(`Hello %s,<br/>
	//Thank you for registering with us!<br/>
	//Your OTP: <h2>%s</h2><br/>
	//`, user.Username, otp)
	//to := []string{user.Email}
	//
	//err := processor.mailer.SendEmail(subject, content, to, nil, nil, nil)
	//if err != nil {
	//	return fmt.Errorf("failed to send verify email: %w", err)
	//}
	//
	//log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
	//	Str("email", user.Email).Msg("processed task")
	return nil
}
