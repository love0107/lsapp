package util

type Response struct {
	Body   string
	Status string
	Code   int
}
type Request struct {
	To     string
	Mobile string
    TemplateName string
}
