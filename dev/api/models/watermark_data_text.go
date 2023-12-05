/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="watermark_data_text.go">
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

// Class for insert watermark text request building.

type IWatermarkDataText interface {
    IsWatermarkDataText() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetColor() IXmlColor
    SetColor(value IXmlColor)
    GetFontFamily() *string
    SetFontFamily(value *string)
    GetFontSize() *float64
    SetFontSize(value *float64)
    GetIsSemitrasparent() *bool
    SetIsSemitrasparent(value *bool)
    GetLayout() *string
    SetLayout(value *string)
    GetText() *string
    SetText(value *string)
}

type WatermarkDataText struct {
    // Class for insert watermark text request building.
    Color IXmlColor `json:"Color,omitempty"`

    // Class for insert watermark text request building.
    FontFamily *string `json:"FontFamily,omitempty"`

    // Class for insert watermark text request building.
    FontSize *float64 `json:"FontSize,omitempty"`

    // Class for insert watermark text request building.
    IsSemitrasparent *bool `json:"IsSemitrasparent,omitempty"`

    // Class for insert watermark text request building.
    Layout *string `json:"Layout,omitempty"`

    // Class for insert watermark text request building.
    Text *string `json:"Text,omitempty"`
}

func (WatermarkDataText) IsWatermarkDataText() bool {
    return true
}

func (WatermarkDataText) IsWatermarkDataBase() bool {
    return true
}

func (obj *WatermarkDataText) Initialize() {
    if (obj.Color != nil) {
        obj.Color.Initialize()
    }


}

func (obj *WatermarkDataText) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["FontFamily"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FontFamily = &parsedValue
        }

    } else if jsonValue, exists := json["fontFamily"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FontFamily = &parsedValue
        }

    }

    if jsonValue, exists := json["FontSize"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.FontSize = &parsedValue
        }

    } else if jsonValue, exists := json["fontSize"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.FontSize = &parsedValue
        }

    }

    if jsonValue, exists := json["IsSemitrasparent"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsSemitrasparent = &parsedValue
        }

    } else if jsonValue, exists := json["isSemitrasparent"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsSemitrasparent = &parsedValue
        }

    }

    if jsonValue, exists := json["Layout"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Layout = &parsedValue
        }

    } else if jsonValue, exists := json["layout"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Layout = &parsedValue
        }

    }

    if jsonValue, exists := json["Text"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Text = &parsedValue
        }

    } else if jsonValue, exists := json["text"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Text = &parsedValue
        }

    }
}

func (obj *WatermarkDataText) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *WatermarkDataText) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Text == nil {
        return errors.New("Property Text in WatermarkDataText is required.")
    }
    if obj.Color != nil {
        if err := obj.Color.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *WatermarkDataText) GetColor() IXmlColor {
    return obj.Color
}

func (obj *WatermarkDataText) SetColor(value IXmlColor) {
    obj.Color = value
}

func (obj *WatermarkDataText) GetFontFamily() *string {
    return obj.FontFamily
}

func (obj *WatermarkDataText) SetFontFamily(value *string) {
    obj.FontFamily = value
}

func (obj *WatermarkDataText) GetFontSize() *float64 {
    return obj.FontSize
}

func (obj *WatermarkDataText) SetFontSize(value *float64) {
    obj.FontSize = value
}

func (obj *WatermarkDataText) GetIsSemitrasparent() *bool {
    return obj.IsSemitrasparent
}

func (obj *WatermarkDataText) SetIsSemitrasparent(value *bool) {
    obj.IsSemitrasparent = value
}

func (obj *WatermarkDataText) GetLayout() *string {
    return obj.Layout
}

func (obj *WatermarkDataText) SetLayout(value *string) {
    obj.Layout = value
}

func (obj *WatermarkDataText) GetText() *string {
    return obj.Text
}

func (obj *WatermarkDataText) SetText(value *string) {
    obj.Text = value
}

