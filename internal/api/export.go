package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"time"
)

//CreateJsonFile создание json файла
func (m *HTTPHeader) CreateJsonFile(data Categories) error {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return errors.New("ошибка: переобразование json")
	}

	h := time.Now().Format("2006-01-02 15:04:05")

	err = ioutil.WriteFile(h+".json", file, 0775)
	if err != nil {
		return errors.New("ошибка: записи файла")
	}
	return nil
}
