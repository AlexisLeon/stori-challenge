package templates

import (
	"html/template"
	"io"
)

type Template interface {
	GetSubject() string
	Compile(wr io.Writer, data any) error
}

var _ Template = (*MailerTemplate)(nil)

type MailerTemplate struct {
	subject string
	path    string
}

func (t *MailerTemplate) Compile(wr io.Writer, data any) error {
	tpl, err := template.ParseFiles(t.path)
	if err != nil {
		return err
	}

	return tpl.Execute(wr, data)
}

func (t *MailerTemplate) GetSubject() string {
	return t.subject
}
