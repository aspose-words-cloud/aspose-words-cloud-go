/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="pdf_save_options_data.go">
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

// Container class for pdf save options.

type IPdfSaveOptionsData interface {
    IsPdfSaveOptionsData() bool
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
    GetAttachmentsEmbeddingMode() *string
    SetAttachmentsEmbeddingMode(value *string)
    GetCacheBackgroundGraphics() *bool
    SetCacheBackgroundGraphics(value *bool)
    GetCompliance() *string
    SetCompliance(value *string)
    GetCreateNoteHyperlinks() *bool
    SetCreateNoteHyperlinks(value *bool)
    GetCustomPropertiesExport() *string
    SetCustomPropertiesExport(value *string)
    GetDigitalSignatureDetails() IPdfDigitalSignatureDetailsData
    SetDigitalSignatureDetails(value IPdfDigitalSignatureDetailsData)
    GetDisplayDocTitle() *bool
    SetDisplayDocTitle(value *bool)
    GetDownsampleOptions() IDownsampleOptionsData
    SetDownsampleOptions(value IDownsampleOptionsData)
    GetEmbedAttachments() *bool
    SetEmbedAttachments(value *bool)
    GetEmbedFullFonts() *bool
    SetEmbedFullFonts(value *bool)
    GetEncryptionDetails() IPdfEncryptionDetailsData
    SetEncryptionDetails(value IPdfEncryptionDetailsData)
    GetExportDocumentStructure() *bool
    SetExportDocumentStructure(value *bool)
    GetExportLanguageToSpanTag() *bool
    SetExportLanguageToSpanTag(value *bool)
    GetFontEmbeddingMode() *string
    SetFontEmbeddingMode(value *string)
    GetHeaderFooterBookmarksExportMode() *string
    SetHeaderFooterBookmarksExportMode(value *string)
    GetImageColorSpaceExportMode() *string
    SetImageColorSpaceExportMode(value *string)
    GetImageCompression() *string
    SetImageCompression(value *string)
    GetInterpolateImages() *bool
    SetInterpolateImages(value *bool)
    GetOpenHyperlinksInNewWindow() *bool
    SetOpenHyperlinksInNewWindow(value *bool)
    GetOutlineOptions() IOutlineOptionsData
    SetOutlineOptions(value IOutlineOptionsData)
    GetPageMode() *string
    SetPageMode(value *string)
    GetPreblendImages() *bool
    SetPreblendImages(value *bool)
    GetPreserveFormFields() *bool
    SetPreserveFormFields(value *bool)
    GetRenderChoiceFormFieldBorder() *bool
    SetRenderChoiceFormFieldBorder(value *bool)
    GetTextCompression() *string
    SetTextCompression(value *string)
    GetUseBookFoldPrintingSettings() *bool
    SetUseBookFoldPrintingSettings(value *bool)
    GetUseCoreFonts() *bool
    SetUseCoreFonts(value *bool)
    GetUseSdtTagAsFormFieldName() *bool
    SetUseSdtTagAsFormFieldName(value *bool)
    GetZoomBehavior() *string
    SetZoomBehavior(value *string)
    GetZoomFactor() *int32
    SetZoomFactor(value *int32)
    GetExportFloatingShapesAsInlineTag() *bool
    SetExportFloatingShapesAsInlineTag(value *bool)
    GetSaveFormat() *string
    SetSaveFormat(value *string)
}

