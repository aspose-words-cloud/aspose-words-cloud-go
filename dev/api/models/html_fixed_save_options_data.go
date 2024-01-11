/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="html_fixed_save_options_data.go">
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

// Container class for fixed html save options.

type IHtmlFixedSaveOptionsData interface {
    IsHtmlFixedSaveOptionsData() bool
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
    GetCssClassNamesPrefix() *string
    SetCssClassNamesPrefix(value *string)
    GetEncoding() *string
    SetEncoding(value *string)
    GetExportEmbeddedCss() *bool
    SetExportEmbeddedCss(value *bool)
    GetExportEmbeddedFonts() *bool
    SetExportEmbeddedFonts(value *bool)
    GetExportEmbeddedImages() *bool
    SetExportEmbeddedImages(value *bool)
    GetExportFormFields() *bool
    SetExportFormFields(value *bool)
    GetFontFormat() *string
    SetFontFormat(value *string)
    GetPageHorizontalAlignment() *string
    SetPageHorizontalAlignment(value *string)
    GetPageMargins() *float64
    SetPageMargins(value *float64)
    GetResourcesFolder() *string
    SetResourcesFolder(value *string)
    GetResourcesFolderAlias() *string
    SetResourcesFolderAlias(value *string)
    GetSaveFontFaceCssSeparately() *bool
    SetSaveFontFaceCssSeparately(value *bool)
    GetShowPageBorder() *bool
    SetShowPageBorder(value *bool)
    GetUseTargetMachineFonts() *bool
    SetUseTargetMachineFonts(value *bool)
    GetSaveFormat() *string
    SetSaveFormat(value *string)
}

