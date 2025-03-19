/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="xps_save_options_data.go">
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

// Container class for xps save options.

type IXpsSaveOptionsData interface {
    IsXpsSaveOptionsData() bool
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
    GetBookmarksOutlineLevel() *int32
    SetBookmarksOutlineLevel(value *int32)
    GetDigitalSignatureDetails() IDigitalSignatureDetails
    SetDigitalSignatureDetails(value IDigitalSignatureDetails)
    GetHeadingsOutlineLevels() *int32
    SetHeadingsOutlineLevels(value *int32)
    GetOutlineOptions() IOutlineOptionsData
    SetOutlineOptions(value IOutlineOptionsData)
    GetUseBookFoldPrintingSettings() *bool
    SetUseBookFoldPrintingSettings(value *bool)
    GetSaveFormat() *string
    SetSaveFormat(value *string)
}

type XpsSaveOptionsData struct {
    // Container class for xps save options.
    AllowEmbeddingPostScriptFonts *bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for xps save options.
    CustomTimeZoneInfoData ITimeZoneInfoData `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for xps save options.
    Dml3DEffectsRenderingMode *string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for xps save options.
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for xps save options.
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // Container class for xps save options.
    FileName *string `json:"FileName,omitempty"`

    // Container class for xps save options.
    ImlRenderingMode *string `json:"ImlRenderingMode,omitempty"`

    // Container class for xps save options.
    UpdateAmbiguousTextFont *bool `json:"UpdateAmbiguousTextFont,omitempty"`

    // Container class for xps save options.
    UpdateCreatedTimeProperty *bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for xps save options.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // Container class for xps save options.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for xps save options.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for xps save options.
    ZipOutput *bool `json:"ZipOutput,omitempty"`

    // Container class for xps save options.
    ColorMode *string `json:"ColorMode,omitempty"`

    // Container class for xps save options.
    JpegQuality *int32 `json:"JpegQuality,omitempty"`

    // Container class for xps save options.
    MetafileRenderingOptions IMetafileRenderingOptionsData `json:"MetafileRenderingOptions,omitempty"`

    // Container class for xps save options.
    NumeralFormat *string `json:"NumeralFormat,omitempty"`

    // Container class for xps save options.
    OptimizeOutput *bool `json:"OptimizeOutput,omitempty"`

    // Container class for xps save options.
    PageCount *int32 `json:"PageCount,omitempty"`

    // Container class for xps save options.
    PageIndex *int32 `json:"PageIndex,omitempty"`

    // Container class for xps save options.
    BookmarksOutlineLevel *int32 `json:"BookmarksOutlineLevel,omitempty"`

    // Container class for xps save options.
    DigitalSignatureDetails IDigitalSignatureDetails `json:"DigitalSignatureDetails,omitempty"`

    // Container class for xps save options.
    HeadingsOutlineLevels *int32 `json:"HeadingsOutlineLevels,omitempty"`

    // Container class for xps save options.
    OutlineOptions IOutlineOptionsData `json:"OutlineOptions,omitempty"`

    // Container class for xps save options.
    UseBookFoldPrintingSettings *bool `json:"UseBookFoldPrintingSettings,omitempty"`

    // Container class for xps save options.
    SaveFormat *string `json:"SaveFormat,omitempty"`
}

func (XpsSaveOptionsData) IsXpsSaveOptionsData() bool {
    return true
}

func (XpsSaveOptionsData) IsFixedPageSaveOptionsData() bool {
    return true
}

func (XpsSaveOptionsData) IsSaveOptionsData() bool {
    return true
}

func (obj *XpsSaveOptionsData) Initialize() {
    var _SaveFormat = "xps"
    obj.SaveFormat = &_SaveFormat


    if (obj.CustomTimeZoneInfoData != nil) {
        obj.CustomTimeZoneInfoData.Initialize()
    }

    if (obj.MetafileRenderingOptions != nil) {
        obj.MetafileRenderingOptions.Initialize()
    }

    if (obj.DigitalSignatureDetails != nil) {
        obj.DigitalSignatureDetails.Initialize()
    }

    if (obj.OutlineOptions != nil) {
        obj.OutlineOptions.Initialize()
    }


}

func (obj *XpsSaveOptionsData) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["BookmarksOutlineLevel"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.BookmarksOutlineLevel = new(int32)
            *obj.BookmarksOutlineLevel = int32(parsedValue)
        }

    } else if jsonValue, exists := json["bookmarksOutlineLevel"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.BookmarksOutlineLevel = new(int32)
            *obj.BookmarksOutlineLevel = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["DigitalSignatureDetails"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IDigitalSignatureDetails = new(DigitalSignatureDetails)
            modelInstance.Deserialize(parsedValue)
            obj.DigitalSignatureDetails = modelInstance
        }

    } else if jsonValue, exists := json["digitalSignatureDetails"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IDigitalSignatureDetails = new(DigitalSignatureDetails)
            modelInstance.Deserialize(parsedValue)
            obj.DigitalSignatureDetails = modelInstance
        }

    }

    if jsonValue, exists := json["HeadingsOutlineLevels"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.HeadingsOutlineLevels = new(int32)
            *obj.HeadingsOutlineLevels = int32(parsedValue)
        }

    } else if jsonValue, exists := json["headingsOutlineLevels"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.HeadingsOutlineLevels = new(int32)
            *obj.HeadingsOutlineLevels = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["OutlineOptions"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IOutlineOptionsData = new(OutlineOptionsData)
            modelInstance.Deserialize(parsedValue)
            obj.OutlineOptions = modelInstance
        }

    } else if jsonValue, exists := json["outlineOptions"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IOutlineOptionsData = new(OutlineOptionsData)
            modelInstance.Deserialize(parsedValue)
            obj.OutlineOptions = modelInstance
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

func (obj *XpsSaveOptionsData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *XpsSaveOptionsData) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.FileName == nil {
        return errors.New("Property FileName in XpsSaveOptionsData is required.")
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
    if obj.DigitalSignatureDetails != nil {
        if err := obj.DigitalSignatureDetails.Validate(); err != nil {
            return err
        }
    }
    if obj.OutlineOptions != nil {
        if err := obj.OutlineOptions.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *XpsSaveOptionsData) GetAllowEmbeddingPostScriptFonts() *bool {
    return obj.AllowEmbeddingPostScriptFonts
}

func (obj *XpsSaveOptionsData) SetAllowEmbeddingPostScriptFonts(value *bool) {
    obj.AllowEmbeddingPostScriptFonts = value
}

func (obj *XpsSaveOptionsData) GetCustomTimeZoneInfoData() ITimeZoneInfoData {
    return obj.CustomTimeZoneInfoData
}

func (obj *XpsSaveOptionsData) SetCustomTimeZoneInfoData(value ITimeZoneInfoData) {
    obj.CustomTimeZoneInfoData = value
}

func (obj *XpsSaveOptionsData) GetDml3DEffectsRenderingMode() *string {
    return obj.Dml3DEffectsRenderingMode
}

func (obj *XpsSaveOptionsData) SetDml3DEffectsRenderingMode(value *string) {
    obj.Dml3DEffectsRenderingMode = value
}

func (obj *XpsSaveOptionsData) GetDmlEffectsRenderingMode() *string {
    return obj.DmlEffectsRenderingMode
}

func (obj *XpsSaveOptionsData) SetDmlEffectsRenderingMode(value *string) {
    obj.DmlEffectsRenderingMode = value
}

func (obj *XpsSaveOptionsData) GetDmlRenderingMode() *string {
    return obj.DmlRenderingMode
}

func (obj *XpsSaveOptionsData) SetDmlRenderingMode(value *string) {
    obj.DmlRenderingMode = value
}

func (obj *XpsSaveOptionsData) GetFileName() *string {
    return obj.FileName
}

func (obj *XpsSaveOptionsData) SetFileName(value *string) {
    obj.FileName = value
}

func (obj *XpsSaveOptionsData) GetImlRenderingMode() *string {
    return obj.ImlRenderingMode
}

func (obj *XpsSaveOptionsData) SetImlRenderingMode(value *string) {
    obj.ImlRenderingMode = value
}

func (obj *XpsSaveOptionsData) GetUpdateAmbiguousTextFont() *bool {
    return obj.UpdateAmbiguousTextFont
}

func (obj *XpsSaveOptionsData) SetUpdateAmbiguousTextFont(value *bool) {
    obj.UpdateAmbiguousTextFont = value
}

func (obj *XpsSaveOptionsData) GetUpdateCreatedTimeProperty() *bool {
    return obj.UpdateCreatedTimeProperty
}

func (obj *XpsSaveOptionsData) SetUpdateCreatedTimeProperty(value *bool) {
    obj.UpdateCreatedTimeProperty = value
}

func (obj *XpsSaveOptionsData) GetUpdateFields() *bool {
    return obj.UpdateFields
}

func (obj *XpsSaveOptionsData) SetUpdateFields(value *bool) {
    obj.UpdateFields = value
}

func (obj *XpsSaveOptionsData) GetUpdateLastPrintedProperty() *bool {
    return obj.UpdateLastPrintedProperty
}

func (obj *XpsSaveOptionsData) SetUpdateLastPrintedProperty(value *bool) {
    obj.UpdateLastPrintedProperty = value
}

func (obj *XpsSaveOptionsData) GetUpdateLastSavedTimeProperty() *bool {
    return obj.UpdateLastSavedTimeProperty
}

func (obj *XpsSaveOptionsData) SetUpdateLastSavedTimeProperty(value *bool) {
    obj.UpdateLastSavedTimeProperty = value
}

func (obj *XpsSaveOptionsData) GetZipOutput() *bool {
    return obj.ZipOutput
}

func (obj *XpsSaveOptionsData) SetZipOutput(value *bool) {
    obj.ZipOutput = value
}

func (obj *XpsSaveOptionsData) GetColorMode() *string {
    return obj.ColorMode
}

func (obj *XpsSaveOptionsData) SetColorMode(value *string) {
    obj.ColorMode = value
}

func (obj *XpsSaveOptionsData) GetJpegQuality() *int32 {
    return obj.JpegQuality
}

func (obj *XpsSaveOptionsData) SetJpegQuality(value *int32) {
    obj.JpegQuality = value
}

func (obj *XpsSaveOptionsData) GetMetafileRenderingOptions() IMetafileRenderingOptionsData {
    return obj.MetafileRenderingOptions
}

func (obj *XpsSaveOptionsData) SetMetafileRenderingOptions(value IMetafileRenderingOptionsData) {
    obj.MetafileRenderingOptions = value
}

func (obj *XpsSaveOptionsData) GetNumeralFormat() *string {
    return obj.NumeralFormat
}

func (obj *XpsSaveOptionsData) SetNumeralFormat(value *string) {
    obj.NumeralFormat = value
}

func (obj *XpsSaveOptionsData) GetOptimizeOutput() *bool {
    return obj.OptimizeOutput
}

func (obj *XpsSaveOptionsData) SetOptimizeOutput(value *bool) {
    obj.OptimizeOutput = value
}

func (obj *XpsSaveOptionsData) GetPageCount() *int32 {
    return obj.PageCount
}

func (obj *XpsSaveOptionsData) SetPageCount(value *int32) {
    obj.PageCount = value
}

func (obj *XpsSaveOptionsData) GetPageIndex() *int32 {
    return obj.PageIndex
}

func (obj *XpsSaveOptionsData) SetPageIndex(value *int32) {
    obj.PageIndex = value
}

func (obj *XpsSaveOptionsData) GetBookmarksOutlineLevel() *int32 {
    return obj.BookmarksOutlineLevel
}

func (obj *XpsSaveOptionsData) SetBookmarksOutlineLevel(value *int32) {
    obj.BookmarksOutlineLevel = value
}

func (obj *XpsSaveOptionsData) GetDigitalSignatureDetails() IDigitalSignatureDetails {
    return obj.DigitalSignatureDetails
}

func (obj *XpsSaveOptionsData) SetDigitalSignatureDetails(value IDigitalSignatureDetails) {
    obj.DigitalSignatureDetails = value
}

func (obj *XpsSaveOptionsData) GetHeadingsOutlineLevels() *int32 {
    return obj.HeadingsOutlineLevels
}

func (obj *XpsSaveOptionsData) SetHeadingsOutlineLevels(value *int32) {
    obj.HeadingsOutlineLevels = value
}

func (obj *XpsSaveOptionsData) GetOutlineOptions() IOutlineOptionsData {
    return obj.OutlineOptions
}

func (obj *XpsSaveOptionsData) SetOutlineOptions(value IOutlineOptionsData) {
    obj.OutlineOptions = value
}

func (obj *XpsSaveOptionsData) GetUseBookFoldPrintingSettings() *bool {
    return obj.UseBookFoldPrintingSettings
}

func (obj *XpsSaveOptionsData) SetUseBookFoldPrintingSettings(value *bool) {
    obj.UseBookFoldPrintingSettings = value
}

func (obj *XpsSaveOptionsData) GetSaveFormat() *string {
    return obj.SaveFormat
}

func (obj *XpsSaveOptionsData) SetSaveFormat(value *string) {
    obj.SaveFormat = value
}