type PdfSaveOptionsData struct {
    // Container class for pdf save options.
    AllowEmbeddingPostScriptFonts *bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for pdf save options.
    CustomTimeZoneInfoData ITimeZoneInfoData `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for pdf save options.
    Dml3DEffectsRenderingMode *string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for pdf save options.
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for pdf save options.
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // Container class for pdf save options.
    FileName *string `json:"FileName,omitempty"`

    // Container class for pdf save options.
    ImlRenderingMode *string `json:"ImlRenderingMode,omitempty"`

    // Container class for pdf save options.
    UpdateAmbiguousTextFont *bool `json:"UpdateAmbiguousTextFont,omitempty"`

    // Container class for pdf save options.
    UpdateCreatedTimeProperty *bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for pdf save options.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // Container class for pdf save options.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for pdf save options.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for pdf save options.
    ZipOutput *bool `json:"ZipOutput,omitempty"`

    // Container class for pdf save options.
    ColorMode *string `json:"ColorMode,omitempty"`

    // Container class for pdf save options.
    JpegQuality *int32 `json:"JpegQuality,omitempty"`

    // Container class for pdf save options.
    MetafileRenderingOptions IMetafileRenderingOptionsData `json:"MetafileRenderingOptions,omitempty"`

    // Container class for pdf save options.
    NumeralFormat *string `json:"NumeralFormat,omitempty"`

    // Container class for pdf save options.
    OptimizeOutput *bool `json:"OptimizeOutput,omitempty"`

    // Container class for pdf save options.
    PageCount *int32 `json:"PageCount,omitempty"`

    // Container class for pdf save options.
    PageIndex *int32 `json:"PageIndex,omitempty"`

    // Container class for pdf save options.
    AttachmentsEmbeddingMode *string `json:"AttachmentsEmbeddingMode,omitempty"`

    // Container class for pdf save options.
    CacheBackgroundGraphics *bool `json:"CacheBackgroundGraphics,omitempty"`

    // Container class for pdf save options.
    Compliance *string `json:"Compliance,omitempty"`

    // Container class for pdf save options.
    CreateNoteHyperlinks *bool `json:"CreateNoteHyperlinks,omitempty"`

    // Container class for pdf save options.
    CustomPropertiesExport *string `json:"CustomPropertiesExport,omitempty"`

    // Container class for pdf save options.
    DigitalSignatureDetails IPdfDigitalSignatureDetailsData `json:"DigitalSignatureDetails,omitempty"`

    // Container class for pdf save options.
    DisplayDocTitle *bool `json:"DisplayDocTitle,omitempty"`

    // Container class for pdf save options.
    DownsampleOptions IDownsampleOptionsData `json:"DownsampleOptions,omitempty"`

    // Container class for pdf save options.
    EmbedAttachments *bool `json:"EmbedAttachments,omitempty"`

    // Container class for pdf save options.
    EmbedFullFonts *bool `json:"EmbedFullFonts,omitempty"`

    // Container class for pdf save options.
    EncryptionDetails IPdfEncryptionDetailsData `json:"EncryptionDetails,omitempty"`

    // Container class for pdf save options.
    ExportDocumentStructure *bool `json:"ExportDocumentStructure,omitempty"`

    // Container class for pdf save options.
    ExportLanguageToSpanTag *bool `json:"ExportLanguageToSpanTag,omitempty"`

    // Container class for pdf save options.
    FontEmbeddingMode *string `json:"FontEmbeddingMode,omitempty"`

    // Container class for pdf save options.
    HeaderFooterBookmarksExportMode *string `json:"HeaderFooterBookmarksExportMode,omitempty"`

    // Container class for pdf save options.
    ImageColorSpaceExportMode *string `json:"ImageColorSpaceExportMode,omitempty"`

    // Container class for pdf save options.
    ImageCompression *string `json:"ImageCompression,omitempty"`

    // Container class for pdf save options.
    InterpolateImages *bool `json:"InterpolateImages,omitempty"`

    // Container class for pdf save options.
    OpenHyperlinksInNewWindow *bool `json:"OpenHyperlinksInNewWindow,omitempty"`

    // Container class for pdf save options.
    OutlineOptions IOutlineOptionsData `json:"OutlineOptions,omitempty"`

    // Container class for pdf save options.
    PageMode *string `json:"PageMode,omitempty"`

    // Container class for pdf save options.
    PreblendImages *bool `json:"PreblendImages,omitempty"`

    // Container class for pdf save options.
    PreserveFormFields *bool `json:"PreserveFormFields,omitempty"`

    // Container class for pdf save options.
    RenderChoiceFormFieldBorder *bool `json:"RenderChoiceFormFieldBorder,omitempty"`

    // Container class for pdf save options.
    TextCompression *string `json:"TextCompression,omitempty"`

    // Container class for pdf save options.
    UseBookFoldPrintingSettings *bool `json:"UseBookFoldPrintingSettings,omitempty"`

    // Container class for pdf save options.
    UseCoreFonts *bool `json:"UseCoreFonts,omitempty"`

    // Container class for pdf save options.
    UseSdtTagAsFormFieldName *bool `json:"UseSdtTagAsFormFieldName,omitempty"`

    // Container class for pdf save options.
    ZoomBehavior *string `json:"ZoomBehavior,omitempty"`

    // Container class for pdf save options.
    ZoomFactor *int32 `json:"ZoomFactor,omitempty"`

    // Container class for pdf save options.
    ExportFloatingShapesAsInlineTag *bool `json:"ExportFloatingShapesAsInlineTag,omitempty"`

    // Container class for pdf save options.
    SaveFormat *string `json:"SaveFormat,omitempty"`
}

func (PdfSaveOptionsData) IsPdfSaveOptionsData() bool {
    return true
}

func (PdfSaveOptionsData) IsFixedPageSaveOptionsData() bool {
    return true
}

func (PdfSaveOptionsData) IsSaveOptionsData() bool {
    return true
}

func (obj *PdfSaveOptionsData) Initialize() {
    var _SaveFormat = "pdf"
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

    if (obj.DownsampleOptions != nil) {
        obj.DownsampleOptions.Initialize()
    }

    if (obj.EncryptionDetails != nil) {
        obj.EncryptionDetails.Initialize()
    }

    if (obj.OutlineOptions != nil) {
        obj.OutlineOptions.Initialize()
    }


}

func (obj *PdfSaveOptionsData) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["AttachmentsEmbeddingMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.AttachmentsEmbeddingMode = &parsedValue
        }

    } else if jsonValue, exists := json["attachmentsEmbeddingMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.AttachmentsEmbeddingMode = &parsedValue
        }

    }

    if jsonValue, exists := json["CacheBackgroundGraphics"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.CacheBackgroundGraphics = &parsedValue
        }

    } else if jsonValue, exists := json["cacheBackgroundGraphics"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.CacheBackgroundGraphics = &parsedValue
        }

    }

    if jsonValue, exists := json["Compliance"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Compliance = &parsedValue
        }

    } else if jsonValue, exists := json["compliance"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Compliance = &parsedValue
        }

    }

    if jsonValue, exists := json["CreateNoteHyperlinks"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.CreateNoteHyperlinks = &parsedValue
        }

    } else if jsonValue, exists := json["createNoteHyperlinks"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.CreateNoteHyperlinks = &parsedValue
        }

    }

    if jsonValue, exists := json["CustomPropertiesExport"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.CustomPropertiesExport = &parsedValue
        }

    } else if jsonValue, exists := json["customPropertiesExport"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.CustomPropertiesExport = &parsedValue
        }

    }

    if jsonValue, exists := json["DigitalSignatureDetails"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IPdfDigitalSignatureDetailsData = new(PdfDigitalSignatureDetailsData)
            modelInstance.Deserialize(parsedValue)
            obj.DigitalSignatureDetails = modelInstance
        }

    } else if jsonValue, exists := json["digitalSignatureDetails"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IPdfDigitalSignatureDetailsData = new(PdfDigitalSignatureDetailsData)
            modelInstance.Deserialize(parsedValue)
            obj.DigitalSignatureDetails = modelInstance
        }

    }

    if jsonValue, exists := json["DisplayDocTitle"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.DisplayDocTitle = &parsedValue
        }

    } else if jsonValue, exists := json["displayDocTitle"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.DisplayDocTitle = &parsedValue
        }

    }

    if jsonValue, exists := json["DownsampleOptions"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IDownsampleOptionsData = new(DownsampleOptionsData)
            modelInstance.Deserialize(parsedValue)
            obj.DownsampleOptions = modelInstance
        }

    } else if jsonValue, exists := json["downsampleOptions"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IDownsampleOptionsData = new(DownsampleOptionsData)
            modelInstance.Deserialize(parsedValue)
            obj.DownsampleOptions = modelInstance
        }

    }

    if jsonValue, exists := json["EmbedAttachments"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.EmbedAttachments = &parsedValue
        }

    } else if jsonValue, exists := json["embedAttachments"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.EmbedAttachments = &parsedValue
        }

    }

    if jsonValue, exists := json["EmbedFullFonts"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.EmbedFullFonts = &parsedValue
        }

    } else if jsonValue, exists := json["embedFullFonts"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.EmbedFullFonts = &parsedValue
        }

    }

    if jsonValue, exists := json["EncryptionDetails"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IPdfEncryptionDetailsData = new(PdfEncryptionDetailsData)
            modelInstance.Deserialize(parsedValue)
            obj.EncryptionDetails = modelInstance
        }

    } else if jsonValue, exists := json["encryptionDetails"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IPdfEncryptionDetailsData = new(PdfEncryptionDetailsData)
            modelInstance.Deserialize(parsedValue)
            obj.EncryptionDetails = modelInstance
        }

    }

    if jsonValue, exists := json["ExportDocumentStructure"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportDocumentStructure = &parsedValue
        }

    } else if jsonValue, exists := json["exportDocumentStructure"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportDocumentStructure = &parsedValue
        }

    }

    if jsonValue, exists := json["ExportLanguageToSpanTag"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportLanguageToSpanTag = &parsedValue
        }

    } else if jsonValue, exists := json["exportLanguageToSpanTag"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportLanguageToSpanTag = &parsedValue
        }

    }

    if jsonValue, exists := json["FontEmbeddingMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FontEmbeddingMode = &parsedValue
        }

    } else if jsonValue, exists := json["fontEmbeddingMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FontEmbeddingMode = &parsedValue
        }

    }

    if jsonValue, exists := json["HeaderFooterBookmarksExportMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.HeaderFooterBookmarksExportMode = &parsedValue
        }

    } else if jsonValue, exists := json["headerFooterBookmarksExportMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.HeaderFooterBookmarksExportMode = &parsedValue
        }

    }

    if jsonValue, exists := json["ImageColorSpaceExportMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ImageColorSpaceExportMode = &parsedValue
        }

    } else if jsonValue, exists := json["imageColorSpaceExportMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ImageColorSpaceExportMode = &parsedValue
        }

    }

    if jsonValue, exists := json["ImageCompression"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ImageCompression = &parsedValue
        }

    } else if jsonValue, exists := json["imageCompression"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ImageCompression = &parsedValue
        }

    }

    if jsonValue, exists := json["InterpolateImages"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.InterpolateImages = &parsedValue
        }

    } else if jsonValue, exists := json["interpolateImages"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.InterpolateImages = &parsedValue
        }

    }

    if jsonValue, exists := json["OpenHyperlinksInNewWindow"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.OpenHyperlinksInNewWindow = &parsedValue
        }

    } else if jsonValue, exists := json["openHyperlinksInNewWindow"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.OpenHyperlinksInNewWindow = &parsedValue
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

    if jsonValue, exists := json["PageMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.PageMode = &parsedValue
        }

    } else if jsonValue, exists := json["pageMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.PageMode = &parsedValue
        }

    }

    if jsonValue, exists := json["PreblendImages"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.PreblendImages = &parsedValue
        }

    } else if jsonValue, exists := json["preblendImages"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.PreblendImages = &parsedValue
        }

    }

    if jsonValue, exists := json["PreserveFormFields"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.PreserveFormFields = &parsedValue
        }

    } else if jsonValue, exists := json["preserveFormFields"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.PreserveFormFields = &parsedValue
        }

    }

    if jsonValue, exists := json["RenderChoiceFormFieldBorder"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.RenderChoiceFormFieldBorder = &parsedValue
        }

    } else if jsonValue, exists := json["renderChoiceFormFieldBorder"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.RenderChoiceFormFieldBorder = &parsedValue
        }

    }

    if jsonValue, exists := json["TextCompression"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.TextCompression = &parsedValue
        }

    } else if jsonValue, exists := json["textCompression"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.TextCompression = &parsedValue
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

    if jsonValue, exists := json["UseCoreFonts"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UseCoreFonts = &parsedValue
        }

    } else if jsonValue, exists := json["useCoreFonts"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UseCoreFonts = &parsedValue
        }

    }

    if jsonValue, exists := json["UseSdtTagAsFormFieldName"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UseSdtTagAsFormFieldName = &parsedValue
        }

    } else if jsonValue, exists := json["useSdtTagAsFormFieldName"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UseSdtTagAsFormFieldName = &parsedValue
        }

    }

    if jsonValue, exists := json["ZoomBehavior"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ZoomBehavior = &parsedValue
        }

    } else if jsonValue, exists := json["zoomBehavior"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ZoomBehavior = &parsedValue
        }

    }

    if jsonValue, exists := json["ZoomFactor"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ZoomFactor = new(int32)
            *obj.ZoomFactor = int32(parsedValue)
        }

    } else if jsonValue, exists := json["zoomFactor"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.ZoomFactor = new(int32)
            *obj.ZoomFactor = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["ExportFloatingShapesAsInlineTag"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportFloatingShapesAsInlineTag = &parsedValue
        }

    } else if jsonValue, exists := json["exportFloatingShapesAsInlineTag"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportFloatingShapesAsInlineTag = &parsedValue
        }

    }
}

func (obj *PdfSaveOptionsData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *PdfSaveOptionsData) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.FileName == nil {
        return errors.New("Property FileName in PdfSaveOptionsData is required.")
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
    if obj.DownsampleOptions != nil {
        if err := obj.DownsampleOptions.Validate(); err != nil {
            return err
        }
    }
    if obj.EncryptionDetails != nil {
        if err := obj.EncryptionDetails.Validate(); err != nil {
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

func (obj *PdfSaveOptionsData) GetAllowEmbeddingPostScriptFonts() *bool {
    return obj.AllowEmbeddingPostScriptFonts
}

func (obj *PdfSaveOptionsData) SetAllowEmbeddingPostScriptFonts(value *bool) {
    obj.AllowEmbeddingPostScriptFonts = value
}

func (obj *PdfSaveOptionsData) GetCustomTimeZoneInfoData() ITimeZoneInfoData {
    return obj.CustomTimeZoneInfoData
}

func (obj *PdfSaveOptionsData) SetCustomTimeZoneInfoData(value ITimeZoneInfoData) {
    obj.CustomTimeZoneInfoData = value
}

func (obj *PdfSaveOptionsData) GetDml3DEffectsRenderingMode() *string {
    return obj.Dml3DEffectsRenderingMode
}

func (obj *PdfSaveOptionsData) SetDml3DEffectsRenderingMode(value *string) {
    obj.Dml3DEffectsRenderingMode = value
}

func (obj *PdfSaveOptionsData) GetDmlEffectsRenderingMode() *string {
    return obj.DmlEffectsRenderingMode
}

func (obj *PdfSaveOptionsData) SetDmlEffectsRenderingMode(value *string) {
    obj.DmlEffectsRenderingMode = value
}

func (obj *PdfSaveOptionsData) GetDmlRenderingMode() *string {
    return obj.DmlRenderingMode
}

func (obj *PdfSaveOptionsData) SetDmlRenderingMode(value *string) {
    obj.DmlRenderingMode = value
}

func (obj *PdfSaveOptionsData) GetFileName() *string {
    return obj.FileName
}

func (obj *PdfSaveOptionsData) SetFileName(value *string) {
    obj.FileName = value
}

func (obj *PdfSaveOptionsData) GetImlRenderingMode() *string {
    return obj.ImlRenderingMode
}

func (obj *PdfSaveOptionsData) SetImlRenderingMode(value *string) {
    obj.ImlRenderingMode = value
}

func (obj *PdfSaveOptionsData) GetUpdateAmbiguousTextFont() *bool {
    return obj.UpdateAmbiguousTextFont
}

func (obj *PdfSaveOptionsData) SetUpdateAmbiguousTextFont(value *bool) {
    obj.UpdateAmbiguousTextFont = value
}

func (obj *PdfSaveOptionsData) GetUpdateCreatedTimeProperty() *bool {
    return obj.UpdateCreatedTimeProperty
}

func (obj *PdfSaveOptionsData) SetUpdateCreatedTimeProperty(value *bool) {
    obj.UpdateCreatedTimeProperty = value
}

func (obj *PdfSaveOptionsData) GetUpdateFields() *bool {
    return obj.UpdateFields
}

func (obj *PdfSaveOptionsData) SetUpdateFields(value *bool) {
    obj.UpdateFields = value
}

func (obj *PdfSaveOptionsData) GetUpdateLastPrintedProperty() *bool {
    return obj.UpdateLastPrintedProperty
}

func (obj *PdfSaveOptionsData) SetUpdateLastPrintedProperty(value *bool) {
    obj.UpdateLastPrintedProperty = value
}

func (obj *PdfSaveOptionsData) GetUpdateLastSavedTimeProperty() *bool {
    return obj.UpdateLastSavedTimeProperty
}

func (obj *PdfSaveOptionsData) SetUpdateLastSavedTimeProperty(value *bool) {
    obj.UpdateLastSavedTimeProperty = value
}

func (obj *PdfSaveOptionsData) GetZipOutput() *bool {
    return obj.ZipOutput
}

func (obj *PdfSaveOptionsData) SetZipOutput(value *bool) {
    obj.ZipOutput = value
}

func (obj *PdfSaveOptionsData) GetColorMode() *string {
    return obj.ColorMode
}

func (obj *PdfSaveOptionsData) SetColorMode(value *string) {
    obj.ColorMode = value
}

func (obj *PdfSaveOptionsData) GetJpegQuality() *int32 {
    return obj.JpegQuality
}

func (obj *PdfSaveOptionsData) SetJpegQuality(value *int32) {
    obj.JpegQuality = value
}

func (obj *PdfSaveOptionsData) GetMetafileRenderingOptions() IMetafileRenderingOptionsData {
    return obj.MetafileRenderingOptions
}

func (obj *PdfSaveOptionsData) SetMetafileRenderingOptions(value IMetafileRenderingOptionsData) {
    obj.MetafileRenderingOptions = value
}

func (obj *PdfSaveOptionsData) GetNumeralFormat() *string {
    return obj.NumeralFormat
}

func (obj *PdfSaveOptionsData) SetNumeralFormat(value *string) {
    obj.NumeralFormat = value
}

func (obj *PdfSaveOptionsData) GetOptimizeOutput() *bool {
    return obj.OptimizeOutput
}

func (obj *PdfSaveOptionsData) SetOptimizeOutput(value *bool) {
    obj.OptimizeOutput = value
}

func (obj *PdfSaveOptionsData) GetPageCount() *int32 {
    return obj.PageCount
}

func (obj *PdfSaveOptionsData) SetPageCount(value *int32) {
    obj.PageCount = value
}

func (obj *PdfSaveOptionsData) GetPageIndex() *int32 {
    return obj.PageIndex
}

func (obj *PdfSaveOptionsData) SetPageIndex(value *int32) {
    obj.PageIndex = value
}

func (obj *PdfSaveOptionsData) GetAttachmentsEmbeddingMode() *string {
    return obj.AttachmentsEmbeddingMode
}

func (obj *PdfSaveOptionsData) SetAttachmentsEmbeddingMode(value *string) {
    obj.AttachmentsEmbeddingMode = value
}

func (obj *PdfSaveOptionsData) GetCacheBackgroundGraphics() *bool {
    return obj.CacheBackgroundGraphics
}

func (obj *PdfSaveOptionsData) SetCacheBackgroundGraphics(value *bool) {
    obj.CacheBackgroundGraphics = value
}

func (obj *PdfSaveOptionsData) GetCompliance() *string {
    return obj.Compliance
}

func (obj *PdfSaveOptionsData) SetCompliance(value *string) {
    obj.Compliance = value
}

func (obj *PdfSaveOptionsData) GetCreateNoteHyperlinks() *bool {
    return obj.CreateNoteHyperlinks
}

func (obj *PdfSaveOptionsData) SetCreateNoteHyperlinks(value *bool) {
    obj.CreateNoteHyperlinks = value
}

func (obj *PdfSaveOptionsData) GetCustomPropertiesExport() *string {
    return obj.CustomPropertiesExport
}

func (obj *PdfSaveOptionsData) SetCustomPropertiesExport(value *string) {
    obj.CustomPropertiesExport = value
}

func (obj *PdfSaveOptionsData) GetDigitalSignatureDetails() IPdfDigitalSignatureDetailsData {
    return obj.DigitalSignatureDetails
}

func (obj *PdfSaveOptionsData) SetDigitalSignatureDetails(value IPdfDigitalSignatureDetailsData) {
    obj.DigitalSignatureDetails = value
}

func (obj *PdfSaveOptionsData) GetDisplayDocTitle() *bool {
    return obj.DisplayDocTitle
}

func (obj *PdfSaveOptionsData) SetDisplayDocTitle(value *bool) {
    obj.DisplayDocTitle = value
}

func (obj *PdfSaveOptionsData) GetDownsampleOptions() IDownsampleOptionsData {
    return obj.DownsampleOptions
}

func (obj *PdfSaveOptionsData) SetDownsampleOptions(value IDownsampleOptionsData) {
    obj.DownsampleOptions = value
}

func (obj *PdfSaveOptionsData) GetEmbedAttachments() *bool {
    return obj.EmbedAttachments
}

func (obj *PdfSaveOptionsData) SetEmbedAttachments(value *bool) {
    obj.EmbedAttachments = value
}

func (obj *PdfSaveOptionsData) GetEmbedFullFonts() *bool {
    return obj.EmbedFullFonts
}

func (obj *PdfSaveOptionsData) SetEmbedFullFonts(value *bool) {
    obj.EmbedFullFonts = value
}

func (obj *PdfSaveOptionsData) GetEncryptionDetails() IPdfEncryptionDetailsData {
    return obj.EncryptionDetails
}

func (obj *PdfSaveOptionsData) SetEncryptionDetails(value IPdfEncryptionDetailsData) {
    obj.EncryptionDetails = value
}

func (obj *PdfSaveOptionsData) GetExportDocumentStructure() *bool {
    return obj.ExportDocumentStructure
}

func (obj *PdfSaveOptionsData) SetExportDocumentStructure(value *bool) {
    obj.ExportDocumentStructure = value
}

func (obj *PdfSaveOptionsData) GetExportLanguageToSpanTag() *bool {
    return obj.ExportLanguageToSpanTag
}

func (obj *PdfSaveOptionsData) SetExportLanguageToSpanTag(value *bool) {
    obj.ExportLanguageToSpanTag = value
}

func (obj *PdfSaveOptionsData) GetFontEmbeddingMode() *string {
    return obj.FontEmbeddingMode
}

func (obj *PdfSaveOptionsData) SetFontEmbeddingMode(value *string) {
    obj.FontEmbeddingMode = value
}

func (obj *PdfSaveOptionsData) GetHeaderFooterBookmarksExportMode() *string {
    return obj.HeaderFooterBookmarksExportMode
}

func (obj *PdfSaveOptionsData) SetHeaderFooterBookmarksExportMode(value *string) {
    obj.HeaderFooterBookmarksExportMode = value
}

func (obj *PdfSaveOptionsData) GetImageColorSpaceExportMode() *string {
    return obj.ImageColorSpaceExportMode
}

func (obj *PdfSaveOptionsData) SetImageColorSpaceExportMode(value *string) {
    obj.ImageColorSpaceExportMode = value
}

func (obj *PdfSaveOptionsData) GetImageCompression() *string {
    return obj.ImageCompression
}

func (obj *PdfSaveOptionsData) SetImageCompression(value *string) {
    obj.ImageCompression = value
}

func (obj *PdfSaveOptionsData) GetInterpolateImages() *bool {
    return obj.InterpolateImages
}

func (obj *PdfSaveOptionsData) SetInterpolateImages(value *bool) {
    obj.InterpolateImages = value
}

func (obj *PdfSaveOptionsData) GetOpenHyperlinksInNewWindow() *bool {
    return obj.OpenHyperlinksInNewWindow
}

func (obj *PdfSaveOptionsData) SetOpenHyperlinksInNewWindow(value *bool) {
    obj.OpenHyperlinksInNewWindow = value
}

func (obj *PdfSaveOptionsData) GetOutlineOptions() IOutlineOptionsData {
    return obj.OutlineOptions
}

func (obj *PdfSaveOptionsData) SetOutlineOptions(value IOutlineOptionsData) {
    obj.OutlineOptions = value
}

func (obj *PdfSaveOptionsData) GetPageMode() *string {
    return obj.PageMode
}

func (obj *PdfSaveOptionsData) SetPageMode(value *string) {
    obj.PageMode = value
}

func (obj *PdfSaveOptionsData) GetPreblendImages() *bool {
    return obj.PreblendImages
}

func (obj *PdfSaveOptionsData) SetPreblendImages(value *bool) {
    obj.PreblendImages = value
}

func (obj *PdfSaveOptionsData) GetPreserveFormFields() *bool {
    return obj.PreserveFormFields
}

func (obj *PdfSaveOptionsData) SetPreserveFormFields(value *bool) {
    obj.PreserveFormFields = value
}

func (obj *PdfSaveOptionsData) GetRenderChoiceFormFieldBorder() *bool {
    return obj.RenderChoiceFormFieldBorder
}

func (obj *PdfSaveOptionsData) SetRenderChoiceFormFieldBorder(value *bool) {
    obj.RenderChoiceFormFieldBorder = value
}

func (obj *PdfSaveOptionsData) GetTextCompression() *string {
    return obj.TextCompression
}

func (obj *PdfSaveOptionsData) SetTextCompression(value *string) {
    obj.TextCompression = value
}

func (obj *PdfSaveOptionsData) GetUseBookFoldPrintingSettings() *bool {
    return obj.UseBookFoldPrintingSettings
}

func (obj *PdfSaveOptionsData) SetUseBookFoldPrintingSettings(value *bool) {
    obj.UseBookFoldPrintingSettings = value
}

func (obj *PdfSaveOptionsData) GetUseCoreFonts() *bool {
    return obj.UseCoreFonts
}

func (obj *PdfSaveOptionsData) SetUseCoreFonts(value *bool) {
    obj.UseCoreFonts = value
}

func (obj *PdfSaveOptionsData) GetUseSdtTagAsFormFieldName() *bool {
    return obj.UseSdtTagAsFormFieldName
}

func (obj *PdfSaveOptionsData) SetUseSdtTagAsFormFieldName(value *bool) {
    obj.UseSdtTagAsFormFieldName = value
}

func (obj *PdfSaveOptionsData) GetZoomBehavior() *string {
    return obj.ZoomBehavior
}

func (obj *PdfSaveOptionsData) SetZoomBehavior(value *string) {
    obj.ZoomBehavior = value
}

func (obj *PdfSaveOptionsData) GetZoomFactor() *int32 {
    return obj.ZoomFactor
}

func (obj *PdfSaveOptionsData) SetZoomFactor(value *int32) {
    obj.ZoomFactor = value
}

func (obj *PdfSaveOptionsData) GetExportFloatingShapesAsInlineTag() *bool {
    return obj.ExportFloatingShapesAsInlineTag
}

func (obj *PdfSaveOptionsData) SetExportFloatingShapesAsInlineTag(value *bool) {
    obj.ExportFloatingShapesAsInlineTag = value
}

func (obj *PdfSaveOptionsData) GetSaveFormat() *string {
    return obj.SaveFormat
}

func (obj *PdfSaveOptionsData) SetSaveFormat(value *string) {
    obj.SaveFormat = value
}

