package apiclient

import (
    "net/http"
    "net/http/cookiejar"
//    "net/url"
    "fmt"
    "time"
    "io"
    "io/ioutil"
//    "os"
    "log"
    "encoding/json"
)

type Client struct {
    BaseURL string
    username string
    password string
    ApiKey string
    HTTPClient *http.Client
    StatusCode int
    ErrorMsg string
}

type Attribute struct {
    Id int
    Name string
    Value string
}

type Object struct {
    Id int
    Objtype string `json:"type"`
    Name string
    Comment string
    Date string
}

type Folder struct {
    Id int
    Objtype string `json:"type"`
    Name string
    Comment string
    Date string
    Attributes []Attribute
}

type Document struct {
    Id int
    Objtype string `json:"type"`
    Name string
    Comment string
    Keywords string
    Date string
    Mimetype string
    Filetype string
    Origfilename string
    Islocked bool
    Expires string
    Version int
    VersionComment string `json:"version_comment"`
    VersionDate string `json:"version_date"`
    Size int
    Attributes []Attribute
    VersionAttributes []Attribute `json:"version_attributes"`
}

type Group struct {
    Id int
    Objtype string `json:"type"`
    Name string
    Comment string
}

type Role struct {
    Id int
    Name string
}

type User struct {
    Id int
    Objtype string `json:"type"`
    Name string
    Comment string
    Login string
    Email string
    Language string
    Theme string
    Role Role
    Hidden bool
    Disabled bool
    Isguest bool
    Isadmin bool
    Groups []Group
}

type Statstotal struct {
    Docstotal int
    Folderstotal int
    Userstotal int
}

type errorResponse struct {
    Success bool
    Message string
    Data string
}

type successResponse struct {
    success bool
    message string
    data interface{}
}

func Connect(baseurl string, apikey string) *Client {
    jar, err := cookiejar.New(nil)
    if err != nil {
        log.Fatalf("Got error while creating cookie jar %s", err.Error())
    }
    return &Client{
        BaseURL: baseurl,
        ApiKey: apikey,
        HTTPClient: &http.Client{
            Timeout: time.Minute,
            Jar: jar, 
        },
    }
}

func (c *Client) getBody(url string) (io.Reader, error) {
    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        return nil, err
    }

    if c.ApiKey != "" {
        req.Header.Set("Authorization", fmt.Sprintf("%s", c.ApiKey))
    }

    res, err := c.HTTPClient.Do(req)
    if err != nil {
        return nil, err
    }

    c.StatusCode = res.StatusCode
    if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
        var errRes errorResponse
        if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
            c.ErrorMsg = errRes.Message
            return nil, fmt.Errorf(errRes.Message)
        }

        return nil, fmt.Errorf("unknown error, status code: %d", res.StatusCode)
    }

    return res.Body, nil
}

func (c *Client) sendRequest(req *http.Request, target interface{}) error {
    if req.Header.Get("Content-Type") == "" {
        req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    }
    req.Header.Set("Accept", "application/json; charset=utf-8")
    if c.ApiKey != "" {
        req.Header.Set("Authorization", fmt.Sprintf("%s", c.ApiKey))
    }

//    urlObj, _ := url.Parse(c.BaseURL)
//    fmt.Print(c.HTTPClient.Jar.Cookies(urlObj))

    res, err := c.HTTPClient.Do(req)
    if err != nil {
        return err
    }

    defer res.Body.Close()

    c.StatusCode = res.StatusCode
    if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
        var errRes errorResponse
        if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
            c.ErrorMsg = errRes.Message
            return fmt.Errorf(errRes.Message)
        }

        return fmt.Errorf("%d, %s", res.StatusCode, errRes.Message)
    }

    if false {
        // since goland 1.16 io.ReadAll should be used
        bodyBytes, err := ioutil.ReadAll(res.Body)
        if err == nil {
            jsonerr := json.Unmarshal([]byte(bodyBytes), target)
            //fmt.Printf("%+v", target)
            return jsonerr
        } else {
            //fmt.Print(err)
            return err
        }
    } else {
        // For some reason reading from the io stream didn't work because 'success'
        // was an unknown field
        //io.Copy(os.Stdout, res.Body)
        decoder := json.NewDecoder(res.Body)
        // Children() returns both, folders and documents and they have different
        // structs. That's why the json is mapped on the struct Object which just
        // contains the fields common to both folders and documents. But that
        // requires to allow unknows fields.
//        decoder.DisallowUnknownFields()
        return decoder.Decode(target)
    }
}


