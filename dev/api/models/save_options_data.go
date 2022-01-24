/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="save_options_data.go">
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

// base container class for save options data.
type SaveOptionsDataResult struct {
    // base container class for save options data.
    AllowEmbeddingPostScriptFonts bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // base container class for save options data.
    CustomTimeZoneInfoData TimeZoneInfoDataResult `json:"CustomTimeZoneInfoData,omitempty"`

    // base container class for save options data.
    Dml3DEffectsRenderingMode string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // base container class for save options data.
    DmlEffectsRenderingMode string `json:"DmlEffectsRenderingMode,omitempty"`

    // base container class for save options data.
    DmlRenderingMode string `json:"DmlRenderingMode,omitempty"`

    // base container class for save options data.
    FileName string `json:"FileName,omitempty"`

    // base container class for save options data.
    FlatOpcXmlMappingOnly bool `json:"FlatOpcXmlMappingOnly,omitempty"`

    // base container class for save options data.
    ImlRenderingMode string `json:"ImlRenderingMode,omitempty"`

    // base container class for save options data.
    SaveFormat string `json:"SaveFormat,omitempty"`

    // base container class for save options data.
    UpdateCreatedTimeProperty bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // base container class for save options data.
    UpdateFields bool `json:"UpdateFields,omitempty"`

    // base container class for save options data.
    UpdateLastPrintedProperty bool `json:"UpdateLastPrintedProperty,omitempty"`

    // base container class for save options data.
    UpdateLastSavedTimeProperty bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // base container class for save options data.
    UpdateSdtContent bool `json:"UpdateSdtContent,omitempty"`

    // base container class for save options data.
    ZipOutput bool `json:"ZipOutput,omitempty"`
}

type SaveOptionsData struct {
    // base container class for save options data.
    AllowEmbeddingPostScriptFonts *bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // base container class for save options data.
    CustomTimeZoneInfoData ITimeZoneInfoData `json:"CustomTimeZoneInfoData,omitempty"`

    // base container class for save options data.
    Dml3DEffectsRenderingMode *string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // base container class for save options data.
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // base container class for save options data.
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // base container class for save options data.
    FileName *string `json:"FileName,omitempty"`

    // base container class for save options data.
    FlatOpcXmlMappingOnly *bool `json:"FlatOpcXmlMappingOnly,omitempty"`

    // base container class for save options data.
    ImlRenderingMode *string `json:"ImlRenderingMode,omitempty"`

    // base container class for save options data.
    const SaveFormat *string `json:"SaveFormat,omitempty"`

    // base container class for save options data.
    UpdateCreatedTimeProperty *bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // base container class for save options data.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // base container class for save options data.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // base container class for save options data.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // base container class for save options data.
    UpdateSdtContent *bool `json:"UpdateSdtContent,omitempty"`

    // base container class for save options data.
    ZipOutput *bool `json:"ZipOutput,omitempty"`
}

type ISaveOptionsData interface {
    IsSaveOptionsData() bool
}
func (SaveOptionsData) IsSaveOptionsData() bool {
    return true
}


