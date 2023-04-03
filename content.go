package apiclient

import (
    "fmt"
    "io"
)

func (c *Client) Content(id int) (io.Reader, error) {
    body, err := c.getBody(fmt.Sprintf("%s/document/%d/content", c.BaseURL, id))
    if err != nil {
        return nil, err
    }

    //defer body.Close()

    return body, nil
}

