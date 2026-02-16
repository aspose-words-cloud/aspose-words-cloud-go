/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="search_results_collection.go">
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

// The collection of search results.

type ISearchResultsCollection interface {
    IsSearchResultsCollection() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetResultsList() []ISearchResult
    SetResultsList(value []ISearchResult)
}

type SearchResultsCollection struct {
    // The collection of search results.
    Link IWordsApiLink `json:"Link,omitempty"`

    // The collection of search results.
    ResultsList []ISearchResult `json:"ResultsList,omitempty"`
}

func (SearchResultsCollection) IsSearchResultsCollection() bool {
    return true
}

func (SearchResultsCollection) IsLinkElement() bool {
    return true
}

func (obj *SearchResultsCollection) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.ResultsList != nil) {
        for _, objElementResultsList := range obj.ResultsList {
            objElementResultsList.Initialize()
        }
    }

}

func (obj *SearchResultsCollection) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["ResultsList"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.ResultsList = make([]ISearchResult, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance ISearchResult = new(SearchResult)
                    modelElementInstance.Deserialize(elementValue)
                    obj.ResultsList = append(obj.ResultsList, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["resultsList"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.ResultsList = make([]ISearchResult, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance ISearchResult = new(SearchResult)
                    modelElementInstance.Deserialize(elementValue)
                    obj.ResultsList = append(obj.ResultsList, modelElementInstance)
                }

            }
        }

    }
}

func (obj *SearchResultsCollection) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *SearchResultsCollection) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Link != nil {
        if err := obj.Link.Validate(); err != nil {
            return err
        }
    }
    if obj.ResultsList != nil {
        for _, elementResultsList := range obj.ResultsList {
            if elementResultsList != nil {
                if err := elementResultsList.Validate(); err != nil {
                    return err
                }
            }
        }
    }

    return nil;
}

func (obj *SearchResultsCollection) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *SearchResultsCollection) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *SearchResultsCollection) GetResultsList() []ISearchResult {
    return obj.ResultsList
}

func (obj *SearchResultsCollection) SetResultsList(value []ISearchResult) {
    obj.ResultsList = value
}

