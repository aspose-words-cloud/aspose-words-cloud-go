/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="api_client.go">
 *   Copyright (c) 2025 Aspose.Words for Cloud
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
    "crypto/rand"
    "crypto/rsa"
    "encoding/base64"
    "encoding/binary"
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "math/big"
    "mime"
    "mime/multipart"
    "net/http"
    "net/http/httputil"
    "net/textproto"
    "net/url"
    "os"
    "path/filepath"
    "reflect"
    "regexp"
    "strconv"
    "strings"
    "time"
    "unicode/utf8"

    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2502/api/models"
    "golang.org/x/oauth2"
)

var (
    jsonCheck = regexp.MustCompile("(?i:[application|text]/json)")
    xmlCheck = regexp.MustCompile("(?i:[application|text]/xml)")
)

// APIClient manages communication with the Aspose.Words for Cloud API Reference API v25.1
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
    cfg     *models.Configuration
    common  service         // Reuse a single struct instead of allocating one for each service on the heap.
    key    *rsa.PublicKey // public key to encrypt data

     // API Services
    WordsApi    *WordsApiService
}

type service struct {
    client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *models.Configuration) (client *APIClient, err error) {
    if cfg.HttpClient == nil {
        cfg.HttpClient = http.DefaultClient
    }

    if cfg.ClientId == "" {
        return nil, errors.New("ClientId must be non-empty string")
    }

    if cfg.ClientSecret == "" {
        return nil, errors.New("ClientSecret must be non-empty string")
    }

    _, urlErr := url.ParseRequestURI(cfg.BaseUrl)

    if urlErr != nil {
        return nil, errors.New("BaseUrl must be valid URL")
    }

    cfg.HttpClient.Timeout = cfg.Timeout.Duration

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

// callAPI do the request. 
func (c *APIClient) callAPI(request *http.Request) (resp *http.Response, err error) {

    defer func() {
        if p := recover(); p != nil {
            panic(fmt.Sprintf("request error: %v", p))
        }
    }()

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
        return response, fmt.Errorf("%s request error: %w", request.URL.String(), err)
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
    tokenUrl.Path = "/v4.0/words/connect/token"

    response, err := http.PostForm(tokenUrl.String(), url.Values{
        "grant_type":    {"client_credentials"},
        "client_id":     {c.cfg.ClientId},
        "client_secret": {c.cfg.ClientSecret},
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
    data models.RequestData) (request *http.Request, err error) {
    var body *bytes.Buffer

    // encrypt passwords in FileReferences
    for _, fileReferenceData := range data.FileReferences {
        if fileReferenceData.Password != nil {
            encrypted, err := c.encrypt(ctx, *fileReferenceData.Password)
            if err == nil {
                fileReferenceData.EncryptedPassword = &encrypted
                fileReferenceData.Password = nil
            }
        }
    }

    // add form parameters and file if available.
    if len(data.FormParams) == 1 {
        body = &bytes.Buffer{}
        formParam := data.FormParams[0]
        data.HeaderParams["Content-Type"] = formParam.MimeType

        if formParam.IsFile {
            _, err = body.Write(formParam.File)
            if err != nil {
                return nil, err
            }
        } else {
            _, err = body.WriteString(formParam.Text)
            if err != nil {
                return nil, err
            }
        }
    } else if len(data.FormParams) > 1 {
        body = &bytes.Buffer{}
        w := multipart.NewWriter(body)
        data.HeaderParams["Content-Type"] = w.FormDataContentType()

        for _, fp := range data.FormParams {
            partDisposition := "form-data; name=\"" + fp.Name + "\""
            if fp.IsFile {
                partDisposition += "; filename=\"" + fp.Name + "\""
            }

            partHeader := make(textproto.MIMEHeader)
            partHeader.Set("Content-Disposition", partDisposition)
            partHeader.Set("Content-Type", fp.MimeType)
            part, err := w.CreatePart(partHeader)
            if err != nil {
                return nil, err
            }

            if fp.IsFile {
                _, err = part.Write(fp.File)
                if err != nil {
                    return nil, err
                }
            } else {
                _, err = part.Write([]byte(fp.Text))
                if err != nil {
                    return nil, err
                }
            }
        }

        // Set Content-Length
        data.HeaderParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
        w.Close()
    }

    // Setup path and query parameters
    url, err := url.Parse(data.Path)
    if err != nil {
        return nil, err
    }

    // Adding Query Param
    query := url.Query()
    for k, v := range data.QueryParams {
        for _, iv := range v {
            if k == "Password" && iv != "" {
                encrypted, err := c.encrypt(ctx, iv)

                if err != nil {
                    return nil, err
                }

                query.Add("encryptedPassword", encrypted)
            } else  {
                query.Add(k, iv)
            }
        }
    }

    // Encode the parameters.
    url.RawQuery = query.Encode()

    // Generate a new request
    if body != nil {
        request, err = http.NewRequest(data.Method, url.String(), body)
    } else {
        request, err = http.NewRequest(data.Method, url.String(), nil)
    }
    if err != nil {
        return nil, err
    }

    // add header parameters, if any
    if len(data.HeaderParams) > 0 {
        headers := http.Header{}
        for h, v := range data.HeaderParams {
            headers.Set(h, v)
        }
        request.Header = headers
    }

    if ctx != nil {
        // add context to the request
        request = request.WithContext(ctx)

        // Walk through any authentication.

        // OAuth2 authentication
        if tok, ok := ctx.Value(models.ContextOAuth2).(oauth2.TokenSource); ok {
            // We were able to grab an oauth2 token from the context
            var latestToken *oauth2.Token
            if latestToken, err = tok.Token(); err != nil {
                return nil, err
            }

            latestToken.SetAuthHeader(request)
        }

        // Basic HTTP Authentication
        if auth, ok := ctx.Value(models.ContextBasicAuth).(models.BasicAuth); ok {
            request.SetBasicAuth(auth.UserName, auth.Password)
        }

        // AccessToken Authentication
        if auth, ok := ctx.Value(models.ContextAccessToken).(string); ok {
            request.Header.Add("Authorization", "Bearer " + auth)
        }
    }

    for header, value := range c.cfg.DefaultHeader {
        request.Header.Add(header, value)
    }

    return request, nil
}

// encrypt string with RSA key from context
func (c *APIClient) encrypt(ctx context.Context, data string) (string, error) {

    if data == "" {
        return data, nil
    }

    if c.key == nil {

        modulus := c.cfg.Modulus
        exponent := c.cfg.Exponent

        if modulus == "" || exponent == "" {
            rsaKeyData, _, err := c.WordsApi.GetPublicKey(ctx, &models.GetPublicKeyRequest{})

            if err != nil {
                return "", err
            }

            if rsaKeyData.Modulus != nil && rsaKeyData.Exponent != nil {
                modulus = *rsaKeyData.Modulus
                exponent = *rsaKeyData.Exponent
            }
        }

        exponentBytes, err := base64.StdEncoding.DecodeString(exponent)

        if err != nil {
            return "", err
        }

        var eBytes []byte
        if len(exponentBytes) < 8 {
            eBytes = make([]byte, 8-len(exponentBytes), 8)
            eBytes = append(eBytes, exponentBytes...)
        } else {
            eBytes = exponentBytes
        }

        modulusBytes, err := base64.StdEncoding.DecodeString(modulus)

        if err != nil {
            return "", err
        }

        c.key = &rsa.PublicKey{
            N: big.NewInt(0).SetBytes(modulusBytes),
            E: int(binary.BigEndian.Uint64(eBytes)),
        }
    }

    encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, c.key, []byte(data))

    if err != nil {
        return "", err
    }

    return base64.StdEncoding.EncodeToString(encrypted), nil
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

func GetBoundary(response *http.Response) string {
    return getBoundary(response.Header.Get("Content-Type"));
}

func GetPartBoundary(response *multipart.Part) string {
    return getBoundary(response.Header.Get("Content-Type"));
}

func getBoundary(contentHeader string) string {
    if contentHeader == "" {
        return ""
    }

    _, params, err := mime.ParseMediaType(contentHeader)

    if err != nil {
        return ""
    }

    return params["boundary"]
}
