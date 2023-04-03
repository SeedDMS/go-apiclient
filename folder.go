package apiclient

import (
    "fmt"
    "net/http"
)

type folderResponse struct {
    Success bool `json:"success"`
    Message string `json:"message"`
    Data Folder
}

func (c *Client) Folder(id int) (*folderResponse, error) {
    req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/folder/%d", c.BaseURL, id), nil)
    if err != nil {
        return nil, err
    }

    res := folderResponse{}
    if err := c.sendRequest(req, &res); err != nil {
        return nil, err
    }

    return &res, nil
}


