/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="bookmark_insert.go">
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

// Represents a bookmark to insert.
type BookmarkInsertResult struct {
    // Represents a bookmark to insert.
    Name string `json:"Name,omitempty"`

    // Represents a bookmark to insert.
    Text string `json:"Text,omitempty"`

    // Represents a bookmark to insert.
    EndRange DocumentPositionResult `json:"EndRange,omitempty"`

    // Represents a bookmark to insert.
    StartRange DocumentPositionResult `json:"StartRange,omitempty"`
}

type BookmarkInsert struct {
    // Represents a bookmark to insert.
    Name *string `json:"Name,omitempty"`

    // Represents a bookmark to insert.
    Text *string `json:"Text,omitempty"`

    // Represents a bookmark to insert.
    EndRange IDocumentPosition `json:"EndRange,omitempty"`

    // Represents a bookmark to insert.
    StartRange IDocumentPosition `json:"StartRange,omitempty"`
}

type IBookmarkInsert interface {
    IsBookmarkInsert() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileContent) []FileContent
}

func (BookmarkInsert) IsBookmarkInsert() bool {
    return true
}

func (BookmarkInsert) IsBookmarkData() bool {
    return true
}

func (obj *BookmarkInsert) Initialize() {
    if (obj.EndRange != nil) {
        obj.EndRange.Initialize()
    }

    if (obj.StartRange != nil) {
        obj.StartRange.Initialize()
    }


}

func (obj *BookmarkInsert) CollectFilesContent(resultFilesContent []FileContent) []FileContent {
    if (obj.EndRange != nil) {
        resultFilesContent = obj.EndRange.CollectFilesContent(resultFilesContent)
    }

    if (obj.StartRange != nil) {
        resultFilesContent = obj.StartRange.CollectFilesContent(resultFilesContent)
    }

    return resultFilesContent
}


