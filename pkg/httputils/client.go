package httputils

import (
	"context"
	"encoding/json"
	"net/http"
)

type HTTPClient interface {
	Get(ctx context.Context, path string) (*http.Response, error)
	ParseJSON(ctx context.Context, res *http.Response, v interface{}) error
}

type httpClient struct {
	client  *http.Client
	baseURL string
}

var _ HTTPClient = (*httpClient)(nil)

func (c httpClient) Get(ctx context.Context, path string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, c.baseURL+path, nil)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req.WithContext(ctx))
}

func (c httpClient) ParseJSON(ctx context.Context, res *http.Response, v interface{}) error {
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(v)
}

func NewHTTPClient(baseURL string, client *http.Client) *httpClient {
	return &httpClient{
		client:  client,
		baseURL: baseURL,
	}
}
