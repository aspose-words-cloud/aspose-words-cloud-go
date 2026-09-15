/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="lists_response.go">
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

// The REST response with a collection of lists, contained in the document.
// This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/lists" REST API requests.

type IListsResponse interface {
    IsListsResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetLists() ILists
    SetLists(value ILists)
}

type ListsResponse struct {
    // The REST response with a collection of lists, contained in the document.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/lists" REST API requests.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a collection of lists, contained in the document.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/lists" REST API requests.
    Lists ILists `json:"Lists,omitempty"`
}

func (ListsResponse) IsListsResponse() bool {
    return true
}

func (ListsResponse) IsWordsResponse() bool {
    return true
}

func (obj *ListsResponse) Initialize() {
    if (obj.Lists != nil) {
        obj.Lists.Initialize()
    }


}

func (obj *ListsResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["Lists"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ILists = new(Lists)
            modelInstance.Deserialize(parsedValue)
            obj.Lists = modelInstance
        }

    } else if jsonValue, exists := json["lists"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ILists = new(Lists)
            modelInstance.Deserialize(parsedValue)
            obj.Lists = modelInstance
        }

    }
}

func (obj *ListsResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *ListsResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Lists != nil {
        if err := obj.Lists.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *ListsResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *ListsResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *ListsResponse) GetLists() ILists {
    return obj.Lists
}

func (obj *ListsResponse) SetLists(value ILists) {
    obj.Lists = value
}

