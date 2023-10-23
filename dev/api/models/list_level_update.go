/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="list_level_update.go">
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

// Represents a document list levels.

type IListLevelUpdate interface {
    IsListLevelUpdate() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetStartAt() *int32
    SetStartAt(value *int32)
    GetNumberStyle() *string
    SetNumberStyle(value *string)
    GetNumberFormat() *string
    SetNumberFormat(value *string)
    GetAlignment() *string
    SetAlignment(value *string)
    GetIsLegal() *bool
    SetIsLegal(value *bool)
    GetRestartAfterLevel() *int32
    SetRestartAfterLevel(value *int32)
    GetTrailingCharacter() *string
    SetTrailingCharacter(value *string)
    GetTabPosition() *float64
    SetTabPosition(value *float64)
    GetNumberPosition() *float64
    SetNumberPosition(value *float64)
    GetTextPosition() *float64
    SetTextPosition(value *float64)
}

type ListLevelUpdate struct {
    // Represents a document list levels.
    StartAt *int32 `json:"StartAt,omitempty"`

    // Represents a document list levels.
    NumberStyle *string `json:"NumberStyle,omitempty"`

    // Represents a document list levels.
    NumberFormat *string `json:"NumberFormat,omitempty"`

    // Represents a document list levels.
    Alignment *string `json:"Alignment,omitempty"`

    // Represents a document list levels.
    IsLegal *bool `json:"IsLegal,omitempty"`

    // Represents a document list levels.
    RestartAfterLevel *int32 `json:"RestartAfterLevel,omitempty"`

    // Represents a document list levels.
    TrailingCharacter *string `json:"TrailingCharacter,omitempty"`

    // Represents a document list levels.
    TabPosition *float64 `json:"TabPosition,omitempty"`

    // Represents a document list levels.
    NumberPosition *float64 `json:"NumberPosition,omitempty"`

    // Represents a document list levels.
    TextPosition *float64 `json:"TextPosition,omitempty"`
}

func (ListLevelUpdate) IsListLevelUpdate() bool {
    return true
}


func (obj *ListLevelUpdate) Initialize() {
}

func (obj *ListLevelUpdate) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["StartAt"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.StartAt = new(int32)
            *obj.StartAt = int32(parsedValue)
        }

    } else if jsonValue, exists := json["startAt"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.StartAt = new(int32)
            *obj.StartAt = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["NumberStyle"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NumberStyle = &parsedValue
        }

    } else if jsonValue, exists := json["numberStyle"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NumberStyle = &parsedValue
        }

    }

    if jsonValue, exists := json["NumberFormat"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NumberFormat = &parsedValue
        }

    } else if jsonValue, exists := json["numberFormat"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NumberFormat = &parsedValue
        }

    }

    if jsonValue, exists := json["Alignment"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Alignment = &parsedValue
        }

    } else if jsonValue, exists := json["alignment"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Alignment = &parsedValue
        }

    }

    if jsonValue, exists := json["IsLegal"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsLegal = &parsedValue
        }

    } else if jsonValue, exists := json["isLegal"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsLegal = &parsedValue
        }

    }

    if jsonValue, exists := json["RestartAfterLevel"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.RestartAfterLevel = new(int32)
            *obj.RestartAfterLevel = int32(parsedValue)
        }

    } else if jsonValue, exists := json["restartAfterLevel"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.RestartAfterLevel = new(int32)
            *obj.RestartAfterLevel = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["TrailingCharacter"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.TrailingCharacter = &parsedValue
        }

    } else if jsonValue, exists := json["trailingCharacter"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.TrailingCharacter = &parsedValue
        }

    }

    if jsonValue, exists := json["TabPosition"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.TabPosition = &parsedValue
        }

    } else if jsonValue, exists := json["tabPosition"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.TabPosition = &parsedValue
        }

    }

    if jsonValue, exists := json["NumberPosition"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.NumberPosition = &parsedValue
        }

    } else if jsonValue, exists := json["numberPosition"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.NumberPosition = &parsedValue
        }

    }

    if jsonValue, exists := json["TextPosition"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.TextPosition = &parsedValue
        }

    } else if jsonValue, exists := json["textPosition"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.TextPosition = &parsedValue
        }

    }
}

