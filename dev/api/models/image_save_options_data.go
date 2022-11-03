/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="image_save_options_data.go">
 *   Copyright (c) 2022 Aspose.Words for Cloud
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

// Container abstract class for image save options.
type ImageSaveOptionsDataResult struct {
    // Container abstract class for image save options.
    AllowEmbeddingPostScriptFonts bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container abstract class for image save options.
    CustomTimeZoneInfoData TimeZoneInfoDataResult `json:"CustomTimeZoneInfoData,omitempty"`

    // Container abstract class for image save options.
    Dml3DEffectsRenderingMode string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container abstract class for image save options.
    DmlEffectsRenderingMode string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container abstract class for image save options.
    DmlRenderingMode string `json:"DmlRenderingMode,omitempty"`

    // Container abstract class for image save options.
    FileName string `json:"FileName,omitempty"`

    // Container abstract class for image save options.
    FlatOpcXmlMappingOnly bool `json:"FlatOpcXmlMappingOnly,omitempty"`

    // Container abstract class for image save options.
    ImlRenderingMode string `json:"ImlRenderingMode,omitempty"`

    // Container abstract class for image save options.
    SaveFormat string `json:"SaveFormat,omitempty"`

    // Container abstract class for image save options.
    UpdateCreatedTimeProperty bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container abstract class for image save options.
    UpdateFields bool `json:"UpdateFields,omitempty"`

    // Container abstract class for image save options.
    UpdateLastPrintedProperty bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container abstract class for image save options.
    UpdateLastSavedTimeProperty bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container abstract class for image save options.
    UpdateSdtContent bool `json:"UpdateSdtContent,omitempty"`

    // Container abstract class for image save options.
    ZipOutput bool `json:"ZipOutput,omitempty"`

    // Container abstract class for image save options.
    ColorMode string `json:"ColorMode,omitempty"`

    // Container abstract class for image save options.
    JpegQuality int32 `json:"JpegQuality,omitempty"`

    // Container abstract class for image save options.
    MetafileRenderingOptions MetafileRenderingOptionsDataResult `json:"MetafileRenderingOptions,omitempty"`

    // Container abstract class for image save options.
    NumeralFormat string `json:"NumeralFormat,omitempty"`

    // Container abstract class for image save options.
    OptimizeOutput bool `json:"OptimizeOutput,omitempty"`

    // Container abstract class for image save options.
    PageCount int32 `json:"PageCount,omitempty"`

    // Container abstract class for image save options.
    PageIndex int32 `json:"PageIndex,omitempty"`

    // Container abstract class for image save options.
    HorizontalResolution float64 `json:"HorizontalResolution,omitempty"`

    // Container abstract class for image save options.
    ImageBrightness float64 `json:"ImageBrightness,omitempty"`

    // Container abstract class for image save options.
    ImageColorMode string `json:"ImageColorMode,omitempty"`

    // Container abstract class for image save options.
    ImageContrast float64 `json:"ImageContrast,omitempty"`

    // Container abstract class for image save options.
    PaperColor string `json:"PaperColor,omitempty"`

    // Container abstract class for image save options.
    PixelFormat string `json:"PixelFormat,omitempty"`

    // Container abstract class for image save options.
    Resolution float64 `json:"Resolution,omitempty"`

    // Container abstract class for image save options.
    Scale float64 `json:"Scale,omitempty"`

    // Container abstract class for image save options.
    UseAntiAliasing bool `json:"UseAntiAliasing,omitempty"`

    // Container abstract class for image save options.
    UseGdiEmfRenderer bool `json:"UseGdiEmfRenderer,omitempty"`

    // Container abstract class for image save options.
    UseHighQualityRendering bool `json:"UseHighQualityRendering,omitempty"`

    // Container abstract class for image save options.
    VerticalResolution float64 `json:"VerticalResolution,omitempty"`
}

