package api

import (
	"errors"
	"net/http"
)

//request запрос в api
func (m *HTTPHeader) request() (*http.Request, error) {
	req, err := http.NewRequest(m.Method, m.URL, nil)
	if err != nil {
		return nil, errors.New("ошибка запроса: " + err.Error())
	}
	return req, nil
}

//clientHeader запрос с параметрами header
func (m *HTTPHeader) RequestClient() (*http.Response, error) {
	req, err := m.request()
	if err != nil {
		return nil, err
	}

	req.Header.Add("Client-Id", m.ClientId)
	req.Header.Add("Api-Key", m.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resClient, err := client.Do(req)
	if err != nil {
		return nil, errors.New("ошибка запроса c клиента: " + err.Error())
	}

	return resClient, nil
}
