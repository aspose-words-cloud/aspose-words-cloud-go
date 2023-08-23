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

// Represents a document list levels.

type IListLevelUpdate interface {
    IsListLevelUpdate() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetAlignment() *string
    SetAlignment(value *string)
    GetIsLegal() *bool
    SetIsLegal(value *bool)
    GetNumberFormat() *string
    SetNumberFormat(value *string)
    GetNumberPosition() *float64
    SetNumberPosition(value *float64)
    GetNumberStyle() *string
    SetNumberStyle(value *string)
    GetRestartAfterLevel() *int32
    SetRestartAfterLevel(value *int32)
    GetStartAt() *int32
    SetStartAt(value *int32)
    GetTabPosition() *float64
    SetTabPosition(value *float64)
    GetTextPosition() *float64
    SetTextPosition(value *float64)
    GetTrailingCharacter() *string
    SetTrailingCharacter(value *string)
}

type ListLevelUpdate struct {
    // Represents a document list levels.
    Alignment *string

    // Represents a document list levels.
    IsLegal *bool

    // Represents a document list levels.
    NumberFormat *string

    // Represents a document list levels.
    NumberPosition *float64

    // Represents a document list levels.
    NumberStyle *string

    // Represents a document list levels.
    RestartAfterLevel *int32

    // Represents a document list levels.
    StartAt *int32

    // Represents a document list levels.
    TabPosition *float64

    // Represents a document list levels.
    TextPosition *float64

    // Represents a document list levels.
    TrailingCharacter *string
}

func (ListLevelUpdate) IsListLevelUpdate() bool {
    return true
}


func (obj *ListLevelUpdate) Initialize() {
}

func (obj *ListLevelUpdate) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["NumberFormat"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NumberFormat = &parsedValue
        }

    } else if jsonValue, exists := json["numberFormat"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NumberFormat = &parsedValue
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

    if jsonValue, exists := json["NumberStyle"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NumberStyle = &parsedValue
        }

    } else if jsonValue, exists := json["numberStyle"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NumberStyle = &parsedValue
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

    if jsonValue, exists := json["TabPosition"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.TabPosition = &parsedValue
        }

    } else if jsonValue, exists := json["tabPosition"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.TabPosition = &parsedValue
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

    if jsonValue, exists := json["TrailingCharacter"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.TrailingCharacter = &parsedValue
        }

    } else if jsonValue, exists := json["trailingCharacter"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.TrailingCharacter = &parsedValue
        }

    }
}

func (obj *ListLevelUpdate) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
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

func (obj *ListLevelUpdate) GetNumberFormat() *string {
    return obj.NumberFormat
}

func (obj *ListLevelUpdate) SetNumberFormat(value *string) {
    obj.NumberFormat = value
}

func (obj *ListLevelUpdate) GetNumberPosition() *float64 {
    return obj.NumberPosition
}

func (obj *ListLevelUpdate) SetNumberPosition(value *float64) {
    obj.NumberPosition = value
}

func (obj *ListLevelUpdate) GetNumberStyle() *string {
    return obj.NumberStyle
}

func (obj *ListLevelUpdate) SetNumberStyle(value *string) {
    obj.NumberStyle = value
}

func (obj *ListLevelUpdate) GetRestartAfterLevel() *int32 {
    return obj.RestartAfterLevel
}

func (obj *ListLevelUpdate) SetRestartAfterLevel(value *int32) {
    obj.RestartAfterLevel = value
}

func (obj *ListLevelUpdate) GetStartAt() *int32 {
    return obj.StartAt
}

func (obj *ListLevelUpdate) SetStartAt(value *int32) {
    obj.StartAt = value
}

func (obj *ListLevelUpdate) GetTabPosition() *float64 {
    return obj.TabPosition
}

func (obj *ListLevelUpdate) SetTabPosition(value *float64) {
    obj.TabPosition = value
}

func (obj *ListLevelUpdate) GetTextPosition() *float64 {
    return obj.TextPosition
}

func (obj *ListLevelUpdate) SetTextPosition(value *float64) {
    obj.TextPosition = value
}

func (obj *ListLevelUpdate) GetTrailingCharacter() *string {
    return obj.TrailingCharacter
}

func (obj *ListLevelUpdate) SetTrailingCharacter(value *string) {
    obj.TrailingCharacter = value
}

