package apiclient

import (
    "fmt"
    "net/http"
)

type childrenResponse struct {
    Success bool `json:"success"`
    Message string `json:"message"`
    Data []Object
}

func (c *Client) Children(id int) (*childrenResponse, error) {
    req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/folder/%d/children", c.BaseURL, id), nil)
    if err != nil {
        return nil, err
    }

    res := childrenResponse{}
    if err := c.sendRequest(req, &res); err != nil {
        return nil, err
    }

    return &res, nil
}

