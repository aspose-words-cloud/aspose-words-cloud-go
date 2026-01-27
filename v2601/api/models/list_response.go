/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="list_response.go">
 *   Copyright (c) 2026 Aspose.Words for Cloud
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

// The REST response with a list information.
// This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/lists/{n}" REST API requests.

type IListResponse interface {
    IsListResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetList() IListInfo
    SetList(value IListInfo)
}

type ListResponse struct {
    // The REST response with a list information.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/lists/{n}" REST API requests.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a list information.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/lists/{n}" REST API requests.
    List IListInfo `json:"List,omitempty"`
}

func (ListResponse) IsListResponse() bool {
    return true
}

func (ListResponse) IsWordsResponse() bool {
    return true
}

func (obj *ListResponse) Initialize() {
    if (obj.List != nil) {
        obj.List.Initialize()
    }


}

func (obj *ListResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["List"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IListInfo = new(ListInfo)
            modelInstance.Deserialize(parsedValue)
            obj.List = modelInstance
        }

    } else if jsonValue, exists := json["list"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IListInfo = new(ListInfo)
            modelInstance.Deserialize(parsedValue)
            obj.List = modelInstance
        }

    }
}

func (obj *ListResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *ListResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.List != nil {
        if err := obj.List.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *ListResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *ListResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *ListResponse) GetList() IListInfo {
    return obj.List
}

func (obj *ListResponse) SetList(value IListInfo) {
    obj.List = value
}

