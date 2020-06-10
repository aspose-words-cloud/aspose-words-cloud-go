/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="api_client.go">
 *   Copyright (c) 2020 Aspose.Words for Cloud
 * </copyright>
 * <summary>
 *   Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 * 
 *  The above copyright notice and this permission notice shall be included in all
 *  copies or substantial portions of the Software.
 * 
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *  SOFTWARE.
 * </summary>
 * --------------------------------------------------------------------------------
 */

package api

import (
    "bytes"
    "context"
    "encoding/json"
    "encoding/xml"
    "errors"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "mime/multipart"
    "net/http"
    "net/http/httputil"
    "net/url"
    "os"
    "path/filepath"
    "reflect"
    "regexp"
    "strconv"
    "strings"
    "time"
    "unicode/utf8"

    "golang.org/x/oauth2"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2007/api/models"
)

var (
    jsonCheck = regexp.MustCompile("(?i:[application|text]/json)")
    xmlCheck = regexp.MustCompile("(?i:[application|text]/xml)")
)

// APIClient manages communication with the Aspose.Words for Cloud API Reference API v20.7
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
    cfg 	*models.Configuration
    common 	service 		// Reuse a single struct instead of allocating one for each service on the heap.

     // API Services
    WordsApi	*WordsApiService
}

type service struct {
    client *APIClient
}

type FormParamContainer struct {
    name string
    text string
    file []byte
    isFile bool
}

func NewTextFormParamContainer(name string, text string) (result FormParamContainer) {
    r := FormParamContainer{}
    r.name = name
    r.text = text
    r.isFile = false
    return r
}

func NewFileFormParamContainer(name string, file []byte) (result FormParamContainer) {
    r := FormParamContainer{}
    r.name = name
    r.file = file
    r.isFile = true
    return r
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *models.Configuration) (client *APIClient, err error) {
    if cfg.HttpClient == nil {
        cfg.HttpClient = http.DefaultClient
    }

    if cfg.AppKey == "" {
        return nil, errors.New("AppKey must be non-empty string")
    }

    if cfg.AppSid == "" {
        return nil, errors.New("AppSid must be non-empty string")
    }

    _, urlErr := url.ParseRequestURI(cfg.BaseUrl)

    if urlErr != nil {
        return nil, errors.New("BaseUrl must be valid URL")
    }

    c := &APIClient{}
    c.cfg = cfg
    c.common.client = c

    // API Services
    c.WordsApi = (*WordsApiService)(&c.common)

    return c, nil
}

func atoi(in string) (int, error) {
    return strconv.Atoi(in)
}


// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
    if len(contentTypes) == 0 {
        return ""
    }
    if contains(contentTypes, "application/json") {
        return "application/json"
    }
    return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
    if len(accepts) == 0 {
        return ""
    }

    if contains(accepts, "application/json") {
        return "application/json"
    }

    return strings.Join(accepts, ",")
}

// contains is a case insenstive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
    for _, a := range haystack {
        if strings.ToLower(a) == strings.ToLower(needle) {
            return true
        }
    }
    return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
    // Make sure there is an object.
    if obj == nil {
        return nil
    }

    // Check the type is as expected.
    if reflect.TypeOf(obj).String() != expected {
        return fmt.Errorf("Expected %s to be of type %s but received %s.", name, expected, reflect.TypeOf(obj).String())
    }
    return nil
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
    var delimiter string

    switch collectionFormat {
    case "pipes":
        delimiter = "|"
    case "ssv":
        delimiter = " "
    case "tsv":
        delimiter = "\t"
    case "csv":
        delimiter = ","
    }

    if reflect.TypeOf(obj).Kind() == reflect.Slice {
        return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
    }

    if reflect.TypeOf(obj).Kind() == reflect.Struct {
        b, _ := json.Marshal(obj)
        return string(b)
    }

    return fmt.Sprintf("%v", obj)
}

