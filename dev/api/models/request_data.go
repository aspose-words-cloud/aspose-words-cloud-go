/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="request_data.go">
 *   Copyright (c) 2021 Aspose.Words for Cloud
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
	"fmt"
	"net/url"
	"reflect"
	"strings"
    "io"
)

// RequestInterface ia implemented all Words API requests
type RequestInterface interface {
	CreateRequestData() (request RequestData, err error)
    CreateResponse(data string) (response interface{}, err error)
}

// RequestData contains requst data
type RequestData struct {
	Path         string
	Method       string
	PostBody     interface{}
	HeaderParams map[string]string
	QueryParams  url.Values
	FormParams   []FormParamContainer
}

type FormParamContainer struct {
    Name string
    Text string
    File []byte
    IsFile bool
}

func NewTextFormParamContainer(name string, text string) (result FormParamContainer) {
    r := FormParamContainer{}
    r.Name = name
    r.Text = text
    r.IsFile = false
    return r
}

func NewFileFormParamContainer(name string, file []byte) (result FormParamContainer) {
    r := FormParamContainer{}
    r.Name = name
    r.File = file
    r.IsFile = true
    return r
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

	if reflect.TypeOf(obj) == reflect.TypeOf((*string)(nil)) {
		sp := obj.(*string)
		return *sp
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