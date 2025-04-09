/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="svg_save_options_data.go">
 *   Copyright (c) 2025 Aspose.Words for Cloud
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

// Container class for svg save options.

type ISvgSaveOptionsData interface {
    IsSvgSaveOptionsData() bool
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
    GetUpdateAmbiguousTextFont() *bool
    SetUpdateAmbiguousTextFont(value *bool)
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
    GetExportEmbeddedImages() *bool
    SetExportEmbeddedImages(value *bool)
    GetFitToViewPort() *bool
    SetFitToViewPort(value *bool)
    GetIdPrefix() *string
    SetIdPrefix(value *string)
    GetMaxImageResolution() *int32
    SetMaxImageResolution(value *int32)
    GetResourcesFolder() *string
    SetResourcesFolder(value *string)
    GetResourcesFolderAlias() *string
    SetResourcesFolderAlias(value *string)
    GetShowPageBorder() *bool
    SetShowPageBorder(value *bool)
    GetTextOutputMode() *string
    SetTextOutputMode(value *string)
    GetSaveFormat() *string
    SetSaveFormat(value *string)
}

type SvgSaveOptionsData struct {
    // Container class for svg save options.
    AllowEmbeddingPostScriptFonts *bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for svg save options.
    CustomTimeZoneInfoData ITimeZoneInfoData `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for svg save options.
    Dml3DEffectsRenderingMode *string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for svg save options.
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for svg save options.
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // Container class for svg save options.
    FileName *string `json:"FileName,omitempty"`

    // Container class for svg save options.
    ImlRenderingMode *string `json:"ImlRenderingMode,omitempty"`

    // Container class for svg save options.
    UpdateAmbiguousTextFont *bool `json:"UpdateAmbiguousTextFont,omitempty"`

    // Container class for svg save options.
    UpdateCreatedTimeProperty *bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for svg save options.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // Container class for svg save options.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for svg save options.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for svg save options.
    ZipOutput *bool `json:"ZipOutput,omitempty"`

    // Container class for svg save options.
    ColorMode *string `json:"ColorMode,omitempty"`

    // Container class for svg save options.
    JpegQuality *int32 `json:"JpegQuality,omitempty"`

    // Container class for svg save options.
    MetafileRenderingOptions IMetafileRenderingOptionsData `json:"MetafileRenderingOptions,omitempty"`

    // Container class for svg save options.
    NumeralFormat *string `json:"NumeralFormat,omitempty"`

    // Container class for svg save options.
    OptimizeOutput *bool `json:"OptimizeOutput,omitempty"`

    // Container class for svg save options.
    PageCount *int32 `json:"PageCount,omitempty"`

    // Container class for svg save options.
    PageIndex *int32 `json:"PageIndex,omitempty"`

    // Container class for svg save options.
    ExportEmbeddedImages *bool `json:"ExportEmbeddedImages,omitempty"`

    // Container class for svg save options.
    FitToViewPort *bool `json:"FitToViewPort,omitempty"`

    // Container class for svg save options.
    IdPrefix *string `json:"IdPrefix,omitempty"`

    // Container class for svg save options.
    MaxImageResolution *int32 `json:"MaxImageResolution,omitempty"`

    // Container class for svg save options.
    ResourcesFolder *string `json:"ResourcesFolder,omitempty"`

    // Container class for svg save options.
    ResourcesFolderAlias *string `json:"ResourcesFolderAlias,omitempty"`

    // Container class for svg save options.
    ShowPageBorder *bool `json:"ShowPageBorder,omitempty"`

    // Container class for svg save options.
    TextOutputMode *string `json:"TextOutputMode,omitempty"`

    // Container class for svg save options.
    SaveFormat *string `json:"SaveFormat,omitempty"`
}

func (SvgSaveOptionsData) IsSvgSaveOptionsData() bool {
    return true
}

func (SvgSaveOptionsData) IsFixedPageSaveOptionsData() bool {
    return true
}

func (SvgSaveOptionsData) IsSaveOptionsData() bool {
    return true
}

func (obj *SvgSaveOptionsData) Initialize() {
    var _SaveFormat = "svg"
    obj.SaveFormat = &_SaveFormat


    if (obj.CustomTimeZoneInfoData != nil) {
        obj.CustomTimeZoneInfoData.Initialize()
    }

    if (obj.MetafileRenderingOptions != nil) {
        obj.MetafileRenderingOptions.Initialize()
    }


}

func (obj *SvgSaveOptionsData) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["UpdateAmbiguousTextFont"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UpdateAmbiguousTextFont = &parsedValue
        }

    } else if jsonValue, exists := json["updateAmbiguousTextFont"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UpdateAmbiguousTextFont = &parsedValue
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

    if jsonValue, exists := json["ExportEmbeddedImages"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportEmbeddedImages = &parsedValue
        }

    } else if jsonValue, exists := json["exportEmbeddedImages"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportEmbeddedImages = &parsedValue
        }

    }

    if jsonValue, exists := json["FitToViewPort"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.FitToViewPort = &parsedValue
        }

    } else if jsonValue, exists := json["fitToViewPort"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.FitToViewPort = &parsedValue
        }

    }

    if jsonValue, exists := json["IdPrefix"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.IdPrefix = &parsedValue
        }

    } else if jsonValue, exists := json["idPrefix"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.IdPrefix = &parsedValue
        }

    }

    if jsonValue, exists := json["MaxImageResolution"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.MaxImageResolution = new(int32)
            *obj.MaxImageResolution = int32(parsedValue)
        }

    } else if jsonValue, exists := json["maxImageResolution"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.MaxImageResolution = new(int32)
            *obj.MaxImageResolution = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["ResourcesFolder"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ResourcesFolder = &parsedValue
        }

    } else if jsonValue, exists := json["resourcesFolder"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ResourcesFolder = &parsedValue
        }

    }

    if jsonValue, exists := json["ResourcesFolderAlias"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ResourcesFolderAlias = &parsedValue
        }

    } else if jsonValue, exists := json["resourcesFolderAlias"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ResourcesFolderAlias = &parsedValue
        }

    }

    if jsonValue, exists := json["ShowPageBorder"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ShowPageBorder = &parsedValue
        }

    } else if jsonValue, exists := json["showPageBorder"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ShowPageBorder = &parsedValue
        }

    }

    if jsonValue, exists := json["TextOutputMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.TextOutputMode = &parsedValue
        }

    } else if jsonValue, exists := json["textOutputMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.TextOutputMode = &parsedValue
        }

    }
}

func (obj *SvgSaveOptionsData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *SvgSaveOptionsData) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.FileName == nil {
        return errors.New("Property FileName in SvgSaveOptionsData is required.")
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

func (obj *SvgSaveOptionsData) GetAllowEmbeddingPostScriptFonts() *bool {
    return obj.AllowEmbeddingPostScriptFonts
}

func (obj *SvgSaveOptionsData) SetAllowEmbeddingPostScriptFonts(value *bool) {
    obj.AllowEmbeddingPostScriptFonts = value
}

func (obj *SvgSaveOptionsData) GetCustomTimeZoneInfoData() ITimeZoneInfoData {
    return obj.CustomTimeZoneInfoData
}

func (obj *SvgSaveOptionsData) SetCustomTimeZoneInfoData(value ITimeZoneInfoData) {
    obj.CustomTimeZoneInfoData = value
}

func (obj *SvgSaveOptionsData) GetDml3DEffectsRenderingMode() *string {
    return obj.Dml3DEffectsRenderingMode
}

func (obj *SvgSaveOptionsData) SetDml3DEffectsRenderingMode(value *string) {
    obj.Dml3DEffectsRenderingMode = value
}

func (obj *SvgSaveOptionsData) GetDmlEffectsRenderingMode() *string {
    return obj.DmlEffectsRenderingMode
}

func (obj *SvgSaveOptionsData) SetDmlEffectsRenderingMode(value *string) {
    obj.DmlEffectsRenderingMode = value
}

func (obj *SvgSaveOptionsData) GetDmlRenderingMode() *string {
    return obj.DmlRenderingMode
}

func (obj *SvgSaveOptionsData) SetDmlRenderingMode(value *string) {
    obj.DmlRenderingMode = value
}

func (obj *SvgSaveOptionsData) GetFileName() *string {
    return obj.FileName
}

func (obj *SvgSaveOptionsData) SetFileName(value *string) {
    obj.FileName = value
}

func (obj *SvgSaveOptionsData) GetImlRenderingMode() *string {
    return obj.ImlRenderingMode
}

func (obj *SvgSaveOptionsData) SetImlRenderingMode(value *string) {
    obj.ImlRenderingMode = value
}

func (obj *SvgSaveOptionsData) GetUpdateAmbiguousTextFont() *bool {
    return obj.UpdateAmbiguousTextFont
}

func (obj *SvgSaveOptionsData) SetUpdateAmbiguousTextFont(value *bool) {
    obj.UpdateAmbiguousTextFont = value
}

func (obj *SvgSaveOptionsData) GetUpdateCreatedTimeProperty() *bool {
    return obj.UpdateCreatedTimeProperty
}

func (obj *SvgSaveOptionsData) SetUpdateCreatedTimeProperty(value *bool) {
    obj.UpdateCreatedTimeProperty = value
}

func (obj *SvgSaveOptionsData) GetUpdateFields() *bool {
    return obj.UpdateFields
}

func (obj *SvgSaveOptionsData) SetUpdateFields(value *bool) {
    obj.UpdateFields = value
}

func (obj *SvgSaveOptionsData) GetUpdateLastPrintedProperty() *bool {
    return obj.UpdateLastPrintedProperty
}

func (obj *SvgSaveOptionsData) SetUpdateLastPrintedProperty(value *bool) {
    obj.UpdateLastPrintedProperty = value
}

func (obj *SvgSaveOptionsData) GetUpdateLastSavedTimeProperty() *bool {
    return obj.UpdateLastSavedTimeProperty
}

func (obj *SvgSaveOptionsData) SetUpdateLastSavedTimeProperty(value *bool) {
    obj.UpdateLastSavedTimeProperty = value
}

func (obj *SvgSaveOptionsData) GetZipOutput() *bool {
    return obj.ZipOutput
}

func (obj *SvgSaveOptionsData) SetZipOutput(value *bool) {
    obj.ZipOutput = value
}

func (obj *SvgSaveOptionsData) GetColorMode() *string {
    return obj.ColorMode
}

func (obj *SvgSaveOptionsData) SetColorMode(value *string) {
    obj.ColorMode = value
}

func (obj *SvgSaveOptionsData) GetJpegQuality() *int32 {
    return obj.JpegQuality
}

func (obj *SvgSaveOptionsData) SetJpegQuality(value *int32) {
    obj.JpegQuality = value
}

func (obj *SvgSaveOptionsData) GetMetafileRenderingOptions() IMetafileRenderingOptionsData {
    return obj.MetafileRenderingOptions
}

func (obj *SvgSaveOptionsData) SetMetafileRenderingOptions(value IMetafileRenderingOptionsData) {
    obj.MetafileRenderingOptions = value
}

func (obj *SvgSaveOptionsData) GetNumeralFormat() *string {
    return obj.NumeralFormat
}

func (obj *SvgSaveOptionsData) SetNumeralFormat(value *string) {
    obj.NumeralFormat = value
}

func (obj *SvgSaveOptionsData) GetOptimizeOutput() *bool {
    return obj.OptimizeOutput
}

func (obj *SvgSaveOptionsData) SetOptimizeOutput(value *bool) {
    obj.OptimizeOutput = value
}

func (obj *SvgSaveOptionsData) GetPageCount() *int32 {
    return obj.PageCount
}

func (obj *SvgSaveOptionsData) SetPageCount(value *int32) {
    obj.PageCount = value
}

func (obj *SvgSaveOptionsData) GetPageIndex() *int32 {
    return obj.PageIndex
}

func (obj *SvgSaveOptionsData) SetPageIndex(value *int32) {
    obj.PageIndex = value
}

func (obj *SvgSaveOptionsData) GetExportEmbeddedImages() *bool {
    return obj.ExportEmbeddedImages
}

func (obj *SvgSaveOptionsData) SetExportEmbeddedImages(value *bool) {
    obj.ExportEmbeddedImages = value
}

func (obj *SvgSaveOptionsData) GetFitToViewPort() *bool {
    return obj.FitToViewPort
}

func (obj *SvgSaveOptionsData) SetFitToViewPort(value *bool) {
    obj.FitToViewPort = value
}

func (obj *SvgSaveOptionsData) GetIdPrefix() *string {
    return obj.IdPrefix
}

func (obj *SvgSaveOptionsData) SetIdPrefix(value *string) {
    obj.IdPrefix = value
}

func (obj *SvgSaveOptionsData) GetMaxImageResolution() *int32 {
    return obj.MaxImageResolution
}

func (obj *SvgSaveOptionsData) SetMaxImageResolution(value *int32) {
    obj.MaxImageResolution = value
}

func (obj *SvgSaveOptionsData) GetResourcesFolder() *string {
    return obj.ResourcesFolder
}

func (obj *SvgSaveOptionsData) SetResourcesFolder(value *string) {
    obj.ResourcesFolder = value
}

func (obj *SvgSaveOptionsData) GetResourcesFolderAlias() *string {
    return obj.ResourcesFolderAlias
}

func (obj *SvgSaveOptionsData) SetResourcesFolderAlias(value *string) {
    obj.ResourcesFolderAlias = value
}

func (obj *SvgSaveOptionsData) GetShowPageBorder() *bool {
    return obj.ShowPageBorder
}

func (obj *SvgSaveOptionsData) SetShowPageBorder(value *bool) {
    obj.ShowPageBorder = value
}

func (obj *SvgSaveOptionsData) GetTextOutputMode() *string {
    return obj.TextOutputMode
}

func (obj *SvgSaveOptionsData) SetTextOutputMode(value *string) {
    obj.TextOutputMode = value
}

func (obj *SvgSaveOptionsData) GetSaveFormat() *string {
    return obj.SaveFormat
}

func (obj *SvgSaveOptionsData) SetSaveFormat(value *string) {
    obj.SaveFormat = value
}

