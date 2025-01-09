/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="search_response.go">
 *   Copyright (c) 2025 Aspose.Words for Cloud
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

// The REST response with a regular expression pattern and a collection of search results.
// This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/search" REST API requests.

type ISearchResponse interface {
    IsSearchResponse() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRequestId() *string
    SetRequestId(value *string)
    GetSearchingPattern() *string
    SetSearchingPattern(value *string)
    GetSearchResults() ISearchResultsCollection
    SetSearchResults(value ISearchResultsCollection)
}

type SearchResponse struct {
    // The REST response with a regular expression pattern and a collection of search results.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/search" REST API requests.
    RequestId *string `json:"RequestId,omitempty"`

    // The REST response with a regular expression pattern and a collection of search results.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/search" REST API requests.
    SearchingPattern *string `json:"SearchingPattern,omitempty"`

    // The REST response with a regular expression pattern and a collection of search results.
    // This response is returned by the Service when handling "GET https://api.aspose.cloud/v4.0/words/Test.doc/search" REST API requests.
    SearchResults ISearchResultsCollection `json:"SearchResults,omitempty"`
}

func (SearchResponse) IsSearchResponse() bool {
    return true
}

func (SearchResponse) IsWordsResponse() bool {
    return true
}

func (obj *SearchResponse) Initialize() {
    if (obj.SearchResults != nil) {
        obj.SearchResults.Initialize()
    }


}

func (obj *SearchResponse) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RequestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    } else if jsonValue, exists := json["requestId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RequestId = &parsedValue
        }

    }

    if jsonValue, exists := json["SearchingPattern"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.SearchingPattern = &parsedValue
        }

    } else if jsonValue, exists := json["searchingPattern"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.SearchingPattern = &parsedValue
        }

    }

    if jsonValue, exists := json["SearchResults"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ISearchResultsCollection = new(SearchResultsCollection)
            modelInstance.Deserialize(parsedValue)
            obj.SearchResults = modelInstance
        }

    } else if jsonValue, exists := json["searchResults"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ISearchResultsCollection = new(SearchResultsCollection)
            modelInstance.Deserialize(parsedValue)
            obj.SearchResults = modelInstance
        }

    }
}

func (obj *SearchResponse) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *SearchResponse) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.SearchResults != nil {
        if err := obj.SearchResults.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *SearchResponse) GetRequestId() *string {
    return obj.RequestId
}

func (obj *SearchResponse) SetRequestId(value *string) {
    obj.RequestId = value
}

func (obj *SearchResponse) GetSearchingPattern() *string {
    return obj.SearchingPattern
}

func (obj *SearchResponse) SetSearchingPattern(value *string) {
    obj.SearchingPattern = value
}

func (obj *SearchResponse) GetSearchResults() ISearchResultsCollection {
    return obj.SearchResults
}

func (obj *SearchResponse) SetSearchResults(value ISearchResultsCollection) {
    obj.SearchResults = value
}

