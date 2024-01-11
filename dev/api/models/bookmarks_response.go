/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="bookmarks_response.go">
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

// The REST response with a collection of bookmarks.
// This response should be returned by the service when handling: GET bookmarks.

type IBookmarksResponse interface {
    IsBookmarksResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetBookmarks() IBookmarks
    SetBookmarks(value IBookmarks)
}

type BookmarksResponse struct {
    // The REST response with a collection of bookmarks.
    // This response should be returned by the service when handling: GET bookmarks.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a collection of bookmarks.
    // This response should be returned by the service when handling: GET bookmarks.
    Bookmarks IBookmarks `json:"Bookmarks,omitempty"`
}

func (BookmarksResponse) IsBookmarksResponse() bool {
    return true
}

func (BookmarksResponse) IsWordsResponse() bool {
    return true
}

func (obj *BookmarksResponse) Initialize() {
    if (obj.Bookmarks != nil) {
        obj.Bookmarks.Initialize()
    }


}

func (obj *BookmarksResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["Bookmarks"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IBookmarks = new(Bookmarks)
            modelInstance.Deserialize(parsedValue)
            obj.Bookmarks = modelInstance
        }

    } else if jsonValue, exists := json["bookmarks"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IBookmarks = new(Bookmarks)
            modelInstance.Deserialize(parsedValue)
            obj.Bookmarks = modelInstance
        }

    }
}

func (obj *BookmarksResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *BookmarksResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Bookmarks != nil {
        if err := obj.Bookmarks.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *BookmarksResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *BookmarksResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *BookmarksResponse) GetBookmarks() IBookmarks {
    return obj.Bookmarks
}

func (obj *BookmarksResponse) SetBookmarks(value IBookmarks) {
    obj.Bookmarks = value
}

