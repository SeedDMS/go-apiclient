package apiclient

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "fmt"
    "net/http"
)

func TestFolder(t *testing.T) {
    // Create the test server and shut it down when the test ends
    c, teardown := setupTestServer()
    defer teardown()

    // Add restapi endpoint to retrieve a folder
    mux.HandleFunc("/folder/8517", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        // ... return the JSON
        fmt.Fprint(w, fixture("folder.json"))
    })

    c.Login("admin", "admin")
    res, err := c.Folder(8517)

    assert.Nil(t, err, "expecting nil error")
    assert.NotNil(t, res, "expecting non-nil result")
    assert.Equal(t, 8517, res.Data.Id, "expecting id=1 as we asked for it")
}


