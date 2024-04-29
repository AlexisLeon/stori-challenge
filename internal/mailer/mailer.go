package mailer

import (
	"bytes"
	"github.com/alexisleon/stori/internal/mailer/provider"
	"github.com/alexisleon/stori/internal/mailer/templates"
	"github.com/alexisleon/stori/internal/models"
)

type Mailer interface {
	SendSummary(summary *models.TransactionSummary) error
}

var _ Mailer = (*MailerImplementation)(nil)

type MailerImplementation struct {
	mailProvider provider.MailProvider

	summaryTemplate templates.Template
}

func NewMailer(p provider.MailProvider) Mailer {
	return &MailerImplementation{
		mailProvider:    p,
		summaryTemplate: templates.NewSummaryMailerTemplate(),
	}
}

func (m *MailerImplementation) SendSummary(summary *models.TransactionSummary) error {
	msg := new(bytes.Buffer)

	err := m.summaryTemplate.Compile(msg, summary)
	if err != nil {
		return err
	}

	return m.mailProvider.Mail(
		summary.User.Email,
		m.summaryTemplate.GetSubject(),
		msg.String(),
	)
}
