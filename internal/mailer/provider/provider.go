package provider

type MailProvider interface {
	Mail(to, subject, message string) error
}
