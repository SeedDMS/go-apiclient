package apiclient

import (
    "fmt"
    "net/http"
)

type statstotalResponse struct {
    Success bool `json:"success"`
    Message string `json:"message"`
    Data Statstotal
}

func (c *Client) Statstotal() (*statstotalResponse, error) {
    req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/statstotal", c.BaseURL), nil)
    if err != nil {
        return nil, err
    }

    res := statstotalResponse{}
    if err := c.sendRequest(req, &res); err != nil {
        return nil, err
    }

    return &res, nil
}

