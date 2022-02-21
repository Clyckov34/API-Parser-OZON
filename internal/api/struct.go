package api

import "time"

//Headers загаловок идентификация пользователя
type HTTPHeader struct {
	ClientId    string
	ApiKey      string
	ContentType string
	URL         string
	Method      string
	WithTimeout time.Duration
}
