/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="bookmark_data.go">
 *   Copyright (c) 2021 Aspose.Words for Cloud
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

// DTO for bookmark updating.
type BookmarkDataResult struct {
    // DTO for bookmark updating.
    Name string `json:"Name,omitempty"`

    // DTO for bookmark updating.
    Text string `json:"Text,omitempty"`
}

type BookmarkData struct {
    // DTO for bookmark updating.
    Name *string `json:"Name,omitempty"`

    // DTO for bookmark updating.
    Text *string `json:"Text,omitempty"`
}

type IBookmarkData interface {
    IsBookmarkData() bool
}
func (BookmarkData) IsBookmarkData() bool {
    return true
}


