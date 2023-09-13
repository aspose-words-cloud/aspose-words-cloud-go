/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="bookmark_response.go">
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

// The REST response with a bookmark.

type IBookmarkResponse interface {
    IsBookmarkResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetRequestId() *string
    SetRequestId(value *string)
    GetBookmark() IBookmark
    SetBookmark(value IBookmark)
}

type BookmarkResponse struct {
    // The REST response with a bookmark.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a bookmark.
    Bookmark IBookmark `json:"Bookmark,omitempty"`
}

func (BookmarkResponse) IsBookmarkResponse() bool {
    return true
}

func (BookmarkResponse) IsWordsResponse() bool {
    return true
}

func (obj *BookmarkResponse) Initialize() {
    if (obj.Bookmark != nil) {
        obj.Bookmark.Initialize()
    }


}

func (obj *BookmarkResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["Bookmark"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IBookmark = new(Bookmark)
            modelInstance.Deserialize(parsedValue)
            obj.Bookmark = modelInstance
        }

    } else if jsonValue, exists := json["bookmark"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IBookmark = new(Bookmark)
            modelInstance.Deserialize(parsedValue)
            obj.Bookmark = modelInstance
        }

    }
}

func (obj *BookmarkResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *BookmarkResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *BookmarkResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *BookmarkResponse) GetBookmark() IBookmark {
    return obj.Bookmark
}

func (obj *BookmarkResponse) SetBookmark(value IBookmark) {
    obj.Bookmark = value
}

