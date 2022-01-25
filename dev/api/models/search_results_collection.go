/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="search_results_collection.go">
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

// The collection of search results.
type SearchResultsCollectionResult struct {
    // The collection of search results.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // The collection of search results.
    ResultsList []SearchResultResult `json:"ResultsList,omitempty"`
}

type SearchResultsCollection struct {
    // The collection of search results.
    Link IWordsApiLink `json:"Link,omitempty"`

    // The collection of search results.
    ResultsList []SearchResult `json:"ResultsList,omitempty"`
}

type ISearchResultsCollection interface {
    IsSearchResultsCollection() bool
    Initialize()
}

func (SearchResultsCollection) IsSearchResultsCollection() bool {
    return true
}

func (SearchResultsCollection) IsLinkElement() bool {
    return true
}

func (obj *SearchResultsCollection) Initialize() {
}