type HtmlFixedSaveOptionsData struct {
    // Container class for fixed html save options.
    AllowEmbeddingPostScriptFonts *bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for fixed html save options.
    CustomTimeZoneInfoData ITimeZoneInfoData `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for fixed html save options.
    Dml3DEffectsRenderingMode *string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for fixed html save options.
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for fixed html save options.
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // Container class for fixed html save options.
    FileName *string `json:"FileName,omitempty"`

    // Container class for fixed html save options.
    ImlRenderingMode *string `json:"ImlRenderingMode,omitempty"`

    // Container class for fixed html save options.
    UpdateCreatedTimeProperty *bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for fixed html save options.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // Container class for fixed html save options.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for fixed html save options.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for fixed html save options.
    ZipOutput *bool `json:"ZipOutput,omitempty"`

    // Container class for fixed html save options.
    ColorMode *string `json:"ColorMode,omitempty"`

    // Container class for fixed html save options.
    JpegQuality *int32 `json:"JpegQuality,omitempty"`

    // Container class for fixed html save options.
    MetafileRenderingOptions IMetafileRenderingOptionsData `json:"MetafileRenderingOptions,omitempty"`

    // Container class for fixed html save options.
    NumeralFormat *string `json:"NumeralFormat,omitempty"`

    // Container class for fixed html save options.
    OptimizeOutput *bool `json:"OptimizeOutput,omitempty"`

    // Container class for fixed html save options.
    PageCount *int32 `json:"PageCount,omitempty"`

    // Container class for fixed html save options.
    PageIndex *int32 `json:"PageIndex,omitempty"`

    // Container class for fixed html save options.
    CssClassNamesPrefix *string `json:"CssClassNamesPrefix,omitempty"`

    // Container class for fixed html save options.
    Encoding *string `json:"Encoding,omitempty"`

    // Container class for fixed html save options.
    ExportEmbeddedCss *bool `json:"ExportEmbeddedCss,omitempty"`

    // Container class for fixed html save options.
    ExportEmbeddedFonts *bool `json:"ExportEmbeddedFonts,omitempty"`

    // Container class for fixed html save options.
    ExportEmbeddedImages *bool `json:"ExportEmbeddedImages,omitempty"`

    // Container class for fixed html save options.
    ExportFormFields *bool `json:"ExportFormFields,omitempty"`

    // Container class for fixed html save options.
    FontFormat *string `json:"FontFormat,omitempty"`

    // Container class for fixed html save options.
    PageHorizontalAlignment *string `json:"PageHorizontalAlignment,omitempty"`

    // Container class for fixed html save options.
    PageMargins *float64 `json:"PageMargins,omitempty"`

    // Container class for fixed html save options.
    ResourcesFolder *string `json:"ResourcesFolder,omitempty"`

    // Container class for fixed html save options.
    ResourcesFolderAlias *string `json:"ResourcesFolderAlias,omitempty"`

    // Container class for fixed html save options.
    SaveFontFaceCssSeparately *bool `json:"SaveFontFaceCssSeparately,omitempty"`

    // Container class for fixed html save options.
    ShowPageBorder *bool `json:"ShowPageBorder,omitempty"`

    // Container class for fixed html save options.
    UseTargetMachineFonts *bool `json:"UseTargetMachineFonts,omitempty"`

    // Container class for fixed html save options.
    SaveFormat *string `json:"SaveFormat,omitempty"`
}

func (HtmlFixedSaveOptionsData) IsHtmlFixedSaveOptionsData() bool {
    return true
}

func (HtmlFixedSaveOptionsData) IsFixedPageSaveOptionsData() bool {
    return true
}

func (HtmlFixedSaveOptionsData) IsSaveOptionsData() bool {
    return true
}

func (obj *HtmlFixedSaveOptionsData) Initialize() {
    var _SaveFormat = "htmlfixed"
    obj.SaveFormat = &_SaveFormat


    if (obj.CustomTimeZoneInfoData != nil) {
        obj.CustomTimeZoneInfoData.Initialize()
    }

    if (obj.MetafileRenderingOptions != nil) {
        obj.MetafileRenderingOptions.Initialize()
    }


}

func (obj *HtmlFixedSaveOptionsData) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["CssClassNamesPrefix"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.CssClassNamesPrefix = &parsedValue
        }

    } else if jsonValue, exists := json["cssClassNamesPrefix"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.CssClassNamesPrefix = &parsedValue
        }

    }

    if jsonValue, exists := json["Encoding"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Encoding = &parsedValue
        }

    } else if jsonValue, exists := json["encoding"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Encoding = &parsedValue
        }

    }

    if jsonValue, exists := json["ExportEmbeddedCss"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportEmbeddedCss = &parsedValue
        }

    } else if jsonValue, exists := json["exportEmbeddedCss"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportEmbeddedCss = &parsedValue
        }

    }

    if jsonValue, exists := json["ExportEmbeddedFonts"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportEmbeddedFonts = &parsedValue
        }

    } else if jsonValue, exists := json["exportEmbeddedFonts"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportEmbeddedFonts = &parsedValue
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

    if jsonValue, exists := json["ExportFormFields"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportFormFields = &parsedValue
        }

    } else if jsonValue, exists := json["exportFormFields"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportFormFields = &parsedValue
        }

    }

    if jsonValue, exists := json["FontFormat"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FontFormat = &parsedValue
        }

    } else if jsonValue, exists := json["fontFormat"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FontFormat = &parsedValue
        }

    }

    if jsonValue, exists := json["PageHorizontalAlignment"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.PageHorizontalAlignment = &parsedValue
        }

    } else if jsonValue, exists := json["pageHorizontalAlignment"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.PageHorizontalAlignment = &parsedValue
        }

    }

    if jsonValue, exists := json["PageMargins"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.PageMargins = &parsedValue
        }

    } else if jsonValue, exists := json["pageMargins"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.PageMargins = &parsedValue
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

    if jsonValue, exists := json["SaveFontFaceCssSeparately"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.SaveFontFaceCssSeparately = &parsedValue
        }

    } else if jsonValue, exists := json["saveFontFaceCssSeparately"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.SaveFontFaceCssSeparately = &parsedValue
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

    if jsonValue, exists := json["UseTargetMachineFonts"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UseTargetMachineFonts = &parsedValue
        }

    } else if jsonValue, exists := json["useTargetMachineFonts"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UseTargetMachineFonts = &parsedValue
        }

    }
}

func (obj *HtmlFixedSaveOptionsData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *HtmlFixedSaveOptionsData) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.FileName == nil {
        return errors.New("Property FileName in HtmlFixedSaveOptionsData is required.")
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

func (obj *HtmlFixedSaveOptionsData) GetAllowEmbeddingPostScriptFonts() *bool {
    return obj.AllowEmbeddingPostScriptFonts
}

func (obj *HtmlFixedSaveOptionsData) SetAllowEmbeddingPostScriptFonts(value *bool) {
    obj.AllowEmbeddingPostScriptFonts = value
}

func (obj *HtmlFixedSaveOptionsData) GetCustomTimeZoneInfoData() ITimeZoneInfoData {
    return obj.CustomTimeZoneInfoData
}

func (obj *HtmlFixedSaveOptionsData) SetCustomTimeZoneInfoData(value ITimeZoneInfoData) {
    obj.CustomTimeZoneInfoData = value
}

func (obj *HtmlFixedSaveOptionsData) GetDml3DEffectsRenderingMode() *string {
    return obj.Dml3DEffectsRenderingMode
}

func (obj *HtmlFixedSaveOptionsData) SetDml3DEffectsRenderingMode(value *string) {
    obj.Dml3DEffectsRenderingMode = value
}

func (obj *HtmlFixedSaveOptionsData) GetDmlEffectsRenderingMode() *string {
    return obj.DmlEffectsRenderingMode
}

func (obj *HtmlFixedSaveOptionsData) SetDmlEffectsRenderingMode(value *string) {
    obj.DmlEffectsRenderingMode = value
}

func (obj *HtmlFixedSaveOptionsData) GetDmlRenderingMode() *string {
    return obj.DmlRenderingMode
}

func (obj *HtmlFixedSaveOptionsData) SetDmlRenderingMode(value *string) {
    obj.DmlRenderingMode = value
}

func (obj *HtmlFixedSaveOptionsData) GetFileName() *string {
    return obj.FileName
}

func (obj *HtmlFixedSaveOptionsData) SetFileName(value *string) {
    obj.FileName = value
}

func (obj *HtmlFixedSaveOptionsData) GetImlRenderingMode() *string {
    return obj.ImlRenderingMode
}

func (obj *HtmlFixedSaveOptionsData) SetImlRenderingMode(value *string) {
    obj.ImlRenderingMode = value
}

func (obj *HtmlFixedSaveOptionsData) GetUpdateCreatedTimeProperty() *bool {
    return obj.UpdateCreatedTimeProperty
}

func (obj *HtmlFixedSaveOptionsData) SetUpdateCreatedTimeProperty(value *bool) {
    obj.UpdateCreatedTimeProperty = value
}

func (obj *HtmlFixedSaveOptionsData) GetUpdateFields() *bool {
    return obj.UpdateFields
}

func (obj *HtmlFixedSaveOptionsData) SetUpdateFields(value *bool) {
    obj.UpdateFields = value
}

func (obj *HtmlFixedSaveOptionsData) GetUpdateLastPrintedProperty() *bool {
    return obj.UpdateLastPrintedProperty
}

func (obj *HtmlFixedSaveOptionsData) SetUpdateLastPrintedProperty(value *bool) {
    obj.UpdateLastPrintedProperty = value
}

func (obj *HtmlFixedSaveOptionsData) GetUpdateLastSavedTimeProperty() *bool {
    return obj.UpdateLastSavedTimeProperty
}

func (obj *HtmlFixedSaveOptionsData) SetUpdateLastSavedTimeProperty(value *bool) {
    obj.UpdateLastSavedTimeProperty = value
}

func (obj *HtmlFixedSaveOptionsData) GetZipOutput() *bool {
    return obj.ZipOutput
}

func (obj *HtmlFixedSaveOptionsData) SetZipOutput(value *bool) {
    obj.ZipOutput = value
}

func (obj *HtmlFixedSaveOptionsData) GetColorMode() *string {
    return obj.ColorMode
}

func (obj *HtmlFixedSaveOptionsData) SetColorMode(value *string) {
    obj.ColorMode = value
}

func (obj *HtmlFixedSaveOptionsData) GetJpegQuality() *int32 {
    return obj.JpegQuality
}

func (obj *HtmlFixedSaveOptionsData) SetJpegQuality(value *int32) {
    obj.JpegQuality = value
}

func (obj *HtmlFixedSaveOptionsData) GetMetafileRenderingOptions() IMetafileRenderingOptionsData {
    return obj.MetafileRenderingOptions
}

func (obj *HtmlFixedSaveOptionsData) SetMetafileRenderingOptions(value IMetafileRenderingOptionsData) {
    obj.MetafileRenderingOptions = value
}

func (obj *HtmlFixedSaveOptionsData) GetNumeralFormat() *string {
    return obj.NumeralFormat
}

func (obj *HtmlFixedSaveOptionsData) SetNumeralFormat(value *string) {
    obj.NumeralFormat = value
}

func (obj *HtmlFixedSaveOptionsData) GetOptimizeOutput() *bool {
    return obj.OptimizeOutput
}

func (obj *HtmlFixedSaveOptionsData) SetOptimizeOutput(value *bool) {
    obj.OptimizeOutput = value
}

func (obj *HtmlFixedSaveOptionsData) GetPageCount() *int32 {
    return obj.PageCount
}

func (obj *HtmlFixedSaveOptionsData) SetPageCount(value *int32) {
    obj.PageCount = value
}

func (obj *HtmlFixedSaveOptionsData) GetPageIndex() *int32 {
    return obj.PageIndex
}

func (obj *HtmlFixedSaveOptionsData) SetPageIndex(value *int32) {
    obj.PageIndex = value
}

func (obj *HtmlFixedSaveOptionsData) GetCssClassNamesPrefix() *string {
    return obj.CssClassNamesPrefix
}

func (obj *HtmlFixedSaveOptionsData) SetCssClassNamesPrefix(value *string) {
    obj.CssClassNamesPrefix = value
}

func (obj *HtmlFixedSaveOptionsData) GetEncoding() *string {
    return obj.Encoding
}

func (obj *HtmlFixedSaveOptionsData) SetEncoding(value *string) {
    obj.Encoding = value
}

func (obj *HtmlFixedSaveOptionsData) GetExportEmbeddedCss() *bool {
    return obj.ExportEmbeddedCss
}

func (obj *HtmlFixedSaveOptionsData) SetExportEmbeddedCss(value *bool) {
    obj.ExportEmbeddedCss = value
}

func (obj *HtmlFixedSaveOptionsData) GetExportEmbeddedFonts() *bool {
    return obj.ExportEmbeddedFonts
}

func (obj *HtmlFixedSaveOptionsData) SetExportEmbeddedFonts(value *bool) {
    obj.ExportEmbeddedFonts = value
}

func (obj *HtmlFixedSaveOptionsData) GetExportEmbeddedImages() *bool {
    return obj.ExportEmbeddedImages
}

func (obj *HtmlFixedSaveOptionsData) SetExportEmbeddedImages(value *bool) {
    obj.ExportEmbeddedImages = value
}

func (obj *HtmlFixedSaveOptionsData) GetExportFormFields() *bool {
    return obj.ExportFormFields
}

func (obj *HtmlFixedSaveOptionsData) SetExportFormFields(value *bool) {
    obj.ExportFormFields = value
}

func (obj *HtmlFixedSaveOptionsData) GetFontFormat() *string {
    return obj.FontFormat
}

func (obj *HtmlFixedSaveOptionsData) SetFontFormat(value *string) {
    obj.FontFormat = value
}

func (obj *HtmlFixedSaveOptionsData) GetPageHorizontalAlignment() *string {
    return obj.PageHorizontalAlignment
}

func (obj *HtmlFixedSaveOptionsData) SetPageHorizontalAlignment(value *string) {
    obj.PageHorizontalAlignment = value
}

func (obj *HtmlFixedSaveOptionsData) GetPageMargins() *float64 {
    return obj.PageMargins
}

func (obj *HtmlFixedSaveOptionsData) SetPageMargins(value *float64) {
    obj.PageMargins = value
}

func (obj *HtmlFixedSaveOptionsData) GetResourcesFolder() *string {
    return obj.ResourcesFolder
}

func (obj *HtmlFixedSaveOptionsData) SetResourcesFolder(value *string) {
    obj.ResourcesFolder = value
}

func (obj *HtmlFixedSaveOptionsData) GetResourcesFolderAlias() *string {
    return obj.ResourcesFolderAlias
}

func (obj *HtmlFixedSaveOptionsData) SetResourcesFolderAlias(value *string) {
    obj.ResourcesFolderAlias = value
}

func (obj *HtmlFixedSaveOptionsData) GetSaveFontFaceCssSeparately() *bool {
    return obj.SaveFontFaceCssSeparately
}

func (obj *HtmlFixedSaveOptionsData) SetSaveFontFaceCssSeparately(value *bool) {
    obj.SaveFontFaceCssSeparately = value
}

func (obj *HtmlFixedSaveOptionsData) GetShowPageBorder() *bool {
    return obj.ShowPageBorder
}

func (obj *HtmlFixedSaveOptionsData) SetShowPageBorder(value *bool) {
    obj.ShowPageBorder = value
}

func (obj *HtmlFixedSaveOptionsData) GetUseTargetMachineFonts() *bool {
    return obj.UseTargetMachineFonts
}

func (obj *HtmlFixedSaveOptionsData) SetUseTargetMachineFonts(value *bool) {
    obj.UseTargetMachineFonts = value
}

func (obj *HtmlFixedSaveOptionsData) GetSaveFormat() *string {
    return obj.SaveFormat
}

func (obj *HtmlFixedSaveOptionsData) SetSaveFormat(value *string) {
    obj.SaveFormat = value
}

