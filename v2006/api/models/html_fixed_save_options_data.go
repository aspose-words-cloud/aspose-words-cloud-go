/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="html_fixed_save_options_data.go">
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

// container class for fixed html save options.
type HtmlFixedSaveOptionsData struct {
    // container class for fixed html save options.
    Dml3DEffectsRenderingMode string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // container class for fixed html save options.
    DmlEffectsRenderingMode string `json:"DmlEffectsRenderingMode,omitempty"`

    // container class for fixed html save options.
    DmlRenderingMode string `json:"DmlRenderingMode,omitempty"`

    // container class for fixed html save options.
    FileName string `json:"FileName,omitempty"`

    // container class for fixed html save options.
    SaveFormat string `json:"SaveFormat,omitempty"`

    // container class for fixed html save options.
    UpdateFields bool `json:"UpdateFields,omitempty"`

    // container class for fixed html save options.
    UpdateLastPrintedProperty bool `json:"UpdateLastPrintedProperty,omitempty"`

    // container class for fixed html save options.
    UpdateLastSavedTimeProperty bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // container class for fixed html save options.
    UpdateSdtContent bool `json:"UpdateSdtContent,omitempty"`

    // container class for fixed html save options.
    ZipOutput bool `json:"ZipOutput,omitempty"`

    // container class for fixed html save options.
    ColorMode string `json:"ColorMode,omitempty"`

    // container class for fixed html save options.
    JpegQuality int32 `json:"JpegQuality,omitempty"`

    // container class for fixed html save options.
    MetafileRenderingOptions *MetafileRenderingOptionsData `json:"MetafileRenderingOptions,omitempty"`

    // container class for fixed html save options.
    NumeralFormat string `json:"NumeralFormat,omitempty"`

    // container class for fixed html save options.
    OptimizeOutput bool `json:"OptimizeOutput,omitempty"`

    // container class for fixed html save options.
    PageCount int32 `json:"PageCount,omitempty"`

    // container class for fixed html save options.
    PageIndex int32 `json:"PageIndex,omitempty"`

    // container class for fixed html save options.
    CssClassNamesPrefix string `json:"CssClassNamesPrefix,omitempty"`

    // container class for fixed html save options.
    Encoding string `json:"Encoding,omitempty"`

    // container class for fixed html save options.
    ExportEmbeddedCss bool `json:"ExportEmbeddedCss,omitempty"`

    // container class for fixed html save options.
    ExportEmbeddedFonts bool `json:"ExportEmbeddedFonts,omitempty"`

    // container class for fixed html save options.
    ExportEmbeddedImages bool `json:"ExportEmbeddedImages,omitempty"`

    // container class for fixed html save options.
    ExportFormFields bool `json:"ExportFormFields,omitempty"`

    // container class for fixed html save options.
    FontFormat string `json:"FontFormat,omitempty"`

    // container class for fixed html save options.
    PageHorizontalAlignment string `json:"PageHorizontalAlignment,omitempty"`

    // container class for fixed html save options.
    PageMargins float64 `json:"PageMargins,omitempty"`

    // container class for fixed html save options.
    ResourcesFolder string `json:"ResourcesFolder,omitempty"`

    // container class for fixed html save options.
    ResourcesFolderAlias string `json:"ResourcesFolderAlias,omitempty"`

    // container class for fixed html save options.
    SaveFontFaceCssSeparately bool `json:"SaveFontFaceCssSeparately,omitempty"`

    // container class for fixed html save options.
    ShowPageBorder bool `json:"ShowPageBorder,omitempty"`

    // container class for fixed html save options.
    UseTargetMachineFonts bool `json:"UseTargetMachineFonts,omitempty"`
}

type IHtmlFixedSaveOptionsData interface {
    IsHtmlFixedSaveOptionsData() bool
}
func (HtmlFixedSaveOptionsData) IsHtmlFixedSaveOptionsData() bool {
    return true
}

func (HtmlFixedSaveOptionsData) IsFixedPageSaveOptionsData() bool {
    return true
}
