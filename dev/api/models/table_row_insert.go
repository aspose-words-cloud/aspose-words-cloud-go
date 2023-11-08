/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="table_row_insert.go">
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

// DTO container with a table row element.

type ITableRowInsert interface {
    IsTableRowInsert() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetColumnsCount() *int32
    SetColumnsCount(value *int32)
    GetInsertAfter() *int32
    SetInsertAfter(value *int32)
}

type TableRowInsert struct {
    // DTO container with a table row element.
    ColumnsCount *int32 `json:"ColumnsCount,omitempty"`

    // DTO container with a table row element.
    InsertAfter *int32 `json:"InsertAfter,omitempty"`
}

func (TableRowInsert) IsTableRowInsert() bool {
    return true
}


func (obj *TableRowInsert) Initialize() {
}

func (obj *TableRowInsert) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["ColumnsCount"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ColumnsCount = new(int32)
            *obj.ColumnsCount = int32(parsedValue)
        }

    } else if jsonValue, exists := json["columnsCount"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ColumnsCount = new(int32)
            *obj.ColumnsCount = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["InsertAfter"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.InsertAfter = new(int32)
            *obj.InsertAfter = int32(parsedValue)
        }

    } else if jsonValue, exists := json["insertAfter"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.InsertAfter = new(int32)
            *obj.InsertAfter = int32(parsedValue)
        }

    }
}

func (obj *TableRowInsert) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *TableRowInsert) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.ColumnsCount == nil {
        return errors.New("Property ColumnsCount in TableRowInsert is required.")
    }
    return nil;
}

func (obj *TableRowInsert) GetColumnsCount() *int32 {
    return obj.ColumnsCount
}

func (obj *TableRowInsert) SetColumnsCount(value *int32) {
    obj.ColumnsCount = value
}

func (obj *TableRowInsert) GetInsertAfter() *int32 {
    return obj.InsertAfter
}

func (obj *TableRowInsert) SetInsertAfter(value *int32) {
    obj.InsertAfter = value
}

