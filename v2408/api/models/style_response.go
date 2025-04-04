/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="style_response.go">
 *   Copyright (c) 2024 Aspose.Words for Cloud
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

// The REST response with a style.
// This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/styles/{0}" REST API requests.

type IStyleResponse interface {
    IsStyleResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetStyle() IStyle
    SetStyle(value IStyle)
}

type StyleResponse struct {
    // The REST response with a style.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/styles/{0}" REST API requests.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a style.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/styles/{0}" REST API requests.
    Style IStyle `json:"Style,omitempty"`
}

func (StyleResponse) IsStyleResponse() bool {
    return true
}

func (StyleResponse) IsWordsResponse() bool {
    return true
}

func (obj *StyleResponse) Initialize() {
    if (obj.Style != nil) {
        obj.Style.Initialize()
    }


}

func (obj *StyleResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["Style"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IStyle = new(Style)
            modelInstance.Deserialize(parsedValue)
            obj.Style = modelInstance
        }

    } else if jsonValue, exists := json["style"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IStyle = new(Style)
            modelInstance.Deserialize(parsedValue)
            obj.Style = modelInstance
        }

    }
}

func (obj *StyleResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *StyleResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Style != nil {
        if err := obj.Style.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *StyleResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *StyleResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *StyleResponse) GetStyle() IStyle {
    return obj.Style
}

func (obj *StyleResponse) SetStyle(value IStyle) {
    obj.Style = value
}

