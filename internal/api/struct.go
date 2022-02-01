package api

import (
	"net/http"
)

//Client
type New interface {
	RequestClient() (*http.Response, error)
	DecoderBodyJson(response *http.Response) (*CategoryModel, error)
	ConvertData(data CategoryModel)
	CategoriesGoods() Categories
	CreateJsonFile(data Categories) error
}

//AppFlag параметры запуска приложения
type AppFlag struct {
	ExpFile    bool
	HTTPHeader HTTPHeader
}

//Headers загаловок идентификация пользователя
type HTTPHeader struct {
	ClientId string
	ApiKey   string
	URL      string
	Method   string
}

type CategoryModel struct {
	Result []CategoryBaseModel `json:"result"`
}

type CategoryBaseModel struct {
	Name     string              `json:"title"`
	Id       int                 `json:"category_id"`
	Children []CategoryBaseModel `json:"children"`
}

type Categories []Category

type Category struct {
	Name    string
	Parents []string
	Id      interface{}
}
