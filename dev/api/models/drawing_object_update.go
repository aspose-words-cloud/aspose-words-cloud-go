/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="drawing_object_update.go">
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

// Drawing object element for update.

type IDrawingObjectUpdate interface {
    IsDrawingObjectUpdate() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRelativeHorizontalPosition() *string
    SetRelativeHorizontalPosition(value *string)
    GetLeft() *float64
    SetLeft(value *float64)
    GetRelativeVerticalPosition() *string
    SetRelativeVerticalPosition(value *string)
    GetTop() *float64
    SetTop(value *float64)
    GetWidth() *float64
    SetWidth(value *float64)
    GetHeight() *float64
    SetHeight(value *float64)
    GetWrapType() *string
    SetWrapType(value *string)
}

type DrawingObjectUpdate struct {
    // Drawing object element for update.
    RelativeHorizontalPosition *string `json:"RelativeHorizontalPosition,omitempty"`

    // Drawing object element for update.
    Left *float64 `json:"Left,omitempty"`

    // Drawing object element for update.
    RelativeVerticalPosition *string `json:"RelativeVerticalPosition,omitempty"`

    // Drawing object element for update.
    Top *float64 `json:"Top,omitempty"`

    // Drawing object element for update.
    Width *float64 `json:"Width,omitempty"`

    // Drawing object element for update.
    Height *float64 `json:"Height,omitempty"`

    // Drawing object element for update.
    WrapType *string `json:"WrapType,omitempty"`
}

func (DrawingObjectUpdate) IsDrawingObjectUpdate() bool {
    return true
}


func (obj *DrawingObjectUpdate) Initialize() {
}

func (obj *DrawingObjectUpdate) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RelativeHorizontalPosition"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RelativeHorizontalPosition = &parsedValue
        }

    } else if jsonValue, exists := json["relativeHorizontalPosition"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RelativeHorizontalPosition = &parsedValue
        }

    }

    if jsonValue, exists := json["Left"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Left = &parsedValue
        }

    } else if jsonValue, exists := json["left"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Left = &parsedValue
        }

    }

    if jsonValue, exists := json["RelativeVerticalPosition"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RelativeVerticalPosition = &parsedValue
        }

    } else if jsonValue, exists := json["relativeVerticalPosition"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RelativeVerticalPosition = &parsedValue
        }

    }

    if jsonValue, exists := json["Top"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Top = &parsedValue
        }

    } else if jsonValue, exists := json["top"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Top = &parsedValue
        }

    }

    if jsonValue, exists := json["Width"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Width = &parsedValue
        }

    } else if jsonValue, exists := json["width"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Width = &parsedValue
        }

    }

    if jsonValue, exists := json["Height"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Height = &parsedValue
        }

    } else if jsonValue, exists := json["height"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Height = &parsedValue
        }

    }

    if jsonValue, exists := json["WrapType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.WrapType = &parsedValue
        }

    } else if jsonValue, exists := json["wrapType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.WrapType = &parsedValue
        }

    }
}

func (obj *DrawingObjectUpdate) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *DrawingObjectUpdate) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.RelativeHorizontalPosition == nil {
        return errors.New("Property RelativeHorizontalPosition in DrawingObjectUpdate is required.")
    }

    if obj.Left == nil {
        return errors.New("Property Left in DrawingObjectUpdate is required.")
    }

    if obj.RelativeVerticalPosition == nil {
        return errors.New("Property RelativeVerticalPosition in DrawingObjectUpdate is required.")
    }

    if obj.Top == nil {
        return errors.New("Property Top in DrawingObjectUpdate is required.")
    }

    if obj.Width == nil {
        return errors.New("Property Width in DrawingObjectUpdate is required.")
    }

    if obj.Height == nil {
        return errors.New("Property Height in DrawingObjectUpdate is required.")
    }

    if obj.WrapType == nil {
        return errors.New("Property WrapType in DrawingObjectUpdate is required.")
    }

    return nil;
}

func (obj *DrawingObjectUpdate) GetRelativeHorizontalPosition() *string {
    return obj.RelativeHorizontalPosition
}

func (obj *DrawingObjectUpdate) SetRelativeHorizontalPosition(value *string) {
    obj.RelativeHorizontalPosition = value
}

func (obj *DrawingObjectUpdate) GetLeft() *float64 {
    return obj.Left
}

func (obj *DrawingObjectUpdate) SetLeft(value *float64) {
    obj.Left = value
}

func (obj *DrawingObjectUpdate) GetRelativeVerticalPosition() *string {
    return obj.RelativeVerticalPosition
}

func (obj *DrawingObjectUpdate) SetRelativeVerticalPosition(value *string) {
    obj.RelativeVerticalPosition = value
}

func (obj *DrawingObjectUpdate) GetTop() *float64 {
    return obj.Top
}

func (obj *DrawingObjectUpdate) SetTop(value *float64) {
    obj.Top = value
}

func (obj *DrawingObjectUpdate) GetWidth() *float64 {
    return obj.Width
}

func (obj *DrawingObjectUpdate) SetWidth(value *float64) {
    obj.Width = value
}

func (obj *DrawingObjectUpdate) GetHeight() *float64 {
    return obj.Height
}

func (obj *DrawingObjectUpdate) SetHeight(value *float64) {
    obj.Height = value
}

func (obj *DrawingObjectUpdate) GetWrapType() *string {
    return obj.WrapType
}

func (obj *DrawingObjectUpdate) SetWrapType(value *string) {
    obj.WrapType = value
}

