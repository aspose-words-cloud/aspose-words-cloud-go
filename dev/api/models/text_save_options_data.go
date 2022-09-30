/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="text_save_options_data.go">
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

// Container class for text save options.
type TextSaveOptionsDataResult struct {
    // Container class for text save options.
    AllowEmbeddingPostScriptFonts bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for text save options.
    CustomTimeZoneInfoData TimeZoneInfoDataResult `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for text save options.
    Dml3DEffectsRenderingMode string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for text save options.
    DmlEffectsRenderingMode string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for text save options.
    DmlRenderingMode string `json:"DmlRenderingMode,omitempty"`

    // Container class for text save options.
    FileName string `json:"FileName,omitempty"`

    // Container class for text save options.
    FlatOpcXmlMappingOnly bool `json:"FlatOpcXmlMappingOnly,omitempty"`

    // Container class for text save options.
    ImlRenderingMode string `json:"ImlRenderingMode,omitempty"`

    // Container class for text save options.
    UpdateCreatedTimeProperty bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for text save options.
    UpdateFields bool `json:"UpdateFields,omitempty"`

    // Container class for text save options.
    UpdateLastPrintedProperty bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for text save options.
    UpdateLastSavedTimeProperty bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for text save options.
    UpdateSdtContent bool `json:"UpdateSdtContent,omitempty"`

    // Container class for text save options.
    ZipOutput bool `json:"ZipOutput,omitempty"`

    // Container class for text save options.
    Encoding string `json:"Encoding,omitempty"`

    // Container class for text save options.
    ExportHeadersFootersMode string `json:"ExportHeadersFootersMode,omitempty"`

    // Container class for text save options.
    ForcePageBreaks bool `json:"ForcePageBreaks,omitempty"`

    // Container class for text save options.
    ParagraphBreak string `json:"ParagraphBreak,omitempty"`

    // Container class for text save options.
    AddBidiMarks bool `json:"AddBidiMarks,omitempty"`

    // Container class for text save options.
    MaxCharactersPerLine int32 `json:"MaxCharactersPerLine,omitempty"`

    // Container class for text save options.
    PreserveTableLayout bool `json:"PreserveTableLayout,omitempty"`

    // Container class for text save options.
    SaveFormat string `json:"SaveFormat,omitempty"`

    // Container class for text save options.
    SimplifyListLabels bool `json:"SimplifyListLabels,omitempty"`
}

type TextSaveOptionsData struct {
    // Container class for text save options.
    AllowEmbeddingPostScriptFonts *bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for text save options.
    CustomTimeZoneInfoData ITimeZoneInfoData `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for text save options.
    Dml3DEffectsRenderingMode *string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for text save options.
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for text save options.
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // Container class for text save options.
    FileName *string `json:"FileName,omitempty"`

    // Container class for text save options.
    FlatOpcXmlMappingOnly *bool `json:"FlatOpcXmlMappingOnly,omitempty"`

    // Container class for text save options.
    ImlRenderingMode *string `json:"ImlRenderingMode,omitempty"`

    // Container class for text save options.
    UpdateCreatedTimeProperty *bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for text save options.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // Container class for text save options.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for text save options.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for text save options.
    UpdateSdtContent *bool `json:"UpdateSdtContent,omitempty"`

    // Container class for text save options.
    ZipOutput *bool `json:"ZipOutput,omitempty"`

    // Container class for text save options.
    Encoding *string `json:"Encoding,omitempty"`

    // Container class for text save options.
    ExportHeadersFootersMode *string `json:"ExportHeadersFootersMode,omitempty"`

    // Container class for text save options.
    ForcePageBreaks *bool `json:"ForcePageBreaks,omitempty"`

    // Container class for text save options.
    ParagraphBreak *string `json:"ParagraphBreak,omitempty"`

    // Container class for text save options.
    AddBidiMarks *bool `json:"AddBidiMarks,omitempty"`

    // Container class for text save options.
    MaxCharactersPerLine *int32 `json:"MaxCharactersPerLine,omitempty"`

    // Container class for text save options.
    PreserveTableLayout *bool `json:"PreserveTableLayout,omitempty"`

    // Container class for text save options.
    SaveFormat *string `json:"SaveFormat,omitempty"`

    // Container class for text save options.
    SimplifyListLabels *bool `json:"SimplifyListLabels,omitempty"`
}

type ITextSaveOptionsData interface {
    IsTextSaveOptionsData() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
}

func (TextSaveOptionsData) IsTextSaveOptionsData() bool {
    return true
}

func (TextSaveOptionsData) IsTxtSaveOptionsBaseData() bool {
    return true
}

func (TextSaveOptionsData) IsSaveOptionsData() bool {
    return true
}

func (obj *TextSaveOptionsData) Initialize() {
    var _SaveFormat = "txt"
    obj.SaveFormat = &_SaveFormat


}

func (obj *TextSaveOptionsData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}


