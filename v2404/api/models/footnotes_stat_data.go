/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="footnotes_stat_data.go">
 *   Copyright (c) 2024 Aspose.Words for Cloud
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

// Container for the footnotes statistical data.

type IFootnotesStatData interface {
    IsFootnotesStatData() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetParagraphCount() *int32
    SetParagraphCount(value *int32)
    GetWordCount() *int32
    SetWordCount(value *int32)
}

type FootnotesStatData struct {
    // Container for the footnotes statistical data.
    ParagraphCount *int32 `json:"ParagraphCount,omitempty"`

    // Container for the footnotes statistical data.
    WordCount *int32 `json:"WordCount,omitempty"`
}

func (FootnotesStatData) IsFootnotesStatData() bool {
    return true
}


func (obj *FootnotesStatData) Initialize() {
}

func (obj *FootnotesStatData) Deserialize(json map[string]interface{}) {
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
}

func (obj *FootnotesStatData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *FootnotesStatData) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.ParagraphCount == nil {
        return errors.New("Property ParagraphCount in FootnotesStatData is required.")
    }
    if obj.WordCount == nil {
        return errors.New("Property WordCount in FootnotesStatData is required.")
    }
    return nil;
}

func (obj *FootnotesStatData) GetParagraphCount() *int32 {
    return obj.ParagraphCount
}

func (obj *FootnotesStatData) SetParagraphCount(value *int32) {
    obj.ParagraphCount = value
}

func (obj *FootnotesStatData) GetWordCount() *int32 {
    return obj.WordCount
}

func (obj *FootnotesStatData) SetWordCount(value *int32) {
    obj.WordCount = value
}

