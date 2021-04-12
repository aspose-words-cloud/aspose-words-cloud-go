/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="fixed_page_save_options_data.go">
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

// Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
type FixedPageSaveOptionsDataResult struct {
    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    AllowEmbeddingPostScriptFonts bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    CustomTimeZoneInfoData TimeZoneInfoDataResult `json:"CustomTimeZoneInfoData,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    Dml3DEffectsRenderingMode string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    DmlEffectsRenderingMode string `json:"DmlEffectsRenderingMode,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    DmlRenderingMode string `json:"DmlRenderingMode,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    FileName string `json:"FileName,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    SaveFormat string `json:"SaveFormat,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    UpdateCreatedTimeProperty bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    UpdateFields bool `json:"UpdateFields,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    UpdateLastPrintedProperty bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    UpdateLastSavedTimeProperty bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    UpdateSdtContent bool `json:"UpdateSdtContent,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    ZipOutput bool `json:"ZipOutput,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    ColorMode string `json:"ColorMode,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    JpegQuality int32 `json:"JpegQuality,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    MetafileRenderingOptions MetafileRenderingOptionsDataResult `json:"MetafileRenderingOptions,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    NumeralFormat string `json:"NumeralFormat,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    OptimizeOutput bool `json:"OptimizeOutput,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    PageCount int32 `json:"PageCount,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    PageIndex int32 `json:"PageIndex,omitempty"`
}

type FixedPageSaveOptionsData struct {
    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    AllowEmbeddingPostScriptFonts *bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    CustomTimeZoneInfoData ITimeZoneInfoData `json:"CustomTimeZoneInfoData,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    Dml3DEffectsRenderingMode *string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    FileName *string `json:"FileName,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    SaveFormat *string `json:"SaveFormat,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    UpdateCreatedTimeProperty *bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    UpdateSdtContent *bool `json:"UpdateSdtContent,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    ZipOutput *bool `json:"ZipOutput,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    ColorMode *string `json:"ColorMode,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    JpegQuality *int32 `json:"JpegQuality,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    MetafileRenderingOptions IMetafileRenderingOptionsData `json:"MetafileRenderingOptions,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    NumeralFormat *string `json:"NumeralFormat,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    OptimizeOutput *bool `json:"OptimizeOutput,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    PageCount *int32 `json:"PageCount,omitempty"`

    // Contains common options that can be specified when saving a document into fixed page formats (PDF, XPS, images etc).
    PageIndex *int32 `json:"PageIndex,omitempty"`
}

type IFixedPageSaveOptionsData interface {
    IsFixedPageSaveOptionsData() bool
}
func (FixedPageSaveOptionsData) IsFixedPageSaveOptionsData() bool {
    return true
}

func (FixedPageSaveOptionsData) IsSaveOptionsData() bool {
    return true
}


