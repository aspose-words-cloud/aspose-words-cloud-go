/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="metafile_rendering_options_data.go">
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

// Container class for options of metafile rendering.

type IMetafileRenderingOptionsData interface {
    IsMetafileRenderingOptionsData() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetEmfPlusDualRenderingMode() *string
    SetEmfPlusDualRenderingMode(value *string)
    GetEmulateRasterOperations() *bool
    SetEmulateRasterOperations(value *bool)
    GetEmulateRenderingToSizeOnPage() *bool
    SetEmulateRenderingToSizeOnPage(value *bool)
    GetEmulateRenderingToSizeOnPageResolution() *int32
    SetEmulateRenderingToSizeOnPageResolution(value *int32)
    GetRenderingMode() *string
    SetRenderingMode(value *string)
    GetUseEmfEmbeddedToWmf() *bool
    SetUseEmfEmbeddedToWmf(value *bool)
}

type MetafileRenderingOptionsData struct {
    // Container class for options of metafile rendering.
    EmfPlusDualRenderingMode *string `json:"EmfPlusDualRenderingMode,omitempty"`

    // Container class for options of metafile rendering.
    EmulateRasterOperations *bool `json:"EmulateRasterOperations,omitempty"`

    // Container class for options of metafile rendering.
    EmulateRenderingToSizeOnPage *bool `json:"EmulateRenderingToSizeOnPage,omitempty"`

    // Container class for options of metafile rendering.
    EmulateRenderingToSizeOnPageResolution *int32 `json:"EmulateRenderingToSizeOnPageResolution,omitempty"`

    // Container class for options of metafile rendering.
    RenderingMode *string `json:"RenderingMode,omitempty"`

    // Container class for options of metafile rendering.
    UseEmfEmbeddedToWmf *bool `json:"UseEmfEmbeddedToWmf,omitempty"`
}

func (MetafileRenderingOptionsData) IsMetafileRenderingOptionsData() bool {
    return true
}


func (obj *MetafileRenderingOptionsData) Initialize() {
}

func (obj *MetafileRenderingOptionsData) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["EmfPlusDualRenderingMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.EmfPlusDualRenderingMode = &parsedValue
        }

    } else if jsonValue, exists := json["emfPlusDualRenderingMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.EmfPlusDualRenderingMode = &parsedValue
        }

    }

    if jsonValue, exists := json["EmulateRasterOperations"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.EmulateRasterOperations = &parsedValue
        }

    } else if jsonValue, exists := json["emulateRasterOperations"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.EmulateRasterOperations = &parsedValue
        }

    }

    if jsonValue, exists := json["EmulateRenderingToSizeOnPage"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.EmulateRenderingToSizeOnPage = &parsedValue
        }

    } else if jsonValue, exists := json["emulateRenderingToSizeOnPage"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.EmulateRenderingToSizeOnPage = &parsedValue
        }

    }

    if jsonValue, exists := json["EmulateRenderingToSizeOnPageResolution"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.EmulateRenderingToSizeOnPageResolution = new(int32)
            *obj.EmulateRenderingToSizeOnPageResolution = int32(parsedValue)
        }

    } else if jsonValue, exists := json["emulateRenderingToSizeOnPageResolution"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.EmulateRenderingToSizeOnPageResolution = new(int32)
            *obj.EmulateRenderingToSizeOnPageResolution = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["RenderingMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RenderingMode = &parsedValue
        }

    } else if jsonValue, exists := json["renderingMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.RenderingMode = &parsedValue
        }

    }

    if jsonValue, exists := json["UseEmfEmbeddedToWmf"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UseEmfEmbeddedToWmf = &parsedValue
        }

    } else if jsonValue, exists := json["useEmfEmbeddedToWmf"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UseEmfEmbeddedToWmf = &parsedValue
        }

    }
}

func (obj *MetafileRenderingOptionsData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *MetafileRenderingOptionsData) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    return nil;
}

func (obj *MetafileRenderingOptionsData) GetEmfPlusDualRenderingMode() *string {
    return obj.EmfPlusDualRenderingMode
}

func (obj *MetafileRenderingOptionsData) SetEmfPlusDualRenderingMode(value *string) {
    obj.EmfPlusDualRenderingMode = value
}

func (obj *MetafileRenderingOptionsData) GetEmulateRasterOperations() *bool {
    return obj.EmulateRasterOperations
}

func (obj *MetafileRenderingOptionsData) SetEmulateRasterOperations(value *bool) {
    obj.EmulateRasterOperations = value
}

func (obj *MetafileRenderingOptionsData) GetEmulateRenderingToSizeOnPage() *bool {
    return obj.EmulateRenderingToSizeOnPage
}

func (obj *MetafileRenderingOptionsData) SetEmulateRenderingToSizeOnPage(value *bool) {
    obj.EmulateRenderingToSizeOnPage = value
}

func (obj *MetafileRenderingOptionsData) GetEmulateRenderingToSizeOnPageResolution() *int32 {
    return obj.EmulateRenderingToSizeOnPageResolution
}

func (obj *MetafileRenderingOptionsData) SetEmulateRenderingToSizeOnPageResolution(value *int32) {
    obj.EmulateRenderingToSizeOnPageResolution = value
}

func (obj *MetafileRenderingOptionsData) GetRenderingMode() *string {
    return obj.RenderingMode
}

func (obj *MetafileRenderingOptionsData) SetRenderingMode(value *string) {
    obj.RenderingMode = value
}

func (obj *MetafileRenderingOptionsData) GetUseEmfEmbeddedToWmf() *bool {
    return obj.UseEmfEmbeddedToWmf
}

func (obj *MetafileRenderingOptionsData) SetUseEmfEmbeddedToWmf(value *bool) {
    obj.UseEmfEmbeddedToWmf = value
}

