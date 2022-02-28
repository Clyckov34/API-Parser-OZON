package api

import (
	"context"
	"errors"
	"net/http"
)

//GetRequest получить запрос с сервера
func (m *API) GetRequest() (*http.Response, error) {
	ctx, _ := context.WithTimeout(context.Background(), m.WithTimeout)

	req, err := http.NewRequestWithContext(ctx, m.Method, m.URL, nil)
	if err != nil {
		return nil, errors.New("error: request server: " + err.Error())
	}

	req.Header.Add("Client-Id", m.ClientId)
	req.Header.Add("Api-Key", m.ApiKey)
	req.Header.Add("Content-Type", m.ContentType)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.New("error: request client: " + err.Error())
	}

	return res, nil
}
