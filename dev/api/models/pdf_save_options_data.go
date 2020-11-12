/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="pdf_save_options_data.go">
 *   Copyright (c) 2020 Aspose.Words for Cloud
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

// container class for pdf save options.
type PdfSaveOptionsData struct {
    // container class for pdf save options.
    Dml3DEffectsRenderingMode string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // container class for pdf save options.
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // container class for pdf save options.
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // container class for pdf save options.
    FileName *string `json:"FileName,omitempty"`

    // container class for pdf save options.
    SaveFormat *string `json:"SaveFormat,omitempty"`

    // container class for pdf save options.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // container class for pdf save options.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // container class for pdf save options.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // container class for pdf save options.
    UpdateSdtContent *bool `json:"UpdateSdtContent,omitempty"`

    // container class for pdf save options.
    ZipOutput *bool `json:"ZipOutput,omitempty"`

    // container class for pdf save options.
    ColorMode *string `json:"ColorMode,omitempty"`

    // container class for pdf save options.
    JpegQuality *int32 `json:"JpegQuality,omitempty"`

    // container class for pdf save options.
    MetafileRenderingOptions *MetafileRenderingOptionsData `json:"MetafileRenderingOptions,omitempty"`

    // container class for pdf save options.
    NumeralFormat *string `json:"NumeralFormat,omitempty"`

    // container class for pdf save options.
    OptimizeOutput *bool `json:"OptimizeOutput,omitempty"`

    // container class for pdf save options.
    PageCount *int32 `json:"PageCount,omitempty"`

    // container class for pdf save options.
    PageIndex *int32 `json:"PageIndex,omitempty"`

    // container class for pdf save options.
    Compliance *string `json:"Compliance,omitempty"`

    // container class for pdf save options.
    CreateNoteHyperlinks *bool `json:"CreateNoteHyperlinks,omitempty"`

    // container class for pdf save options.
    CustomPropertiesExport *string `json:"CustomPropertiesExport,omitempty"`

    // container class for pdf save options.
    DigitalSignatureDetails *PdfDigitalSignatureDetailsData `json:"DigitalSignatureDetails,omitempty"`

    // container class for pdf save options.
    DisplayDocTitle *bool `json:"DisplayDocTitle,omitempty"`

    // container class for pdf save options.
    DownsampleOptions *DownsampleOptionsData `json:"DownsampleOptions,omitempty"`

    // container class for pdf save options.
    EmbedFullFonts *bool `json:"EmbedFullFonts,omitempty"`

    // container class for pdf save options.
    EncryptionDetails *PdfEncryptionDetailsData `json:"EncryptionDetails,omitempty"`

    // container class for pdf save options.
    EscapeUri *bool `json:"EscapeUri,omitempty"`

    // container class for pdf save options.
    ExportDocumentStructure *bool `json:"ExportDocumentStructure,omitempty"`

    // container class for pdf save options.
    FontEmbeddingMode *string `json:"FontEmbeddingMode,omitempty"`

    // container class for pdf save options.
    HeaderFooterBookmarksExportMode string `json:"HeaderFooterBookmarksExportMode,omitempty"`

    // container class for pdf save options.
    ImageColorSpaceExportMode *string `json:"ImageColorSpaceExportMode,omitempty"`

    // container class for pdf save options.
    ImageCompression *string `json:"ImageCompression,omitempty"`

    // container class for pdf save options.
    InterpolateImages *bool `json:"InterpolateImages,omitempty"`

    // container class for pdf save options.
    OpenHyperlinksInNewWindow *bool `json:"OpenHyperlinksInNewWindow,omitempty"`

    // container class for pdf save options.
    OutlineOptions *OutlineOptionsData `json:"OutlineOptions,omitempty"`

    // container class for pdf save options.
    PageMode *string `json:"PageMode,omitempty"`

    // container class for pdf save options.
    PreblendImages *bool `json:"PreblendImages,omitempty"`

    // container class for pdf save options.
    PreserveFormFields *bool `json:"PreserveFormFields,omitempty"`

    // container class for pdf save options.
    TextCompression *string `json:"TextCompression,omitempty"`

    // container class for pdf save options.
    UseBookFoldPrintingSettings *bool `json:"UseBookFoldPrintingSettings,omitempty"`

    // container class for pdf save options.
    UseCoreFonts *bool `json:"UseCoreFonts,omitempty"`

    // container class for pdf save options.
    ZoomBehavior *string `json:"ZoomBehavior,omitempty"`

    // container class for pdf save options.
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
