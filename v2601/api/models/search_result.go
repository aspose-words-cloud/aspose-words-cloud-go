/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="search_result.go">
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

// Result of search operation.

type ISearchResult interface {
    IsSearchResult() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRangeStart() IDocumentPosition
    SetRangeStart(value IDocumentPosition)
    GetRangeEnd() IDocumentPosition
    SetRangeEnd(value IDocumentPosition)
}

type SearchResult struct {
    // Result of search operation.
    RangeStart IDocumentPosition `json:"RangeStart,omitempty"`

    // Result of search operation.
    RangeEnd IDocumentPosition `json:"RangeEnd,omitempty"`
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

func (obj *SearchResult) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RangeStart"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IDocumentPosition = new(DocumentPosition)
            modelInstance.Deserialize(parsedValue)
            obj.RangeStart = modelInstance
        }

    } else if jsonValue, exists := json["rangeStart"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IDocumentPosition = new(DocumentPosition)
            modelInstance.Deserialize(parsedValue)
            obj.RangeStart = modelInstance
        }

    }

    if jsonValue, exists := json["RangeEnd"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IDocumentPosition = new(DocumentPosition)
            modelInstance.Deserialize(parsedValue)
            obj.RangeEnd = modelInstance
        }

    } else if jsonValue, exists := json["rangeEnd"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IDocumentPosition = new(DocumentPosition)
            modelInstance.Deserialize(parsedValue)
            obj.RangeEnd = modelInstance
        }

    }
}

func (obj *SearchResult) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *SearchResult) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.RangeStart != nil {
        if err := obj.RangeStart.Validate(); err != nil {
            return err
        }
    }
    if obj.RangeEnd != nil {
        if err := obj.RangeEnd.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *SearchResult) GetRangeStart() IDocumentPosition {
    return obj.RangeStart
}

func (obj *SearchResult) SetRangeStart(value IDocumentPosition) {
    obj.RangeStart = value
}

func (obj *SearchResult) GetRangeEnd() IDocumentPosition {
    return obj.RangeEnd
}

func (obj *SearchResult) SetRangeEnd(value IDocumentPosition) {
    obj.RangeEnd = value
}

