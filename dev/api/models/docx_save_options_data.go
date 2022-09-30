/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="docx_save_options_data.go">
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

// Container class for docx save options.
type DocxSaveOptionsDataResult struct {
    // Container class for docx save options.
    AllowEmbeddingPostScriptFonts bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for docx save options.
    CustomTimeZoneInfoData TimeZoneInfoDataResult `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for docx save options.
    Dml3DEffectsRenderingMode string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for docx save options.
    DmlEffectsRenderingMode string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for docx save options.
    DmlRenderingMode string `json:"DmlRenderingMode,omitempty"`

    // Container class for docx save options.
    FileName string `json:"FileName,omitempty"`

    // Container class for docx save options.
    FlatOpcXmlMappingOnly bool `json:"FlatOpcXmlMappingOnly,omitempty"`

    // Container class for docx save options.
    ImlRenderingMode string `json:"ImlRenderingMode,omitempty"`

    // Container class for docx save options.
    UpdateCreatedTimeProperty bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for docx save options.
    UpdateFields bool `json:"UpdateFields,omitempty"`

    // Container class for docx save options.
    UpdateLastPrintedProperty bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for docx save options.
    UpdateLastSavedTimeProperty bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for docx save options.
    UpdateSdtContent bool `json:"UpdateSdtContent,omitempty"`

    // Container class for docx save options.
    ZipOutput bool `json:"ZipOutput,omitempty"`

    // Container class for docx save options.
    Compliance string `json:"Compliance,omitempty"`

    // Container class for docx save options.
    CompressionLevel string `json:"CompressionLevel,omitempty"`

    // Container class for docx save options.
    Password string `json:"Password,omitempty"`

    // Container class for docx save options.
    PrettyFormat bool `json:"PrettyFormat,omitempty"`

    // Container class for docx save options.
    SaveFormat string `json:"SaveFormat,omitempty"`
}

type DocxSaveOptionsData struct {
    // Container class for docx save options.
    AllowEmbeddingPostScriptFonts *bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for docx save options.
    CustomTimeZoneInfoData ITimeZoneInfoData `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for docx save options.
    Dml3DEffectsRenderingMode *string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for docx save options.
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for docx save options.
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // Container class for docx save options.
    FileName *string `json:"FileName,omitempty"`

    // Container class for docx save options.
    FlatOpcXmlMappingOnly *bool `json:"FlatOpcXmlMappingOnly,omitempty"`

    // Container class for docx save options.
    ImlRenderingMode *string `json:"ImlRenderingMode,omitempty"`

    // Container class for docx save options.
    UpdateCreatedTimeProperty *bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for docx save options.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // Container class for docx save options.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for docx save options.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for docx save options.
    UpdateSdtContent *bool `json:"UpdateSdtContent,omitempty"`

    // Container class for docx save options.
    ZipOutput *bool `json:"ZipOutput,omitempty"`

    // Container class for docx save options.
    Compliance *string `json:"Compliance,omitempty"`

    // Container class for docx save options.
    CompressionLevel *string `json:"CompressionLevel,omitempty"`

    // Container class for docx save options.
    Password *string `json:"Password,omitempty"`

    // Container class for docx save options.
    PrettyFormat *bool `json:"PrettyFormat,omitempty"`

    // Container class for docx save options.
    SaveFormat *string `json:"SaveFormat,omitempty"`
}

type IDocxSaveOptionsData interface {
    IsDocxSaveOptionsData() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
}

func (DocxSaveOptionsData) IsDocxSaveOptionsData() bool {
    return true
}

func (DocxSaveOptionsData) IsOoxmlSaveOptionsData() bool {
    return true
}

func (DocxSaveOptionsData) IsSaveOptionsData() bool {
    return true
}

func (obj *DocxSaveOptionsData) Initialize() {
    var _SaveFormat = "docx"
    obj.SaveFormat = &_SaveFormat


}

func (obj *DocxSaveOptionsData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}