// callAPI do the request. 
func (c *APIClient) callAPI(request *http.Request) (resp *http.Response, err error) {

    // log request
    if c.cfg.DebugMode {
        dumpRequest, err := httputil.DumpRequest(request, true)
        if err != nil {
            return nil, err
        }

        log.Print(string(dumpRequest))
    }

    response, err := c.cfg.HttpClient.Do(request)

    if err != nil {
        return response, err
    }

    if c.cfg.DebugMode {
        dumpResponse, err := httputil.DumpResponse(response, true)
        if err != nil {
            return nil, err
        }

        log.Print(string(dumpResponse))
    }

    return response, err
}

// Change base path to allow switching to mocks
func (c *APIClient) ChangeBasePath(path string) {
    c.cfg.BaseUrl = path
}

// create context with token
func (c *APIClient) NewContextWithToken(ctx context.Context) (ctxWithToken context.Context, err error) {

    type authResponse struct {
        AccessToken  string `json:"access_token"`
        RefreshToken string `json:"refresh_token"`
    }

    if ctx == nil {
        ctx = context.Background()
    }

    tokenUrl, _ := url.Parse(c.cfg.BaseUrl)
    tokenUrl.Path = "/connect/token"

    response, err := http.PostForm(tokenUrl.String(), url.Values{
        "grant_type":    {"client_credentials"},
        "client_id":     {c.cfg.AppSid},
        "client_secret": {c.cfg.AppKey},
    })

    // http error
    if err != nil {
        return nil, err
    }

    if response.StatusCode != 200 {
        return nil, errors.New(fmt.Sprintf("Get a token request failed: %s returned %d status", tokenUrl.String(), response.StatusCode))
    }

    defer response.Body.Close()
    jsonData, err := ioutil.ReadAll(response.Body)

    if err != nil {
        return nil, err
    }

    var result authResponse
    parseErr := json.Unmarshal(jsonData, &result)

    if parseErr != nil {
        return nil, err
    }

    return context.WithValue(ctx, models.ContextAccessToken, result.AccessToken), nil
}

// prepareRequest build the request
func (c *APIClient) prepareRequest (
    ctx context.Context,
    path string, method string,
    postBody interface{},
    headerParams map[string]string,
    queryParams url.Values,
    formParams []FormParamContainer) (localVarRequest *http.Request, err error) {

    var body *bytes.Buffer

    // Detect postBody type and post.
    if postBody != nil {
        contentType := headerParams["Content-Type"]
        if contentType == "" {
            contentType = detectContentType(postBody)
            headerParams["Content-Type"] = contentType
        }

        body, err = setBody(postBody, contentType)
        if err != nil {
            return nil, err
        }
    }

    // add form parameters and file if available.
    if len(formParams) > 0 {
        if body != nil {
            return nil, errors.New("Cannot specify postBody and multipart form at the same time.")
        }
        body = &bytes.Buffer{}
        w := multipart.NewWriter(body)
        headerParams["Content-Type"] = w.FormDataContentType()

        for _, fp := range formParams {
            if fp.isFile {
                if len(fp.file) > 0 && fp.name != "" {
                    w.Boundary()
                    part, err := w.CreateFormFile("file", filepath.Base(fp.name))
                    if err != nil {
                        return nil, err
                    }
                    _, err = part.Write(fp.file)
                    if err != nil {
                        return nil, err
                    }
                }
            } else {
                w.WriteField(fp.name, fp.text)
            }
        }

        // Set Content-Length
        headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
        w.Close()
    }

    // Setup path and query parameters
    url, err := url.Parse(path)
    if err != nil {
        return nil, err
    }

    // Adding Query Param
    query := url.Query()
    for k, v := range queryParams {
        for _, iv := range v {
            query.Add(k, iv)
        }
    }

    // Encode the parameters.
    url.RawQuery = query.Encode()

    // Generate a new request
    if body != nil {
        localVarRequest, err = http.NewRequest(method, url.String(), body)
    } else {
        localVarRequest, err = http.NewRequest(method, url.String(), nil)
    }
    if err != nil {
        return nil, err
    }

    // add header parameters, if any
    if len(headerParams) > 0 {
        headers := http.Header{}
        for h, v := range headerParams {
            headers.Set(h, v)
        }
        localVarRequest.Header = headers
    }

    if ctx != nil {
        // add context to the request
        localVarRequest = localVarRequest.WithContext(ctx)

        // Walk through any authentication.

        // OAuth2 authentication
        if tok, ok := ctx.Value(models.ContextOAuth2).(oauth2.TokenSource); ok {
            // We were able to grab an oauth2 token from the context
            var latestToken *oauth2.Token
            if latestToken, err = tok.Token(); err != nil {
                return nil, err
            }

            latestToken.SetAuthHeader(localVarRequest)
        }

        // Basic HTTP Authentication
        if auth, ok := ctx.Value(models.ContextBasicAuth).(models.BasicAuth); ok {
            localVarRequest.SetBasicAuth(auth.UserName, auth.Password)
        }

        // AccessToken Authentication
        if auth, ok := ctx.Value(models.ContextAccessToken).(string); ok {
            localVarRequest.Header.Add("Authorization", "Bearer " + auth)
        }
    }

    for header, value := range c.cfg.DefaultHeader {
        localVarRequest.Header.Add(header, value)
    }

    return localVarRequest, nil
}


// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close()

    part, err := w.CreateFormFile(fieldName, filepath.Base(path))
    if err != nil {
        return err
    }
    _, err = io.Copy(part, file)

    return err
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) (error) {
    return fmt.Errorf(format, a...)
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
    if bodyBuf == nil {
        bodyBuf = &bytes.Buffer{}
    }

    if reader, ok := body.(io.Reader); ok {
        _, err = bodyBuf.ReadFrom(reader)
    } else if b, ok := body.([]byte); ok {
        _, err = bodyBuf.Write(b)
    } else if s, ok := body.(string); ok {
        _, err = bodyBuf.WriteString(s)
    } else if jsonCheck.MatchString(contentType) {
        err = json.NewEncoder(bodyBuf).Encode(body)
    } else if xmlCheck.MatchString(contentType) {
        xml.NewEncoder(bodyBuf).Encode(body)
    }

    if err != nil {
        return nil, err
    }

    if bodyBuf.Len() == 0 {
        err = fmt.Errorf("Invalid body type %s\n", contentType)
        return nil, err
    }
    return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
    contentType := "text/plain; charset=utf-8"
    kind := reflect.TypeOf(body).Kind()

    switch kind {
    case reflect.Struct, reflect.Map, reflect.Ptr:
        contentType = "application/json; charset=utf-8"
    case reflect.String:
        contentType = "text/plain; charset=utf-8"
    default:
        if b, ok := body.([]byte); ok {
            contentType = http.DetectContentType(b)
        } else if kind == reflect.Slice {
            contentType = "application/json; charset=utf-8"
        }
    }

    return contentType
}


// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
    cc := cacheControl{}
    ccHeader := headers.Get("Cache-Control")
    for _, part := range strings.Split(ccHeader, ",") {
        part = strings.Trim(part, " ")
        if part == "" {
            continue
        }
        if strings.ContainsRune(part, '=') {
            keyval := strings.Split(part, "=")
            cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
        } else {
            cc[part] = ""
        }
    }
    return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) (time.Time) {
    // Figure out when the cache expires.
    var expires time.Time
    now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
    if err != nil {
        return time.Now()
    }
    respCacheControl := parseCacheControl(r.Header)

    if maxAge, ok := respCacheControl["max-age"]; ok {
        lifetime, err := time.ParseDuration(maxAge + "s")
        if err != nil {
            expires = now
        }
        expires = now.Add(lifetime)
    } else {
        expiresHeader := r.Header.Get("Expires")
        if expiresHeader != "" {
            expires, err = time.Parse(time.RFC1123, expiresHeader)
            if err != nil {
                expires = now
            }
        }
    }
    return expires
}

func strlen(s string) (int) {
    return utf8.RuneCountInString(s)
}

func CreateWordsApi(config *models.Configuration) (wordsApi *WordsApiService, ctxWithToken context.Context, err error) {
    client, err := NewAPIClient(config)

    if err != nil {
        return nil, nil, err
    }

    ctx, err := client.NewContextWithToken(nil)

    if err != nil {
        return nil, nil, err
    }

    return client.WordsApi, ctx, err
}
