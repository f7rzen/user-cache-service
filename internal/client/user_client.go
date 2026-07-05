package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/f7rzen/user-cache-service/internal/model"
)

type UserClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewUserClient(baseURL string, httpClient *http.Client) *UserClient {
	return &UserClient{
		baseURL:    baseURL,
		httpClient: httpClient,
	}
}

func (c *UserClient) GetUsers(ctx context.Context) ([]model.ExternalUser, error) {
	url := c.baseURL + "/users"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("external API returned status: %d", resp.StatusCode)
	}

	var users []model.ExternalUser

	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, err
	}

	return users, nil
}

func (c *UserClient) GetUserByID(ctx context.Context, id int64) (model.ExternalUser, error) {
	url := fmt.Sprintf("%s/users/%d", c.baseURL, id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return model.ExternalUser{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return model.ExternalUser{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.ExternalUser{}, fmt.Errorf("external API returned status: %d", resp.StatusCode)
	}

	var user model.ExternalUser

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return model.ExternalUser{}, err
	}

	return user, nil
}
