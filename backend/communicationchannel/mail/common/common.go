package common

type Request struct {
	TemplateName string
	Subject      string
	To           []Recipient
	Cc           []Recipient
	Bcc          []Recipient
	Sender       Recipient
	Body         string
	ReplyTo      Recipient
}

type Response struct {
}

type Recipient struct {
	Name  string
	Email string
	Type  string
}
