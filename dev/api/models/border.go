/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="border.go">
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

// Represents a border of an object.
// Borders can be applied to various document elements including paragraph, run of text inside a paragraph or a table cell.

type IBorder interface {
    IsBorder() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetBorderType() *string
    SetBorderType(value *string)
    GetColor() IXmlColor
    SetColor(value IXmlColor)
    GetDistanceFromText() *float64
    SetDistanceFromText(value *float64)
    GetLineStyle() *string
    SetLineStyle(value *string)
    GetLineWidth() *float64
    SetLineWidth(value *float64)
    GetShadow() *bool
    SetShadow(value *bool)
}

type Border struct {
    // Represents a border of an object.
    // Borders can be applied to various document elements including paragraph, run of text inside a paragraph or a table cell.
    Link IWordsApiLink `json:"Link,omitempty"`

    // Represents a border of an object.
    // Borders can be applied to various document elements including paragraph, run of text inside a paragraph or a table cell.
    BorderType *string `json:"BorderType,omitempty"`

    // Represents a border of an object.
    // Borders can be applied to various document elements including paragraph, run of text inside a paragraph or a table cell.
    Color IXmlColor `json:"Color,omitempty"`

    // Represents a border of an object.
    // Borders can be applied to various document elements including paragraph, run of text inside a paragraph or a table cell.
    DistanceFromText *float64 `json:"DistanceFromText,omitempty"`

    // Represents a border of an object.
    // Borders can be applied to various document elements including paragraph, run of text inside a paragraph or a table cell.
    LineStyle *string `json:"LineStyle,omitempty"`

    // Represents a border of an object.
    // Borders can be applied to various document elements including paragraph, run of text inside a paragraph or a table cell.
    LineWidth *float64 `json:"LineWidth,omitempty"`

    // Represents a border of an object.
    // Borders can be applied to various document elements including paragraph, run of text inside a paragraph or a table cell.
    Shadow *bool `json:"Shadow,omitempty"`
}

func (Border) IsBorder() bool {
    return true
}

func (Border) IsLinkElement() bool {
    return true
}

func (obj *Border) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.Color != nil) {
        obj.Color.Initialize()
    }


}

func (obj *Border) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["BorderType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.BorderType = &parsedValue
        }

    } else if jsonValue, exists := json["borderType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.BorderType = &parsedValue
        }

    }

    if jsonValue, exists := json["Color"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IXmlColor = new(XmlColor)
            modelInstance.Deserialize(parsedValue)
            obj.Color = modelInstance
        }

    } else if jsonValue, exists := json["color"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IXmlColor = new(XmlColor)
            modelInstance.Deserialize(parsedValue)
            obj.Color = modelInstance
        }

    }

    if jsonValue, exists := json["DistanceFromText"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.DistanceFromText = &parsedValue
        }

    } else if jsonValue, exists := json["distanceFromText"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.DistanceFromText = &parsedValue
        }

    }

    if jsonValue, exists := json["LineStyle"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.LineStyle = &parsedValue
        }

    } else if jsonValue, exists := json["lineStyle"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.LineStyle = &parsedValue
        }

    }

    if jsonValue, exists := json["LineWidth"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.LineWidth = &parsedValue
        }

    } else if jsonValue, exists := json["lineWidth"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.LineWidth = &parsedValue
        }

    }

    if jsonValue, exists := json["Shadow"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Shadow = &parsedValue
        }

    } else if jsonValue, exists := json["shadow"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Shadow = &parsedValue
        }

    }
}

func (obj *Border) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *Border) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Link != nil {
        if err := obj.Link.Validate(); err != nil {
            return err
        }
    }
    if obj.Color != nil {
        if err := obj.Color.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *Border) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *Border) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *Border) GetBorderType() *string {
    return obj.BorderType
}

func (obj *Border) SetBorderType(value *string) {
    obj.BorderType = value
}

func (obj *Border) GetColor() IXmlColor {
    return obj.Color
}

func (obj *Border) SetColor(value IXmlColor) {
    obj.Color = value
}

func (obj *Border) GetDistanceFromText() *float64 {
    return obj.DistanceFromText
}

func (obj *Border) SetDistanceFromText(value *float64) {
    obj.DistanceFromText = value
}

func (obj *Border) GetLineStyle() *string {
    return obj.LineStyle
}

func (obj *Border) SetLineStyle(value *string) {
    obj.LineStyle = value
}

func (obj *Border) GetLineWidth() *float64 {
    return obj.LineWidth
}

func (obj *Border) SetLineWidth(value *float64) {
    obj.LineWidth = value
}

func (obj *Border) GetShadow() *bool {
    return obj.Shadow
}

func (obj *Border) SetShadow(value *bool) {
    obj.Shadow = value
}

