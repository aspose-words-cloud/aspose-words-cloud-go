/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="search_result.go">
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

// Result of search operation.
type SearchResultResult struct {
    // Result of search operation.
    RangeStart DocumentPositionResult `json:"RangeStart,omitempty"`

    // Result of search operation.
    RangeEnd DocumentPositionResult `json:"RangeEnd,omitempty"`
}

type SearchResult struct {
    // Result of search operation.
    RangeStart IDocumentPosition `json:"RangeStart,omitempty"`

    // Result of search operation.
    RangeEnd IDocumentPosition `json:"RangeEnd,omitempty"`
}

type ISearchResult interface {
    IsSearchResult() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
}

func (SearchResult) IsSearchResult() bool {
    return true
}


func (obj *SearchResult) Initialize() {
    if (obj.RangeStart != nil) {
        obj.RangeStart.Initialize()
    }

    if (obj.RangeEnd != nil) {
        obj.RangeEnd.Initialize()
    }


}

func (obj *SearchResult) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}


