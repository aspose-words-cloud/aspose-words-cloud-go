/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="tiff_save_options_data.go">
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

// Container class for tiff save options.

type ITiffSaveOptionsData interface {
    IsTiffSaveOptionsData() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
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
    GetUseGdiEmfRenderer() *bool
    SetUseGdiEmfRenderer(value *bool)
    GetThresholdForFloydSteinbergDithering() *int32
    SetThresholdForFloydSteinbergDithering(value *int32)
    GetTiffBinarizationMethod() *string
    SetTiffBinarizationMethod(value *string)
    GetTiffCompression() *string
    SetTiffCompression(value *string)
    GetSaveFormat() *string
    SetSaveFormat(value *string)
}

type TiffSaveOptionsData struct {
    // Container class for tiff save options.
    AllowEmbeddingPostScriptFonts *bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for tiff save options.
    CustomTimeZoneInfoData ITimeZoneInfoData `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for tiff save options.
    Dml3DEffectsRenderingMode *string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for tiff save options.
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for tiff save options.
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // Container class for tiff save options.
    FileName *string `json:"FileName,omitempty"`

    // Container class for tiff save options.
    ImlRenderingMode *string `json:"ImlRenderingMode,omitempty"`

    // Container class for tiff save options.
    UpdateCreatedTimeProperty *bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for tiff save options.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // Container class for tiff save options.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for tiff save options.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for tiff save options.
    ZipOutput *bool `json:"ZipOutput,omitempty"`

    // Container class for tiff save options.
    ColorMode *string `json:"ColorMode,omitempty"`

    // Container class for tiff save options.
    JpegQuality *int32 `json:"JpegQuality,omitempty"`

    // Container class for tiff save options.
    MetafileRenderingOptions IMetafileRenderingOptionsData `json:"MetafileRenderingOptions,omitempty"`

    // Container class for tiff save options.
    NumeralFormat *string `json:"NumeralFormat,omitempty"`

    // Container class for tiff save options.
    OptimizeOutput *bool `json:"OptimizeOutput,omitempty"`

    // Container class for tiff save options.
    PageCount *int32 `json:"PageCount,omitempty"`

    // Container class for tiff save options.
    PageIndex *int32 `json:"PageIndex,omitempty"`

    // Container class for tiff save options.
    HorizontalResolution *float64 `json:"HorizontalResolution,omitempty"`

    // Container class for tiff save options.
    ImageBrightness *float64 `json:"ImageBrightness,omitempty"`

    // Container class for tiff save options.
    ImageColorMode *string `json:"ImageColorMode,omitempty"`

    // Container class for tiff save options.
    ImageContrast *float64 `json:"ImageContrast,omitempty"`

    // Container class for tiff save options.
    PaperColor *string `json:"PaperColor,omitempty"`

    // Container class for tiff save options.
    PixelFormat *string `json:"PixelFormat,omitempty"`

    // Container class for tiff save options.
    Resolution *float64 `json:"Resolution,omitempty"`

    // Container class for tiff save options.
    Scale *float64 `json:"Scale,omitempty"`

    // Container class for tiff save options.
    UseAntiAliasing *bool `json:"UseAntiAliasing,omitempty"`

    // Container class for tiff save options.
    UseHighQualityRendering *bool `json:"UseHighQualityRendering,omitempty"`

    // Container class for tiff save options.
    VerticalResolution *float64 `json:"VerticalResolution,omitempty"`

    // Container class for tiff save options.
    UseGdiEmfRenderer *bool `json:"UseGdiEmfRenderer,omitempty"`

    // Container class for tiff save options.
    ThresholdForFloydSteinbergDithering *int32 `json:"ThresholdForFloydSteinbergDithering,omitempty"`

    // Container class for tiff save options.
    TiffBinarizationMethod *string `json:"TiffBinarizationMethod,omitempty"`

    // Container class for tiff save options.
    TiffCompression *string `json:"TiffCompression,omitempty"`

    // Container class for tiff save options.
    SaveFormat *string `json:"SaveFormat,omitempty"`
}

func (TiffSaveOptionsData) IsTiffSaveOptionsData() bool {
    return true
}

func (TiffSaveOptionsData) IsImageSaveOptionsData() bool {
    return true
}

func (TiffSaveOptionsData) IsFixedPageSaveOptionsData() bool {
    return true
}

func (TiffSaveOptionsData) IsSaveOptionsData() bool {
    return true
}

func (obj *TiffSaveOptionsData) Initialize() {
    var _SaveFormat = "tiff"
    obj.SaveFormat = &_SaveFormat


    if (obj.CustomTimeZoneInfoData != nil) {
        obj.CustomTimeZoneInfoData.Initialize()
    }

    if (obj.MetafileRenderingOptions != nil) {
        obj.MetafileRenderingOptions.Initialize()
    }


}

func (obj *TiffSaveOptionsData) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["UseGdiEmfRenderer"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UseGdiEmfRenderer = &parsedValue
        }

    } else if jsonValue, exists := json["useGdiEmfRenderer"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UseGdiEmfRenderer = &parsedValue
        }

    }

    if jsonValue, exists := json["ThresholdForFloydSteinbergDithering"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ThresholdForFloydSteinbergDithering = new(int32)
            *obj.ThresholdForFloydSteinbergDithering = int32(parsedValue)
        }

    } else if jsonValue, exists := json["thresholdForFloydSteinbergDithering"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ThresholdForFloydSteinbergDithering = new(int32)
            *obj.ThresholdForFloydSteinbergDithering = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["TiffBinarizationMethod"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.TiffBinarizationMethod = &parsedValue
        }

    } else if jsonValue, exists := json["tiffBinarizationMethod"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.TiffBinarizationMethod = &parsedValue
        }

    }

    if jsonValue, exists := json["TiffCompression"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.TiffCompression = &parsedValue
        }

    } else if jsonValue, exists := json["tiffCompression"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.TiffCompression = &parsedValue
        }

    }
}

func (obj *TiffSaveOptionsData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *TiffSaveOptionsData) GetAllowEmbeddingPostScriptFonts() *bool {
    return obj.AllowEmbeddingPostScriptFonts
}

func (obj *TiffSaveOptionsData) SetAllowEmbeddingPostScriptFonts(value *bool) {
    obj.AllowEmbeddingPostScriptFonts = value
}

func (obj *TiffSaveOptionsData) GetCustomTimeZoneInfoData() ITimeZoneInfoData {
    return obj.CustomTimeZoneInfoData
}

func (obj *TiffSaveOptionsData) SetCustomTimeZoneInfoData(value ITimeZoneInfoData) {
    obj.CustomTimeZoneInfoData = value
}

func (obj *TiffSaveOptionsData) GetDml3DEffectsRenderingMode() *string {
    return obj.Dml3DEffectsRenderingMode
}

func (obj *TiffSaveOptionsData) SetDml3DEffectsRenderingMode(value *string) {
    obj.Dml3DEffectsRenderingMode = value
}

func (obj *TiffSaveOptionsData) GetDmlEffectsRenderingMode() *string {
    return obj.DmlEffectsRenderingMode
}

func (obj *TiffSaveOptionsData) SetDmlEffectsRenderingMode(value *string) {
    obj.DmlEffectsRenderingMode = value
}

func (obj *TiffSaveOptionsData) GetDmlRenderingMode() *string {
    return obj.DmlRenderingMode
}

func (obj *TiffSaveOptionsData) SetDmlRenderingMode(value *string) {
    obj.DmlRenderingMode = value
}

func (obj *TiffSaveOptionsData) GetFileName() *string {
    return obj.FileName
}

func (obj *TiffSaveOptionsData) SetFileName(value *string) {
    obj.FileName = value
}

func (obj *TiffSaveOptionsData) GetImlRenderingMode() *string {
    return obj.ImlRenderingMode
}

func (obj *TiffSaveOptionsData) SetImlRenderingMode(value *string) {
    obj.ImlRenderingMode = value
}

func (obj *TiffSaveOptionsData) GetUpdateCreatedTimeProperty() *bool {
    return obj.UpdateCreatedTimeProperty
}

func (obj *TiffSaveOptionsData) SetUpdateCreatedTimeProperty(value *bool) {
    obj.UpdateCreatedTimeProperty = value
}

func (obj *TiffSaveOptionsData) GetUpdateFields() *bool {
    return obj.UpdateFields
}

func (obj *TiffSaveOptionsData) SetUpdateFields(value *bool) {
    obj.UpdateFields = value
}

func (obj *TiffSaveOptionsData) GetUpdateLastPrintedProperty() *bool {
    return obj.UpdateLastPrintedProperty
}

func (obj *TiffSaveOptionsData) SetUpdateLastPrintedProperty(value *bool) {
    obj.UpdateLastPrintedProperty = value
}

func (obj *TiffSaveOptionsData) GetUpdateLastSavedTimeProperty() *bool {
    return obj.UpdateLastSavedTimeProperty
}

func (obj *TiffSaveOptionsData) SetUpdateLastSavedTimeProperty(value *bool) {
    obj.UpdateLastSavedTimeProperty = value
}

func (obj *TiffSaveOptionsData) GetZipOutput() *bool {
    return obj.ZipOutput
}

func (obj *TiffSaveOptionsData) SetZipOutput(value *bool) {
    obj.ZipOutput = value
}

func (obj *TiffSaveOptionsData) GetColorMode() *string {
    return obj.ColorMode
}

func (obj *TiffSaveOptionsData) SetColorMode(value *string) {
    obj.ColorMode = value
}

func (obj *TiffSaveOptionsData) GetJpegQuality() *int32 {
    return obj.JpegQuality
}

func (obj *TiffSaveOptionsData) SetJpegQuality(value *int32) {
    obj.JpegQuality = value
}

func (obj *TiffSaveOptionsData) GetMetafileRenderingOptions() IMetafileRenderingOptionsData {
    return obj.MetafileRenderingOptions
}

func (obj *TiffSaveOptionsData) SetMetafileRenderingOptions(value IMetafileRenderingOptionsData) {
    obj.MetafileRenderingOptions = value
}

func (obj *TiffSaveOptionsData) GetNumeralFormat() *string {
    return obj.NumeralFormat
}

func (obj *TiffSaveOptionsData) SetNumeralFormat(value *string) {
    obj.NumeralFormat = value
}

func (obj *TiffSaveOptionsData) GetOptimizeOutput() *bool {
    return obj.OptimizeOutput
}

func (obj *TiffSaveOptionsData) SetOptimizeOutput(value *bool) {
    obj.OptimizeOutput = value
}

func (obj *TiffSaveOptionsData) GetPageCount() *int32 {
    return obj.PageCount
}

func (obj *TiffSaveOptionsData) SetPageCount(value *int32) {
    obj.PageCount = value
}

func (obj *TiffSaveOptionsData) GetPageIndex() *int32 {
    return obj.PageIndex
}

func (obj *TiffSaveOptionsData) SetPageIndex(value *int32) {
    obj.PageIndex = value
}

func (obj *TiffSaveOptionsData) GetHorizontalResolution() *float64 {
    return obj.HorizontalResolution
}

func (obj *TiffSaveOptionsData) SetHorizontalResolution(value *float64) {
    obj.HorizontalResolution = value
}

func (obj *TiffSaveOptionsData) GetImageBrightness() *float64 {
    return obj.ImageBrightness
}

func (obj *TiffSaveOptionsData) SetImageBrightness(value *float64) {
    obj.ImageBrightness = value
}

func (obj *TiffSaveOptionsData) GetImageColorMode() *string {
    return obj.ImageColorMode
}

func (obj *TiffSaveOptionsData) SetImageColorMode(value *string) {
    obj.ImageColorMode = value
}

func (obj *TiffSaveOptionsData) GetImageContrast() *float64 {
    return obj.ImageContrast
}

func (obj *TiffSaveOptionsData) SetImageContrast(value *float64) {
    obj.ImageContrast = value
}

func (obj *TiffSaveOptionsData) GetPaperColor() *string {
    return obj.PaperColor
}

func (obj *TiffSaveOptionsData) SetPaperColor(value *string) {
    obj.PaperColor = value
}

func (obj *TiffSaveOptionsData) GetPixelFormat() *string {
    return obj.PixelFormat
}

func (obj *TiffSaveOptionsData) SetPixelFormat(value *string) {
    obj.PixelFormat = value
}

func (obj *TiffSaveOptionsData) GetResolution() *float64 {
    return obj.Resolution
}

func (obj *TiffSaveOptionsData) SetResolution(value *float64) {
    obj.Resolution = value
}

func (obj *TiffSaveOptionsData) GetScale() *float64 {
    return obj.Scale
}

func (obj *TiffSaveOptionsData) SetScale(value *float64) {
    obj.Scale = value
}

func (obj *TiffSaveOptionsData) GetUseAntiAliasing() *bool {
    return obj.UseAntiAliasing
}

func (obj *TiffSaveOptionsData) SetUseAntiAliasing(value *bool) {
    obj.UseAntiAliasing = value
}

func (obj *TiffSaveOptionsData) GetUseHighQualityRendering() *bool {
    return obj.UseHighQualityRendering
}

func (obj *TiffSaveOptionsData) SetUseHighQualityRendering(value *bool) {
    obj.UseHighQualityRendering = value
}

func (obj *TiffSaveOptionsData) GetVerticalResolution() *float64 {
    return obj.VerticalResolution
}

func (obj *TiffSaveOptionsData) SetVerticalResolution(value *float64) {
    obj.VerticalResolution = value
}

func (obj *TiffSaveOptionsData) GetUseGdiEmfRenderer() *bool {
    return obj.UseGdiEmfRenderer
}

func (obj *TiffSaveOptionsData) SetUseGdiEmfRenderer(value *bool) {
    obj.UseGdiEmfRenderer = value
}

func (obj *TiffSaveOptionsData) GetThresholdForFloydSteinbergDithering() *int32 {
    return obj.ThresholdForFloydSteinbergDithering
}

func (obj *TiffSaveOptionsData) SetThresholdForFloydSteinbergDithering(value *int32) {
    obj.ThresholdForFloydSteinbergDithering = value
}

func (obj *TiffSaveOptionsData) GetTiffBinarizationMethod() *string {
    return obj.TiffBinarizationMethod
}

func (obj *TiffSaveOptionsData) SetTiffBinarizationMethod(value *string) {
    obj.TiffBinarizationMethod = value
}

func (obj *TiffSaveOptionsData) GetTiffCompression() *string {
    return obj.TiffCompression
}

func (obj *TiffSaveOptionsData) SetTiffCompression(value *string) {
    obj.TiffCompression = value
}

func (obj *TiffSaveOptionsData) GetSaveFormat() *string {
    return obj.SaveFormat
}

func (obj *TiffSaveOptionsData) SetSaveFormat(value *string) {
    obj.SaveFormat = value
}

