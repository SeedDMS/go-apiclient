package apiclient

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
    // Create the test server and shut it down when the test ends
    c, teardown := setupTestServer()
    defer teardown()

    res, err := c.Login("admin", "admin")

    assert.Nil(t, err, "expecting nil error")
    assert.NotNil(t, res, "expecting non-nil result")

}
