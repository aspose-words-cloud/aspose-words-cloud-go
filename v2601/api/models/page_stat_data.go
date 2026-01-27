/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="page_stat_data.go">
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

// Container for the page's statistical data.

type IPageStatData interface {
    IsPageStatData() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetFootnotesStatData() IFootnotesStatData
    SetFootnotesStatData(value IFootnotesStatData)
    GetParagraphCount() *int32
    SetParagraphCount(value *int32)
    GetWordCount() *int32
    SetWordCount(value *int32)
    GetPageNumber() *int32
    SetPageNumber(value *int32)
}

type PageStatData struct {
    // Container for the page's statistical data.
    FootnotesStatData IFootnotesStatData `json:"FootnotesStatData,omitempty"`

    // Container for the page's statistical data.
    ParagraphCount *int32 `json:"ParagraphCount,omitempty"`

    // Container for the page's statistical data.
    WordCount *int32 `json:"WordCount,omitempty"`

    // Container for the page's statistical data.
    PageNumber *int32 `json:"PageNumber,omitempty"`
}

func (PageStatData) IsPageStatData() bool {
    return true
}


func (obj *PageStatData) Initialize() {
    if (obj.FootnotesStatData != nil) {
        obj.FootnotesStatData.Initialize()
    }


}

func (obj *PageStatData) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["FootnotesStatData"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFootnotesStatData = new(FootnotesStatData)
            modelInstance.Deserialize(parsedValue)
            obj.FootnotesStatData = modelInstance
        }

    } else if jsonValue, exists := json["footnotesStatData"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFootnotesStatData = new(FootnotesStatData)
            modelInstance.Deserialize(parsedValue)
            obj.FootnotesStatData = modelInstance
        }

    }

    if jsonValue, exists := json["ParagraphCount"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ParagraphCount = new(int32)
            *obj.ParagraphCount = int32(parsedValue)
        }

    } else if jsonValue, exists := json["paragraphCount"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ParagraphCount = new(int32)
            *obj.ParagraphCount = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["WordCount"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.WordCount = new(int32)
            *obj.WordCount = int32(parsedValue)
        }

    } else if jsonValue, exists := json["wordCount"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.WordCount = new(int32)
            *obj.WordCount = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["PageNumber"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.PageNumber = new(int32)
            *obj.PageNumber = int32(parsedValue)
        }

    } else if jsonValue, exists := json["pageNumber"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.PageNumber = new(int32)
            *obj.PageNumber = int32(parsedValue)
        }

    }
}

func (obj *PageStatData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *PageStatData) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.ParagraphCount == nil {
        return errors.New("Property ParagraphCount in PageStatData is required.")
    }
    if obj.WordCount == nil {
        return errors.New("Property WordCount in PageStatData is required.")
    }
    if obj.PageNumber == nil {
        return errors.New("Property PageNumber in PageStatData is required.")
    }
    if obj.FootnotesStatData != nil {
        if err := obj.FootnotesStatData.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *PageStatData) GetFootnotesStatData() IFootnotesStatData {
    return obj.FootnotesStatData
}

func (obj *PageStatData) SetFootnotesStatData(value IFootnotesStatData) {
    obj.FootnotesStatData = value
}

func (obj *PageStatData) GetParagraphCount() *int32 {
    return obj.ParagraphCount
}

func (obj *PageStatData) SetParagraphCount(value *int32) {
    obj.ParagraphCount = value
}

func (obj *PageStatData) GetWordCount() *int32 {
    return obj.WordCount
}

func (obj *PageStatData) SetWordCount(value *int32) {
    obj.WordCount = value
}

func (obj *PageStatData) GetPageNumber() *int32 {
    return obj.PageNumber
}

func (obj *PageStatData) SetPageNumber(value *int32) {
    obj.PageNumber = value
}

