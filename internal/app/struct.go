package app

import (
	"myapiproject/internal/api"
)

//AppFlag параметры запуска приложения
type AppFlag struct {
	ExportFile bool
	API        api.API
}
