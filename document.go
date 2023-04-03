package apiclient

import (
    "fmt"
    "net/http"
)

type documentResponse struct {
    Success bool `json:"success"`
    Message string `json:"message"`
    Data Document
}

func (c *Client) Document(id int) (*documentResponse, error) {
    req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/document/%d", c.BaseURL, id), nil)
    if err != nil {
        return nil, err
    }

    res := documentResponse{}
    if err := c.sendRequest(req, &res); err != nil {
        return nil, err
    }

    return &res, nil
}


