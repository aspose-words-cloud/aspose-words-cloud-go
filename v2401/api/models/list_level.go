/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="list_level.go">
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

// DTO container with a document list level.

type IListLevel interface {
    IsListLevel() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
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
    GetFont() IFont
    SetFont(value IFont)
    GetTabPosition() *float64
    SetTabPosition(value *float64)
    GetNumberPosition() *float64
    SetNumberPosition(value *float64)
    GetTextPosition() *float64
    SetTextPosition(value *float64)
    GetLinkedStyle() IStyle
    SetLinkedStyle(value IStyle)
}

type ListLevel struct {
    // DTO container with a document list level.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with a document list level.
    StartAt *int32 `json:"StartAt,omitempty"`

    // DTO container with a document list level.
    NumberStyle *string `json:"NumberStyle,omitempty"`

    // DTO container with a document list level.
    NumberFormat *string `json:"NumberFormat,omitempty"`

    // DTO container with a document list level.
    Alignment *string `json:"Alignment,omitempty"`

    // DTO container with a document list level.
    IsLegal *bool `json:"IsLegal,omitempty"`

    // DTO container with a document list level.
    RestartAfterLevel *int32 `json:"RestartAfterLevel,omitempty"`

    // DTO container with a document list level.
    TrailingCharacter *string `json:"TrailingCharacter,omitempty"`

    // DTO container with a document list level.
    Font IFont `json:"Font,omitempty"`

    // DTO container with a document list level.
    TabPosition *float64 `json:"TabPosition,omitempty"`

    // DTO container with a document list level.
    NumberPosition *float64 `json:"NumberPosition,omitempty"`

    // DTO container with a document list level.
    TextPosition *float64 `json:"TextPosition,omitempty"`

    // DTO container with a document list level.
    LinkedStyle IStyle `json:"LinkedStyle,omitempty"`
}

func (ListLevel) IsListLevel() bool {
    return true
}

func (ListLevel) IsLinkElement() bool {
    return true
}

func (obj *ListLevel) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.Font != nil) {
        obj.Font.Initialize()
    }

    if (obj.LinkedStyle != nil) {
        obj.LinkedStyle.Initialize()
    }


}

func (obj *ListLevel) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Link"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IWordsApiLink = new(WordsApiLink)
            modelInstance.Deserialize(parsedValue)
            obj.Link = modelInstance
        }

    } else if jsonValue, exists := json["link"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IWordsApiLink = new(WordsApiLink)
            modelInstance.Deserialize(parsedValue)
            obj.Link = modelInstance
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

    if jsonValue, exists := json["Font"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFont = new(Font)
            modelInstance.Deserialize(parsedValue)
            obj.Font = modelInstance
        }

    } else if jsonValue, exists := json["font"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFont = new(Font)
            modelInstance.Deserialize(parsedValue)
            obj.Font = modelInstance
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

    if jsonValue, exists := json["LinkedStyle"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IStyle = new(Style)
            modelInstance.Deserialize(parsedValue)
            obj.LinkedStyle = modelInstance
        }

    } else if jsonValue, exists := json["linkedStyle"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IStyle = new(Style)
            modelInstance.Deserialize(parsedValue)
            obj.LinkedStyle = modelInstance
        }

    }
}

func (obj *ListLevel) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *ListLevel) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.StartAt == nil {
        return errors.New("Property StartAt in ListLevel is required.")
    }
    if obj.NumberStyle == nil {
        return errors.New("Property NumberStyle in ListLevel is required.")
    }
    if obj.Alignment == nil {
        return errors.New("Property Alignment in ListLevel is required.")
    }
    if obj.IsLegal == nil {
        return errors.New("Property IsLegal in ListLevel is required.")
    }
    if obj.RestartAfterLevel == nil {
        return errors.New("Property RestartAfterLevel in ListLevel is required.")
    }
    if obj.TrailingCharacter == nil {
        return errors.New("Property TrailingCharacter in ListLevel is required.")
    }
    if obj.TabPosition == nil {
        return errors.New("Property TabPosition in ListLevel is required.")
    }
    if obj.NumberPosition == nil {
        return errors.New("Property NumberPosition in ListLevel is required.")
    }
    if obj.TextPosition == nil {
        return errors.New("Property TextPosition in ListLevel is required.")
    }
    if obj.Link != nil {
        if err := obj.Link.Validate(); err != nil {
            return err
        }
    }
    if obj.Font != nil {
        if err := obj.Font.Validate(); err != nil {
            return err
        }
    }
    if obj.LinkedStyle != nil {
        if err := obj.LinkedStyle.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *ListLevel) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *ListLevel) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *ListLevel) GetStartAt() *int32 {
    return obj.StartAt
}

func (obj *ListLevel) SetStartAt(value *int32) {
    obj.StartAt = value
}

func (obj *ListLevel) GetNumberStyle() *string {
    return obj.NumberStyle
}

func (obj *ListLevel) SetNumberStyle(value *string) {
    obj.NumberStyle = value
}

func (obj *ListLevel) GetNumberFormat() *string {
    return obj.NumberFormat
}

func (obj *ListLevel) SetNumberFormat(value *string) {
    obj.NumberFormat = value
}

func (obj *ListLevel) GetAlignment() *string {
    return obj.Alignment
}

func (obj *ListLevel) SetAlignment(value *string) {
    obj.Alignment = value
}

func (obj *ListLevel) GetIsLegal() *bool {
    return obj.IsLegal
}

func (obj *ListLevel) SetIsLegal(value *bool) {
    obj.IsLegal = value
}

func (obj *ListLevel) GetRestartAfterLevel() *int32 {
    return obj.RestartAfterLevel
}

func (obj *ListLevel) SetRestartAfterLevel(value *int32) {
    obj.RestartAfterLevel = value
}

func (obj *ListLevel) GetTrailingCharacter() *string {
    return obj.TrailingCharacter
}

func (obj *ListLevel) SetTrailingCharacter(value *string) {
    obj.TrailingCharacter = value
}

func (obj *ListLevel) GetFont() IFont {
    return obj.Font
}

func (obj *ListLevel) SetFont(value IFont) {
    obj.Font = value
}

func (obj *ListLevel) GetTabPosition() *float64 {
    return obj.TabPosition
}

func (obj *ListLevel) SetTabPosition(value *float64) {
    obj.TabPosition = value
}

func (obj *ListLevel) GetNumberPosition() *float64 {
    return obj.NumberPosition
}

func (obj *ListLevel) SetNumberPosition(value *float64) {
    obj.NumberPosition = value
}

func (obj *ListLevel) GetTextPosition() *float64 {
    return obj.TextPosition
}

func (obj *ListLevel) SetTextPosition(value *float64) {
    obj.TextPosition = value
}

func (obj *ListLevel) GetLinkedStyle() IStyle {
    return obj.LinkedStyle
}

func (obj *ListLevel) SetLinkedStyle(value IStyle) {
    obj.LinkedStyle = value
}

