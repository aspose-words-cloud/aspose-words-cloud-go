/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="xaml_flow_save_options_data.go">
 *   Copyright (c) 2020 Aspose.Words for Cloud
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

// container class for xaml flow save options.
type XamlFlowSaveOptionsData struct {
    // container class for xaml flow save options.
    Dml3DEffectsRenderingMode string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // container class for xaml flow save options.
    DmlEffectsRenderingMode string `json:"DmlEffectsRenderingMode,omitempty"`

    // container class for xaml flow save options.
    DmlRenderingMode string `json:"DmlRenderingMode,omitempty"`

    // container class for xaml flow save options.
    FileName string `json:"FileName,omitempty"`

    // container class for xaml flow save options.
    SaveFormat string `json:"SaveFormat,omitempty"`

    // container class for xaml flow save options.
    UpdateFields bool `json:"UpdateFields,omitempty"`

    // container class for xaml flow save options.
    UpdateLastPrintedProperty bool `json:"UpdateLastPrintedProperty,omitempty"`

    // container class for xaml flow save options.
    UpdateLastSavedTimeProperty bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // container class for xaml flow save options.
    UpdateSdtContent bool `json:"UpdateSdtContent,omitempty"`

    // container class for xaml flow save options.
    ZipOutput bool `json:"ZipOutput,omitempty"`

    // container class for xaml flow save options.
    ImagesFolder string `json:"ImagesFolder,omitempty"`

    // container class for xaml flow save options.
    ImagesFolderAlias string `json:"ImagesFolderAlias,omitempty"`
}

type IXamlFlowSaveOptionsData interface {
    IsXamlFlowSaveOptionsData() bool
}
func (XamlFlowSaveOptionsData) IsXamlFlowSaveOptionsData() bool {
    return true
}

func (XamlFlowSaveOptionsData) IsSaveOptionsData() bool {
    return true
}
