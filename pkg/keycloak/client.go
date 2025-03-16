package keycloak

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	baseURL string
	realm   string
	client  *resty.Client
	token   string
}

func NewClient() *Client {
	return &Client{
		baseURL: os.Getenv("KEYCLOAK_BASE_URL"),
		realm:   os.Getenv("KEYCLOAK_REALM"),
		client:  resty.New().SetTimeout(10 * time.Second),
	}
}

func (c *Client) getAdminToken(ctx context.Context) error {
	resp, err := c.client.R().
		SetContext(ctx).
		SetFormData(
			map[string]string{
				"grant_type": "password",
				"client_id":  "admin-cli",
				"username":   os.Getenv("KEYCLOAK_ADMIN_USERNAME"),
				"password":   os.Getenv("KEYCLOAK_ADMIN_PASSWORD"),
			},
		).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetResult(map[string]interface{}{}).
		Post(fmt.Sprintf("%s/realms/master/protocol/openid-connect/token", c.baseURL))

	if err != nil {
		return err
	}

	result := resp.Result().(*map[string]interface{})
	if token, ok := (*result)["access_token"].(string); ok {
		c.token = token
		return nil
	}

	return fmt.Errorf("failed to get admin token")
}

func (c *Client) CreateUser(ctx context.Context, name, email, password string) error {
	if err := c.getAdminToken(ctx); err != nil {
		return err
	}

	user := map[string]interface{}{
		"username":  email,
		"email":     email,
		"firstName": name,
		"lastName":  name,
		"enabled":   true,
		"credentials": []map[string]interface{}{
			{
				"type":      "password",
				"value":     password,
				"temporary": false,
			},
		},
	}

	_, err := c.client.R().
		SetContext(ctx).
		SetHeader("Authorization", "Bearer "+c.token).
		SetHeader("Content-Type", "application/json").
		SetBody(user).
		Post(fmt.Sprintf("%s/admin/realms/%s/users", c.baseURL, c.realm))

	return err
}

func (c *Client) Login(ctx context.Context, email, password string) (string, error) {
	resp, err := c.client.R().
		SetContext(ctx).
		SetFormData(
			map[string]string{
				"grant_type": "password",
				"client_id":  os.Getenv("KEYCLOAK_CLIENT_ID"),
				"username":   email,
				"password":   password,
			},
		).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetResult(map[string]interface{}{}).
		Post(fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token", c.baseURL, c.realm))

	if err != nil {
		return "", err
	}

	result := resp.Result().(*map[string]interface{})
	if token, ok := (*result)["access_token"].(string); ok {
		return token, nil
	}

	return "", fmt.Errorf("invalid credentials")
}
