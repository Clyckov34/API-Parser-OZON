package api

import (
	"net/http"
)

//Client
type Client interface {
	RequestClient() (*http.Response, error)
	DecoderBodyJson(response *http.Response) (*CategoryModel, error)
	ConvertData(data CategoryModel)
	CreateJsonFile(data Categories) error
}

//AppFlag параметры запуска приложения
type AppFlag struct {
	ExpFile bool
	Header  Headers
}

//Headers загаловок идентификация пользователя
type Headers struct {
	ClientId string
	ApiKey   string
	URL      string
	Method   string
}

type CategoryBaseModel struct {
	Name     string              `json:"title"`
	Id       int                 `json:"category_id"`
	Children []CategoryBaseModel `json:"children"`
}

type CategoryModel struct {
	Result []CategoryBaseModel `json:"result"`
}

type Categories []Category

type Category struct {
	Name    string
	Parents []string
	Id      interface{}
}
