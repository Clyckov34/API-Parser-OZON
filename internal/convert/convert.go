package convert

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

var (
	title []string //Хранит категории товара до уровня в котором есть товары
	goods Categories
)

//DecoderBodyJson парсинг из json в структуру
func (m *Categories) DecoderBodyJson(response *http.Response) (*CategoryModel, error) {
	res, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("ошибка получения BODY: " + err.Error())
	}

	var ctgMod CategoryModel
	json.Unmarshal(res, &ctgMod)

	return &ctgMod, nil
}

//ConvertData конвертирование данных в плоскую структуру. ОБРАЗЕЦ 1
func (m *Categories) ConvertData(data CategoryModel) {
	for _, v := range data.Result {
		if len(v.Children) > 0 {
			title = append(title, v.Name)
			m.ConvertData(CategoryModel{Result: v.Children})
		} else {
			goods = append(goods, Category{Id: v.Id, Name: v.Name, Parents: title})
		}
	}
	title = nil //Обнуляет гатегории
}

func (m *Categories) Goods() Categories {
	return goods
}
