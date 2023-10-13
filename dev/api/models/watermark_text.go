/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="watermark_text.go">
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

type IWatermarkText interface {
    IsWatermarkText() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetRotationAngle() *float64
    SetRotationAngle(value *float64)
    GetText() *string
    SetText(value *string)
}

type WatermarkText struct {
    // Class for insert watermark text request building.
    RotationAngle *float64 `json:"RotationAngle,omitempty"`

    // Class for insert watermark text request building.
    Text *string `json:"Text,omitempty"`
}

func (WatermarkText) IsWatermarkText() bool {
    return true
}


func (obj *WatermarkText) Initialize() {
}

func (obj *WatermarkText) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["RotationAngle"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.RotationAngle = &parsedValue
        }

    } else if jsonValue, exists := json["rotationAngle"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.RotationAngle = &parsedValue
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

func (obj *WatermarkText) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *WatermarkText) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.RotationAngle == nil {
        return errors.New("Property RotationAngle in WatermarkText is required.")
    }

    if obj.Text == nil {
        return errors.New("Property Text in WatermarkText is required.")
    }

    return nil;
}

func (obj *WatermarkText) GetRotationAngle() *float64 {
    return obj.RotationAngle
}

func (obj *WatermarkText) SetRotationAngle(value *float64) {
    obj.RotationAngle = value
}

func (obj *WatermarkText) GetText() *string {
    return obj.Text
}

func (obj *WatermarkText) SetText(value *string) {
    obj.Text = value
}

