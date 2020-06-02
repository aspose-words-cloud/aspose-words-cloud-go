/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="configuration.go">
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

package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextOAuth2 takes a oauth2.TokenSource as authentication for the request.
	ContextOAuth2    	= contextKey("token")

	// ContextBasicAuth takes BasicAuth as authentication for the request.
	ContextBasicAuth 	= contextKey("basic")

	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken 	= contextKey("accesstoken")

	// ContextAPIKey takes an APIKey as authentication for the request
 	ContextAPIKey 		= contextKey("apikey")
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth 
type BasicAuth struct {
	UserName      string            `json:"userName,omitempty"`
	Password      string            `json:"password,omitempty"`	
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key 	string
	Prefix	string
}

type Configuration struct {
	BaseUrl       string            `json:"BaseUrl,omitempty"`
	AppKey        string            `json:"AppKey"`
	AppSid        string            `json:"AppSid"`
	DebugMode     bool              `json:"DebugMode,omitempty"`
	DefaultHeader map[string]string `json:"DefaultHeader,omitempty"`
	HttpClient    *http.Client
}

func NewConfiguration(configFilePath string) (pConfig *Configuration, err error) {

	data, err := ioutil.ReadFile(configFilePath)

	if err != nil {
		return nil, err
	}

	cfg := Configuration{
		BaseUrl:       "https://api.aspose.cloud",
		DebugMode:     false,
		DefaultHeader: map[string]string{"x-aspose-client": "go sdk", "x-aspose-client-version": "20.6"},
	}
	err = json.Unmarshal(data, &cfg)

	if err != nil {
		return nil, err
	}

	u, err := url.Parse(cfg.BaseUrl)
	u.Path = path.Join(u.Path, "v4.0")
	cfg.BaseUrl = u.String()

	return &cfg, nil
}

func (w *WordsApiErrorResponse) Error() string {
	out, err := json.Marshal(w)
	if err != nil {
		return err.Error()
	}

	return string(out)
}