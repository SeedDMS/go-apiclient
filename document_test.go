package apiclient

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "fmt"
    "net/http"
)

func TestDocument(t *testing.T) {
    // Create the test server and shut it down when the test ends
    c, teardown := setupTestServer()
    defer teardown()

    // Add restapi endpoint to retrieve a document
    mux.HandleFunc("/document/22545", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        // ... return the JSON
        fmt.Fprint(w, fixture("document.json"))
    })

    c.Login("admin", "admin")

    res, err := c.Document(22545)

    assert.Nil(t, err, "expecting nil error")
    assert.NotNil(t, res, "expecting non-nil result")
    assert.Equal(t, 22545, res.Data.Id, "expecting id=22545 as we asked for it")
}


