package apiclient

import (
    "fmt"
    "net/http"
)

type versionResponse struct {
    Success bool `json:"success"`
    Message string `json:"message"`
    Data Version
}

func (c *Client) Version() (*versionResponse, error) {
    req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/version", c.BaseURL), nil)
    if err != nil {
        return nil, err
    }

    res := versionResponse{}
    if err := c.sendRequest(req, &res); err != nil {
        return nil, err
    }

    return &res, nil
}

