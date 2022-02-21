package api

import (
	"context"
	"errors"
	"net/http"
)

//clientHeader запрос с параметрами header
func (m *HTTPHeader) RequestClient() (*http.Response, error) {
	ctx, _ := context.WithTimeout(context.Background(), m.WithTimeout)

	req, err := http.NewRequestWithContext(ctx, m.Method, m.URL, nil)
	if err != nil {
		return nil, errors.New("ошибка запроса: " + err.Error())
	}

	req.Header.Add("Client-Id", m.ClientId)
	req.Header.Add("Api-Key", m.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.New("ошибка запроса от клиента: " + err.Error())
	}

	return res, nil
}
