package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

var (
	title      []string       //Хранит категории товара до уровня в котором есть товары
	categories = Categories{} //Хранит товары
)

//DecoderBodyJson парсинг из json в структуру
func (m *HTTPHeader) DecoderBodyJson(response *http.Response) (*CategoryModel, error) {
	res, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("ошибка получения BODY: " + err.Error())
	}

	var ctgMod CategoryModel
	json.Unmarshal(res, &ctgMod)

	return &ctgMod, nil
}

//ConvertData конвертирование данных в плоскую структуру. ОБРАЗЕЦ 1
func (m *HTTPHeader) ConvertData(data CategoryModel) {
	for _, v := range data.Result {
		if len(v.Children) != 0 {
			title = append(title, v.Name)
			m.ConvertData(CategoryModel{Result: v.Children})
		} else {
			categories = append(categories, Category{Id: v.Id, Name: v.Name, Parents: title})
		}
	}
	title = nil //Обнуляет гатегории
}

//CategoriesGoods каталог товаров
func (m *HTTPHeader) CategoriesGoods() Categories {
	return categories
}
