package api

import "time"

//API загаловок идентификация пользователя
type API struct {
	ClientId    string
	ApiKey      string
	ContentType string
	URL         string
	Method      string
	WithTimeout time.Duration
}
