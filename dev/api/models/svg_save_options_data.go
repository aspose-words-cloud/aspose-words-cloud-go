/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="svg_save_options_data.go">
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

// Container class for svg save options.
type SvgSaveOptionsDataResult struct {
    // Container class for svg save options.
    AllowEmbeddingPostScriptFonts bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for svg save options.
    Dml3DEffectsRenderingMode string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for svg save options.
    DmlEffectsRenderingMode string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for svg save options.
    DmlRenderingMode string `json:"DmlRenderingMode,omitempty"`

    // Container class for svg save options.
    FileName string `json:"FileName,omitempty"`

    // Container class for svg save options.
    SaveFormat string `json:"SaveFormat,omitempty"`

    // Container class for svg save options.
    UpdateCreatedTimeProperty bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for svg save options.
    UpdateFields bool `json:"UpdateFields,omitempty"`

    // Container class for svg save options.
    UpdateLastPrintedProperty bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for svg save options.
    UpdateLastSavedTimeProperty bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for svg save options.
    UpdateSdtContent bool `json:"UpdateSdtContent,omitempty"`

    // Container class for svg save options.
    ZipOutput bool `json:"ZipOutput,omitempty"`

    // Container class for svg save options.
    ColorMode string `json:"ColorMode,omitempty"`

    // Container class for svg save options.
    JpegQuality int32 `json:"JpegQuality,omitempty"`

    // Container class for svg save options.
    MetafileRenderingOptions MetafileRenderingOptionsDataResult `json:"MetafileRenderingOptions,omitempty"`

    // Container class for svg save options.
    NumeralFormat string `json:"NumeralFormat,omitempty"`

    // Container class for svg save options.
    OptimizeOutput bool `json:"OptimizeOutput,omitempty"`

    // Container class for svg save options.
    PageCount int32 `json:"PageCount,omitempty"`

    // Container class for svg save options.
    PageIndex int32 `json:"PageIndex,omitempty"`

    // Container class for svg save options.
    ExportEmbeddedImages bool `json:"ExportEmbeddedImages,omitempty"`

    // Container class for svg save options.
    FitToViewPort bool `json:"FitToViewPort,omitempty"`

    // Container class for svg save options.
    ResourcesFolder string `json:"ResourcesFolder,omitempty"`

    // Container class for svg save options.
    ResourcesFolderAlias string `json:"ResourcesFolderAlias,omitempty"`

    // Container class for svg save options.
    ShowPageBorder bool `json:"ShowPageBorder,omitempty"`

    // Container class for svg save options.
    TextOutputMode string `json:"TextOutputMode,omitempty"`
}

type SvgSaveOptionsData struct {
    // Container class for svg save options.
    AllowEmbeddingPostScriptFonts *bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for svg save options.
    Dml3DEffectsRenderingMode *string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for svg save options.
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for svg save options.
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // Container class for svg save options.
    FileName *string `json:"FileName,omitempty"`

    // Container class for svg save options.
    SaveFormat *string `json:"SaveFormat,omitempty"`

    // Container class for svg save options.
    UpdateCreatedTimeProperty *bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for svg save options.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // Container class for svg save options.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for svg save options.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for svg save options.
    UpdateSdtContent *bool `json:"UpdateSdtContent,omitempty"`

    // Container class for svg save options.
    ZipOutput *bool `json:"ZipOutput,omitempty"`

    // Container class for svg save options.
    ColorMode *string `json:"ColorMode,omitempty"`

    // Container class for svg save options.
    JpegQuality *int32 `json:"JpegQuality,omitempty"`

    // Container class for svg save options.
    MetafileRenderingOptions IMetafileRenderingOptionsData `json:"MetafileRenderingOptions,omitempty"`

    // Container class for svg save options.
    NumeralFormat *string `json:"NumeralFormat,omitempty"`

    // Container class for svg save options.
    OptimizeOutput *bool `json:"OptimizeOutput,omitempty"`

    // Container class for svg save options.
    PageCount *int32 `json:"PageCount,omitempty"`

    // Container class for svg save options.
    PageIndex *int32 `json:"PageIndex,omitempty"`

    // Container class for svg save options.
    ExportEmbeddedImages *bool `json:"ExportEmbeddedImages,omitempty"`

    // Container class for svg save options.
    FitToViewPort *bool `json:"FitToViewPort,omitempty"`

    // Container class for svg save options.
    ResourcesFolder *string `json:"ResourcesFolder,omitempty"`

    // Container class for svg save options.
    ResourcesFolderAlias *string `json:"ResourcesFolderAlias,omitempty"`

    // Container class for svg save options.
    ShowPageBorder *bool `json:"ShowPageBorder,omitempty"`

    // Container class for svg save options.
    TextOutputMode *string `json:"TextOutputMode,omitempty"`
}

type ISvgSaveOptionsData interface {
    IsSvgSaveOptionsData() bool
}
func (SvgSaveOptionsData) IsSvgSaveOptionsData() bool {
    return true
}

func (SvgSaveOptionsData) IsFixedPageSaveOptionsData() bool {
    return true
}


