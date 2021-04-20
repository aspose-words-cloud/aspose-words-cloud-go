/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="pcl_save_options_data.go">
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

// Container class for pcl save options.
type PclSaveOptionsDataResult struct {
    // Container class for pcl save options.
    AllowEmbeddingPostScriptFonts bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for pcl save options.
    CustomTimeZoneInfoData TimeZoneInfoDataResult `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for pcl save options.
    Dml3DEffectsRenderingMode string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for pcl save options.
    DmlEffectsRenderingMode string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for pcl save options.
    DmlRenderingMode string `json:"DmlRenderingMode,omitempty"`

    // Container class for pcl save options.
    FileName string `json:"FileName,omitempty"`

    // Container class for pcl save options.
    SaveFormat string `json:"SaveFormat,omitempty"`

    // Container class for pcl save options.
    UpdateCreatedTimeProperty bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for pcl save options.
    UpdateFields bool `json:"UpdateFields,omitempty"`

    // Container class for pcl save options.
    UpdateLastPrintedProperty bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for pcl save options.
    UpdateLastSavedTimeProperty bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for pcl save options.
    UpdateSdtContent bool `json:"UpdateSdtContent,omitempty"`

    // Container class for pcl save options.
    ZipOutput bool `json:"ZipOutput,omitempty"`

    // Container class for pcl save options.
    ColorMode string `json:"ColorMode,omitempty"`

    // Container class for pcl save options.
    JpegQuality int32 `json:"JpegQuality,omitempty"`

    // Container class for pcl save options.
    MetafileRenderingOptions MetafileRenderingOptionsDataResult `json:"MetafileRenderingOptions,omitempty"`

    // Container class for pcl save options.
    NumeralFormat string `json:"NumeralFormat,omitempty"`

    // Container class for pcl save options.
    OptimizeOutput bool `json:"OptimizeOutput,omitempty"`

    // Container class for pcl save options.
    PageCount int32 `json:"PageCount,omitempty"`

    // Container class for pcl save options.
    PageIndex int32 `json:"PageIndex,omitempty"`

    // Container class for pcl save options.
    FalllbackFontName string `json:"FalllbackFontName,omitempty"`

    // Container class for pcl save options.
    RasterizeTransformedElements bool `json:"RasterizeTransformedElements,omitempty"`
}

type PclSaveOptionsData struct {
    // Container class for pcl save options.
    AllowEmbeddingPostScriptFonts *bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for pcl save options.
    CustomTimeZoneInfoData ITimeZoneInfoData `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for pcl save options.
    Dml3DEffectsRenderingMode *string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for pcl save options.
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for pcl save options.
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // Container class for pcl save options.
    FileName *string `json:"FileName,omitempty"`

    // Container class for pcl save options.
    SaveFormat *string `json:"SaveFormat,omitempty"`

    // Container class for pcl save options.
    UpdateCreatedTimeProperty *bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for pcl save options.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // Container class for pcl save options.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for pcl save options.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for pcl save options.
    UpdateSdtContent *bool `json:"UpdateSdtContent,omitempty"`

    // Container class for pcl save options.
    ZipOutput *bool `json:"ZipOutput,omitempty"`

    // Container class for pcl save options.
    ColorMode *string `json:"ColorMode,omitempty"`

    // Container class for pcl save options.
    JpegQuality *int32 `json:"JpegQuality,omitempty"`

    // Container class for pcl save options.
    MetafileRenderingOptions IMetafileRenderingOptionsData `json:"MetafileRenderingOptions,omitempty"`

    // Container class for pcl save options.
    NumeralFormat *string `json:"NumeralFormat,omitempty"`

    // Container class for pcl save options.
    OptimizeOutput *bool `json:"OptimizeOutput,omitempty"`

    // Container class for pcl save options.
    PageCount *int32 `json:"PageCount,omitempty"`

    // Container class for pcl save options.
    PageIndex *int32 `json:"PageIndex,omitempty"`

    // Container class for pcl save options.
    FalllbackFontName *string `json:"FalllbackFontName,omitempty"`

    // Container class for pcl save options.
    RasterizeTransformedElements *bool `json:"RasterizeTransformedElements,omitempty"`
}

type IPclSaveOptionsData interface {
    IsPclSaveOptionsData() bool
}
func (PclSaveOptionsData) IsPclSaveOptionsData() bool {
    return true
}

func (PclSaveOptionsData) IsFixedPageSaveOptionsData() bool {
    return true
}


