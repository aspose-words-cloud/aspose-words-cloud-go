/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="table_cell_insert.go">
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

// DTO container with a table cell.

type ITableCellInsert interface {
    IsTableCellInsert() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetInsertAfter() *int32
    SetInsertAfter(value *int32)
}

type TableCellInsert struct {
    // DTO container with a table cell.
    InsertAfter *int32 `json:"InsertAfter,omitempty"`
}

func (TableCellInsert) IsTableCellInsert() bool {
    return true
}


func (obj *TableCellInsert) Initialize() {
}

func (obj *TableCellInsert) Deserialize(json map[string]interface{}) {
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

func (obj *TableCellInsert) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *TableCellInsert) GetInsertAfter() *int32 {
    return obj.InsertAfter
}

func (obj *TableCellInsert) SetInsertAfter(value *int32) {
    obj.InsertAfter = value
}

