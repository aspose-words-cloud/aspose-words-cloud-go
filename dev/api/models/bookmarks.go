/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="bookmarks.go">
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

// Represents an array of bookmarks.

type IBookmarks interface {
    IsBookmarks() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetBookmarkList() []IBookmark
    SetBookmarkList(value []IBookmark)
}

type Bookmarks struct {
    // Represents an array of bookmarks.
    Link IWordsApiLink

    // Represents an array of bookmarks.
    BookmarkList []IBookmark
}

func (Bookmarks) IsBookmarks() bool {
    return true
}

func (Bookmarks) IsLinkElement() bool {
    return true
}

func (obj *Bookmarks) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.BookmarkList != nil) {
        for _, objElementBookmarkList := range obj.BookmarkList {
            objElementBookmarkList.Initialize()
        }
    }

}

func (obj *Bookmarks) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Link"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IWordsApiLink = new(WordsApiLink)
            modelInstance.Deserialize(parsedValue)
            obj.Link = modelInstance
        }

    } else if jsonValue, exists := json["link"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IWordsApiLink = new(WordsApiLink)
            modelInstance.Deserialize(parsedValue)
            obj.Link = modelInstance
        }

    }

    if jsonValue, exists := json["BookmarkList"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IBookmark = new(Bookmark)
                    modelElementInstance.Deserialize(elementValue)
                    obj.BookmarkList = append(obj.BookmarkList, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["bookmarkList"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IBookmark = new(Bookmark)
                    modelElementInstance.Deserialize(elementValue)
                    obj.BookmarkList = append(obj.BookmarkList, modelElementInstance)
                }

            }
        }

    }
}

func (obj *Bookmarks) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *Bookmarks) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *Bookmarks) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *Bookmarks) GetBookmarkList() []IBookmark {
    return obj.BookmarkList
}

func (obj *Bookmarks) SetBookmarkList(value []IBookmark) {
    obj.BookmarkList = value
}

