/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="downsample_options_data.go">
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

// Container class for Downsample options.

type IDownsampleOptionsData interface {
    IsDownsampleOptionsData() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetDownsampleImages() *bool
    SetDownsampleImages(value *bool)
    GetResolution() *int32
    SetResolution(value *int32)
    GetResolutionThreshold() *int32
    SetResolutionThreshold(value *int32)
}

type DownsampleOptionsData struct {
    // Container class for Downsample options.
    DownsampleImages *bool `json:"DownsampleImages,omitempty"`

    // Container class for Downsample options.
    Resolution *int32 `json:"Resolution,omitempty"`

    // Container class for Downsample options.
    ResolutionThreshold *int32 `json:"ResolutionThreshold,omitempty"`
}

func (DownsampleOptionsData) IsDownsampleOptionsData() bool {
    return true
}


func (obj *DownsampleOptionsData) Initialize() {
}

func (obj *DownsampleOptionsData) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["DownsampleImages"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.DownsampleImages = &parsedValue
        }

    } else if jsonValue, exists := json["downsampleImages"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.DownsampleImages = &parsedValue
        }

    }

    if jsonValue, exists := json["Resolution"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Resolution = new(int32)
            *obj.Resolution = int32(parsedValue)
        }

    } else if jsonValue, exists := json["resolution"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Resolution = new(int32)
            *obj.Resolution = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["ResolutionThreshold"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ResolutionThreshold = new(int32)
            *obj.ResolutionThreshold = int32(parsedValue)
        }

    } else if jsonValue, exists := json["resolutionThreshold"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ResolutionThreshold = new(int32)
            *obj.ResolutionThreshold = int32(parsedValue)
        }

    }
}

func (obj *DownsampleOptionsData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *DownsampleOptionsData) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    return nil;
}

func (obj *DownsampleOptionsData) GetDownsampleImages() *bool {
    return obj.DownsampleImages
}

func (obj *DownsampleOptionsData) SetDownsampleImages(value *bool) {
    obj.DownsampleImages = value
}

func (obj *DownsampleOptionsData) GetResolution() *int32 {
    return obj.Resolution
}

func (obj *DownsampleOptionsData) SetResolution(value *int32) {
    obj.Resolution = value
}

func (obj *DownsampleOptionsData) GetResolutionThreshold() *int32 {
    return obj.ResolutionThreshold
}

func (obj *DownsampleOptionsData) SetResolutionThreshold(value *int32) {
    obj.ResolutionThreshold = value
}

