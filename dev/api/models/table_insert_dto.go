/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="table_insert_dto.go">
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

// DTO container with a table element.

type ITableInsertDto interface {
    IsTableInsertDto() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetColumnsCount() *int32
    SetColumnsCount(value *int32)
    GetPosition() INewDocumentPosition
    SetPosition(value INewDocumentPosition)
    GetRowsCount() *int32
    SetRowsCount(value *int32)
}

type TableInsertDto struct {
    // DTO container with a table element.
    ColumnsCount *int32 `json:"ColumnsCount,omitempty"`

    // DTO container with a table element.
    Position INewDocumentPosition `json:"Position,omitempty"`

    // DTO container with a table element.
    RowsCount *int32 `json:"RowsCount,omitempty"`
}

func (TableInsertDto) IsTableInsertDto() bool {
    return true
}


func (obj *TableInsertDto) Initialize() {
    if (obj.Position != nil) {
        obj.Position.Initialize()
    }


}

func (obj *TableInsertDto) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["Position"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance INewDocumentPosition = new(NewDocumentPosition)
            modelInstance.Deserialize(parsedValue)
            obj.Position = modelInstance
        }

    } else if jsonValue, exists := json["position"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance INewDocumentPosition = new(NewDocumentPosition)
            modelInstance.Deserialize(parsedValue)
            obj.Position = modelInstance
        }

    }

    if jsonValue, exists := json["RowsCount"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.RowsCount = new(int32)
            *obj.RowsCount = int32(parsedValue)
        }

    } else if jsonValue, exists := json["rowsCount"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.RowsCount = new(int32)
            *obj.RowsCount = int32(parsedValue)
        }

    }
}

func (obj *TableInsertDto) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *TableInsertDto) GetColumnsCount() *int32 {
    return obj.ColumnsCount
}

func (obj *TableInsertDto) SetColumnsCount(value *int32) {
    obj.ColumnsCount = value
}

func (obj *TableInsertDto) GetPosition() INewDocumentPosition {
    return obj.Position
}

func (obj *TableInsertDto) SetPosition(value INewDocumentPosition) {
    obj.Position = value
}

func (obj *TableInsertDto) GetRowsCount() *int32 {
    return obj.RowsCount
}

func (obj *TableInsertDto) SetRowsCount(value *int32) {
    obj.RowsCount = value
}