func (obj *ListLevelUpdate) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *ListLevelUpdate) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.StartAt == nil {
        return errors.New("Property StartAt in ListLevelUpdate is required.")
    }

    if obj.NumberStyle == nil {
        return errors.New("Property NumberStyle in ListLevelUpdate is required.")
    }

    if obj.NumberFormat == nil {
        return errors.New("Property NumberFormat in ListLevelUpdate is required.")
    }

    if obj.Alignment == nil {
        return errors.New("Property Alignment in ListLevelUpdate is required.")
    }

    if obj.IsLegal == nil {
        return errors.New("Property IsLegal in ListLevelUpdate is required.")
    }

    if obj.RestartAfterLevel == nil {
        return errors.New("Property RestartAfterLevel in ListLevelUpdate is required.")
    }

    if obj.TrailingCharacter == nil {
        return errors.New("Property TrailingCharacter in ListLevelUpdate is required.")
    }

    if obj.TabPosition == nil {
        return errors.New("Property TabPosition in ListLevelUpdate is required.")
    }

    if obj.NumberPosition == nil {
        return errors.New("Property NumberPosition in ListLevelUpdate is required.")
    }

    if obj.TextPosition == nil {
        return errors.New("Property TextPosition in ListLevelUpdate is required.")
    }

    return nil;
}

func (obj *ListLevelUpdate) GetStartAt() *int32 {
    return obj.StartAt
}

func (obj *ListLevelUpdate) SetStartAt(value *int32) {
    obj.StartAt = value
}

func (obj *ListLevelUpdate) GetNumberStyle() *string {
    return obj.NumberStyle
}

func (obj *ListLevelUpdate) SetNumberStyle(value *string) {
    obj.NumberStyle = value
}

func (obj *ListLevelUpdate) GetNumberFormat() *string {
    return obj.NumberFormat
}

func (obj *ListLevelUpdate) SetNumberFormat(value *string) {
    obj.NumberFormat = value
}

func (obj *ListLevelUpdate) GetAlignment() *string {
    return obj.Alignment
}

func (obj *ListLevelUpdate) SetAlignment(value *string) {
    obj.Alignment = value
}

func (obj *ListLevelUpdate) GetIsLegal() *bool {
    return obj.IsLegal
}

func (obj *ListLevelUpdate) SetIsLegal(value *bool) {
    obj.IsLegal = value
}

func (obj *ListLevelUpdate) GetRestartAfterLevel() *int32 {
    return obj.RestartAfterLevel
}

func (obj *ListLevelUpdate) SetRestartAfterLevel(value *int32) {
    obj.RestartAfterLevel = value
}

func (obj *ListLevelUpdate) GetTrailingCharacter() *string {
    return obj.TrailingCharacter
}

func (obj *ListLevelUpdate) SetTrailingCharacter(value *string) {
    obj.TrailingCharacter = value
}

func (obj *ListLevelUpdate) GetTabPosition() *float64 {
    return obj.TabPosition
}

func (obj *ListLevelUpdate) SetTabPosition(value *float64) {
    obj.TabPosition = value
}

func (obj *ListLevelUpdate) GetNumberPosition() *float64 {
    return obj.NumberPosition
}

func (obj *ListLevelUpdate) SetNumberPosition(value *float64) {
    obj.NumberPosition = value
}

func (obj *ListLevelUpdate) GetTextPosition() *float64 {
    return obj.TextPosition
}

func (obj *ListLevelUpdate) SetTextPosition(value *float64) {
    obj.TextPosition = value
}

