/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="watermark_data_image.go">
 *   Copyright (c) 2026 Aspose.Words for Cloud
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

// Class for insert watermark image request building.

type IWatermarkDataImage interface {
    IsWatermarkDataImage() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetImage() IFileReference
    SetImage(value IFileReference)
    GetIsWashout() *bool
    SetIsWashout(value *bool)
    GetScale() *float64
    SetScale(value *float64)
}

type WatermarkDataImage struct {
    // Class for insert watermark image request building.
    Image IFileReference `json:"Image,omitempty"`

    // Class for insert watermark image request building.
    IsWashout *bool `json:"IsWashout,omitempty"`

    // Class for insert watermark image request building.
    Scale *float64 `json:"Scale,omitempty"`
}

func (WatermarkDataImage) IsWatermarkDataImage() bool {
    return true
}

func (WatermarkDataImage) IsWatermarkDataBase() bool {
    return true
}

func (obj *WatermarkDataImage) Initialize() {
    if (obj.Image != nil) {
        obj.Image.Initialize()
    }


}

func (obj *WatermarkDataImage) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Image"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFileReference = new(FileReference)
            modelInstance.Deserialize(parsedValue)
            obj.Image = modelInstance
        }

    } else if jsonValue, exists := json["image"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IFileReference = new(FileReference)
            modelInstance.Deserialize(parsedValue)
            obj.Image = modelInstance
        }

    }

    if jsonValue, exists := json["IsWashout"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsWashout = &parsedValue
        }

    } else if jsonValue, exists := json["isWashout"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsWashout = &parsedValue
        }

    }

    if jsonValue, exists := json["Scale"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Scale = &parsedValue
        }

    } else if jsonValue, exists := json["scale"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Scale = &parsedValue
        }

    }
}

func (obj *WatermarkDataImage) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    if (obj.Image != nil) {
        resultFilesContent = obj.Image.CollectFilesContent(resultFilesContent)
    }



    return resultFilesContent
}

func (obj *WatermarkDataImage) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Image == nil {
        return errors.New("Property Image in WatermarkDataImage is required.")
    }
    if obj.Image != nil {
        if err := obj.Image.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *WatermarkDataImage) GetImage() IFileReference {
    return obj.Image
}

func (obj *WatermarkDataImage) SetImage(value IFileReference) {
    obj.Image = value
}

func (obj *WatermarkDataImage) GetIsWashout() *bool {
    return obj.IsWashout
}

func (obj *WatermarkDataImage) SetIsWashout(value *bool) {
    obj.IsWashout = value
}

func (obj *WatermarkDataImage) GetScale() *float64 {
    return obj.Scale
}

func (obj *WatermarkDataImage) SetScale(value *float64) {
    obj.Scale = value
}

