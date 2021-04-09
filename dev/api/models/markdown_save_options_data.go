/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="markdown_save_options_data.go">
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

// Container class for markdown save options.
type MarkdownSaveOptionsDataResult struct {
    // Container class for markdown save options.
    AllowEmbeddingPostScriptFonts bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for markdown save options.
    CustomTimeZoneInfoData TimeZoneInfoDataResult `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for markdown save options.
    Dml3DEffectsRenderingMode string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for markdown save options.
    DmlEffectsRenderingMode string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for markdown save options.
    DmlRenderingMode string `json:"DmlRenderingMode,omitempty"`

    // Container class for markdown save options.
    FileName string `json:"FileName,omitempty"`

    // Container class for markdown save options.
    SaveFormat string `json:"SaveFormat,omitempty"`

    // Container class for markdown save options.
    UpdateCreatedTimeProperty bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for markdown save options.
    UpdateFields bool `json:"UpdateFields,omitempty"`

    // Container class for markdown save options.
    UpdateLastPrintedProperty bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for markdown save options.
    UpdateLastSavedTimeProperty bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for markdown save options.
    UpdateSdtContent bool `json:"UpdateSdtContent,omitempty"`

    // Container class for markdown save options.
    ZipOutput bool `json:"ZipOutput,omitempty"`

    // Container class for markdown save options.
    Encoding string `json:"Encoding,omitempty"`

    // Container class for markdown save options.
    ExportHeadersFootersMode string `json:"ExportHeadersFootersMode,omitempty"`

    // Container class for markdown save options.
    ForcePageBreaks bool `json:"ForcePageBreaks,omitempty"`

    // Container class for markdown save options.
    ParagraphBreak string `json:"ParagraphBreak,omitempty"`

    // Container class for markdown save options.
    TableContentAlignment string `json:"TableContentAlignment,omitempty"`
}

type MarkdownSaveOptionsData struct {
    // Container class for markdown save options.
    AllowEmbeddingPostScriptFonts *bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for markdown save options.
    CustomTimeZoneInfoData ITimeZoneInfoData `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for markdown save options.
    Dml3DEffectsRenderingMode *string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for markdown save options.
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for markdown save options.
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // Container class for markdown save options.
    FileName *string `json:"FileName,omitempty"`

    // Container class for markdown save options.
    SaveFormat *string `json:"SaveFormat,omitempty"`

    // Container class for markdown save options.
    UpdateCreatedTimeProperty *bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for markdown save options.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // Container class for markdown save options.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for markdown save options.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for markdown save options.
    UpdateSdtContent *bool `json:"UpdateSdtContent,omitempty"`

    // Container class for markdown save options.
    ZipOutput *bool `json:"ZipOutput,omitempty"`

    // Container class for markdown save options.
    Encoding *string `json:"Encoding,omitempty"`

    // Container class for markdown save options.
    ExportHeadersFootersMode *string `json:"ExportHeadersFootersMode,omitempty"`

    // Container class for markdown save options.
    ForcePageBreaks *bool `json:"ForcePageBreaks,omitempty"`

    // Container class for markdown save options.
    ParagraphBreak *string `json:"ParagraphBreak,omitempty"`

    // Container class for markdown save options.
    TableContentAlignment *string `json:"TableContentAlignment,omitempty"`
}

type IMarkdownSaveOptionsData interface {
    IsMarkdownSaveOptionsData() bool
}
func (MarkdownSaveOptionsData) IsMarkdownSaveOptionsData() bool {
    return true
}

func (MarkdownSaveOptionsData) IsTxtSaveOptionsBaseData() bool {
    return true
}


