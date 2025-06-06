/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="ott_save_options_data.go">
 *   Copyright (c) 2023 Aspose.Words for Cloud
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

// Container class for ott save options.
type OttSaveOptionsDataResult struct {
    // Container class for ott save options.
    AllowEmbeddingPostScriptFonts bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for ott save options.
    CustomTimeZoneInfoData TimeZoneInfoDataResult `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for ott save options.
    Dml3DEffectsRenderingMode string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for ott save options.
    DmlEffectsRenderingMode string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for ott save options.
    DmlRenderingMode string `json:"DmlRenderingMode,omitempty"`

    // Container class for ott save options.
    FileName string `json:"FileName,omitempty"`

    // Container class for ott save options.
    ImlRenderingMode string `json:"ImlRenderingMode,omitempty"`

    // Container class for ott save options.
    UpdateCreatedTimeProperty bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for ott save options.
    UpdateFields bool `json:"UpdateFields,omitempty"`

    // Container class for ott save options.
    UpdateLastPrintedProperty bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for ott save options.
    UpdateLastSavedTimeProperty bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for ott save options.
    ZipOutput bool `json:"ZipOutput,omitempty"`

    // Container class for ott save options.
    IsStrictSchema11 bool `json:"IsStrictSchema11,omitempty"`

    // Container class for ott save options.
    MeasureUnit string `json:"MeasureUnit,omitempty"`

    // Container class for ott save options.
    Password string `json:"Password,omitempty"`

    // Container class for ott save options.
    PrettyFormat bool `json:"PrettyFormat,omitempty"`

    // Container class for ott save options.
    SaveFormat string `json:"SaveFormat,omitempty"`
}

type OttSaveOptionsData struct {
    // Container class for ott save options.
    AllowEmbeddingPostScriptFonts *bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for ott save options.
    CustomTimeZoneInfoData ITimeZoneInfoData `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for ott save options.
    Dml3DEffectsRenderingMode *string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for ott save options.
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for ott save options.
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // Container class for ott save options.
    FileName *string `json:"FileName,omitempty"`

    // Container class for ott save options.
    ImlRenderingMode *string `json:"ImlRenderingMode,omitempty"`

    // Container class for ott save options.
    UpdateCreatedTimeProperty *bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for ott save options.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // Container class for ott save options.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for ott save options.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for ott save options.
    ZipOutput *bool `json:"ZipOutput,omitempty"`

    // Container class for ott save options.
    IsStrictSchema11 *bool `json:"IsStrictSchema11,omitempty"`

    // Container class for ott save options.
    MeasureUnit *string `json:"MeasureUnit,omitempty"`

    // Container class for ott save options.
    Password *string `json:"Password,omitempty"`

    // Container class for ott save options.
    PrettyFormat *bool `json:"PrettyFormat,omitempty"`

    // Container class for ott save options.
    SaveFormat *string `json:"SaveFormat,omitempty"`
}

type IOttSaveOptionsData interface {
    IsOttSaveOptionsData() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
}

func (OttSaveOptionsData) IsOttSaveOptionsData() bool {
    return true
}

func (OttSaveOptionsData) IsOdtSaveOptionsData() bool {
    return true
}

func (OttSaveOptionsData) IsSaveOptionsData() bool {
    return true
}

func (obj *OttSaveOptionsData) Initialize() {
    var _SaveFormat = "ott"
    obj.SaveFormat = &_SaveFormat


    if (obj.CustomTimeZoneInfoData != nil) {
        obj.CustomTimeZoneInfoData.Initialize()
    }


}

func (obj *OttSaveOptionsData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}