type ImageSaveOptionsData struct {
    // Container abstract class for image save options.
    AllowEmbeddingPostScriptFonts *bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container abstract class for image save options.
    CustomTimeZoneInfoData ITimeZoneInfoData `json:"CustomTimeZoneInfoData,omitempty"`

    // Container abstract class for image save options.
    Dml3DEffectsRenderingMode *string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container abstract class for image save options.
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container abstract class for image save options.
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // Container abstract class for image save options.
    FileName *string `json:"FileName,omitempty"`

    // Container abstract class for image save options.
    FlatOpcXmlMappingOnly *bool `json:"FlatOpcXmlMappingOnly,omitempty"`

    // Container abstract class for image save options.
    ImlRenderingMode *string `json:"ImlRenderingMode,omitempty"`

    // Container abstract class for image save options.
    SaveFormat *string `json:"SaveFormat,omitempty"`

    // Container abstract class for image save options.
    UpdateCreatedTimeProperty *bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container abstract class for image save options.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // Container abstract class for image save options.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container abstract class for image save options.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container abstract class for image save options.
    UpdateSdtContent *bool `json:"UpdateSdtContent,omitempty"`

    // Container abstract class for image save options.
    ZipOutput *bool `json:"ZipOutput,omitempty"`

    // Container abstract class for image save options.
    ColorMode *string `json:"ColorMode,omitempty"`

    // Container abstract class for image save options.
    JpegQuality *int32 `json:"JpegQuality,omitempty"`

    // Container abstract class for image save options.
    MetafileRenderingOptions IMetafileRenderingOptionsData `json:"MetafileRenderingOptions,omitempty"`

    // Container abstract class for image save options.
    NumeralFormat *string `json:"NumeralFormat,omitempty"`

    // Container abstract class for image save options.
    OptimizeOutput *bool `json:"OptimizeOutput,omitempty"`

    // Container abstract class for image save options.
    PageCount *int32 `json:"PageCount,omitempty"`

    // Container abstract class for image save options.
    PageIndex *int32 `json:"PageIndex,omitempty"`

    // Container abstract class for image save options.
    HorizontalResolution *float64 `json:"HorizontalResolution,omitempty"`

    // Container abstract class for image save options.
    ImageBrightness *float64 `json:"ImageBrightness,omitempty"`

    // Container abstract class for image save options.
    ImageColorMode *string `json:"ImageColorMode,omitempty"`

    // Container abstract class for image save options.
    ImageContrast *float64 `json:"ImageContrast,omitempty"`

    // Container abstract class for image save options.
    PaperColor *string `json:"PaperColor,omitempty"`

    // Container abstract class for image save options.
    PixelFormat *string `json:"PixelFormat,omitempty"`

    // Container abstract class for image save options.
    Resolution *float64 `json:"Resolution,omitempty"`

    // Container abstract class for image save options.
    Scale *float64 `json:"Scale,omitempty"`

    // Container abstract class for image save options.
    UseAntiAliasing *bool `json:"UseAntiAliasing,omitempty"`

    // Container abstract class for image save options.
    UseGdiEmfRenderer *bool `json:"UseGdiEmfRenderer,omitempty"`

    // Container abstract class for image save options.
    UseHighQualityRendering *bool `json:"UseHighQualityRendering,omitempty"`

    // Container abstract class for image save options.
    VerticalResolution *float64 `json:"VerticalResolution,omitempty"`
}

type IImageSaveOptionsData interface {
    IsImageSaveOptionsData() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
}

func (ImageSaveOptionsData) IsImageSaveOptionsData() bool {
    return true
}

func (ImageSaveOptionsData) IsFixedPageSaveOptionsData() bool {
    return true
}

func (ImageSaveOptionsData) IsSaveOptionsData() bool {
    return true
}

func (obj *ImageSaveOptionsData) Initialize() {
    if (obj.CustomTimeZoneInfoData != nil) {
        obj.CustomTimeZoneInfoData.Initialize()
    }

    if (obj.MetafileRenderingOptions != nil) {
        obj.MetafileRenderingOptions.Initialize()
    }


}

func (obj *ImageSaveOptionsData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}


