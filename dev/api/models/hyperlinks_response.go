/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="hyperlinks_response.go">
 *   Copyright (c) 2023 Aspose.Words for Cloud
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
    "errors"
)

// The REST response with a collection of hyperlinks.
// This response should be returned by the service when handling "GET /{name}/hyperlinks" REST API calls.

type IHyperlinksResponse interface {
    IsHyperlinksResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetHyperlinks() IHyperlinks
    SetHyperlinks(value IHyperlinks)
}

type HyperlinksResponse struct {
    // The REST response with a collection of hyperlinks.
    // This response should be returned by the service when handling "GET /{name}/hyperlinks" REST API calls.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a collection of hyperlinks.
    // This response should be returned by the service when handling "GET /{name}/hyperlinks" REST API calls.
    Hyperlinks IHyperlinks `json:"Hyperlinks,omitempty"`
}

func (HyperlinksResponse) IsHyperlinksResponse() bool {
    return true
}

func (HyperlinksResponse) IsWordsResponse() bool {
    return true
}

func (obj *HyperlinksResponse) Initialize() {
    if (obj.Hyperlinks != nil) {
        obj.Hyperlinks.Initialize()
    }


}

func (obj *HyperlinksResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["Hyperlinks"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IHyperlinks = new(Hyperlinks)
            modelInstance.Deserialize(parsedValue)
            obj.Hyperlinks = modelInstance
        }

    } else if jsonValue, exists := json["hyperlinks"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IHyperlinks = new(Hyperlinks)
            modelInstance.Deserialize(parsedValue)
            obj.Hyperlinks = modelInstance
        }

    }
}

func (obj *HyperlinksResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *HyperlinksResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Hyperlinks != nil {
        if err := obj.Hyperlinks.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *HyperlinksResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *HyperlinksResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *HyperlinksResponse) GetHyperlinks() IHyperlinks {
    return obj.Hyperlinks
}

func (obj *HyperlinksResponse) SetHyperlinks(value IHyperlinks) {
    obj.Hyperlinks = value
}

