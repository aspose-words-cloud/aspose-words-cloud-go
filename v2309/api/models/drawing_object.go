/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="drawing_object.go">
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

// DTO container with a DrawingObject.

type IDrawingObject interface {
    IsDrawingObject() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetNodeId() *string
    SetNodeId(value *string)
    GetRenderLinks() []IWordsApiLink
    SetRenderLinks(value []IWordsApiLink)
    GetWidth() *float64
    SetWidth(value *float64)
    GetHeight() *float64
    SetHeight(value *float64)
    GetOleDataLink() IWordsApiLink
    SetOleDataLink(value IWordsApiLink)
    GetImageDataLink() IWordsApiLink
    SetImageDataLink(value IWordsApiLink)
    GetRelativeHorizontalPosition() *string
    SetRelativeHorizontalPosition(value *string)
    GetLeft() *float64
    SetLeft(value *float64)
    GetRelativeVerticalPosition() *string
    SetRelativeVerticalPosition(value *string)
    GetTop() *float64
    SetTop(value *float64)
    GetWrapType() *string
    SetWrapType(value *string)
}

type DrawingObject struct {
    // DTO container with a DrawingObject.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with a DrawingObject.
    NodeId *string `json:"NodeId,omitempty"`

    // DTO container with a DrawingObject.
    RenderLinks []IWordsApiLink `json:"RenderLinks,omitempty"`

    // DTO container with a DrawingObject.
    Width *float64 `json:"Width,omitempty"`

    // DTO container with a DrawingObject.
    Height *float64 `json:"Height,omitempty"`

    // DTO container with a DrawingObject.
    OleDataLink IWordsApiLink `json:"OleDataLink,omitempty"`

    // DTO container with a DrawingObject.
    ImageDataLink IWordsApiLink `json:"ImageDataLink,omitempty"`

    // DTO container with a DrawingObject.
    RelativeHorizontalPosition *string `json:"RelativeHorizontalPosition,omitempty"`

    // DTO container with a DrawingObject.
    Left *float64 `json:"Left,omitempty"`

    // DTO container with a DrawingObject.
    RelativeVerticalPosition *string `json:"RelativeVerticalPosition,omitempty"`

    // DTO container with a DrawingObject.
    Top *float64 `json:"Top,omitempty"`

    // DTO container with a DrawingObject.
    WrapType *string `json:"WrapType,omitempty"`
}

func (DrawingObject) IsDrawingObject() bool {
    return true
}

func (DrawingObject) IsDrawingObjectLink() bool {
    return true
}

func (DrawingObject) IsNodeLink() bool {
    return true
}

func (DrawingObject) IsLinkElement() bool {
    return true
}

func (obj *DrawingObject) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.RenderLinks != nil) {
        for _, objElementRenderLinks := range obj.RenderLinks {
            objElementRenderLinks.Initialize()
        }
    }
    if (obj.OleDataLink != nil) {
        obj.OleDataLink.Initialize()
    }

    if (obj.ImageDataLink != nil) {
        obj.ImageDataLink.Initialize()
    }


}

func (obj *DrawingObject) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["NodeId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NodeId = &parsedValue
        }

    } else if jsonValue, exists := json["nodeId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NodeId = &parsedValue
        }

    }

    if jsonValue, exists := json["RenderLinks"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.RenderLinks = make([]IWordsApiLink, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IWordsApiLink = new(WordsApiLink)
                    modelElementInstance.Deserialize(elementValue)
                    obj.RenderLinks = append(obj.RenderLinks, modelElementInstance)
                }

            }
        }

    } else if jsonValue, exists := json["renderLinks"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.RenderLinks = make([]IWordsApiLink, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(map[string]interface{}); valid {
                    var modelElementInstance IWordsApiLink = new(WordsApiLink)
                    modelElementInstance.Deserialize(elementValue)
                    obj.RenderLinks = append(obj.RenderLinks, modelElementInstance)
                }

            }
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

    if jsonValue, exists := json["OleDataLink"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IWordsApiLink = new(WordsApiLink)
            modelInstance.Deserialize(parsedValue)
            obj.OleDataLink = modelInstance
        }

    } else if jsonValue, exists := json["oleDataLink"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IWordsApiLink = new(WordsApiLink)
            modelInstance.Deserialize(parsedValue)
            obj.OleDataLink = modelInstance
        }

    }

    if jsonValue, exists := json["ImageDataLink"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IWordsApiLink = new(WordsApiLink)
            modelInstance.Deserialize(parsedValue)
            obj.ImageDataLink = modelInstance
        }

    } else if jsonValue, exists := json["imageDataLink"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IWordsApiLink = new(WordsApiLink)
            modelInstance.Deserialize(parsedValue)
            obj.ImageDataLink = modelInstance
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

func (obj *DrawingObject) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *DrawingObject) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *DrawingObject) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *DrawingObject) GetNodeId() *string {
    return obj.NodeId
}

func (obj *DrawingObject) SetNodeId(value *string) {
    obj.NodeId = value
}

func (obj *DrawingObject) GetRenderLinks() []IWordsApiLink {
    return obj.RenderLinks
}

func (obj *DrawingObject) SetRenderLinks(value []IWordsApiLink) {
    obj.RenderLinks = value
}

func (obj *DrawingObject) GetWidth() *float64 {
    return obj.Width
}

func (obj *DrawingObject) SetWidth(value *float64) {
    obj.Width = value
}

func (obj *DrawingObject) GetHeight() *float64 {
    return obj.Height
}

func (obj *DrawingObject) SetHeight(value *float64) {
    obj.Height = value
}

func (obj *DrawingObject) GetOleDataLink() IWordsApiLink {
    return obj.OleDataLink
}

func (obj *DrawingObject) SetOleDataLink(value IWordsApiLink) {
    obj.OleDataLink = value
}

func (obj *DrawingObject) GetImageDataLink() IWordsApiLink {
    return obj.ImageDataLink
}

func (obj *DrawingObject) SetImageDataLink(value IWordsApiLink) {
    obj.ImageDataLink = value
}

func (obj *DrawingObject) GetRelativeHorizontalPosition() *string {
    return obj.RelativeHorizontalPosition
}

func (obj *DrawingObject) SetRelativeHorizontalPosition(value *string) {
    obj.RelativeHorizontalPosition = value
}

func (obj *DrawingObject) GetLeft() *float64 {
    return obj.Left
}

func (obj *DrawingObject) SetLeft(value *float64) {
    obj.Left = value
}

func (obj *DrawingObject) GetRelativeVerticalPosition() *string {
    return obj.RelativeVerticalPosition
}

func (obj *DrawingObject) SetRelativeVerticalPosition(value *string) {
    obj.RelativeVerticalPosition = value
}

func (obj *DrawingObject) GetTop() *float64 {
    return obj.Top
}

func (obj *DrawingObject) SetTop(value *float64) {
    obj.Top = value
}

func (obj *DrawingObject) GetWrapType() *string {
    return obj.WrapType
}

func (obj *DrawingObject) SetWrapType(value *string) {
    obj.WrapType = value
}

