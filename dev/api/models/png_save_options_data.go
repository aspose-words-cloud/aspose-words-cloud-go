/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="png_save_options_data.go">
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

// container class for png save options.
type PngSaveOptionsData struct {
    // container class for png save options.
    Dml3DEffectsRenderingMode string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // container class for png save options.
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // container class for png save options.
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // container class for png save options.
    FileName *string `json:"FileName,omitempty"`

    // container class for png save options.
    SaveFormat *string `json:"SaveFormat,omitempty"`

    // container class for png save options.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // container class for png save options.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // container class for png save options.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // container class for png save options.
    UpdateSdtContent *bool `json:"UpdateSdtContent,omitempty"`

    // container class for png save options.
    ZipOutput *bool `json:"ZipOutput,omitempty"`

    // container class for png save options.
    ColorMode *string `json:"ColorMode,omitempty"`

    // container class for png save options.
    JpegQuality *int32 `json:"JpegQuality,omitempty"`

    // container class for png save options.
    MetafileRenderingOptions *MetafileRenderingOptionsData `json:"MetafileRenderingOptions,omitempty"`

    // container class for png save options.
    NumeralFormat *string `json:"NumeralFormat,omitempty"`

    // container class for png save options.
    OptimizeOutput *bool `json:"OptimizeOutput,omitempty"`

    // container class for png save options.
    PageCount *int32 `json:"PageCount,omitempty"`

    // container class for png save options.
    PageIndex *int32 `json:"PageIndex,omitempty"`

    // container class for png save options.
    GraphicsQualityOptions *GraphicsQualityOptionsData `json:"GraphicsQualityOptions,omitempty"`

    // container class for png save options.
    HorizontalResolution *float64 `json:"HorizontalResolution,omitempty"`

    // container class for png save options.
    ImageBrightness *float64 `json:"ImageBrightness,omitempty"`

    // container class for png save options.
    ImageColorMode *string `json:"ImageColorMode,omitempty"`

    // container class for png save options.
    ImageContrast *float64 `json:"ImageContrast,omitempty"`

    // container class for png save options.
    PaperColor *string `json:"PaperColor,omitempty"`

    // container class for png save options.
    PixelFormat *string `json:"PixelFormat,omitempty"`

    // container class for png save options.
    Resolution *float64 `json:"Resolution,omitempty"`

    // container class for png save options.
    Scale *float64 `json:"Scale,omitempty"`

    // container class for png save options.
    UseAntiAliasing *bool `json:"UseAntiAliasing,omitempty"`

    // container class for png save options.
    UseGdiEmfRenderer *bool `json:"UseGdiEmfRenderer,omitempty"`

    // container class for png save options.
    UseHighQualityRendering *bool `json:"UseHighQualityRendering,omitempty"`

    // container class for png save options.
    VerticalResolution *float64 `json:"VerticalResolution,omitempty"`
}

type IPngSaveOptionsData interface {
    IsPngSaveOptionsData() bool
}
func (PngSaveOptionsData) IsPngSaveOptionsData() bool {
    return true
}

func (PngSaveOptionsData) IsImageSaveOptionsData() bool {
    return true
}
