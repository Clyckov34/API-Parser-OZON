package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

//CreateJsonFile создание json файла
func (m *Headers) CreateJsonFile(data Categories) error {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return errors.New("ошибка: переобразование json")
	}

	err = ioutil.WriteFile("ozon.json", file, 0775)
	if err != nil {
		return errors.New("ошибка: записи файла")
	}
	return nil
}
