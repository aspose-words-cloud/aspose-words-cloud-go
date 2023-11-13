/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="gif_save_options_data.go">
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

// Container class for gif save options.

type IGifSaveOptionsData interface {
    IsGifSaveOptionsData() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetAllowEmbeddingPostScriptFonts() *bool
    SetAllowEmbeddingPostScriptFonts(value *bool)
    GetCustomTimeZoneInfoData() ITimeZoneInfoData
    SetCustomTimeZoneInfoData(value ITimeZoneInfoData)
    GetDml3DEffectsRenderingMode() *string
    SetDml3DEffectsRenderingMode(value *string)
    GetDmlEffectsRenderingMode() *string
    SetDmlEffectsRenderingMode(value *string)
    GetDmlRenderingMode() *string
    SetDmlRenderingMode(value *string)
    GetFileName() *string
    SetFileName(value *string)
    GetImlRenderingMode() *string
    SetImlRenderingMode(value *string)
    GetUpdateCreatedTimeProperty() *bool
    SetUpdateCreatedTimeProperty(value *bool)
    GetUpdateFields() *bool
    SetUpdateFields(value *bool)
    GetUpdateLastPrintedProperty() *bool
    SetUpdateLastPrintedProperty(value *bool)
    GetUpdateLastSavedTimeProperty() *bool
    SetUpdateLastSavedTimeProperty(value *bool)
    GetZipOutput() *bool
    SetZipOutput(value *bool)
    GetColorMode() *string
    SetColorMode(value *string)
    GetJpegQuality() *int32
    SetJpegQuality(value *int32)
    GetMetafileRenderingOptions() IMetafileRenderingOptionsData
    SetMetafileRenderingOptions(value IMetafileRenderingOptionsData)
    GetNumeralFormat() *string
    SetNumeralFormat(value *string)
    GetOptimizeOutput() *bool
    SetOptimizeOutput(value *bool)
    GetPageCount() *int32
    SetPageCount(value *int32)
    GetPageIndex() *int32
    SetPageIndex(value *int32)
    GetHorizontalResolution() *float64
    SetHorizontalResolution(value *float64)
    GetImageBrightness() *float64
    SetImageBrightness(value *float64)
    GetImageColorMode() *string
    SetImageColorMode(value *string)
    GetImageContrast() *float64
    SetImageContrast(value *float64)
    GetPaperColor() *string
    SetPaperColor(value *string)
    GetPixelFormat() *string
    SetPixelFormat(value *string)
    GetResolution() *float64
    SetResolution(value *float64)
    GetScale() *float64
    SetScale(value *float64)
    GetUseAntiAliasing() *bool
    SetUseAntiAliasing(value *bool)
    GetUseHighQualityRendering() *bool
    SetUseHighQualityRendering(value *bool)
    GetVerticalResolution() *float64
    SetVerticalResolution(value *float64)
    GetImageHeight() *int32
    SetImageHeight(value *int32)
    GetImageWidth() *int32
    SetImageWidth(value *int32)
    GetUseGdiEmfRenderer() *bool
    SetUseGdiEmfRenderer(value *bool)
    GetSaveFormat() *string
    SetSaveFormat(value *string)
}

