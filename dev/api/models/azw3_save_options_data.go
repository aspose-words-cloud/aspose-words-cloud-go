/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="azw3_save_options_data.go">
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

// Container class for azw3 save options.

type IAzw3SaveOptionsData interface {
    IsAzw3SaveOptionsData() bool
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
    GetAllowNegativeIndent() *bool
    SetAllowNegativeIndent(value *bool)
    GetCssClassNamePrefix() *string
    SetCssClassNamePrefix(value *string)
    GetCssStyleSheetFileName() *string
    SetCssStyleSheetFileName(value *string)
    GetCssStyleSheetType() *string
    SetCssStyleSheetType(value *string)
    GetDocumentSplitCriteria() *string
    SetDocumentSplitCriteria(value *string)
    GetDocumentSplitHeadingLevel() *int32
    SetDocumentSplitHeadingLevel(value *int32)
    GetEncoding() *string
    SetEncoding(value *string)
    GetExportDocumentProperties() *bool
    SetExportDocumentProperties(value *bool)
    GetExportDropDownFormFieldAsText() *bool
    SetExportDropDownFormFieldAsText(value *bool)
    GetExportFontResources() *bool
    SetExportFontResources(value *bool)
    GetExportFontsAsBase64() *bool
    SetExportFontsAsBase64(value *bool)
    GetExportHeadersFootersMode() *string
    SetExportHeadersFootersMode(value *string)
    GetExportImagesAsBase64() *bool
    SetExportImagesAsBase64(value *bool)
    GetExportLanguageInformation() *bool
    SetExportLanguageInformation(value *bool)
    GetExportListLabels() *string
    SetExportListLabels(value *string)
    GetExportOriginalUrlForLinkedImages() *bool
    SetExportOriginalUrlForLinkedImages(value *bool)
    GetExportPageMargins() *bool
    SetExportPageMargins(value *bool)
    GetExportPageSetup() *bool
    SetExportPageSetup(value *bool)
    GetExportRelativeFontSize() *bool
    SetExportRelativeFontSize(value *bool)
    GetExportRoundtripInformation() *bool
    SetExportRoundtripInformation(value *bool)
    GetExportTextInputFormFieldAsText() *bool
    SetExportTextInputFormFieldAsText(value *bool)
    GetExportTocPageNumbers() *bool
    SetExportTocPageNumbers(value *bool)
    GetExportXhtmlTransitional() *bool
    SetExportXhtmlTransitional(value *bool)
    GetFontResourcesSubsettingSizeThreshold() *int32
    SetFontResourcesSubsettingSizeThreshold(value *int32)
    GetFontsFolder() *string
    SetFontsFolder(value *string)
    GetFontsFolderAlias() *string
    SetFontsFolderAlias(value *string)
    GetHtmlVersion() *string
    SetHtmlVersion(value *string)
    GetImageResolution() *int32
    SetImageResolution(value *int32)
    GetImagesFolder() *string
    SetImagesFolder(value *string)
    GetImagesFolderAlias() *string
    SetImagesFolderAlias(value *string)
    GetMetafileFormat() *string
    SetMetafileFormat(value *string)
    GetOfficeMathOutputMode() *string
    SetOfficeMathOutputMode(value *string)
    GetPrettyFormat() *bool
    SetPrettyFormat(value *bool)
    GetReplaceBackslashWithYenSign() *bool
    SetReplaceBackslashWithYenSign(value *bool)
    GetResolveFontNames() *bool
    SetResolveFontNames(value *bool)
    GetResourceFolder() *string
    SetResourceFolder(value *string)
    GetResourceFolderAlias() *string
    SetResourceFolderAlias(value *string)
    GetScaleImageToShapeSize() *bool
    SetScaleImageToShapeSize(value *bool)
    GetTableWidthOutputMode() *string
    SetTableWidthOutputMode(value *string)
    GetNavigationMapLevel() *int32
    SetNavigationMapLevel(value *int32)
    GetSaveFormat() *string
    SetSaveFormat(value *string)
}

