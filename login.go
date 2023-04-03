package apiclient

import (
    "fmt"
    "net/http"
    "net/url"
    "strings"
)

type loginResponse struct {
    Success bool `json:"success"`
    Message string `json:"message"`
    Data User
}

func (c *Client) Login(username string, password string) (*loginResponse, error) {
    if c.ApiKey == "" {
        data := url.Values{}
        data.Set("user", username)
        data.Set("pass", password)

        req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/login", c.BaseURL), strings.NewReader(data.Encode()))
        if err != nil {
            return nil, err
        }

        res := loginResponse{}
        if err := c.sendRequest(req, &res); err != nil {
            return nil, err
        }

        return &res, nil
    } else {
        req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/account", c.BaseURL), nil)
        if err != nil {
            return nil, err
        }

        res := loginResponse{}
        if err := c.sendRequest(req, &res); err != nil {
            return nil, err
        }

        return &res, nil
    }
}
