/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="eps_save_options_data.go">
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

// Container class for eps save options.

type IEpsSaveOptionsData interface {
    IsEpsSaveOptionsData() bool
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

type EpsSaveOptionsData struct {
    // Container class for eps save options.
    AllowEmbeddingPostScriptFonts *bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for eps save options.
    CustomTimeZoneInfoData ITimeZoneInfoData `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for eps save options.
    Dml3DEffectsRenderingMode *string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for eps save options.
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for eps save options.
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // Container class for eps save options.
    FileName *string `json:"FileName,omitempty"`

    // Container class for eps save options.
    ImlRenderingMode *string `json:"ImlRenderingMode,omitempty"`

    // Container class for eps save options.
    UpdateCreatedTimeProperty *bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for eps save options.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // Container class for eps save options.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for eps save options.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for eps save options.
    ZipOutput *bool `json:"ZipOutput,omitempty"`

    // Container class for eps save options.
    ColorMode *string `json:"ColorMode,omitempty"`

    // Container class for eps save options.
    JpegQuality *int32 `json:"JpegQuality,omitempty"`

    // Container class for eps save options.
    MetafileRenderingOptions IMetafileRenderingOptionsData `json:"MetafileRenderingOptions,omitempty"`

    // Container class for eps save options.
    NumeralFormat *string `json:"NumeralFormat,omitempty"`

    // Container class for eps save options.
    OptimizeOutput *bool `json:"OptimizeOutput,omitempty"`

    // Container class for eps save options.
    PageCount *int32 `json:"PageCount,omitempty"`

    // Container class for eps save options.
    PageIndex *int32 `json:"PageIndex,omitempty"`

    // Container class for eps save options.
    HorizontalResolution *float64 `json:"HorizontalResolution,omitempty"`

    // Container class for eps save options.
    ImageBrightness *float64 `json:"ImageBrightness,omitempty"`

    // Container class for eps save options.
    ImageColorMode *string `json:"ImageColorMode,omitempty"`

    // Container class for eps save options.
    ImageContrast *float64 `json:"ImageContrast,omitempty"`

    // Container class for eps save options.
    PaperColor *string `json:"PaperColor,omitempty"`

    // Container class for eps save options.
    PixelFormat *string `json:"PixelFormat,omitempty"`

    // Container class for eps save options.
    Resolution *float64 `json:"Resolution,omitempty"`

    // Container class for eps save options.
    Scale *float64 `json:"Scale,omitempty"`

    // Container class for eps save options.
    UseAntiAliasing *bool `json:"UseAntiAliasing,omitempty"`

    // Container class for eps save options.
    UseHighQualityRendering *bool `json:"UseHighQualityRendering,omitempty"`

    // Container class for eps save options.
    VerticalResolution *float64 `json:"VerticalResolution,omitempty"`

    // Container class for eps save options.
    ImageHeight *int32 `json:"ImageHeight,omitempty"`

    // Container class for eps save options.
    ImageWidth *int32 `json:"ImageWidth,omitempty"`

    // Container class for eps save options.
    UseGdiEmfRenderer *bool `json:"UseGdiEmfRenderer,omitempty"`

    // Container class for eps save options.
    SaveFormat *string `json:"SaveFormat,omitempty"`
}

func (EpsSaveOptionsData) IsEpsSaveOptionsData() bool {
    return true
}

func (EpsSaveOptionsData) IsImageSaveOptionsData() bool {
    return true
}

func (EpsSaveOptionsData) IsFixedPageSaveOptionsData() bool {
    return true
}

func (EpsSaveOptionsData) IsSaveOptionsData() bool {
    return true
}

func (obj *EpsSaveOptionsData) Initialize() {
    var _SaveFormat = "eps"
    obj.SaveFormat = &_SaveFormat


    if (obj.CustomTimeZoneInfoData != nil) {
        obj.CustomTimeZoneInfoData.Initialize()
    }

    if (obj.MetafileRenderingOptions != nil) {
        obj.MetafileRenderingOptions.Initialize()
    }


}

func (obj *EpsSaveOptionsData) Deserialize(json map[string]interface{}) {
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

func (obj *EpsSaveOptionsData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *EpsSaveOptionsData) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.FileName == nil {
        return errors.New("Property FileName in EpsSaveOptionsData is required.")
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

func (obj *EpsSaveOptionsData) GetAllowEmbeddingPostScriptFonts() *bool {
    return obj.AllowEmbeddingPostScriptFonts
}

func (obj *EpsSaveOptionsData) SetAllowEmbeddingPostScriptFonts(value *bool) {
    obj.AllowEmbeddingPostScriptFonts = value
}

func (obj *EpsSaveOptionsData) GetCustomTimeZoneInfoData() ITimeZoneInfoData {
    return obj.CustomTimeZoneInfoData
}

func (obj *EpsSaveOptionsData) SetCustomTimeZoneInfoData(value ITimeZoneInfoData) {
    obj.CustomTimeZoneInfoData = value
}

func (obj *EpsSaveOptionsData) GetDml3DEffectsRenderingMode() *string {
    return obj.Dml3DEffectsRenderingMode
}

func (obj *EpsSaveOptionsData) SetDml3DEffectsRenderingMode(value *string) {
    obj.Dml3DEffectsRenderingMode = value
}

func (obj *EpsSaveOptionsData) GetDmlEffectsRenderingMode() *string {
    return obj.DmlEffectsRenderingMode
}

func (obj *EpsSaveOptionsData) SetDmlEffectsRenderingMode(value *string) {
    obj.DmlEffectsRenderingMode = value
}

func (obj *EpsSaveOptionsData) GetDmlRenderingMode() *string {
    return obj.DmlRenderingMode
}

func (obj *EpsSaveOptionsData) SetDmlRenderingMode(value *string) {
    obj.DmlRenderingMode = value
}

func (obj *EpsSaveOptionsData) GetFileName() *string {
    return obj.FileName
}

func (obj *EpsSaveOptionsData) SetFileName(value *string) {
    obj.FileName = value
}

func (obj *EpsSaveOptionsData) GetImlRenderingMode() *string {
    return obj.ImlRenderingMode
}

func (obj *EpsSaveOptionsData) SetImlRenderingMode(value *string) {
    obj.ImlRenderingMode = value
}

func (obj *EpsSaveOptionsData) GetUpdateCreatedTimeProperty() *bool {
    return obj.UpdateCreatedTimeProperty
}

func (obj *EpsSaveOptionsData) SetUpdateCreatedTimeProperty(value *bool) {
    obj.UpdateCreatedTimeProperty = value
}

func (obj *EpsSaveOptionsData) GetUpdateFields() *bool {
    return obj.UpdateFields
}

func (obj *EpsSaveOptionsData) SetUpdateFields(value *bool) {
    obj.UpdateFields = value
}

func (obj *EpsSaveOptionsData) GetUpdateLastPrintedProperty() *bool {
    return obj.UpdateLastPrintedProperty
}

func (obj *EpsSaveOptionsData) SetUpdateLastPrintedProperty(value *bool) {
    obj.UpdateLastPrintedProperty = value
}

func (obj *EpsSaveOptionsData) GetUpdateLastSavedTimeProperty() *bool {
    return obj.UpdateLastSavedTimeProperty
}

func (obj *EpsSaveOptionsData) SetUpdateLastSavedTimeProperty(value *bool) {
    obj.UpdateLastSavedTimeProperty = value
}

func (obj *EpsSaveOptionsData) GetZipOutput() *bool {
    return obj.ZipOutput
}

func (obj *EpsSaveOptionsData) SetZipOutput(value *bool) {
    obj.ZipOutput = value
}

func (obj *EpsSaveOptionsData) GetColorMode() *string {
    return obj.ColorMode
}

func (obj *EpsSaveOptionsData) SetColorMode(value *string) {
    obj.ColorMode = value
}

func (obj *EpsSaveOptionsData) GetJpegQuality() *int32 {
    return obj.JpegQuality
}

func (obj *EpsSaveOptionsData) SetJpegQuality(value *int32) {
    obj.JpegQuality = value
}

func (obj *EpsSaveOptionsData) GetMetafileRenderingOptions() IMetafileRenderingOptionsData {
    return obj.MetafileRenderingOptions
}

func (obj *EpsSaveOptionsData) SetMetafileRenderingOptions(value IMetafileRenderingOptionsData) {
    obj.MetafileRenderingOptions = value
}

func (obj *EpsSaveOptionsData) GetNumeralFormat() *string {
    return obj.NumeralFormat
}

func (obj *EpsSaveOptionsData) SetNumeralFormat(value *string) {
    obj.NumeralFormat = value
}

func (obj *EpsSaveOptionsData) GetOptimizeOutput() *bool {
    return obj.OptimizeOutput
}

func (obj *EpsSaveOptionsData) SetOptimizeOutput(value *bool) {
    obj.OptimizeOutput = value
}

func (obj *EpsSaveOptionsData) GetPageCount() *int32 {
    return obj.PageCount
}

func (obj *EpsSaveOptionsData) SetPageCount(value *int32) {
    obj.PageCount = value
}

func (obj *EpsSaveOptionsData) GetPageIndex() *int32 {
    return obj.PageIndex
}

func (obj *EpsSaveOptionsData) SetPageIndex(value *int32) {
    obj.PageIndex = value
}

func (obj *EpsSaveOptionsData) GetHorizontalResolution() *float64 {
    return obj.HorizontalResolution
}

func (obj *EpsSaveOptionsData) SetHorizontalResolution(value *float64) {
    obj.HorizontalResolution = value
}

func (obj *EpsSaveOptionsData) GetImageBrightness() *float64 {
    return obj.ImageBrightness
}

func (obj *EpsSaveOptionsData) SetImageBrightness(value *float64) {
    obj.ImageBrightness = value
}

func (obj *EpsSaveOptionsData) GetImageColorMode() *string {
    return obj.ImageColorMode
}

func (obj *EpsSaveOptionsData) SetImageColorMode(value *string) {
    obj.ImageColorMode = value
}

func (obj *EpsSaveOptionsData) GetImageContrast() *float64 {
    return obj.ImageContrast
}

func (obj *EpsSaveOptionsData) SetImageContrast(value *float64) {
    obj.ImageContrast = value
}

func (obj *EpsSaveOptionsData) GetPaperColor() *string {
    return obj.PaperColor
}

func (obj *EpsSaveOptionsData) SetPaperColor(value *string) {
    obj.PaperColor = value
}

func (obj *EpsSaveOptionsData) GetPixelFormat() *string {
    return obj.PixelFormat
}

func (obj *EpsSaveOptionsData) SetPixelFormat(value *string) {
    obj.PixelFormat = value
}

func (obj *EpsSaveOptionsData) GetResolution() *float64 {
    return obj.Resolution
}

func (obj *EpsSaveOptionsData) SetResolution(value *float64) {
    obj.Resolution = value
}

func (obj *EpsSaveOptionsData) GetScale() *float64 {
    return obj.Scale
}

func (obj *EpsSaveOptionsData) SetScale(value *float64) {
    obj.Scale = value
}

func (obj *EpsSaveOptionsData) GetUseAntiAliasing() *bool {
    return obj.UseAntiAliasing
}

func (obj *EpsSaveOptionsData) SetUseAntiAliasing(value *bool) {
    obj.UseAntiAliasing = value
}

func (obj *EpsSaveOptionsData) GetUseHighQualityRendering() *bool {
    return obj.UseHighQualityRendering
}

func (obj *EpsSaveOptionsData) SetUseHighQualityRendering(value *bool) {
    obj.UseHighQualityRendering = value
}

func (obj *EpsSaveOptionsData) GetVerticalResolution() *float64 {
    return obj.VerticalResolution
}

func (obj *EpsSaveOptionsData) SetVerticalResolution(value *float64) {
    obj.VerticalResolution = value
}

func (obj *EpsSaveOptionsData) GetImageHeight() *int32 {
    return obj.ImageHeight
}

func (obj *EpsSaveOptionsData) SetImageHeight(value *int32) {
    obj.ImageHeight = value
}

func (obj *EpsSaveOptionsData) GetImageWidth() *int32 {
    return obj.ImageWidth
}

func (obj *EpsSaveOptionsData) SetImageWidth(value *int32) {
    obj.ImageWidth = value
}

func (obj *EpsSaveOptionsData) GetUseGdiEmfRenderer() *bool {
    return obj.UseGdiEmfRenderer
}

func (obj *EpsSaveOptionsData) SetUseGdiEmfRenderer(value *bool) {
    obj.UseGdiEmfRenderer = value
}

func (obj *EpsSaveOptionsData) GetSaveFormat() *string {
    return obj.SaveFormat
}

func (obj *EpsSaveOptionsData) SetSaveFormat(value *string) {
    obj.SaveFormat = value
}