type Azw3SaveOptionsData struct {
    // Container class for azw3 save options.
    AllowEmbeddingPostScriptFonts *bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for azw3 save options.
    CustomTimeZoneInfoData ITimeZoneInfoData `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for azw3 save options.
    Dml3DEffectsRenderingMode *string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for azw3 save options.
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for azw3 save options.
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // Container class for azw3 save options.
    FileName *string `json:"FileName,omitempty"`

    // Container class for azw3 save options.
    ImlRenderingMode *string `json:"ImlRenderingMode,omitempty"`

    // Container class for azw3 save options.
    UpdateCreatedTimeProperty *bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for azw3 save options.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // Container class for azw3 save options.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for azw3 save options.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for azw3 save options.
    ZipOutput *bool `json:"ZipOutput,omitempty"`

    // Container class for azw3 save options.
    AllowNegativeIndent *bool `json:"AllowNegativeIndent,omitempty"`

    // Container class for azw3 save options.
    CssClassNamePrefix *string `json:"CssClassNamePrefix,omitempty"`

    // Container class for azw3 save options.
    CssStyleSheetFileName *string `json:"CssStyleSheetFileName,omitempty"`

    // Container class for azw3 save options.
    CssStyleSheetType *string `json:"CssStyleSheetType,omitempty"`

    // Container class for azw3 save options.
    DocumentSplitCriteria *string `json:"DocumentSplitCriteria,omitempty"`

    // Container class for azw3 save options.
    DocumentSplitHeadingLevel *int32 `json:"DocumentSplitHeadingLevel,omitempty"`

    // Container class for azw3 save options.
    Encoding *string `json:"Encoding,omitempty"`

    // Container class for azw3 save options.
    ExportDocumentProperties *bool `json:"ExportDocumentProperties,omitempty"`

    // Container class for azw3 save options.
    ExportDropDownFormFieldAsText *bool `json:"ExportDropDownFormFieldAsText,omitempty"`

    // Container class for azw3 save options.
    ExportFontResources *bool `json:"ExportFontResources,omitempty"`

    // Container class for azw3 save options.
    ExportFontsAsBase64 *bool `json:"ExportFontsAsBase64,omitempty"`

    // Container class for azw3 save options.
    ExportHeadersFootersMode *string `json:"ExportHeadersFootersMode,omitempty"`

    // Container class for azw3 save options.
    ExportImagesAsBase64 *bool `json:"ExportImagesAsBase64,omitempty"`

    // Container class for azw3 save options.
    ExportLanguageInformation *bool `json:"ExportLanguageInformation,omitempty"`

    // Container class for azw3 save options.
    ExportListLabels *string `json:"ExportListLabels,omitempty"`

    // Container class for azw3 save options.
    ExportOriginalUrlForLinkedImages *bool `json:"ExportOriginalUrlForLinkedImages,omitempty"`

    // Container class for azw3 save options.
    ExportPageMargins *bool `json:"ExportPageMargins,omitempty"`

    // Container class for azw3 save options.
    ExportPageSetup *bool `json:"ExportPageSetup,omitempty"`

    // Container class for azw3 save options.
    ExportRelativeFontSize *bool `json:"ExportRelativeFontSize,omitempty"`

    // Container class for azw3 save options.
    ExportRoundtripInformation *bool `json:"ExportRoundtripInformation,omitempty"`

    // Container class for azw3 save options.
    ExportTextInputFormFieldAsText *bool `json:"ExportTextInputFormFieldAsText,omitempty"`

    // Container class for azw3 save options.
    ExportTocPageNumbers *bool `json:"ExportTocPageNumbers,omitempty"`

    // Container class for azw3 save options.
    ExportXhtmlTransitional *bool `json:"ExportXhtmlTransitional,omitempty"`

    // Container class for azw3 save options.
    FontResourcesSubsettingSizeThreshold *int32 `json:"FontResourcesSubsettingSizeThreshold,omitempty"`

    // Container class for azw3 save options.
    FontsFolder *string `json:"FontsFolder,omitempty"`

    // Container class for azw3 save options.
    FontsFolderAlias *string `json:"FontsFolderAlias,omitempty"`

    // Container class for azw3 save options.
    HtmlVersion *string `json:"HtmlVersion,omitempty"`

    // Container class for azw3 save options.
    ImageResolution *int32 `json:"ImageResolution,omitempty"`

    // Container class for azw3 save options.
    ImagesFolder *string `json:"ImagesFolder,omitempty"`

    // Container class for azw3 save options.
    ImagesFolderAlias *string `json:"ImagesFolderAlias,omitempty"`

    // Container class for azw3 save options.
    MetafileFormat *string `json:"MetafileFormat,omitempty"`

    // Container class for azw3 save options.
    OfficeMathOutputMode *string `json:"OfficeMathOutputMode,omitempty"`

    // Container class for azw3 save options.
    PrettyFormat *bool `json:"PrettyFormat,omitempty"`

    // Container class for azw3 save options.
    ReplaceBackslashWithYenSign *bool `json:"ReplaceBackslashWithYenSign,omitempty"`

    // Container class for azw3 save options.
    ResolveFontNames *bool `json:"ResolveFontNames,omitempty"`

    // Container class for azw3 save options.
    ResourceFolder *string `json:"ResourceFolder,omitempty"`

    // Container class for azw3 save options.
    ResourceFolderAlias *string `json:"ResourceFolderAlias,omitempty"`

    // Container class for azw3 save options.
    ScaleImageToShapeSize *bool `json:"ScaleImageToShapeSize,omitempty"`

    // Container class for azw3 save options.
    TableWidthOutputMode *string `json:"TableWidthOutputMode,omitempty"`

    // Container class for azw3 save options.
    NavigationMapLevel *int32 `json:"NavigationMapLevel,omitempty"`

    // Container class for azw3 save options.
    SaveFormat *string `json:"SaveFormat,omitempty"`
}

func (Azw3SaveOptionsData) IsAzw3SaveOptionsData() bool {
    return true
}

func (Azw3SaveOptionsData) IsHtmlSaveOptionsData() bool {
    return true
}

func (Azw3SaveOptionsData) IsSaveOptionsData() bool {
    return true
}

func (obj *Azw3SaveOptionsData) Initialize() {
    var _SaveFormat = "azw3"
    obj.SaveFormat = &_SaveFormat


    if (obj.CustomTimeZoneInfoData != nil) {
        obj.CustomTimeZoneInfoData.Initialize()
    }


}

func (obj *Azw3SaveOptionsData) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["AllowNegativeIndent"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.AllowNegativeIndent = &parsedValue
        }

    } else if jsonValue, exists := json["allowNegativeIndent"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.AllowNegativeIndent = &parsedValue
        }

    }

    if jsonValue, exists := json["CssClassNamePrefix"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.CssClassNamePrefix = &parsedValue
        }

    } else if jsonValue, exists := json["cssClassNamePrefix"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.CssClassNamePrefix = &parsedValue
        }

    }

    if jsonValue, exists := json["CssStyleSheetFileName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.CssStyleSheetFileName = &parsedValue
        }

    } else if jsonValue, exists := json["cssStyleSheetFileName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.CssStyleSheetFileName = &parsedValue
        }

    }

    if jsonValue, exists := json["CssStyleSheetType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.CssStyleSheetType = &parsedValue
        }

    } else if jsonValue, exists := json["cssStyleSheetType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.CssStyleSheetType = &parsedValue
        }

    }

    if jsonValue, exists := json["DocumentSplitCriteria"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DocumentSplitCriteria = &parsedValue
        }

    } else if jsonValue, exists := json["documentSplitCriteria"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DocumentSplitCriteria = &parsedValue
        }

    }

    if jsonValue, exists := json["DocumentSplitHeadingLevel"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.DocumentSplitHeadingLevel = new(int32)
            *obj.DocumentSplitHeadingLevel = int32(parsedValue)
        }

    } else if jsonValue, exists := json["documentSplitHeadingLevel"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.DocumentSplitHeadingLevel = new(int32)
            *obj.DocumentSplitHeadingLevel = int32(parsedValue)
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

    if jsonValue, exists := json["ExportDocumentProperties"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportDocumentProperties = &parsedValue
        }

    } else if jsonValue, exists := json["exportDocumentProperties"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportDocumentProperties = &parsedValue
        }

    }

    if jsonValue, exists := json["ExportDropDownFormFieldAsText"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportDropDownFormFieldAsText = &parsedValue
        }

    } else if jsonValue, exists := json["exportDropDownFormFieldAsText"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportDropDownFormFieldAsText = &parsedValue
        }

    }

    if jsonValue, exists := json["ExportFontResources"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportFontResources = &parsedValue
        }

    } else if jsonValue, exists := json["exportFontResources"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportFontResources = &parsedValue
        }

    }

    if jsonValue, exists := json["ExportFontsAsBase64"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportFontsAsBase64 = &parsedValue
        }

    } else if jsonValue, exists := json["exportFontsAsBase64"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportFontsAsBase64 = &parsedValue
        }

    }

    if jsonValue, exists := json["ExportHeadersFootersMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ExportHeadersFootersMode = &parsedValue
        }

    } else if jsonValue, exists := json["exportHeadersFootersMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ExportHeadersFootersMode = &parsedValue
        }

    }

    if jsonValue, exists := json["ExportImagesAsBase64"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportImagesAsBase64 = &parsedValue
        }

    } else if jsonValue, exists := json["exportImagesAsBase64"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportImagesAsBase64 = &parsedValue
        }

    }

    if jsonValue, exists := json["ExportLanguageInformation"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportLanguageInformation = &parsedValue
        }

    } else if jsonValue, exists := json["exportLanguageInformation"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportLanguageInformation = &parsedValue
        }

    }

    if jsonValue, exists := json["ExportListLabels"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ExportListLabels = &parsedValue
        }

    } else if jsonValue, exists := json["exportListLabels"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ExportListLabels = &parsedValue
        }

    }

    if jsonValue, exists := json["ExportOriginalUrlForLinkedImages"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportOriginalUrlForLinkedImages = &parsedValue
        }

    } else if jsonValue, exists := json["exportOriginalUrlForLinkedImages"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportOriginalUrlForLinkedImages = &parsedValue
        }

    }

    if jsonValue, exists := json["ExportPageMargins"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportPageMargins = &parsedValue
        }

    } else if jsonValue, exists := json["exportPageMargins"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportPageMargins = &parsedValue
        }

    }

    if jsonValue, exists := json["ExportPageSetup"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportPageSetup = &parsedValue
        }

    } else if jsonValue, exists := json["exportPageSetup"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportPageSetup = &parsedValue
        }

    }

    if jsonValue, exists := json["ExportRelativeFontSize"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportRelativeFontSize = &parsedValue
        }

    } else if jsonValue, exists := json["exportRelativeFontSize"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportRelativeFontSize = &parsedValue
        }

    }

    if jsonValue, exists := json["ExportRoundtripInformation"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportRoundtripInformation = &parsedValue
        }

    } else if jsonValue, exists := json["exportRoundtripInformation"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportRoundtripInformation = &parsedValue
        }

    }

    if jsonValue, exists := json["ExportTextInputFormFieldAsText"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportTextInputFormFieldAsText = &parsedValue
        }

    } else if jsonValue, exists := json["exportTextInputFormFieldAsText"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportTextInputFormFieldAsText = &parsedValue
        }

    }

    if jsonValue, exists := json["ExportTocPageNumbers"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportTocPageNumbers = &parsedValue
        }

    } else if jsonValue, exists := json["exportTocPageNumbers"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportTocPageNumbers = &parsedValue
        }

    }

    if jsonValue, exists := json["ExportXhtmlTransitional"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportXhtmlTransitional = &parsedValue
        }

    } else if jsonValue, exists := json["exportXhtmlTransitional"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportXhtmlTransitional = &parsedValue
        }

    }

    if jsonValue, exists := json["FontResourcesSubsettingSizeThreshold"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.FontResourcesSubsettingSizeThreshold = new(int32)
            *obj.FontResourcesSubsettingSizeThreshold = int32(parsedValue)
        }

    } else if jsonValue, exists := json["fontResourcesSubsettingSizeThreshold"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.FontResourcesSubsettingSizeThreshold = new(int32)
            *obj.FontResourcesSubsettingSizeThreshold = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["FontsFolder"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FontsFolder = &parsedValue
        }

    } else if jsonValue, exists := json["fontsFolder"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FontsFolder = &parsedValue
        }

    }

    if jsonValue, exists := json["FontsFolderAlias"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FontsFolderAlias = &parsedValue
        }

    } else if jsonValue, exists := json["fontsFolderAlias"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FontsFolderAlias = &parsedValue
        }

    }

    if jsonValue, exists := json["HtmlVersion"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.HtmlVersion = &parsedValue
        }

    } else if jsonValue, exists := json["htmlVersion"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.HtmlVersion = &parsedValue
        }

    }

    if jsonValue, exists := json["ImageResolution"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ImageResolution = new(int32)
            *obj.ImageResolution = int32(parsedValue)
        }

    } else if jsonValue, exists := json["imageResolution"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ImageResolution = new(int32)
            *obj.ImageResolution = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["ImagesFolder"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ImagesFolder = &parsedValue
        }

    } else if jsonValue, exists := json["imagesFolder"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ImagesFolder = &parsedValue
        }

    }

    if jsonValue, exists := json["ImagesFolderAlias"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ImagesFolderAlias = &parsedValue
        }

    } else if jsonValue, exists := json["imagesFolderAlias"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ImagesFolderAlias = &parsedValue
        }

    }

    if jsonValue, exists := json["MetafileFormat"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.MetafileFormat = &parsedValue
        }

    } else if jsonValue, exists := json["metafileFormat"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.MetafileFormat = &parsedValue
        }

    }

    if jsonValue, exists := json["OfficeMathOutputMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.OfficeMathOutputMode = &parsedValue
        }

    } else if jsonValue, exists := json["officeMathOutputMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.OfficeMathOutputMode = &parsedValue
        }

    }

    if jsonValue, exists := json["PrettyFormat"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.PrettyFormat = &parsedValue
        }

    } else if jsonValue, exists := json["prettyFormat"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.PrettyFormat = &parsedValue
        }

    }

    if jsonValue, exists := json["ReplaceBackslashWithYenSign"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ReplaceBackslashWithYenSign = &parsedValue
        }

    } else if jsonValue, exists := json["replaceBackslashWithYenSign"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ReplaceBackslashWithYenSign = &parsedValue
        }

    }

    if jsonValue, exists := json["ResolveFontNames"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ResolveFontNames = &parsedValue
        }

    } else if jsonValue, exists := json["resolveFontNames"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ResolveFontNames = &parsedValue
        }

    }

    if jsonValue, exists := json["ResourceFolder"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ResourceFolder = &parsedValue
        }

    } else if jsonValue, exists := json["resourceFolder"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ResourceFolder = &parsedValue
        }

    }

    if jsonValue, exists := json["ResourceFolderAlias"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ResourceFolderAlias = &parsedValue
        }

    } else if jsonValue, exists := json["resourceFolderAlias"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ResourceFolderAlias = &parsedValue
        }

    }

    if jsonValue, exists := json["ScaleImageToShapeSize"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ScaleImageToShapeSize = &parsedValue
        }

    } else if jsonValue, exists := json["scaleImageToShapeSize"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ScaleImageToShapeSize = &parsedValue
        }

    }

    if jsonValue, exists := json["TableWidthOutputMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.TableWidthOutputMode = &parsedValue
        }

    } else if jsonValue, exists := json["tableWidthOutputMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.TableWidthOutputMode = &parsedValue
        }

    }

    if jsonValue, exists := json["NavigationMapLevel"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.NavigationMapLevel = new(int32)
            *obj.NavigationMapLevel = int32(parsedValue)
        }

    } else if jsonValue, exists := json["navigationMapLevel"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.NavigationMapLevel = new(int32)
            *obj.NavigationMapLevel = int32(parsedValue)
        }

    }
}

func (obj *Azw3SaveOptionsData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *Azw3SaveOptionsData) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.FileName == nil {
        return errors.New("Property FileName in Azw3SaveOptionsData is required.")
    }
    if obj.CustomTimeZoneInfoData != nil {
        if err := obj.CustomTimeZoneInfoData.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *Azw3SaveOptionsData) GetAllowEmbeddingPostScriptFonts() *bool {
    return obj.AllowEmbeddingPostScriptFonts
}

func (obj *Azw3SaveOptionsData) SetAllowEmbeddingPostScriptFonts(value *bool) {
    obj.AllowEmbeddingPostScriptFonts = value
}

func (obj *Azw3SaveOptionsData) GetCustomTimeZoneInfoData() ITimeZoneInfoData {
    return obj.CustomTimeZoneInfoData
}

func (obj *Azw3SaveOptionsData) SetCustomTimeZoneInfoData(value ITimeZoneInfoData) {
    obj.CustomTimeZoneInfoData = value
}

func (obj *Azw3SaveOptionsData) GetDml3DEffectsRenderingMode() *string {
    return obj.Dml3DEffectsRenderingMode
}

func (obj *Azw3SaveOptionsData) SetDml3DEffectsRenderingMode(value *string) {
    obj.Dml3DEffectsRenderingMode = value
}

func (obj *Azw3SaveOptionsData) GetDmlEffectsRenderingMode() *string {
    return obj.DmlEffectsRenderingMode
}

func (obj *Azw3SaveOptionsData) SetDmlEffectsRenderingMode(value *string) {
    obj.DmlEffectsRenderingMode = value
}

func (obj *Azw3SaveOptionsData) GetDmlRenderingMode() *string {
    return obj.DmlRenderingMode
}

func (obj *Azw3SaveOptionsData) SetDmlRenderingMode(value *string) {
    obj.DmlRenderingMode = value
}

func (obj *Azw3SaveOptionsData) GetFileName() *string {
    return obj.FileName
}

func (obj *Azw3SaveOptionsData) SetFileName(value *string) {
    obj.FileName = value
}

func (obj *Azw3SaveOptionsData) GetImlRenderingMode() *string {
    return obj.ImlRenderingMode
}

func (obj *Azw3SaveOptionsData) SetImlRenderingMode(value *string) {
    obj.ImlRenderingMode = value
}

func (obj *Azw3SaveOptionsData) GetUpdateCreatedTimeProperty() *bool {
    return obj.UpdateCreatedTimeProperty
}

func (obj *Azw3SaveOptionsData) SetUpdateCreatedTimeProperty(value *bool) {
    obj.UpdateCreatedTimeProperty = value
}

func (obj *Azw3SaveOptionsData) GetUpdateFields() *bool {
    return obj.UpdateFields
}

func (obj *Azw3SaveOptionsData) SetUpdateFields(value *bool) {
    obj.UpdateFields = value
}

func (obj *Azw3SaveOptionsData) GetUpdateLastPrintedProperty() *bool {
    return obj.UpdateLastPrintedProperty
}

func (obj *Azw3SaveOptionsData) SetUpdateLastPrintedProperty(value *bool) {
    obj.UpdateLastPrintedProperty = value
}

func (obj *Azw3SaveOptionsData) GetUpdateLastSavedTimeProperty() *bool {
    return obj.UpdateLastSavedTimeProperty
}

func (obj *Azw3SaveOptionsData) SetUpdateLastSavedTimeProperty(value *bool) {
    obj.UpdateLastSavedTimeProperty = value
}

func (obj *Azw3SaveOptionsData) GetZipOutput() *bool {
    return obj.ZipOutput
}

func (obj *Azw3SaveOptionsData) SetZipOutput(value *bool) {
    obj.ZipOutput = value
}

func (obj *Azw3SaveOptionsData) GetAllowNegativeIndent() *bool {
    return obj.AllowNegativeIndent
}

func (obj *Azw3SaveOptionsData) SetAllowNegativeIndent(value *bool) {
    obj.AllowNegativeIndent = value
}

func (obj *Azw3SaveOptionsData) GetCssClassNamePrefix() *string {
    return obj.CssClassNamePrefix
}

func (obj *Azw3SaveOptionsData) SetCssClassNamePrefix(value *string) {
    obj.CssClassNamePrefix = value
}

func (obj *Azw3SaveOptionsData) GetCssStyleSheetFileName() *string {
    return obj.CssStyleSheetFileName
}

func (obj *Azw3SaveOptionsData) SetCssStyleSheetFileName(value *string) {
    obj.CssStyleSheetFileName = value
}

func (obj *Azw3SaveOptionsData) GetCssStyleSheetType() *string {
    return obj.CssStyleSheetType
}

func (obj *Azw3SaveOptionsData) SetCssStyleSheetType(value *string) {
    obj.CssStyleSheetType = value
}

func (obj *Azw3SaveOptionsData) GetDocumentSplitCriteria() *string {
    return obj.DocumentSplitCriteria
}

func (obj *Azw3SaveOptionsData) SetDocumentSplitCriteria(value *string) {
    obj.DocumentSplitCriteria = value
}

func (obj *Azw3SaveOptionsData) GetDocumentSplitHeadingLevel() *int32 {
    return obj.DocumentSplitHeadingLevel
}

func (obj *Azw3SaveOptionsData) SetDocumentSplitHeadingLevel(value *int32) {
    obj.DocumentSplitHeadingLevel = value
}

func (obj *Azw3SaveOptionsData) GetEncoding() *string {
    return obj.Encoding
}

func (obj *Azw3SaveOptionsData) SetEncoding(value *string) {
    obj.Encoding = value
}

func (obj *Azw3SaveOptionsData) GetExportDocumentProperties() *bool {
    return obj.ExportDocumentProperties
}

func (obj *Azw3SaveOptionsData) SetExportDocumentProperties(value *bool) {
    obj.ExportDocumentProperties = value
}

func (obj *Azw3SaveOptionsData) GetExportDropDownFormFieldAsText() *bool {
    return obj.ExportDropDownFormFieldAsText
}

func (obj *Azw3SaveOptionsData) SetExportDropDownFormFieldAsText(value *bool) {
    obj.ExportDropDownFormFieldAsText = value
}

func (obj *Azw3SaveOptionsData) GetExportFontResources() *bool {
    return obj.ExportFontResources
}

func (obj *Azw3SaveOptionsData) SetExportFontResources(value *bool) {
    obj.ExportFontResources = value
}

func (obj *Azw3SaveOptionsData) GetExportFontsAsBase64() *bool {
    return obj.ExportFontsAsBase64
}

func (obj *Azw3SaveOptionsData) SetExportFontsAsBase64(value *bool) {
    obj.ExportFontsAsBase64 = value
}

func (obj *Azw3SaveOptionsData) GetExportHeadersFootersMode() *string {
    return obj.ExportHeadersFootersMode
}

func (obj *Azw3SaveOptionsData) SetExportHeadersFootersMode(value *string) {
    obj.ExportHeadersFootersMode = value
}

func (obj *Azw3SaveOptionsData) GetExportImagesAsBase64() *bool {
    return obj.ExportImagesAsBase64
}

func (obj *Azw3SaveOptionsData) SetExportImagesAsBase64(value *bool) {
    obj.ExportImagesAsBase64 = value
}

func (obj *Azw3SaveOptionsData) GetExportLanguageInformation() *bool {
    return obj.ExportLanguageInformation
}

func (obj *Azw3SaveOptionsData) SetExportLanguageInformation(value *bool) {
    obj.ExportLanguageInformation = value
}

func (obj *Azw3SaveOptionsData) GetExportListLabels() *string {
    return obj.ExportListLabels
}

func (obj *Azw3SaveOptionsData) SetExportListLabels(value *string) {
    obj.ExportListLabels = value
}

func (obj *Azw3SaveOptionsData) GetExportOriginalUrlForLinkedImages() *bool {
    return obj.ExportOriginalUrlForLinkedImages
}

func (obj *Azw3SaveOptionsData) SetExportOriginalUrlForLinkedImages(value *bool) {
    obj.ExportOriginalUrlForLinkedImages = value
}

func (obj *Azw3SaveOptionsData) GetExportPageMargins() *bool {
    return obj.ExportPageMargins
}

func (obj *Azw3SaveOptionsData) SetExportPageMargins(value *bool) {
    obj.ExportPageMargins = value
}

func (obj *Azw3SaveOptionsData) GetExportPageSetup() *bool {
    return obj.ExportPageSetup
}

func (obj *Azw3SaveOptionsData) SetExportPageSetup(value *bool) {
    obj.ExportPageSetup = value
}

func (obj *Azw3SaveOptionsData) GetExportRelativeFontSize() *bool {
    return obj.ExportRelativeFontSize
}

func (obj *Azw3SaveOptionsData) SetExportRelativeFontSize(value *bool) {
    obj.ExportRelativeFontSize = value
}

func (obj *Azw3SaveOptionsData) GetExportRoundtripInformation() *bool {
    return obj.ExportRoundtripInformation
}

func (obj *Azw3SaveOptionsData) SetExportRoundtripInformation(value *bool) {
    obj.ExportRoundtripInformation = value
}

func (obj *Azw3SaveOptionsData) GetExportTextInputFormFieldAsText() *bool {
    return obj.ExportTextInputFormFieldAsText
}

func (obj *Azw3SaveOptionsData) SetExportTextInputFormFieldAsText(value *bool) {
    obj.ExportTextInputFormFieldAsText = value
}

func (obj *Azw3SaveOptionsData) GetExportTocPageNumbers() *bool {
    return obj.ExportTocPageNumbers
}

func (obj *Azw3SaveOptionsData) SetExportTocPageNumbers(value *bool) {
    obj.ExportTocPageNumbers = value
}

func (obj *Azw3SaveOptionsData) GetExportXhtmlTransitional() *bool {
    return obj.ExportXhtmlTransitional
}

func (obj *Azw3SaveOptionsData) SetExportXhtmlTransitional(value *bool) {
    obj.ExportXhtmlTransitional = value
}

func (obj *Azw3SaveOptionsData) GetFontResourcesSubsettingSizeThreshold() *int32 {
    return obj.FontResourcesSubsettingSizeThreshold
}

func (obj *Azw3SaveOptionsData) SetFontResourcesSubsettingSizeThreshold(value *int32) {
    obj.FontResourcesSubsettingSizeThreshold = value
}

func (obj *Azw3SaveOptionsData) GetFontsFolder() *string {
    return obj.FontsFolder
}

func (obj *Azw3SaveOptionsData) SetFontsFolder(value *string) {
    obj.FontsFolder = value
}

func (obj *Azw3SaveOptionsData) GetFontsFolderAlias() *string {
    return obj.FontsFolderAlias
}

func (obj *Azw3SaveOptionsData) SetFontsFolderAlias(value *string) {
    obj.FontsFolderAlias = value
}

func (obj *Azw3SaveOptionsData) GetHtmlVersion() *string {
    return obj.HtmlVersion
}

func (obj *Azw3SaveOptionsData) SetHtmlVersion(value *string) {
    obj.HtmlVersion = value
}

func (obj *Azw3SaveOptionsData) GetImageResolution() *int32 {
    return obj.ImageResolution
}

func (obj *Azw3SaveOptionsData) SetImageResolution(value *int32) {
    obj.ImageResolution = value
}

func (obj *Azw3SaveOptionsData) GetImagesFolder() *string {
    return obj.ImagesFolder
}

func (obj *Azw3SaveOptionsData) SetImagesFolder(value *string) {
    obj.ImagesFolder = value
}

func (obj *Azw3SaveOptionsData) GetImagesFolderAlias() *string {
    return obj.ImagesFolderAlias
}

func (obj *Azw3SaveOptionsData) SetImagesFolderAlias(value *string) {
    obj.ImagesFolderAlias = value
}

func (obj *Azw3SaveOptionsData) GetMetafileFormat() *string {
    return obj.MetafileFormat
}

func (obj *Azw3SaveOptionsData) SetMetafileFormat(value *string) {
    obj.MetafileFormat = value
}

func (obj *Azw3SaveOptionsData) GetOfficeMathOutputMode() *string {
    return obj.OfficeMathOutputMode
}

func (obj *Azw3SaveOptionsData) SetOfficeMathOutputMode(value *string) {
    obj.OfficeMathOutputMode = value
}

func (obj *Azw3SaveOptionsData) GetPrettyFormat() *bool {
    return obj.PrettyFormat
}

func (obj *Azw3SaveOptionsData) SetPrettyFormat(value *bool) {
    obj.PrettyFormat = value
}

func (obj *Azw3SaveOptionsData) GetReplaceBackslashWithYenSign() *bool {
    return obj.ReplaceBackslashWithYenSign
}

func (obj *Azw3SaveOptionsData) SetReplaceBackslashWithYenSign(value *bool) {
    obj.ReplaceBackslashWithYenSign = value
}

func (obj *Azw3SaveOptionsData) GetResolveFontNames() *bool {
    return obj.ResolveFontNames
}

func (obj *Azw3SaveOptionsData) SetResolveFontNames(value *bool) {
    obj.ResolveFontNames = value
}

func (obj *Azw3SaveOptionsData) GetResourceFolder() *string {
    return obj.ResourceFolder
}

func (obj *Azw3SaveOptionsData) SetResourceFolder(value *string) {
    obj.ResourceFolder = value
}

func (obj *Azw3SaveOptionsData) GetResourceFolderAlias() *string {
    return obj.ResourceFolderAlias
}

func (obj *Azw3SaveOptionsData) SetResourceFolderAlias(value *string) {
    obj.ResourceFolderAlias = value
}

func (obj *Azw3SaveOptionsData) GetScaleImageToShapeSize() *bool {
    return obj.ScaleImageToShapeSize
}

func (obj *Azw3SaveOptionsData) SetScaleImageToShapeSize(value *bool) {
    obj.ScaleImageToShapeSize = value
}

func (obj *Azw3SaveOptionsData) GetTableWidthOutputMode() *string {
    return obj.TableWidthOutputMode
}

func (obj *Azw3SaveOptionsData) SetTableWidthOutputMode(value *string) {
    obj.TableWidthOutputMode = value
}

func (obj *Azw3SaveOptionsData) GetNavigationMapLevel() *int32 {
    return obj.NavigationMapLevel
}

func (obj *Azw3SaveOptionsData) SetNavigationMapLevel(value *int32) {
    obj.NavigationMapLevel = value
}

func (obj *Azw3SaveOptionsData) GetSaveFormat() *string {
    return obj.SaveFormat
}

func (obj *Azw3SaveOptionsData) SetSaveFormat(value *string) {
    obj.SaveFormat = value
}

