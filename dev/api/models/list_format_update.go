/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="list_format_update.go">
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

import (
    "errors"
)

// Paragraph list format element for update.

type IListFormatUpdate interface {
    IsListFormatUpdate() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetListLevelNumber() *int32
    SetListLevelNumber(value *int32)
    GetListId() *int32
    SetListId(value *int32)
}

type ListFormatUpdate struct {
    // Paragraph list format element for update.
    ListLevelNumber *int32 `json:"ListLevelNumber,omitempty"`

    // Paragraph list format element for update.
    ListId *int32 `json:"ListId,omitempty"`
}

func (ListFormatUpdate) IsListFormatUpdate() bool {
    return true
}


func (obj *ListFormatUpdate) Initialize() {
}

func (obj *ListFormatUpdate) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["ListLevelNumber"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ListLevelNumber = new(int32)
            *obj.ListLevelNumber = int32(parsedValue)
        }

    } else if jsonValue, exists := json["listLevelNumber"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ListLevelNumber = new(int32)
            *obj.ListLevelNumber = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["ListId"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ListId = new(int32)
            *obj.ListId = int32(parsedValue)
        }

    } else if jsonValue, exists := json["listId"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ListId = new(int32)
            *obj.ListId = int32(parsedValue)
        }

    }
}

func (obj *ListFormatUpdate) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *ListFormatUpdate) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.ListLevelNumber == nil {
        return errors.New("Property ListLevelNumber in ListFormatUpdate is required.")
    }

    if obj.ListId == nil {
        return errors.New("Property ListId in ListFormatUpdate is required.")
    }

    return nil;
}

func (obj *ListFormatUpdate) GetListLevelNumber() *int32 {
    return obj.ListLevelNumber
}

func (obj *ListFormatUpdate) SetListLevelNumber(value *int32) {
    obj.ListLevelNumber = value
}

func (obj *ListFormatUpdate) GetListId() *int32 {
    return obj.ListId
}

func (obj *ListFormatUpdate) SetListId(value *int32) {
    obj.ListId = value
}

