package apiclient

import (
    "fmt"
    "net/http"
)

type echoResponse struct {
    Success bool `json:"success"`
    Message string `json:"message"`
    Data string `json:"data"`
}

func (c *Client) Echo(msg string) (*echoResponse, error) {
    req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/echo/%s", c.BaseURL, msg), nil)
    if err != nil {
        return nil, err
    }

    res := echoResponse{}
    if err := c.sendRequest(req, &res); err != nil {
        return nil, err
    }

    return &res, nil
}

