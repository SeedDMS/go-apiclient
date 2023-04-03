package apiclient

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "fmt"
    "net/http"
)

func TestChildren(t *testing.T) {
    // Create the test server and shut it down when the test ends
    c, teardown := setupTestServer()
    defer teardown()

    // Add restapi endpoint to retrieve the children of a folder
    mux.HandleFunc("/folder/1/children", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        // ... return the JSON
        fmt.Fprint(w, fixture("children.json"))
    })

    c.Login("admin", "admin")

    res, err := c.Children(1)

    assert.Nil(t, err, "expecting nil error")
    assert.NotNil(t, res, "expecting non-nil result")
    assert.NotNil(t, res.Data, "expecting array of objects")
}


