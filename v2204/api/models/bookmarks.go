/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="bookmarks.go">
 *   Copyright (c) 2022 Aspose.Words for Cloud
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
type BookmarksResult struct {
    // Represents an array of bookmarks.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // Represents an array of bookmarks.
    BookmarkList []BookmarkResult `json:"BookmarkList,omitempty"`
}

type Bookmarks struct {
    // Represents an array of bookmarks.
    Link IWordsApiLink `json:"Link,omitempty"`

    // Represents an array of bookmarks.
    BookmarkList []Bookmark `json:"BookmarkList,omitempty"`
}

type IBookmarks interface {
    IsBookmarks() bool
    Initialize()
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



}


