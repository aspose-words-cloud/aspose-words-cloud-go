/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="drawing_object_insert.go">
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

// Drawing object element for insert.

type IDrawingObjectInsert interface {
    IsDrawingObjectInsert() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetHeight() *float64
    SetHeight(value *float64)
    GetLeft() *float64
    SetLeft(value *float64)
    GetPosition() INewDocumentPosition
    SetPosition(value INewDocumentPosition)
    GetRelativeHorizontalPosition() *string
    SetRelativeHorizontalPosition(value *string)
    GetRelativeVerticalPosition() *string
    SetRelativeVerticalPosition(value *string)
    GetTop() *float64
    SetTop(value *float64)
    GetWidth() *float64
    SetWidth(value *float64)
    GetWrapType() *string
    SetWrapType(value *string)
}

type DrawingObjectInsert struct {
    // Drawing object element for insert.
    Height *float64

    // Drawing object element for insert.
    Left *float64

    // Drawing object element for insert.
    Position INewDocumentPosition

    // Drawing object element for insert.
    RelativeHorizontalPosition *string

    // Drawing object element for insert.
    RelativeVerticalPosition *string

    // Drawing object element for insert.
    Top *float64

    // Drawing object element for insert.
    Width *float64

    // Drawing object element for insert.
    WrapType *string
}

func (DrawingObjectInsert) IsDrawingObjectInsert() bool {
    return true
}


func (obj *DrawingObjectInsert) Initialize() {
    if (obj.Position != nil) {
        obj.Position.Initialize()
    }


}

func (obj *DrawingObjectInsert) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Height"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Height = &parsedValue
        }

    } else if jsonValue, exists := json["height"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Height = &parsedValue
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

    if jsonValue, exists := json["RelativeHorizontalPosition"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RelativeHorizontalPosition = &parsedValue
        }

    } else if jsonValue, exists := json["relativeHorizontalPosition"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RelativeHorizontalPosition = &parsedValue
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

func (obj *DrawingObjectInsert) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *DrawingObjectInsert) GetHeight() *float64 {
    return obj.Height
}

func (obj *DrawingObjectInsert) SetHeight(value *float64) {
    obj.Height = value
}

func (obj *DrawingObjectInsert) GetLeft() *float64 {
    return obj.Left
}

func (obj *DrawingObjectInsert) SetLeft(value *float64) {
    obj.Left = value
}

func (obj *DrawingObjectInsert) GetPosition() INewDocumentPosition {
    return obj.Position
}

func (obj *DrawingObjectInsert) SetPosition(value INewDocumentPosition) {
    obj.Position = value
}

func (obj *DrawingObjectInsert) GetRelativeHorizontalPosition() *string {
    return obj.RelativeHorizontalPosition
}

func (obj *DrawingObjectInsert) SetRelativeHorizontalPosition(value *string) {
    obj.RelativeHorizontalPosition = value
}

func (obj *DrawingObjectInsert) GetRelativeVerticalPosition() *string {
    return obj.RelativeVerticalPosition
}

func (obj *DrawingObjectInsert) SetRelativeVerticalPosition(value *string) {
    obj.RelativeVerticalPosition = value
}

func (obj *DrawingObjectInsert) GetTop() *float64 {
    return obj.Top
}

func (obj *DrawingObjectInsert) SetTop(value *float64) {
    obj.Top = value
}

func (obj *DrawingObjectInsert) GetWidth() *float64 {
    return obj.Width
}

func (obj *DrawingObjectInsert) SetWidth(value *float64) {
    obj.Width = value
}

func (obj *DrawingObjectInsert) GetWrapType() *string {
    return obj.WrapType
}

func (obj *DrawingObjectInsert) SetWrapType(value *string) {
    obj.WrapType = value
}

