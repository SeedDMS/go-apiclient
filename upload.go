package apiclient

import (
    "io"
    "fmt"
    "bytes"
    "mime/multipart"
    "net/http"
)

type UploadResponse struct {
    Success bool `json:"success"`
    Message string `json:"message"`
    Data Document
}

func (c *Client) Upload(file io.Reader, params map[string]string, parentid int) (*UploadResponse, error) {
    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)
    part, err := writer.CreateFormFile("text/plain", params["filename"])
    if err != nil {
        return nil, err
    }

    io.Copy(part, file)
    for key, val := range params {
		_ = writer.WriteField(key, val)
	}
    writer.Close()

    req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/folder/%d/document", c.BaseURL, parentid), body)
    if err != nil {
        return nil, err
    }

    res := UploadResponse{}
    req.Header.Set("Content-Type", writer.FormDataContentType())
    if err := c.sendRequest(req, &res); err != nil {
        return nil, err
    }

    return &res, nil
}
