package tests

import (
	"github.com/go_sample/core/config"
	"github.com/go_sample/core/mail"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	config := config.All("../.")
	sender := mail.NewGmailSender(config.Email.SenderName, config.Email.SenderAddress, config.Email.SenderPassword)

	subject := "A test email"
	content := `
	<h1>Hello world</h1>
	<p>This is a test message from <a href="http://google.com">CV Sorting Project</a></p>
	`
	to := []string{"riadul.arisaftech@gmail.com"}
	//attachFiles := []string{"../README.md"}

	err := sender.SendEmail(subject, content, to, nil, nil, nil)
	require.NoError(t, err)
}
