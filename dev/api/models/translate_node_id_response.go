/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="translate_node_id_response.go">
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

package models

import (
    "errors"
)

// The REST response with a node path.
// This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/translate/{0}" REST API requests.

type ITranslateNodeIdResponse interface {
    IsTranslateNodeIdResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetPath() *string
    SetPath(value *string)
}

type TranslateNodeIdResponse struct {
    // The REST response with a node path.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/translate/{0}" REST API requests.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a node path.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/translate/{0}" REST API requests.
    Path *string `json:"Path,omitempty"`
}

func (TranslateNodeIdResponse) IsTranslateNodeIdResponse() bool {
    return true
}

func (TranslateNodeIdResponse) IsWordsResponse() bool {
    return true
}

func (obj *TranslateNodeIdResponse) Initialize() {
}

func (obj *TranslateNodeIdResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["Path"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Path = &parsedValue
        }

    } else if jsonValue, exists := json["path"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Path = &parsedValue
        }

    }
}

func (obj *TranslateNodeIdResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *TranslateNodeIdResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    return nil;
}

func (obj *TranslateNodeIdResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *TranslateNodeIdResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *TranslateNodeIdResponse) GetPath() *string {
    return obj.Path
}

func (obj *TranslateNodeIdResponse) SetPath(value *string) {
    obj.Path = value
}

