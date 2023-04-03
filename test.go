package apiclient

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
)

var (
    mux    *http.ServeMux
    server *httptest.Server
)

// setupTestServer create a local http server which mimics the endpoints
// of the SeedDMS RestAPI
// This function creates to endpoints which are use in many other test
// All other endpoints need to be created when they are used for testing
// The function returns the api client and teardown function to shutdown
// the server
func setupTestServer() (*Client, func()) {
    mux = http.NewServeMux()
    server = httptest.NewServer(mux)

    // Create a client using the test server and a random apikey
    client := Connect(server.URL, "apikey")
//    fmt.Print(server.URL)

    mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        // ... return the JSON
        fmt.Fprint(w, fixture("login.json"))
    })

    mux.HandleFunc("/account", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        // ... return the JSON
        fmt.Fprint(w, fixture("account.json"))
    })

    return client, func() {
        server.Close()
    }
}

func fixture(path string) string {
    b, err := ioutil.ReadFile("testdata/fixtures/" + path)
    if err != nil {
        panic(err)
    }
    return string(b)
}