type GifSaveOptionsData struct {
    // Container class for gif save options.
    AllowEmbeddingPostScriptFonts *bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for gif save options.
    CustomTimeZoneInfoData ITimeZoneInfoData `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for gif save options.
    Dml3DEffectsRenderingMode *string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for gif save options.
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for gif save options.
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // Container class for gif save options.
    FileName *string `json:"FileName,omitempty"`

    // Container class for gif save options.
    ImlRenderingMode *string `json:"ImlRenderingMode,omitempty"`

    // Container class for gif save options.
    UpdateCreatedTimeProperty *bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for gif save options.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // Container class for gif save options.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for gif save options.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for gif save options.
    ZipOutput *bool `json:"ZipOutput,omitempty"`

    // Container class for gif save options.
    ColorMode *string `json:"ColorMode,omitempty"`

    // Container class for gif save options.
    JpegQuality *int32 `json:"JpegQuality,omitempty"`

    // Container class for gif save options.
    MetafileRenderingOptions IMetafileRenderingOptionsData `json:"MetafileRenderingOptions,omitempty"`

    // Container class for gif save options.
    NumeralFormat *string `json:"NumeralFormat,omitempty"`

    // Container class for gif save options.
    OptimizeOutput *bool `json:"OptimizeOutput,omitempty"`

    // Container class for gif save options.
    PageCount *int32 `json:"PageCount,omitempty"`

    // Container class for gif save options.
    PageIndex *int32 `json:"PageIndex,omitempty"`

    // Container class for gif save options.
    HorizontalResolution *float64 `json:"HorizontalResolution,omitempty"`

    // Container class for gif save options.
    ImageBrightness *float64 `json:"ImageBrightness,omitempty"`

    // Container class for gif save options.
    ImageColorMode *string `json:"ImageColorMode,omitempty"`

    // Container class for gif save options.
    ImageContrast *float64 `json:"ImageContrast,omitempty"`

    // Container class for gif save options.
    PaperColor *string `json:"PaperColor,omitempty"`

    // Container class for gif save options.
    PixelFormat *string `json:"PixelFormat,omitempty"`

    // Container class for gif save options.
    Resolution *float64 `json:"Resolution,omitempty"`

    // Container class for gif save options.
    Scale *float64 `json:"Scale,omitempty"`

    // Container class for gif save options.
    UseAntiAliasing *bool `json:"UseAntiAliasing,omitempty"`

    // Container class for gif save options.
    UseHighQualityRendering *bool `json:"UseHighQualityRendering,omitempty"`

    // Container class for gif save options.
    VerticalResolution *float64 `json:"VerticalResolution,omitempty"`

    // Container class for gif save options.
    ImageHeight *int32 `json:"ImageHeight,omitempty"`

    // Container class for gif save options.
    ImageWidth *int32 `json:"ImageWidth,omitempty"`

    // Container class for gif save options.
    UseGdiEmfRenderer *bool `json:"UseGdiEmfRenderer,omitempty"`

    // Container class for gif save options.
    SaveFormat *string `json:"SaveFormat,omitempty"`
}

func (GifSaveOptionsData) IsGifSaveOptionsData() bool {
    return true
}

func (GifSaveOptionsData) IsImageSaveOptionsData() bool {
    return true
}

func (GifSaveOptionsData) IsFixedPageSaveOptionsData() bool {
    return true
}

func (GifSaveOptionsData) IsSaveOptionsData() bool {
    return true
}

func (obj *GifSaveOptionsData) Initialize() {
    var _SaveFormat = "gif"
    obj.SaveFormat = &_SaveFormat


    if (obj.CustomTimeZoneInfoData != nil) {
        obj.CustomTimeZoneInfoData.Initialize()
    }

    if (obj.MetafileRenderingOptions != nil) {
        obj.MetafileRenderingOptions.Initialize()
    }


}

func (obj *GifSaveOptionsData) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["AllowEmbeddingPostScriptFonts"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.AllowEmbeddingPostScriptFonts = &parsedValue
        }

    } else if jsonValue, exists := json["allowEmbeddingPostScriptFonts"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.AllowEmbeddingPostScriptFonts = &parsedValue
        }

    }

    if jsonValue, exists := json["CustomTimeZoneInfoData"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ITimeZoneInfoData = new(TimeZoneInfoData)
            modelInstance.Deserialize(parsedValue)
            obj.CustomTimeZoneInfoData = modelInstance
        }

    } else if jsonValue, exists := json["customTimeZoneInfoData"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ITimeZoneInfoData = new(TimeZoneInfoData)
            modelInstance.Deserialize(parsedValue)
            obj.CustomTimeZoneInfoData = modelInstance
        }

    }

    if jsonValue, exists := json["Dml3DEffectsRenderingMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Dml3DEffectsRenderingMode = &parsedValue
        }

    } else if jsonValue, exists := json["dml3DEffectsRenderingMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Dml3DEffectsRenderingMode = &parsedValue
        }

    }

    if jsonValue, exists := json["DmlEffectsRenderingMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DmlEffectsRenderingMode = &parsedValue
        }

    } else if jsonValue, exists := json["dmlEffectsRenderingMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DmlEffectsRenderingMode = &parsedValue
        }

    }

    if jsonValue, exists := json["DmlRenderingMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DmlRenderingMode = &parsedValue
        }

    } else if jsonValue, exists := json["dmlRenderingMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DmlRenderingMode = &parsedValue
        }

    }

    if jsonValue, exists := json["FileName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FileName = &parsedValue
        }

    } else if jsonValue, exists := json["fileName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FileName = &parsedValue
        }

    }

    if jsonValue, exists := json["ImlRenderingMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ImlRenderingMode = &parsedValue
        }

    } else if jsonValue, exists := json["imlRenderingMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ImlRenderingMode = &parsedValue
        }

    }

    if jsonValue, exists := json["UpdateCreatedTimeProperty"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UpdateCreatedTimeProperty = &parsedValue
        }

    } else if jsonValue, exists := json["updateCreatedTimeProperty"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UpdateCreatedTimeProperty = &parsedValue
        }

    }

    if jsonValue, exists := json["UpdateFields"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UpdateFields = &parsedValue
        }

    } else if jsonValue, exists := json["updateFields"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UpdateFields = &parsedValue
        }

    }

    if jsonValue, exists := json["UpdateLastPrintedProperty"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UpdateLastPrintedProperty = &parsedValue
        }

    } else if jsonValue, exists := json["updateLastPrintedProperty"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UpdateLastPrintedProperty = &parsedValue
        }

    }

    if jsonValue, exists := json["UpdateLastSavedTimeProperty"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UpdateLastSavedTimeProperty = &parsedValue
        }

    } else if jsonValue, exists := json["updateLastSavedTimeProperty"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UpdateLastSavedTimeProperty = &parsedValue
        }

    }

    if jsonValue, exists := json["ZipOutput"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ZipOutput = &parsedValue
        }

    } else if jsonValue, exists := json["zipOutput"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ZipOutput = &parsedValue
        }

    }

    if jsonValue, exists := json["ColorMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ColorMode = &parsedValue
        }

    } else if jsonValue, exists := json["colorMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ColorMode = &parsedValue
        }

    }

    if jsonValue, exists := json["JpegQuality"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.JpegQuality = new(int32)
            *obj.JpegQuality = int32(parsedValue)
        }

    } else if jsonValue, exists := json["jpegQuality"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.JpegQuality = new(int32)
            *obj.JpegQuality = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["MetafileRenderingOptions"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IMetafileRenderingOptionsData = new(MetafileRenderingOptionsData)
            modelInstance.Deserialize(parsedValue)
            obj.MetafileRenderingOptions = modelInstance
        }

    } else if jsonValue, exists := json["metafileRenderingOptions"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IMetafileRenderingOptionsData = new(MetafileRenderingOptionsData)
            modelInstance.Deserialize(parsedValue)
            obj.MetafileRenderingOptions = modelInstance
        }

    }

    if jsonValue, exists := json["NumeralFormat"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NumeralFormat = &parsedValue
        }

    } else if jsonValue, exists := json["numeralFormat"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NumeralFormat = &parsedValue
        }

    }

    if jsonValue, exists := json["OptimizeOutput"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.OptimizeOutput = &parsedValue
        }

    } else if jsonValue, exists := json["optimizeOutput"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.OptimizeOutput = &parsedValue
        }

    }

    if jsonValue, exists := json["PageCount"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.PageCount = new(int32)
            *obj.PageCount = int32(parsedValue)
        }

    } else if jsonValue, exists := json["pageCount"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.PageCount = new(int32)
            *obj.PageCount = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["PageIndex"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.PageIndex = new(int32)
            *obj.PageIndex = int32(parsedValue)
        }

    } else if jsonValue, exists := json["pageIndex"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.PageIndex = new(int32)
            *obj.PageIndex = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["HorizontalResolution"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.HorizontalResolution = &parsedValue
        }

    } else if jsonValue, exists := json["horizontalResolution"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.HorizontalResolution = &parsedValue
        }

    }

    if jsonValue, exists := json["ImageBrightness"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ImageBrightness = &parsedValue
        }

    } else if jsonValue, exists := json["imageBrightness"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ImageBrightness = &parsedValue
        }

    }

    if jsonValue, exists := json["ImageColorMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ImageColorMode = &parsedValue
        }

    } else if jsonValue, exists := json["imageColorMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ImageColorMode = &parsedValue
        }

    }

    if jsonValue, exists := json["ImageContrast"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ImageContrast = &parsedValue
        }

    } else if jsonValue, exists := json["imageContrast"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ImageContrast = &parsedValue
        }

    }

    if jsonValue, exists := json["PaperColor"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.PaperColor = &parsedValue
        }

    } else if jsonValue, exists := json["paperColor"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.PaperColor = &parsedValue
        }

    }

    if jsonValue, exists := json["PixelFormat"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.PixelFormat = &parsedValue
        }

    } else if jsonValue, exists := json["pixelFormat"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.PixelFormat = &parsedValue
        }

    }

    if jsonValue, exists := json["Resolution"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Resolution = &parsedValue
        }

    } else if jsonValue, exists := json["resolution"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.Resolution = &parsedValue
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

    if jsonValue, exists := json["UseAntiAliasing"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UseAntiAliasing = &parsedValue
        }

    } else if jsonValue, exists := json["useAntiAliasing"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UseAntiAliasing = &parsedValue
        }

    }

    if jsonValue, exists := json["UseHighQualityRendering"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UseHighQualityRendering = &parsedValue
        }

    } else if jsonValue, exists := json["useHighQualityRendering"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UseHighQualityRendering = &parsedValue
        }

    }

    if jsonValue, exists := json["VerticalResolution"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.VerticalResolution = &parsedValue
        }

    } else if jsonValue, exists := json["verticalResolution"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.VerticalResolution = &parsedValue
        }

    }

    if jsonValue, exists := json["ImageHeight"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ImageHeight = new(int32)
            *obj.ImageHeight = int32(parsedValue)
        }

    } else if jsonValue, exists := json["imageHeight"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ImageHeight = new(int32)
            *obj.ImageHeight = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["ImageWidth"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ImageWidth = new(int32)
            *obj.ImageWidth = int32(parsedValue)
        }

    } else if jsonValue, exists := json["imageWidth"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ImageWidth = new(int32)
            *obj.ImageWidth = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["UseGdiEmfRenderer"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UseGdiEmfRenderer = &parsedValue
        }

    } else if jsonValue, exists := json["useGdiEmfRenderer"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UseGdiEmfRenderer = &parsedValue
        }

    }
}

func (obj *GifSaveOptionsData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *GifSaveOptionsData) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.FileName == nil {
        return errors.New("Property FileName in GifSaveOptionsData is required.")
    }
    if obj.CustomTimeZoneInfoData != nil {
        if err := obj.CustomTimeZoneInfoData.Validate(); err != nil {
            return err
        }
    }
    if obj.MetafileRenderingOptions != nil {
        if err := obj.MetafileRenderingOptions.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *GifSaveOptionsData) GetAllowEmbeddingPostScriptFonts() *bool {
    return obj.AllowEmbeddingPostScriptFonts
}

func (obj *GifSaveOptionsData) SetAllowEmbeddingPostScriptFonts(value *bool) {
    obj.AllowEmbeddingPostScriptFonts = value
}

func (obj *GifSaveOptionsData) GetCustomTimeZoneInfoData() ITimeZoneInfoData {
    return obj.CustomTimeZoneInfoData
}

func (obj *GifSaveOptionsData) SetCustomTimeZoneInfoData(value ITimeZoneInfoData) {
    obj.CustomTimeZoneInfoData = value
}

func (obj *GifSaveOptionsData) GetDml3DEffectsRenderingMode() *string {
    return obj.Dml3DEffectsRenderingMode
}

func (obj *GifSaveOptionsData) SetDml3DEffectsRenderingMode(value *string) {
    obj.Dml3DEffectsRenderingMode = value
}

func (obj *GifSaveOptionsData) GetDmlEffectsRenderingMode() *string {
    return obj.DmlEffectsRenderingMode
}

func (obj *GifSaveOptionsData) SetDmlEffectsRenderingMode(value *string) {
    obj.DmlEffectsRenderingMode = value
}

func (obj *GifSaveOptionsData) GetDmlRenderingMode() *string {
    return obj.DmlRenderingMode
}

func (obj *GifSaveOptionsData) SetDmlRenderingMode(value *string) {
    obj.DmlRenderingMode = value
}

func (obj *GifSaveOptionsData) GetFileName() *string {
    return obj.FileName
}

func (obj *GifSaveOptionsData) SetFileName(value *string) {
    obj.FileName = value
}

func (obj *GifSaveOptionsData) GetImlRenderingMode() *string {
    return obj.ImlRenderingMode
}

func (obj *GifSaveOptionsData) SetImlRenderingMode(value *string) {
    obj.ImlRenderingMode = value
}

func (obj *GifSaveOptionsData) GetUpdateCreatedTimeProperty() *bool {
    return obj.UpdateCreatedTimeProperty
}

func (obj *GifSaveOptionsData) SetUpdateCreatedTimeProperty(value *bool) {
    obj.UpdateCreatedTimeProperty = value
}

func (obj *GifSaveOptionsData) GetUpdateFields() *bool {
    return obj.UpdateFields
}

func (obj *GifSaveOptionsData) SetUpdateFields(value *bool) {
    obj.UpdateFields = value
}

func (obj *GifSaveOptionsData) GetUpdateLastPrintedProperty() *bool {
    return obj.UpdateLastPrintedProperty
}

func (obj *GifSaveOptionsData) SetUpdateLastPrintedProperty(value *bool) {
    obj.UpdateLastPrintedProperty = value
}

func (obj *GifSaveOptionsData) GetUpdateLastSavedTimeProperty() *bool {
    return obj.UpdateLastSavedTimeProperty
}

func (obj *GifSaveOptionsData) SetUpdateLastSavedTimeProperty(value *bool) {
    obj.UpdateLastSavedTimeProperty = value
}

func (obj *GifSaveOptionsData) GetZipOutput() *bool {
    return obj.ZipOutput
}

func (obj *GifSaveOptionsData) SetZipOutput(value *bool) {
    obj.ZipOutput = value
}

func (obj *GifSaveOptionsData) GetColorMode() *string {
    return obj.ColorMode
}

func (obj *GifSaveOptionsData) SetColorMode(value *string) {
    obj.ColorMode = value
}

func (obj *GifSaveOptionsData) GetJpegQuality() *int32 {
    return obj.JpegQuality
}

func (obj *GifSaveOptionsData) SetJpegQuality(value *int32) {
    obj.JpegQuality = value
}

func (obj *GifSaveOptionsData) GetMetafileRenderingOptions() IMetafileRenderingOptionsData {
    return obj.MetafileRenderingOptions
}

func (obj *GifSaveOptionsData) SetMetafileRenderingOptions(value IMetafileRenderingOptionsData) {
    obj.MetafileRenderingOptions = value
}

func (obj *GifSaveOptionsData) GetNumeralFormat() *string {
    return obj.NumeralFormat
}

func (obj *GifSaveOptionsData) SetNumeralFormat(value *string) {
    obj.NumeralFormat = value
}

func (obj *GifSaveOptionsData) GetOptimizeOutput() *bool {
    return obj.OptimizeOutput
}

func (obj *GifSaveOptionsData) SetOptimizeOutput(value *bool) {
    obj.OptimizeOutput = value
}

func (obj *GifSaveOptionsData) GetPageCount() *int32 {
    return obj.PageCount
}

func (obj *GifSaveOptionsData) SetPageCount(value *int32) {
    obj.PageCount = value
}

func (obj *GifSaveOptionsData) GetPageIndex() *int32 {
    return obj.PageIndex
}

func (obj *GifSaveOptionsData) SetPageIndex(value *int32) {
    obj.PageIndex = value
}

func (obj *GifSaveOptionsData) GetHorizontalResolution() *float64 {
    return obj.HorizontalResolution
}

func (obj *GifSaveOptionsData) SetHorizontalResolution(value *float64) {
    obj.HorizontalResolution = value
}

func (obj *GifSaveOptionsData) GetImageBrightness() *float64 {
    return obj.ImageBrightness
}

func (obj *GifSaveOptionsData) SetImageBrightness(value *float64) {
    obj.ImageBrightness = value
}

func (obj *GifSaveOptionsData) GetImageColorMode() *string {
    return obj.ImageColorMode
}

func (obj *GifSaveOptionsData) SetImageColorMode(value *string) {
    obj.ImageColorMode = value
}

func (obj *GifSaveOptionsData) GetImageContrast() *float64 {
    return obj.ImageContrast
}

func (obj *GifSaveOptionsData) SetImageContrast(value *float64) {
    obj.ImageContrast = value
}

func (obj *GifSaveOptionsData) GetPaperColor() *string {
    return obj.PaperColor
}

func (obj *GifSaveOptionsData) SetPaperColor(value *string) {
    obj.PaperColor = value
}

func (obj *GifSaveOptionsData) GetPixelFormat() *string {
    return obj.PixelFormat
}

func (obj *GifSaveOptionsData) SetPixelFormat(value *string) {
    obj.PixelFormat = value
}

func (obj *GifSaveOptionsData) GetResolution() *float64 {
    return obj.Resolution
}

func (obj *GifSaveOptionsData) SetResolution(value *float64) {
    obj.Resolution = value
}

func (obj *GifSaveOptionsData) GetScale() *float64 {
    return obj.Scale
}

func (obj *GifSaveOptionsData) SetScale(value *float64) {
    obj.Scale = value
}

func (obj *GifSaveOptionsData) GetUseAntiAliasing() *bool {
    return obj.UseAntiAliasing
}

func (obj *GifSaveOptionsData) SetUseAntiAliasing(value *bool) {
    obj.UseAntiAliasing = value
}

func (obj *GifSaveOptionsData) GetUseHighQualityRendering() *bool {
    return obj.UseHighQualityRendering
}

func (obj *GifSaveOptionsData) SetUseHighQualityRendering(value *bool) {
    obj.UseHighQualityRendering = value
}

func (obj *GifSaveOptionsData) GetVerticalResolution() *float64 {
    return obj.VerticalResolution
}

func (obj *GifSaveOptionsData) SetVerticalResolution(value *float64) {
    obj.VerticalResolution = value
}

func (obj *GifSaveOptionsData) GetImageHeight() *int32 {
    return obj.ImageHeight
}

func (obj *GifSaveOptionsData) SetImageHeight(value *int32) {
    obj.ImageHeight = value
}

func (obj *GifSaveOptionsData) GetImageWidth() *int32 {
    return obj.ImageWidth
}

func (obj *GifSaveOptionsData) SetImageWidth(value *int32) {
    obj.ImageWidth = value
}

func (obj *GifSaveOptionsData) GetUseGdiEmfRenderer() *bool {
    return obj.UseGdiEmfRenderer
}

func (obj *GifSaveOptionsData) SetUseGdiEmfRenderer(value *bool) {
    obj.UseGdiEmfRenderer = value
}

func (obj *GifSaveOptionsData) GetSaveFormat() *string {
    return obj.SaveFormat
}

func (obj *GifSaveOptionsData) SetSaveFormat(value *string) {
    obj.SaveFormat = value
}

