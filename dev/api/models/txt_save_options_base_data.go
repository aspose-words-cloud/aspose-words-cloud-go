/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="txt_save_options_base_data.go">
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

// Base class for save options of text formats.
type TxtSaveOptionsBaseDataResult struct {
    // Base class for save options of text formats.
    Dml3DEffectsRenderingMode string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Base class for save options of text formats.
    DmlEffectsRenderingMode string `json:"DmlEffectsRenderingMode,omitempty"`

    // Base class for save options of text formats.
    DmlRenderingMode string `json:"DmlRenderingMode,omitempty"`

    // Base class for save options of text formats.
    FileName string `json:"FileName,omitempty"`

    // Base class for save options of text formats.
    SaveFormat string `json:"SaveFormat,omitempty"`

    // Base class for save options of text formats.
    UpdateFields bool `json:"UpdateFields,omitempty"`

    // Base class for save options of text formats.
    UpdateLastPrintedProperty bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Base class for save options of text formats.
    UpdateLastSavedTimeProperty bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Base class for save options of text formats.
    UpdateSdtContent bool `json:"UpdateSdtContent,omitempty"`

    // Base class for save options of text formats.
    ZipOutput bool `json:"ZipOutput,omitempty"`

    // Base class for save options of text formats.
    Encoding string `json:"Encoding,omitempty"`

    // Base class for save options of text formats.
    ExportHeadersFootersMode string `json:"ExportHeadersFootersMode,omitempty"`

    // Base class for save options of text formats.
    ForcePageBreaks bool `json:"ForcePageBreaks,omitempty"`

    // Base class for save options of text formats.
    ParagraphBreak string `json:"ParagraphBreak,omitempty"`
}

type TxtSaveOptionsBaseData struct {
    // Base class for save options of text formats.
    Dml3DEffectsRenderingMode *string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Base class for save options of text formats.
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // Base class for save options of text formats.
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // Base class for save options of text formats.
    FileName *string `json:"FileName,omitempty"`

    // Base class for save options of text formats.
    SaveFormat *string `json:"SaveFormat,omitempty"`

    // Base class for save options of text formats.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // Base class for save options of text formats.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Base class for save options of text formats.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Base class for save options of text formats.
    UpdateSdtContent *bool `json:"UpdateSdtContent,omitempty"`

    // Base class for save options of text formats.
    ZipOutput *bool `json:"ZipOutput,omitempty"`

    // Base class for save options of text formats.
    Encoding *string `json:"Encoding,omitempty"`

    // Base class for save options of text formats.
    ExportHeadersFootersMode *string `json:"ExportHeadersFootersMode,omitempty"`

    // Base class for save options of text formats.
    ForcePageBreaks *bool `json:"ForcePageBreaks,omitempty"`

    // Base class for save options of text formats.
    ParagraphBreak *string `json:"ParagraphBreak,omitempty"`
}

type ITxtSaveOptionsBaseData interface {
    IsTxtSaveOptionsBaseData() bool
}
func (TxtSaveOptionsBaseData) IsTxtSaveOptionsBaseData() bool {
    return true
}

func (TxtSaveOptionsBaseData) IsSaveOptionsData() bool {
    return true
}


