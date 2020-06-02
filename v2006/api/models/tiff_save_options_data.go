/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="tiff_save_options_data.go">
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

// Container class for tiff save options.
type TiffSaveOptionsData struct {
    // Container class for tiff save options.
    Dml3DEffectsRenderingMode string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for tiff save options.
    DmlEffectsRenderingMode string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for tiff save options.
    DmlRenderingMode string `json:"DmlRenderingMode,omitempty"`

    // Container class for tiff save options.
    FileName string `json:"FileName,omitempty"`

    // Container class for tiff save options.
    SaveFormat string `json:"SaveFormat,omitempty"`

    // Container class for tiff save options.
    UpdateFields bool `json:"UpdateFields,omitempty"`

    // Container class for tiff save options.
    UpdateLastPrintedProperty bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for tiff save options.
    UpdateLastSavedTimeProperty bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for tiff save options.
    UpdateSdtContent bool `json:"UpdateSdtContent,omitempty"`

    // Container class for tiff save options.
    ZipOutput bool `json:"ZipOutput,omitempty"`

    // Container class for tiff save options.
    ColorMode string `json:"ColorMode,omitempty"`

    // Container class for tiff save options.
    JpegQuality int32 `json:"JpegQuality,omitempty"`

    // Container class for tiff save options.
    MetafileRenderingOptions *MetafileRenderingOptionsData `json:"MetafileRenderingOptions,omitempty"`

    // Container class for tiff save options.
    NumeralFormat string `json:"NumeralFormat,omitempty"`

    // Container class for tiff save options.
    OptimizeOutput bool `json:"OptimizeOutput,omitempty"`

    // Container class for tiff save options.
    PageCount int32 `json:"PageCount,omitempty"`

    // Container class for tiff save options.
    PageIndex int32 `json:"PageIndex,omitempty"`

    // Container class for tiff save options.
    GraphicsQualityOptions *GraphicsQualityOptionsData `json:"GraphicsQualityOptions,omitempty"`

    // Container class for tiff save options.
    HorizontalResolution float64 `json:"HorizontalResolution,omitempty"`

    // Container class for tiff save options.
    ImageBrightness float64 `json:"ImageBrightness,omitempty"`

    // Container class for tiff save options.
    ImageColorMode string `json:"ImageColorMode,omitempty"`

    // Container class for tiff save options.
    ImageContrast float64 `json:"ImageContrast,omitempty"`

    // Container class for tiff save options.
    PaperColor string `json:"PaperColor,omitempty"`

    // Container class for tiff save options.
    PixelFormat string `json:"PixelFormat,omitempty"`

    // Container class for tiff save options.
    Resolution float64 `json:"Resolution,omitempty"`

    // Container class for tiff save options.
    Scale float64 `json:"Scale,omitempty"`

    // Container class for tiff save options.
    UseAntiAliasing bool `json:"UseAntiAliasing,omitempty"`

    // Container class for tiff save options.
    UseGdiEmfRenderer bool `json:"UseGdiEmfRenderer,omitempty"`

    // Container class for tiff save options.
    UseHighQualityRendering bool `json:"UseHighQualityRendering,omitempty"`

    // Container class for tiff save options.
    VerticalResolution float64 `json:"VerticalResolution,omitempty"`

    // Container class for tiff save options.
    ThresholdForFloydSteinbergDithering int32 `json:"ThresholdForFloydSteinbergDithering,omitempty"`

    // Container class for tiff save options.
    TiffBinarizationMethod string `json:"TiffBinarizationMethod,omitempty"`

    // Container class for tiff save options.
    TiffCompression string `json:"TiffCompression,omitempty"`
}

type ITiffSaveOptionsData interface {
    IsTiffSaveOptionsData() bool
}
func (TiffSaveOptionsData) IsTiffSaveOptionsData() bool {
    return true
}

func (TiffSaveOptionsData) IsImageSaveOptionsData() bool {
    return true
}
