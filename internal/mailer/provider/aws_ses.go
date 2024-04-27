package provider

import (
	"github.com/alexisleon/stori/internal/conf"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

var _ MailProvider = (*AwsSESProviderImplementation)(nil)

type AwsSESProviderImplementation struct {
	awsSesSession *ses.SES

	sourceEmail  string
	replyToEmail string
}

func NewAwsSESProvider(c *conf.Config) MailProvider {
	sess := session.Must(session.NewSession())

	return &AwsSESProviderImplementation{
		awsSesSession: ses.New(sess),
		sourceEmail:   c.Mailer.Source,
		replyToEmail:  c.Mailer.ReplyTo,
	}
}

func (m *AwsSESProviderImplementation) Mail(to, subject, message string) error {
	sesEmailInput := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{&to},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Data: aws.String(message),
				},
			},
			Subject: &ses.Content{
				Data: aws.String(subject),
			},
		},
		Source: aws.String(m.sourceEmail),
		ReplyToAddresses: []*string{
			aws.String(m.replyToEmail),
		},
	}

	_, err := m.awsSesSession.SendEmail(sesEmailInput)
	if err != nil {
		return err
	}

	return nil
}
