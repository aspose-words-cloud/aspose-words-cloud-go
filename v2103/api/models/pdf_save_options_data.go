/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="pdf_save_options_data.go">
 *   Copyright (c) 2021 Aspose.Words for Cloud
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

// Container class for pdf save options.
type PdfSaveOptionsDataResult struct {
    // Container class for pdf save options.
    AllowEmbeddingPostScriptFonts bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for pdf save options.
    Dml3DEffectsRenderingMode string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for pdf save options.
    DmlEffectsRenderingMode string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for pdf save options.
    DmlRenderingMode string `json:"DmlRenderingMode,omitempty"`

    // Container class for pdf save options.
    FileName string `json:"FileName,omitempty"`

    // Container class for pdf save options.
    SaveFormat string `json:"SaveFormat,omitempty"`

    // Container class for pdf save options.
    UpdateCreatedTimeProperty bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for pdf save options.
    UpdateFields bool `json:"UpdateFields,omitempty"`

    // Container class for pdf save options.
    UpdateLastPrintedProperty bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for pdf save options.
    UpdateLastSavedTimeProperty bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for pdf save options.
    UpdateSdtContent bool `json:"UpdateSdtContent,omitempty"`

    // Container class for pdf save options.
    ZipOutput bool `json:"ZipOutput,omitempty"`

    // Container class for pdf save options.
    ColorMode string `json:"ColorMode,omitempty"`

    // Container class for pdf save options.
    JpegQuality int32 `json:"JpegQuality,omitempty"`

    // Container class for pdf save options.
    MetafileRenderingOptions MetafileRenderingOptionsDataResult `json:"MetafileRenderingOptions,omitempty"`

    // Container class for pdf save options.
    NumeralFormat string `json:"NumeralFormat,omitempty"`

    // Container class for pdf save options.
    OptimizeOutput bool `json:"OptimizeOutput,omitempty"`

    // Container class for pdf save options.
    PageCount int32 `json:"PageCount,omitempty"`

    // Container class for pdf save options.
    PageIndex int32 `json:"PageIndex,omitempty"`

    // Container class for pdf save options.
    Compliance string `json:"Compliance,omitempty"`

    // Container class for pdf save options.
    CreateNoteHyperlinks bool `json:"CreateNoteHyperlinks,omitempty"`

    // Container class for pdf save options.
    CustomPropertiesExport string `json:"CustomPropertiesExport,omitempty"`

    // Container class for pdf save options.
    DigitalSignatureDetails PdfDigitalSignatureDetailsDataResult `json:"DigitalSignatureDetails,omitempty"`

    // Container class for pdf save options.
    DisplayDocTitle bool `json:"DisplayDocTitle,omitempty"`

    // Container class for pdf save options.
    DownsampleOptions DownsampleOptionsDataResult `json:"DownsampleOptions,omitempty"`

    // Container class for pdf save options.
    EmbedFullFonts bool `json:"EmbedFullFonts,omitempty"`

    // Container class for pdf save options.
    EncryptionDetails PdfEncryptionDetailsDataResult `json:"EncryptionDetails,omitempty"`

    // Container class for pdf save options.
    EscapeUri bool `json:"EscapeUri,omitempty"`

    // Container class for pdf save options.
    ExportDocumentStructure bool `json:"ExportDocumentStructure,omitempty"`

    // Container class for pdf save options.
    FontEmbeddingMode string `json:"FontEmbeddingMode,omitempty"`

    // Container class for pdf save options.
    HeaderFooterBookmarksExportMode string `json:"HeaderFooterBookmarksExportMode,omitempty"`

    // Container class for pdf save options.
    ImageColorSpaceExportMode string `json:"ImageColorSpaceExportMode,omitempty"`

    // Container class for pdf save options.
    ImageCompression string `json:"ImageCompression,omitempty"`

    // Container class for pdf save options.
    InterpolateImages bool `json:"InterpolateImages,omitempty"`

    // Container class for pdf save options.
    OpenHyperlinksInNewWindow bool `json:"OpenHyperlinksInNewWindow,omitempty"`

    // Container class for pdf save options.
    OutlineOptions OutlineOptionsDataResult `json:"OutlineOptions,omitempty"`

    // Container class for pdf save options.
    PageMode string `json:"PageMode,omitempty"`

    // Container class for pdf save options.
    PreblendImages bool `json:"PreblendImages,omitempty"`

    // Container class for pdf save options.
    PreserveFormFields bool `json:"PreserveFormFields,omitempty"`

    // Container class for pdf save options.
    TextCompression string `json:"TextCompression,omitempty"`

    // Container class for pdf save options.
    UseBookFoldPrintingSettings bool `json:"UseBookFoldPrintingSettings,omitempty"`

    // Container class for pdf save options.
    UseCoreFonts bool `json:"UseCoreFonts,omitempty"`

    // Container class for pdf save options.
    ZoomBehavior string `json:"ZoomBehavior,omitempty"`

    // Container class for pdf save options.
    ZoomFactor int32 `json:"ZoomFactor,omitempty"`
}

type PdfSaveOptionsData struct {
    // Container class for pdf save options.
    AllowEmbeddingPostScriptFonts *bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for pdf save options.
    Dml3DEffectsRenderingMode *string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for pdf save options.
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for pdf save options.
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // Container class for pdf save options.
    FileName *string `json:"FileName,omitempty"`

    // Container class for pdf save options.
    SaveFormat *string `json:"SaveFormat,omitempty"`

    // Container class for pdf save options.
    UpdateCreatedTimeProperty *bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for pdf save options.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // Container class for pdf save options.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for pdf save options.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for pdf save options.
    UpdateSdtContent *bool `json:"UpdateSdtContent,omitempty"`

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
    EmbedFullFonts *bool `json:"EmbedFullFonts,omitempty"`

    // Container class for pdf save options.
    EncryptionDetails IPdfEncryptionDetailsData `json:"EncryptionDetails,omitempty"`

    // Container class for pdf save options.
    EscapeUri *bool `json:"EscapeUri,omitempty"`

    // Container class for pdf save options.
    ExportDocumentStructure *bool `json:"ExportDocumentStructure,omitempty"`

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
    TextCompression *string `json:"TextCompression,omitempty"`

    // Container class for pdf save options.
    UseBookFoldPrintingSettings *bool `json:"UseBookFoldPrintingSettings,omitempty"`

    // Container class for pdf save options.
    UseCoreFonts *bool `json:"UseCoreFonts,omitempty"`

    // Container class for pdf save options.
    ZoomBehavior *string `json:"ZoomBehavior,omitempty"`

    // Container class for pdf save options.
    ZoomFactor *int32 `json:"ZoomFactor,omitempty"`
}

type IPdfSaveOptionsData interface {
    IsPdfSaveOptionsData() bool
}
func (PdfSaveOptionsData) IsPdfSaveOptionsData() bool {
    return true
}

func (PdfSaveOptionsData) IsFixedPageSaveOptionsData() bool {
    return true
}


