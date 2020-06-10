/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="doc_save_options_data.go">
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

// container class for doc/dot save options.
type DocSaveOptionsData struct {
    // container class for doc/dot save options.
    Dml3DEffectsRenderingMode string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // container class for doc/dot save options.
    DmlEffectsRenderingMode string `json:"DmlEffectsRenderingMode,omitempty"`

    // container class for doc/dot save options.
    DmlRenderingMode string `json:"DmlRenderingMode,omitempty"`

    // container class for doc/dot save options.
    FileName string `json:"FileName,omitempty"`

    // container class for doc/dot save options.
    SaveFormat string `json:"SaveFormat,omitempty"`

    // container class for doc/dot save options.
    UpdateFields bool `json:"UpdateFields,omitempty"`

    // container class for doc/dot save options.
    UpdateLastPrintedProperty bool `json:"UpdateLastPrintedProperty,omitempty"`

    // container class for doc/dot save options.
    UpdateLastSavedTimeProperty bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // container class for doc/dot save options.
    UpdateSdtContent bool `json:"UpdateSdtContent,omitempty"`

    // container class for doc/dot save options.
    ZipOutput bool `json:"ZipOutput,omitempty"`

    // container class for doc/dot save options.
    AlwaysCompressMetafiles bool `json:"AlwaysCompressMetafiles,omitempty"`

    // container class for doc/dot save options.
    Password string `json:"Password,omitempty"`

    // container class for doc/dot save options.
    SavePictureBullet bool `json:"SavePictureBullet,omitempty"`

    // container class for doc/dot save options.
    SaveRoutingSlip bool `json:"SaveRoutingSlip,omitempty"`
}

type IDocSaveOptionsData interface {
    IsDocSaveOptionsData() bool
}
func (DocSaveOptionsData) IsDocSaveOptionsData() bool {
    return true
}

func (DocSaveOptionsData) IsSaveOptionsData() bool {
    return true
}
