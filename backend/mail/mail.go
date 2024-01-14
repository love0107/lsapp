package mail

import "time"

// Mail represents an email structure
type Mail struct {
	Sender      string
	Recipient   string
	Subject     string
	Body        string
	SentTime   time.Time
	Attachments []string
	CC          []Recipient
	BCC         []Recipient
	Headers     map[string]string
}

type Recipient struct{
	Name string
	Mail string
	Type string
}
