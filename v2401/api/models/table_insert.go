/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="table_insert.go">
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

// DTO container with a table element.

type ITableInsert interface {
    IsTableInsert() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetColumnsCount() *int32
    SetColumnsCount(value *int32)
    GetPosition() IPosition
    SetPosition(value IPosition)
    GetRowsCount() *int32
    SetRowsCount(value *int32)
}

type TableInsert struct {
    // DTO container with a table element.
    ColumnsCount *int32 `json:"ColumnsCount,omitempty"`

    // DTO container with a table element.
    Position IPosition `json:"Position,omitempty"`

    // DTO container with a table element.
    RowsCount *int32 `json:"RowsCount,omitempty"`
}

func (TableInsert) IsTableInsert() bool {
    return true
}


func (obj *TableInsert) Initialize() {
    if (obj.Position != nil) {
        obj.Position.Initialize()
    }


}

func (obj *TableInsert) Deserialize(json map[string]interface{}) {
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
            var modelInstance IPosition = nil
            if jsonType, found := parsedValue["$type"]; found {
                jsonTypeStr := jsonType.(string)
                if jsonTypeStr == "PositionAfterNode, _" { modelInstance = new(PositionAfterNode) }
                if jsonTypeStr == "PositionBeforeNode, _" { modelInstance = new(PositionBeforeNode) }
                if jsonTypeStr == "PositionInsideNode, _" { modelInstance = new(PositionInsideNode) }
            }

            modelInstance.Deserialize(parsedValue)
            obj.Position = modelInstance
        }

    } else if jsonValue, exists := json["position"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IPosition = nil
            if jsonType, found := parsedValue["$type"]; found {
                jsonTypeStr := jsonType.(string)
                if jsonTypeStr == "PositionAfterNode, _" { modelInstance = new(PositionAfterNode) }
                if jsonTypeStr == "PositionBeforeNode, _" { modelInstance = new(PositionBeforeNode) }
                if jsonTypeStr == "PositionInsideNode, _" { modelInstance = new(PositionInsideNode) }
            }

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

func (obj *TableInsert) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *TableInsert) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.ColumnsCount == nil {
        return errors.New("Property ColumnsCount in TableInsert is required.")
    }
    if obj.RowsCount == nil {
        return errors.New("Property RowsCount in TableInsert is required.")
    }
    if obj.Position != nil {
        if err := obj.Position.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *TableInsert) GetColumnsCount() *int32 {
    return obj.ColumnsCount
}

func (obj *TableInsert) SetColumnsCount(value *int32) {
    obj.ColumnsCount = value
}

func (obj *TableInsert) GetPosition() IPosition {
    return obj.Position
}

func (obj *TableInsert) SetPosition(value IPosition) {
    obj.Position = value
}

func (obj *TableInsert) GetRowsCount() *int32 {
    return obj.RowsCount
}

func (obj *TableInsert) SetRowsCount(value *int32) {
    obj.RowsCount = value
}

