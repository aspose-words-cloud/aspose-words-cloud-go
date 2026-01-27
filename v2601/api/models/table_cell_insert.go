/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="table_cell_insert.go">
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

// DTO container with a table cell.

type ITableCellInsert interface {
    IsTableCellInsert() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetExistingCellPosition() IPosition
    SetExistingCellPosition(value IPosition)
    GetInsertAfter() *int32
    SetInsertAfter(value *int32)
}

type TableCellInsert struct {
    // DTO container with a table cell.
    ExistingCellPosition IPosition `json:"ExistingCellPosition,omitempty"`

    // DTO container with a table cell.
    InsertAfter *int32 `json:"InsertAfter,omitempty"`
}

func (TableCellInsert) IsTableCellInsert() bool {
    return true
}


func (obj *TableCellInsert) Initialize() {
    if (obj.ExistingCellPosition != nil) {
        obj.ExistingCellPosition.Initialize()
    }


}

func (obj *TableCellInsert) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["ExistingCellPosition"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IPosition = nil
            if jsonType, found := parsedValue["$type"]; found {
                jsonTypeStr := jsonType.(string)
                if jsonTypeStr == "PositionAfterNode, _" { modelInstance = new(PositionAfterNode) }
                if jsonTypeStr == "PositionBeforeNode, _" { modelInstance = new(PositionBeforeNode) }
                if jsonTypeStr == "PositionInsideNode, _" { modelInstance = new(PositionInsideNode) }
            }

            modelInstance.Deserialize(parsedValue)
            obj.ExistingCellPosition = modelInstance
        }

    } else if jsonValue, exists := json["existingCellPosition"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IPosition = nil
            if jsonType, found := parsedValue["$type"]; found {
                jsonTypeStr := jsonType.(string)
                if jsonTypeStr == "PositionAfterNode, _" { modelInstance = new(PositionAfterNode) }
                if jsonTypeStr == "PositionBeforeNode, _" { modelInstance = new(PositionBeforeNode) }
                if jsonTypeStr == "PositionInsideNode, _" { modelInstance = new(PositionInsideNode) }
            }

            modelInstance.Deserialize(parsedValue)
            obj.ExistingCellPosition = modelInstance
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

func (obj *TableCellInsert) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *TableCellInsert) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.ExistingCellPosition != nil {
        if err := obj.ExistingCellPosition.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *TableCellInsert) GetExistingCellPosition() IPosition {
    return obj.ExistingCellPosition
}

func (obj *TableCellInsert) SetExistingCellPosition(value IPosition) {
    obj.ExistingCellPosition = value
}

func (obj *TableCellInsert) GetInsertAfter() *int32 {
    return obj.InsertAfter
}

func (obj *TableCellInsert) SetInsertAfter(value *int32) {
    obj.InsertAfter = value
}

