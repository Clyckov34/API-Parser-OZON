package app

import (
	"myapiproject/internal/api"
)

//AppFlag параметры запуска приложения
type AppFlag struct {
	ExpFile    bool
	HTTPHeader api.HTTPHeader
}
