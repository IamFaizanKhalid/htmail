package htmail

import (
	"bytes"
	"fmt"
	"html/template"
	"net/mail"
)

type HTMaiL struct {
	to       mail.Address
	from     mail.Address
	template *template.Template
	sections sections
}

type sections struct {
	Subject   string
	PreHeader string
	Body      template.HTML
}

// MailComponents is used as a parameter for NewHTMaiL()
type MailComponents struct {
	Template  template.Template
	Subject   string
	PreHeader string
	To        mail.Address
	From      mail.Address
}

// NewHTMaiL can be used to get a new HTMaiL object
func NewHTMaiL(c MailComponents) HTMaiL {
	return HTMaiL{
		template: &c.Template,
		sections: sections{
			Subject:   c.Subject,
			PreHeader: c.PreHeader,
		},
		to:   c.To,
		from: c.From,
	}
}

// AppendElement appends an html element to the message
func (t *HTMaiL) AppendElement(elem template.HTML) {
	t.sections.Body += elem
}

// GenerateMessage can be used to generate html to send in email
func (t *HTMaiL) GenerateMessage() (bytes.Buffer, error) {
	var message bytes.Buffer
	message.Write(getHeader(t.to.Address, t.from.String(), t.sections.Subject))

	err := t.template.Execute(&message, t.sections)
	if err != nil {
		return bytes.Buffer{}, err
	}

	return message, nil
}

func getHeader(to, from, subject string) []byte {
	return []byte(fmt.Sprintf(
		"To: %s\nFrom: %s \nSubject: %s\nMIME-version: 1.0;\nContent-Type: text/html;\r\n\r\n",
		to, from, subject,
	))
}
