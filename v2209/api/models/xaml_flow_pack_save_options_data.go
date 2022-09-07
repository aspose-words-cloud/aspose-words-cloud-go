/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="xaml_flow_pack_save_options_data.go">
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

// Container class for xamlflow_pack save options.
type XamlFlowPackSaveOptionsDataResult struct {
    // Container class for xamlflow_pack save options.
    AllowEmbeddingPostScriptFonts bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for xamlflow_pack save options.
    CustomTimeZoneInfoData TimeZoneInfoDataResult `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for xamlflow_pack save options.
    Dml3DEffectsRenderingMode string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for xamlflow_pack save options.
    DmlEffectsRenderingMode string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for xamlflow_pack save options.
    DmlRenderingMode string `json:"DmlRenderingMode,omitempty"`

    // Container class for xamlflow_pack save options.
    FileName string `json:"FileName,omitempty"`

    // Container class for xamlflow_pack save options.
    FlatOpcXmlMappingOnly bool `json:"FlatOpcXmlMappingOnly,omitempty"`

    // Container class for xamlflow_pack save options.
    ImlRenderingMode string `json:"ImlRenderingMode,omitempty"`

    // Container class for xamlflow_pack save options.
    UpdateCreatedTimeProperty bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for xamlflow_pack save options.
    UpdateFields bool `json:"UpdateFields,omitempty"`

    // Container class for xamlflow_pack save options.
    UpdateLastPrintedProperty bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for xamlflow_pack save options.
    UpdateLastSavedTimeProperty bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for xamlflow_pack save options.
    UpdateSdtContent bool `json:"UpdateSdtContent,omitempty"`

    // Container class for xamlflow_pack save options.
    ZipOutput bool `json:"ZipOutput,omitempty"`

    // Container class for xamlflow_pack save options.
    ImagesFolder string `json:"ImagesFolder,omitempty"`

    // Container class for xamlflow_pack save options.
    ImagesFolderAlias string `json:"ImagesFolderAlias,omitempty"`

    // Container class for xamlflow_pack save options.
    SaveFormat string `json:"SaveFormat,omitempty"`
}

type XamlFlowPackSaveOptionsData struct {
    // Container class for xamlflow_pack save options.
    AllowEmbeddingPostScriptFonts *bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for xamlflow_pack save options.
    CustomTimeZoneInfoData ITimeZoneInfoData `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for xamlflow_pack save options.
    Dml3DEffectsRenderingMode *string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for xamlflow_pack save options.
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for xamlflow_pack save options.
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // Container class for xamlflow_pack save options.
    FileName *string `json:"FileName,omitempty"`

    // Container class for xamlflow_pack save options.
    FlatOpcXmlMappingOnly *bool `json:"FlatOpcXmlMappingOnly,omitempty"`

    // Container class for xamlflow_pack save options.
    ImlRenderingMode *string `json:"ImlRenderingMode,omitempty"`

    // Container class for xamlflow_pack save options.
    UpdateCreatedTimeProperty *bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for xamlflow_pack save options.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // Container class for xamlflow_pack save options.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for xamlflow_pack save options.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for xamlflow_pack save options.
    UpdateSdtContent *bool `json:"UpdateSdtContent,omitempty"`

    // Container class for xamlflow_pack save options.
    ZipOutput *bool `json:"ZipOutput,omitempty"`

    // Container class for xamlflow_pack save options.
    ImagesFolder *string `json:"ImagesFolder,omitempty"`

    // Container class for xamlflow_pack save options.
    ImagesFolderAlias *string `json:"ImagesFolderAlias,omitempty"`

    // Container class for xamlflow_pack save options.
    SaveFormat *string `json:"SaveFormat,omitempty"`
}

type IXamlFlowPackSaveOptionsData interface {
    IsXamlFlowPackSaveOptionsData() bool
    Initialize()
}

func (XamlFlowPackSaveOptionsData) IsXamlFlowPackSaveOptionsData() bool {
    return true
}

func (XamlFlowPackSaveOptionsData) IsXamlFlowSaveOptionsData() bool {
    return true
}

func (XamlFlowPackSaveOptionsData) IsSaveOptionsData() bool {
    return true
}

func (obj *XamlFlowPackSaveOptionsData) Initialize() {
    var _SaveFormat = "xamlflow_pack"
    obj.SaveFormat = &_SaveFormat

    if (obj.CustomTimeZoneInfoData != nil) {
        obj.CustomTimeZoneInfoData.Initialize()
    }





























}


