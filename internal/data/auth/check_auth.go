package auth

import (
	"context"
	"net/http"
	"job-order-be/internal/entity/auth"
	"job-order-be/pkg/errors"
	"job-order-be/pkg/httpclient"
)

// Data ...
type Data struct {
	client  *httpclient.Client
	baseURL string
}

// New ...
func New(client *httpclient.Client, baseURL string) Data {
	d := Data{
		client:  client,
		baseURL: baseURL,
	}
	return d
}

// CheckAuth ...
func (d Data) CheckAuth(ctx context.Context, _token, code string) (auth.Auth, error) {
	var auth auth.Auth
	var endpoint = "/checkrights"
	var url = d.baseURL + endpoint
	var body = map[string]interface{}{
		"code": code,
	}

	headers := make(http.Header)
	headers.Set("Authorization", _token)
	headers.Set("Content-Type", "application/json")

	_, err := d.client.PostJSON(ctx, url, endpoint, headers, body, &auth)
	if err != nil {
		return auth, errors.Wrap(err, "[DATA][CheckAuth]")
	}

	return auth, nil
}
