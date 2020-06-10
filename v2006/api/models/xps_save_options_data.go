/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="xps_save_options_data.go">
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

// Container class for xps save options.
type XpsSaveOptionsData struct {
    // Container class for xps save options.
    Dml3DEffectsRenderingMode string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for xps save options.
    DmlEffectsRenderingMode string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for xps save options.
    DmlRenderingMode string `json:"DmlRenderingMode,omitempty"`

    // Container class for xps save options.
    FileName string `json:"FileName,omitempty"`

    // Container class for xps save options.
    SaveFormat string `json:"SaveFormat,omitempty"`

    // Container class for xps save options.
    UpdateFields bool `json:"UpdateFields,omitempty"`

    // Container class for xps save options.
    UpdateLastPrintedProperty bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for xps save options.
    UpdateLastSavedTimeProperty bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for xps save options.
    UpdateSdtContent bool `json:"UpdateSdtContent,omitempty"`

    // Container class for xps save options.
    ZipOutput bool `json:"ZipOutput,omitempty"`

    // Container class for xps save options.
    ColorMode string `json:"ColorMode,omitempty"`

    // Container class for xps save options.
    JpegQuality int32 `json:"JpegQuality,omitempty"`

    // Container class for xps save options.
    MetafileRenderingOptions *MetafileRenderingOptionsData `json:"MetafileRenderingOptions,omitempty"`

    // Container class for xps save options.
    NumeralFormat string `json:"NumeralFormat,omitempty"`

    // Container class for xps save options.
    OptimizeOutput bool `json:"OptimizeOutput,omitempty"`

    // Container class for xps save options.
    PageCount int32 `json:"PageCount,omitempty"`

    // Container class for xps save options.
    PageIndex int32 `json:"PageIndex,omitempty"`

    // Container class for xps save options.
    BookmarksOutlineLevel int32 `json:"BookmarksOutlineLevel,omitempty"`

    // Container class for xps save options.
    HeadingsOutlineLevels int32 `json:"HeadingsOutlineLevels,omitempty"`

    // Container class for xps save options.
    OutlineOptions *OutlineOptionsData `json:"OutlineOptions,omitempty"`

    // Container class for xps save options.
    UseBookFoldPrintingSettings bool `json:"UseBookFoldPrintingSettings,omitempty"`
}

type IXpsSaveOptionsData interface {
    IsXpsSaveOptionsData() bool
}
func (XpsSaveOptionsData) IsXpsSaveOptionsData() bool {
    return true
}

func (XpsSaveOptionsData) IsFixedPageSaveOptionsData() bool {
    return true
}
