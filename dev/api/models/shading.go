/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="shading.go">
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

// DTO container with a paragraph format shading element.

type IShading interface {
    IsShading() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetBackgroundPatternColor() IXmlColor
    SetBackgroundPatternColor(value IXmlColor)
    GetForegroundPatternColor() IXmlColor
    SetForegroundPatternColor(value IXmlColor)
    GetTexture() *string
    SetTexture(value *string)
}

type Shading struct {
    // DTO container with a paragraph format shading element.
    BackgroundPatternColor IXmlColor `json:"BackgroundPatternColor,omitempty"`

    // DTO container with a paragraph format shading element.
    ForegroundPatternColor IXmlColor `json:"ForegroundPatternColor,omitempty"`

    // DTO container with a paragraph format shading element.
    Texture *string `json:"Texture,omitempty"`
}

func (Shading) IsShading() bool {
    return true
}


func (obj *Shading) Initialize() {
    if (obj.BackgroundPatternColor != nil) {
        obj.BackgroundPatternColor.Initialize()
    }

    if (obj.ForegroundPatternColor != nil) {
        obj.ForegroundPatternColor.Initialize()
    }


}

func (obj *Shading) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["BackgroundPatternColor"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IXmlColor = new(XmlColor)
            modelInstance.Deserialize(parsedValue)
            obj.BackgroundPatternColor = modelInstance
        }

    } else if jsonValue, exists := json["backgroundPatternColor"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IXmlColor = new(XmlColor)
            modelInstance.Deserialize(parsedValue)
            obj.BackgroundPatternColor = modelInstance
        }

    }

    if jsonValue, exists := json["ForegroundPatternColor"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IXmlColor = new(XmlColor)
            modelInstance.Deserialize(parsedValue)
            obj.ForegroundPatternColor = modelInstance
        }

    } else if jsonValue, exists := json["foregroundPatternColor"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IXmlColor = new(XmlColor)
            modelInstance.Deserialize(parsedValue)
            obj.ForegroundPatternColor = modelInstance
        }

    }

    if jsonValue, exists := json["Texture"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Texture = &parsedValue
        }

    } else if jsonValue, exists := json["texture"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Texture = &parsedValue
        }

    }
}

func (obj *Shading) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *Shading) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.BackgroundPatternColor == nil {
        return errors.New("Property BackgroundPatternColor in Shading is required.")
    }

    if obj.BackgroundPatternColor != nil {
        if err := obj.BackgroundPatternColor.Validate(); err != nil {
            return err
        }
    }

    if obj.ForegroundPatternColor == nil {
        return errors.New("Property ForegroundPatternColor in Shading is required.")
    }

    if obj.ForegroundPatternColor != nil {
        if err := obj.ForegroundPatternColor.Validate(); err != nil {
            return err
        }
    }

    if obj.Texture == nil {
        return errors.New("Property Texture in Shading is required.")
    }

    return nil;
}

func (obj *Shading) GetBackgroundPatternColor() IXmlColor {
    return obj.BackgroundPatternColor
}

func (obj *Shading) SetBackgroundPatternColor(value IXmlColor) {
    obj.BackgroundPatternColor = value
}

func (obj *Shading) GetForegroundPatternColor() IXmlColor {
    return obj.ForegroundPatternColor
}

func (obj *Shading) SetForegroundPatternColor(value IXmlColor) {
    obj.ForegroundPatternColor = value
}

func (obj *Shading) GetTexture() *string {
    return obj.Texture
}

func (obj *Shading) SetTexture(value *string) {
    obj.Texture = value
}

