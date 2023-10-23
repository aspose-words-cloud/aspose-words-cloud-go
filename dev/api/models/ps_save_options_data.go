/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="ps_save_options_data.go">
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

// Container class for ps save options.

type IPsSaveOptionsData interface {
    IsPsSaveOptionsData() bool
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
    GetUseBookFoldPrintingSettings() *bool
    SetUseBookFoldPrintingSettings(value *bool)
    GetSaveFormat() *string
    SetSaveFormat(value *string)
}

type PsSaveOptionsData struct {
    // Container class for ps save options.
    AllowEmbeddingPostScriptFonts *bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for ps save options.
    CustomTimeZoneInfoData ITimeZoneInfoData `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for ps save options.
    Dml3DEffectsRenderingMode *string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for ps save options.
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for ps save options.
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // Container class for ps save options.
    FileName *string `json:"FileName,omitempty"`

    // Container class for ps save options.
    ImlRenderingMode *string `json:"ImlRenderingMode,omitempty"`

    // Container class for ps save options.
    UpdateCreatedTimeProperty *bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for ps save options.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // Container class for ps save options.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for ps save options.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for ps save options.
    ZipOutput *bool `json:"ZipOutput,omitempty"`

    // Container class for ps save options.
    ColorMode *string `json:"ColorMode,omitempty"`

    // Container class for ps save options.
    JpegQuality *int32 `json:"JpegQuality,omitempty"`

    // Container class for ps save options.
    MetafileRenderingOptions IMetafileRenderingOptionsData `json:"MetafileRenderingOptions,omitempty"`

    // Container class for ps save options.
    NumeralFormat *string `json:"NumeralFormat,omitempty"`

    // Container class for ps save options.
    OptimizeOutput *bool `json:"OptimizeOutput,omitempty"`

    // Container class for ps save options.
    PageCount *int32 `json:"PageCount,omitempty"`

    // Container class for ps save options.
    PageIndex *int32 `json:"PageIndex,omitempty"`

    // Container class for ps save options.
    UseBookFoldPrintingSettings *bool `json:"UseBookFoldPrintingSettings,omitempty"`

    // Container class for ps save options.
    SaveFormat *string `json:"SaveFormat,omitempty"`
}

func (PsSaveOptionsData) IsPsSaveOptionsData() bool {
    return true
}

func (PsSaveOptionsData) IsFixedPageSaveOptionsData() bool {
    return true
}

func (PsSaveOptionsData) IsSaveOptionsData() bool {
    return true
}

func (obj *PsSaveOptionsData) Initialize() {
    var _SaveFormat = "ps"
    obj.SaveFormat = &_SaveFormat


    if (obj.CustomTimeZoneInfoData != nil) {
        obj.CustomTimeZoneInfoData.Initialize()
    }

    if (obj.MetafileRenderingOptions != nil) {
        obj.MetafileRenderingOptions.Initialize()
    }


}

func (obj *PsSaveOptionsData) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["UseBookFoldPrintingSettings"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UseBookFoldPrintingSettings = &parsedValue
        }

    } else if jsonValue, exists := json["useBookFoldPrintingSettings"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UseBookFoldPrintingSettings = &parsedValue
        }

    }
}

func (obj *PsSaveOptionsData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *PsSaveOptionsData) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.FileName == nil {
        return errors.New("Property FileName in PsSaveOptionsData is required.")
    }
    if obj.AllowEmbeddingPostScriptFonts == nil {
        return errors.New("Property AllowEmbeddingPostScriptFonts in PsSaveOptionsData is required.")
    }

    if obj.CustomTimeZoneInfoData == nil {
        return errors.New("Property CustomTimeZoneInfoData in PsSaveOptionsData is required.")
    }

    if obj.CustomTimeZoneInfoData != nil {
        if err := obj.CustomTimeZoneInfoData.Validate(); err != nil {
            return err
        }
    }

    if obj.Dml3DEffectsRenderingMode == nil {
        return errors.New("Property Dml3DEffectsRenderingMode in PsSaveOptionsData is required.")
    }

    if obj.DmlEffectsRenderingMode == nil {
        return errors.New("Property DmlEffectsRenderingMode in PsSaveOptionsData is required.")
    }

    if obj.DmlRenderingMode == nil {
        return errors.New("Property DmlRenderingMode in PsSaveOptionsData is required.")
    }

    if obj.FileName == nil {
        return errors.New("Property FileName in PsSaveOptionsData is required.")
    }

    if obj.ImlRenderingMode == nil {
        return errors.New("Property ImlRenderingMode in PsSaveOptionsData is required.")
    }

    if obj.UpdateCreatedTimeProperty == nil {
        return errors.New("Property UpdateCreatedTimeProperty in PsSaveOptionsData is required.")
    }

    if obj.UpdateFields == nil {
        return errors.New("Property UpdateFields in PsSaveOptionsData is required.")
    }

    if obj.UpdateLastPrintedProperty == nil {
        return errors.New("Property UpdateLastPrintedProperty in PsSaveOptionsData is required.")
    }

    if obj.UpdateLastSavedTimeProperty == nil {
        return errors.New("Property UpdateLastSavedTimeProperty in PsSaveOptionsData is required.")
    }

    if obj.ZipOutput == nil {
        return errors.New("Property ZipOutput in PsSaveOptionsData is required.")
    }

    if obj.ColorMode == nil {
        return errors.New("Property ColorMode in PsSaveOptionsData is required.")
    }

    if obj.JpegQuality == nil {
        return errors.New("Property JpegQuality in PsSaveOptionsData is required.")
    }

    if obj.MetafileRenderingOptions == nil {
        return errors.New("Property MetafileRenderingOptions in PsSaveOptionsData is required.")
    }

    if obj.MetafileRenderingOptions != nil {
        if err := obj.MetafileRenderingOptions.Validate(); err != nil {
            return err
        }
    }

    if obj.NumeralFormat == nil {
        return errors.New("Property NumeralFormat in PsSaveOptionsData is required.")
    }

    if obj.OptimizeOutput == nil {
        return errors.New("Property OptimizeOutput in PsSaveOptionsData is required.")
    }

    if obj.PageCount == nil {
        return errors.New("Property PageCount in PsSaveOptionsData is required.")
    }

    if obj.PageIndex == nil {
        return errors.New("Property PageIndex in PsSaveOptionsData is required.")
    }

    if obj.UseBookFoldPrintingSettings == nil {
        return errors.New("Property UseBookFoldPrintingSettings in PsSaveOptionsData is required.")
    }

    if obj.SaveFormat == nil {
        return errors.New("Property SaveFormat in PsSaveOptionsData is required.")
    }

    return nil;
}

func (obj *PsSaveOptionsData) GetAllowEmbeddingPostScriptFonts() *bool {
    return obj.AllowEmbeddingPostScriptFonts
}

func (obj *PsSaveOptionsData) SetAllowEmbeddingPostScriptFonts(value *bool) {
    obj.AllowEmbeddingPostScriptFonts = value
}

func (obj *PsSaveOptionsData) GetCustomTimeZoneInfoData() ITimeZoneInfoData {
    return obj.CustomTimeZoneInfoData
}

func (obj *PsSaveOptionsData) SetCustomTimeZoneInfoData(value ITimeZoneInfoData) {
    obj.CustomTimeZoneInfoData = value
}

func (obj *PsSaveOptionsData) GetDml3DEffectsRenderingMode() *string {
    return obj.Dml3DEffectsRenderingMode
}

func (obj *PsSaveOptionsData) SetDml3DEffectsRenderingMode(value *string) {
    obj.Dml3DEffectsRenderingMode = value
}

func (obj *PsSaveOptionsData) GetDmlEffectsRenderingMode() *string {
    return obj.DmlEffectsRenderingMode
}

func (obj *PsSaveOptionsData) SetDmlEffectsRenderingMode(value *string) {
    obj.DmlEffectsRenderingMode = value
}

func (obj *PsSaveOptionsData) GetDmlRenderingMode() *string {
    return obj.DmlRenderingMode
}

func (obj *PsSaveOptionsData) SetDmlRenderingMode(value *string) {
    obj.DmlRenderingMode = value
}

func (obj *PsSaveOptionsData) GetFileName() *string {
    return obj.FileName
}

func (obj *PsSaveOptionsData) SetFileName(value *string) {
    obj.FileName = value
}

func (obj *PsSaveOptionsData) GetImlRenderingMode() *string {
    return obj.ImlRenderingMode
}

func (obj *PsSaveOptionsData) SetImlRenderingMode(value *string) {
    obj.ImlRenderingMode = value
}

func (obj *PsSaveOptionsData) GetUpdateCreatedTimeProperty() *bool {
    return obj.UpdateCreatedTimeProperty
}

func (obj *PsSaveOptionsData) SetUpdateCreatedTimeProperty(value *bool) {
    obj.UpdateCreatedTimeProperty = value
}

func (obj *PsSaveOptionsData) GetUpdateFields() *bool {
    return obj.UpdateFields
}

func (obj *PsSaveOptionsData) SetUpdateFields(value *bool) {
    obj.UpdateFields = value
}

func (obj *PsSaveOptionsData) GetUpdateLastPrintedProperty() *bool {
    return obj.UpdateLastPrintedProperty
}

func (obj *PsSaveOptionsData) SetUpdateLastPrintedProperty(value *bool) {
    obj.UpdateLastPrintedProperty = value
}

func (obj *PsSaveOptionsData) GetUpdateLastSavedTimeProperty() *bool {
    return obj.UpdateLastSavedTimeProperty
}

func (obj *PsSaveOptionsData) SetUpdateLastSavedTimeProperty(value *bool) {
    obj.UpdateLastSavedTimeProperty = value
}

func (obj *PsSaveOptionsData) GetZipOutput() *bool {
    return obj.ZipOutput
}

func (obj *PsSaveOptionsData) SetZipOutput(value *bool) {
    obj.ZipOutput = value
}

func (obj *PsSaveOptionsData) GetColorMode() *string {
    return obj.ColorMode
}

func (obj *PsSaveOptionsData) SetColorMode(value *string) {
    obj.ColorMode = value
}

func (obj *PsSaveOptionsData) GetJpegQuality() *int32 {
    return obj.JpegQuality
}

func (obj *PsSaveOptionsData) SetJpegQuality(value *int32) {
    obj.JpegQuality = value
}

func (obj *PsSaveOptionsData) GetMetafileRenderingOptions() IMetafileRenderingOptionsData {
    return obj.MetafileRenderingOptions
}

func (obj *PsSaveOptionsData) SetMetafileRenderingOptions(value IMetafileRenderingOptionsData) {
    obj.MetafileRenderingOptions = value
}

func (obj *PsSaveOptionsData) GetNumeralFormat() *string {
    return obj.NumeralFormat
}

func (obj *PsSaveOptionsData) SetNumeralFormat(value *string) {
    obj.NumeralFormat = value
}

func (obj *PsSaveOptionsData) GetOptimizeOutput() *bool {
    return obj.OptimizeOutput
}

func (obj *PsSaveOptionsData) SetOptimizeOutput(value *bool) {
    obj.OptimizeOutput = value
}

func (obj *PsSaveOptionsData) GetPageCount() *int32 {
    return obj.PageCount
}

func (obj *PsSaveOptionsData) SetPageCount(value *int32) {
    obj.PageCount = value
}

func (obj *PsSaveOptionsData) GetPageIndex() *int32 {
    return obj.PageIndex
}

func (obj *PsSaveOptionsData) SetPageIndex(value *int32) {
    obj.PageIndex = value
}

func (obj *PsSaveOptionsData) GetUseBookFoldPrintingSettings() *bool {
    return obj.UseBookFoldPrintingSettings
}

func (obj *PsSaveOptionsData) SetUseBookFoldPrintingSettings(value *bool) {
    obj.UseBookFoldPrintingSettings = value
}

func (obj *PsSaveOptionsData) GetSaveFormat() *string {
    return obj.SaveFormat
}

func (obj *PsSaveOptionsData) SetSaveFormat(value *string) {
    obj.SaveFormat = value
}

