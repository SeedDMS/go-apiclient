package apiclient

import (
    "os"
    "testing"
    "github.com/stretchr/testify/assert"
    "fmt"
    "net/http"
)

func TestUpload(t *testing.T) {
    // Create the test server and shut it down when the test ends
    c, teardown := setupTestServer()
    defer teardown()

    // Add restapi endpoint to retrieve a document
    mux.HandleFunc("/folder/1/document", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        // ... return the JSON
        fmt.Fprint(w, fixture("upload.json"))
    })

    c.Login("admin", "admin")
    extraParams := map[string]string{
		"name":        "Document uploaded with go api-client",
		"keywords":    "go restapi",
        "filename":    "upload.txt",
	}
    file, err := os.Open("upload.go")
    if err != nil {
        return
    }
    defer file.Close()

    res, err := c.Upload(file, extraParams, 1)

    assert.Nil(t, err, "expecting nil error")
    assert.NotNil(t, res, "expecting non-nil result")
    if res != nil {
        assert.Equal(t, extraParams["name"], res.Data.Name, "expecting name of uploaded document")
    }
}


