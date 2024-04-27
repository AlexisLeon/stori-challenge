package templates

func NewSummaryMailerTemplate() *MailerTemplate {
	return &MailerTemplate{
		subject: "Tu resumen de transacciones esta listo",
		path:    "internal/mailer/templates/html/summary.html",
	}
}
