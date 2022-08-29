/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="rtf_save_options_data.go">
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

// Container class for rtf save options.
type RtfSaveOptionsDataResult struct {
    // Container class for rtf save options.
    AllowEmbeddingPostScriptFonts bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for rtf save options.
    CustomTimeZoneInfoData TimeZoneInfoDataResult `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for rtf save options.
    Dml3DEffectsRenderingMode string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for rtf save options.
    DmlEffectsRenderingMode string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for rtf save options.
    DmlRenderingMode string `json:"DmlRenderingMode,omitempty"`

    // Container class for rtf save options.
    FileName string `json:"FileName,omitempty"`

    // Container class for rtf save options.
    FlatOpcXmlMappingOnly bool `json:"FlatOpcXmlMappingOnly,omitempty"`

    // Container class for rtf save options.
    ImlRenderingMode string `json:"ImlRenderingMode,omitempty"`

    // Container class for rtf save options.
    UpdateCreatedTimeProperty bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for rtf save options.
    UpdateFields bool `json:"UpdateFields,omitempty"`

    // Container class for rtf save options.
    UpdateLastPrintedProperty bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for rtf save options.
    UpdateLastSavedTimeProperty bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for rtf save options.
    UpdateSdtContent bool `json:"UpdateSdtContent,omitempty"`

    // Container class for rtf save options.
    ZipOutput bool `json:"ZipOutput,omitempty"`

    // Container class for rtf save options.
    ExportCompactSize bool `json:"ExportCompactSize,omitempty"`

    // Container class for rtf save options.
    ExportImagesForOldReaders bool `json:"ExportImagesForOldReaders,omitempty"`

    // Container class for rtf save options.
    PrettyFormat bool `json:"PrettyFormat,omitempty"`

    // Container class for rtf save options.
    SaveFormat string `json:"SaveFormat,omitempty"`

    // Container class for rtf save options.
    SaveImagesAsWmf bool `json:"SaveImagesAsWmf,omitempty"`
}

type RtfSaveOptionsData struct {
    // Container class for rtf save options.
    AllowEmbeddingPostScriptFonts *bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for rtf save options.
    CustomTimeZoneInfoData ITimeZoneInfoData `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for rtf save options.
    Dml3DEffectsRenderingMode *string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for rtf save options.
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for rtf save options.
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // Container class for rtf save options.
    FileName *string `json:"FileName,omitempty"`

    // Container class for rtf save options.
    FlatOpcXmlMappingOnly *bool `json:"FlatOpcXmlMappingOnly,omitempty"`

    // Container class for rtf save options.
    ImlRenderingMode *string `json:"ImlRenderingMode,omitempty"`

    // Container class for rtf save options.
    UpdateCreatedTimeProperty *bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for rtf save options.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // Container class for rtf save options.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for rtf save options.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for rtf save options.
    UpdateSdtContent *bool `json:"UpdateSdtContent,omitempty"`

    // Container class for rtf save options.
    ZipOutput *bool `json:"ZipOutput,omitempty"`

    // Container class for rtf save options.
    ExportCompactSize *bool `json:"ExportCompactSize,omitempty"`

    // Container class for rtf save options.
    ExportImagesForOldReaders *bool `json:"ExportImagesForOldReaders,omitempty"`

    // Container class for rtf save options.
    PrettyFormat *bool `json:"PrettyFormat,omitempty"`

    // Container class for rtf save options.
    SaveFormat *string `json:"SaveFormat,omitempty"`

    // Container class for rtf save options.
    SaveImagesAsWmf *bool `json:"SaveImagesAsWmf,omitempty"`
}

type IRtfSaveOptionsData interface {
    IsRtfSaveOptionsData() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileContent) []FileContent
}

func (RtfSaveOptionsData) IsRtfSaveOptionsData() bool {
    return true
}

func (RtfSaveOptionsData) IsSaveOptionsData() bool {
    return true
}

func (obj *RtfSaveOptionsData) Initialize() {
    var _SaveFormat = "rtf"
    obj.SaveFormat = &_SaveFormat


    if (obj.CustomTimeZoneInfoData != nil) {
        obj.CustomTimeZoneInfoData.Initialize()
    }


}

func (obj *RtfSaveOptionsData) CollectFilesContent(resultFilesContent []FileContent) []FileContent {
    return resultFilesContent
}


