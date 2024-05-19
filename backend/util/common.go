package util

import "time"

type Response struct {
	Body   string
	Status string
	Code   int
}
type Request struct {
	To           string
	Mobile       string
	TemplateName string
}

func GetCurrentTimeIn(timezone string) (time.Time, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, err
	}
	return time.Now().In(location), nil
}
