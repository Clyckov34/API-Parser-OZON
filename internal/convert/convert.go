package convert

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

var (
	category []string //Хранит категории товара до уровня в котором есть товары
	goods    Categories
)

//SetDecoderBodyJson парсинг из json в структуру
func (m *Categories) SetDecoderBodyJson(response *http.Response) (*CategoryModel, error) {
	res, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("ошибка получения BODY: " + err.Error())
	}

	var ctgMod CategoryModel
	json.Unmarshal(res, &ctgMod)

	return &ctgMod, nil
}

//SetConvert конвертирование данных в плоскую структуру. ОБРАЗЕЦ 1
func (m *Categories) SetConvert(data CategoryModel) {
	for _, v := range data.Result {
		if len(v.Children) > 0 {
			category = append(category, v.Name)
			m.SetConvert(CategoryModel{Result: v.Children})
		} else {
			goods = append(goods, Category{Id: v.Id, Name: v.Name, Parents: category})
		}
	}
	category = nil //Обнуляет гатегории
}

//GetGoods получить данные товара
func (m *Categories) GetGoods() Categories {
	return goods
}
