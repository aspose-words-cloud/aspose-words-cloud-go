/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="bmp_save_options_data.go">
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

// Container class for bmp save options.

type IBmpSaveOptionsData interface {
    IsBmpSaveOptionsData() bool
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

type BmpSaveOptionsData struct {
    // Container class for bmp save options.
    AllowEmbeddingPostScriptFonts *bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for bmp save options.
    CustomTimeZoneInfoData ITimeZoneInfoData `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for bmp save options.
    Dml3DEffectsRenderingMode *string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for bmp save options.
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for bmp save options.
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // Container class for bmp save options.
    FileName *string `json:"FileName,omitempty"`

    // Container class for bmp save options.
    ImlRenderingMode *string `json:"ImlRenderingMode,omitempty"`

    // Container class for bmp save options.
    UpdateCreatedTimeProperty *bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for bmp save options.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // Container class for bmp save options.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for bmp save options.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for bmp save options.
    ZipOutput *bool `json:"ZipOutput,omitempty"`

    // Container class for bmp save options.
    ColorMode *string `json:"ColorMode,omitempty"`

    // Container class for bmp save options.
    JpegQuality *int32 `json:"JpegQuality,omitempty"`

    // Container class for bmp save options.
    MetafileRenderingOptions IMetafileRenderingOptionsData `json:"MetafileRenderingOptions,omitempty"`

    // Container class for bmp save options.
    NumeralFormat *string `json:"NumeralFormat,omitempty"`

    // Container class for bmp save options.
    OptimizeOutput *bool `json:"OptimizeOutput,omitempty"`

    // Container class for bmp save options.
    PageCount *int32 `json:"PageCount,omitempty"`

    // Container class for bmp save options.
    PageIndex *int32 `json:"PageIndex,omitempty"`

    // Container class for bmp save options.
    HorizontalResolution *float64 `json:"HorizontalResolution,omitempty"`

    // Container class for bmp save options.
    ImageBrightness *float64 `json:"ImageBrightness,omitempty"`

    // Container class for bmp save options.
    ImageColorMode *string `json:"ImageColorMode,omitempty"`

    // Container class for bmp save options.
    ImageContrast *float64 `json:"ImageContrast,omitempty"`

    // Container class for bmp save options.
    PaperColor *string `json:"PaperColor,omitempty"`

    // Container class for bmp save options.
    PixelFormat *string `json:"PixelFormat,omitempty"`

    // Container class for bmp save options.
    Resolution *float64 `json:"Resolution,omitempty"`

    // Container class for bmp save options.
    Scale *float64 `json:"Scale,omitempty"`

    // Container class for bmp save options.
    UseAntiAliasing *bool `json:"UseAntiAliasing,omitempty"`

    // Container class for bmp save options.
    UseHighQualityRendering *bool `json:"UseHighQualityRendering,omitempty"`

    // Container class for bmp save options.
    VerticalResolution *float64 `json:"VerticalResolution,omitempty"`

    // Container class for bmp save options.
    ImageHeight *int32 `json:"ImageHeight,omitempty"`

    // Container class for bmp save options.
    ImageWidth *int32 `json:"ImageWidth,omitempty"`

    // Container class for bmp save options.
    UseGdiEmfRenderer *bool `json:"UseGdiEmfRenderer,omitempty"`

    // Container class for bmp save options.
    SaveFormat *string `json:"SaveFormat,omitempty"`
}

func (BmpSaveOptionsData) IsBmpSaveOptionsData() bool {
    return true
}

func (BmpSaveOptionsData) IsImageSaveOptionsData() bool {
    return true
}

func (BmpSaveOptionsData) IsFixedPageSaveOptionsData() bool {
    return true
}

func (BmpSaveOptionsData) IsSaveOptionsData() bool {
    return true
}

func (obj *BmpSaveOptionsData) Initialize() {
    var _SaveFormat = "bmp"
    obj.SaveFormat = &_SaveFormat


    if (obj.CustomTimeZoneInfoData != nil) {
        obj.CustomTimeZoneInfoData.Initialize()
    }

    if (obj.MetafileRenderingOptions != nil) {
        obj.MetafileRenderingOptions.Initialize()
    }


}

func (obj *BmpSaveOptionsData) Deserialize(json map[string]interface{}) {
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

func (obj *BmpSaveOptionsData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *BmpSaveOptionsData) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.FileName == nil {
        return errors.New("Property FileName in BmpSaveOptionsData is required.")
    }
    if obj.AllowEmbeddingPostScriptFonts == nil {
        return errors.New("Property AllowEmbeddingPostScriptFonts in BmpSaveOptionsData is required.")
    }

    if obj.CustomTimeZoneInfoData == nil {
        return errors.New("Property CustomTimeZoneInfoData in BmpSaveOptionsData is required.")
    }

    if obj.CustomTimeZoneInfoData != nil {
        if err := obj.CustomTimeZoneInfoData.Validate(); err != nil {
            return err
        }
    }

    if obj.Dml3DEffectsRenderingMode == nil {
        return errors.New("Property Dml3DEffectsRenderingMode in BmpSaveOptionsData is required.")
    }

    if obj.DmlEffectsRenderingMode == nil {
        return errors.New("Property DmlEffectsRenderingMode in BmpSaveOptionsData is required.")
    }

    if obj.DmlRenderingMode == nil {
        return errors.New("Property DmlRenderingMode in BmpSaveOptionsData is required.")
    }

    if obj.FileName == nil {
        return errors.New("Property FileName in BmpSaveOptionsData is required.")
    }

    if obj.ImlRenderingMode == nil {
        return errors.New("Property ImlRenderingMode in BmpSaveOptionsData is required.")
    }

    if obj.UpdateCreatedTimeProperty == nil {
        return errors.New("Property UpdateCreatedTimeProperty in BmpSaveOptionsData is required.")
    }

    if obj.UpdateFields == nil {
        return errors.New("Property UpdateFields in BmpSaveOptionsData is required.")
    }

    if obj.UpdateLastPrintedProperty == nil {
        return errors.New("Property UpdateLastPrintedProperty in BmpSaveOptionsData is required.")
    }

    if obj.UpdateLastSavedTimeProperty == nil {
        return errors.New("Property UpdateLastSavedTimeProperty in BmpSaveOptionsData is required.")
    }

    if obj.ZipOutput == nil {
        return errors.New("Property ZipOutput in BmpSaveOptionsData is required.")
    }

    if obj.ColorMode == nil {
        return errors.New("Property ColorMode in BmpSaveOptionsData is required.")
    }

    if obj.JpegQuality == nil {
        return errors.New("Property JpegQuality in BmpSaveOptionsData is required.")
    }

    if obj.MetafileRenderingOptions == nil {
        return errors.New("Property MetafileRenderingOptions in BmpSaveOptionsData is required.")
    }

    if obj.MetafileRenderingOptions != nil {
        if err := obj.MetafileRenderingOptions.Validate(); err != nil {
            return err
        }
    }

    if obj.NumeralFormat == nil {
        return errors.New("Property NumeralFormat in BmpSaveOptionsData is required.")
    }

    if obj.OptimizeOutput == nil {
        return errors.New("Property OptimizeOutput in BmpSaveOptionsData is required.")
    }

    if obj.PageCount == nil {
        return errors.New("Property PageCount in BmpSaveOptionsData is required.")
    }

    if obj.PageIndex == nil {
        return errors.New("Property PageIndex in BmpSaveOptionsData is required.")
    }

    if obj.HorizontalResolution == nil {
        return errors.New("Property HorizontalResolution in BmpSaveOptionsData is required.")
    }

    if obj.ImageBrightness == nil {
        return errors.New("Property ImageBrightness in BmpSaveOptionsData is required.")
    }

    if obj.ImageColorMode == nil {
        return errors.New("Property ImageColorMode in BmpSaveOptionsData is required.")
    }

    if obj.ImageContrast == nil {
        return errors.New("Property ImageContrast in BmpSaveOptionsData is required.")
    }

    if obj.PaperColor == nil {
        return errors.New("Property PaperColor in BmpSaveOptionsData is required.")
    }

    if obj.PixelFormat == nil {
        return errors.New("Property PixelFormat in BmpSaveOptionsData is required.")
    }

    if obj.Resolution == nil {
        return errors.New("Property Resolution in BmpSaveOptionsData is required.")
    }

    if obj.Scale == nil {
        return errors.New("Property Scale in BmpSaveOptionsData is required.")
    }

    if obj.UseAntiAliasing == nil {
        return errors.New("Property UseAntiAliasing in BmpSaveOptionsData is required.")
    }

    if obj.UseHighQualityRendering == nil {
        return errors.New("Property UseHighQualityRendering in BmpSaveOptionsData is required.")
    }

    if obj.VerticalResolution == nil {
        return errors.New("Property VerticalResolution in BmpSaveOptionsData is required.")
    }

    if obj.ImageHeight == nil {
        return errors.New("Property ImageHeight in BmpSaveOptionsData is required.")
    }

    if obj.ImageWidth == nil {
        return errors.New("Property ImageWidth in BmpSaveOptionsData is required.")
    }

    if obj.UseGdiEmfRenderer == nil {
        return errors.New("Property UseGdiEmfRenderer in BmpSaveOptionsData is required.")
    }

    if obj.SaveFormat == nil {
        return errors.New("Property SaveFormat in BmpSaveOptionsData is required.")
    }

    return nil;
}

func (obj *BmpSaveOptionsData) GetAllowEmbeddingPostScriptFonts() *bool {
    return obj.AllowEmbeddingPostScriptFonts
}

func (obj *BmpSaveOptionsData) SetAllowEmbeddingPostScriptFonts(value *bool) {
    obj.AllowEmbeddingPostScriptFonts = value
}

func (obj *BmpSaveOptionsData) GetCustomTimeZoneInfoData() ITimeZoneInfoData {
    return obj.CustomTimeZoneInfoData
}

func (obj *BmpSaveOptionsData) SetCustomTimeZoneInfoData(value ITimeZoneInfoData) {
    obj.CustomTimeZoneInfoData = value
}

func (obj *BmpSaveOptionsData) GetDml3DEffectsRenderingMode() *string {
    return obj.Dml3DEffectsRenderingMode
}

func (obj *BmpSaveOptionsData) SetDml3DEffectsRenderingMode(value *string) {
    obj.Dml3DEffectsRenderingMode = value
}

func (obj *BmpSaveOptionsData) GetDmlEffectsRenderingMode() *string {
    return obj.DmlEffectsRenderingMode
}

func (obj *BmpSaveOptionsData) SetDmlEffectsRenderingMode(value *string) {
    obj.DmlEffectsRenderingMode = value
}

func (obj *BmpSaveOptionsData) GetDmlRenderingMode() *string {
    return obj.DmlRenderingMode
}

func (obj *BmpSaveOptionsData) SetDmlRenderingMode(value *string) {
    obj.DmlRenderingMode = value
}

func (obj *BmpSaveOptionsData) GetFileName() *string {
    return obj.FileName
}

func (obj *BmpSaveOptionsData) SetFileName(value *string) {
    obj.FileName = value
}

func (obj *BmpSaveOptionsData) GetImlRenderingMode() *string {
    return obj.ImlRenderingMode
}

func (obj *BmpSaveOptionsData) SetImlRenderingMode(value *string) {
    obj.ImlRenderingMode = value
}

func (obj *BmpSaveOptionsData) GetUpdateCreatedTimeProperty() *bool {
    return obj.UpdateCreatedTimeProperty
}

func (obj *BmpSaveOptionsData) SetUpdateCreatedTimeProperty(value *bool) {
    obj.UpdateCreatedTimeProperty = value
}

func (obj *BmpSaveOptionsData) GetUpdateFields() *bool {
    return obj.UpdateFields
}

func (obj *BmpSaveOptionsData) SetUpdateFields(value *bool) {
    obj.UpdateFields = value
}

func (obj *BmpSaveOptionsData) GetUpdateLastPrintedProperty() *bool {
    return obj.UpdateLastPrintedProperty
}

func (obj *BmpSaveOptionsData) SetUpdateLastPrintedProperty(value *bool) {
    obj.UpdateLastPrintedProperty = value
}

func (obj *BmpSaveOptionsData) GetUpdateLastSavedTimeProperty() *bool {
    return obj.UpdateLastSavedTimeProperty
}

func (obj *BmpSaveOptionsData) SetUpdateLastSavedTimeProperty(value *bool) {
    obj.UpdateLastSavedTimeProperty = value
}

func (obj *BmpSaveOptionsData) GetZipOutput() *bool {
    return obj.ZipOutput
}

func (obj *BmpSaveOptionsData) SetZipOutput(value *bool) {
    obj.ZipOutput = value
}

func (obj *BmpSaveOptionsData) GetColorMode() *string {
    return obj.ColorMode
}

func (obj *BmpSaveOptionsData) SetColorMode(value *string) {
    obj.ColorMode = value
}

func (obj *BmpSaveOptionsData) GetJpegQuality() *int32 {
    return obj.JpegQuality
}

func (obj *BmpSaveOptionsData) SetJpegQuality(value *int32) {
    obj.JpegQuality = value
}

func (obj *BmpSaveOptionsData) GetMetafileRenderingOptions() IMetafileRenderingOptionsData {
    return obj.MetafileRenderingOptions
}

func (obj *BmpSaveOptionsData) SetMetafileRenderingOptions(value IMetafileRenderingOptionsData) {
    obj.MetafileRenderingOptions = value
}

func (obj *BmpSaveOptionsData) GetNumeralFormat() *string {
    return obj.NumeralFormat
}

func (obj *BmpSaveOptionsData) SetNumeralFormat(value *string) {
    obj.NumeralFormat = value
}

func (obj *BmpSaveOptionsData) GetOptimizeOutput() *bool {
    return obj.OptimizeOutput
}

func (obj *BmpSaveOptionsData) SetOptimizeOutput(value *bool) {
    obj.OptimizeOutput = value
}

func (obj *BmpSaveOptionsData) GetPageCount() *int32 {
    return obj.PageCount
}

func (obj *BmpSaveOptionsData) SetPageCount(value *int32) {
    obj.PageCount = value
}

func (obj *BmpSaveOptionsData) GetPageIndex() *int32 {
    return obj.PageIndex
}

func (obj *BmpSaveOptionsData) SetPageIndex(value *int32) {
    obj.PageIndex = value
}

func (obj *BmpSaveOptionsData) GetHorizontalResolution() *float64 {
    return obj.HorizontalResolution
}

func (obj *BmpSaveOptionsData) SetHorizontalResolution(value *float64) {
    obj.HorizontalResolution = value
}

func (obj *BmpSaveOptionsData) GetImageBrightness() *float64 {
    return obj.ImageBrightness
}

func (obj *BmpSaveOptionsData) SetImageBrightness(value *float64) {
    obj.ImageBrightness = value
}

func (obj *BmpSaveOptionsData) GetImageColorMode() *string {
    return obj.ImageColorMode
}

func (obj *BmpSaveOptionsData) SetImageColorMode(value *string) {
    obj.ImageColorMode = value
}

func (obj *BmpSaveOptionsData) GetImageContrast() *float64 {
    return obj.ImageContrast
}

func (obj *BmpSaveOptionsData) SetImageContrast(value *float64) {
    obj.ImageContrast = value
}

func (obj *BmpSaveOptionsData) GetPaperColor() *string {
    return obj.PaperColor
}

func (obj *BmpSaveOptionsData) SetPaperColor(value *string) {
    obj.PaperColor = value
}

func (obj *BmpSaveOptionsData) GetPixelFormat() *string {
    return obj.PixelFormat
}

func (obj *BmpSaveOptionsData) SetPixelFormat(value *string) {
    obj.PixelFormat = value
}

func (obj *BmpSaveOptionsData) GetResolution() *float64 {
    return obj.Resolution
}

func (obj *BmpSaveOptionsData) SetResolution(value *float64) {
    obj.Resolution = value
}

func (obj *BmpSaveOptionsData) GetScale() *float64 {
    return obj.Scale
}

func (obj *BmpSaveOptionsData) SetScale(value *float64) {
    obj.Scale = value
}

func (obj *BmpSaveOptionsData) GetUseAntiAliasing() *bool {
    return obj.UseAntiAliasing
}

func (obj *BmpSaveOptionsData) SetUseAntiAliasing(value *bool) {
    obj.UseAntiAliasing = value
}

func (obj *BmpSaveOptionsData) GetUseHighQualityRendering() *bool {
    return obj.UseHighQualityRendering
}

func (obj *BmpSaveOptionsData) SetUseHighQualityRendering(value *bool) {
    obj.UseHighQualityRendering = value
}

func (obj *BmpSaveOptionsData) GetVerticalResolution() *float64 {
    return obj.VerticalResolution
}

func (obj *BmpSaveOptionsData) SetVerticalResolution(value *float64) {
    obj.VerticalResolution = value
}

func (obj *BmpSaveOptionsData) GetImageHeight() *int32 {
    return obj.ImageHeight
}

func (obj *BmpSaveOptionsData) SetImageHeight(value *int32) {
    obj.ImageHeight = value
}

func (obj *BmpSaveOptionsData) GetImageWidth() *int32 {
    return obj.ImageWidth
}

func (obj *BmpSaveOptionsData) SetImageWidth(value *int32) {
    obj.ImageWidth = value
}

func (obj *BmpSaveOptionsData) GetUseGdiEmfRenderer() *bool {
    return obj.UseGdiEmfRenderer
}

func (obj *BmpSaveOptionsData) SetUseGdiEmfRenderer(value *bool) {
    obj.UseGdiEmfRenderer = value
}

func (obj *BmpSaveOptionsData) GetSaveFormat() *string {
    return obj.SaveFormat
}

func (obj *BmpSaveOptionsData) SetSaveFormat(value *string) {
    obj.SaveFormat = value
}

