/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="xaml_flow_pack_save_options_data.go">
 *   Copyright (c) 2026 Aspose.Words for Cloud
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

import (
    "errors"
)

// Container class for xamlflow_pack save options.

type IXamlFlowPackSaveOptionsData interface {
    IsXamlFlowPackSaveOptionsData() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetAllowEmbeddingPostScriptFonts() *bool
    SetAllowEmbeddingPostScriptFonts(value *bool)
    GetCustomTimeZoneInfoData() ITimeZoneInfoData
    SetCustomTimeZoneInfoData(value ITimeZoneInfoData)
    GetDml3DEffectsRenderingMode() *string
    SetDml3DEffectsRenderingMode(value *string)
    GetDmlEffectsRenderingMode() *string
    SetDmlEffectsRenderingMode(value *string)
    GetDmlRenderingMode() *string
    SetDmlRenderingMode(value *string)
    GetFileName() *string
    SetFileName(value *string)
    GetImlRenderingMode() *string
    SetImlRenderingMode(value *string)
    GetUpdateAmbiguousTextFont() *bool
    SetUpdateAmbiguousTextFont(value *bool)
    GetUpdateCreatedTimeProperty() *bool
    SetUpdateCreatedTimeProperty(value *bool)
    GetUpdateFields() *bool
    SetUpdateFields(value *bool)
    GetUpdateLastPrintedProperty() *bool
    SetUpdateLastPrintedProperty(value *bool)
    GetUpdateLastSavedTimeProperty() *bool
    SetUpdateLastSavedTimeProperty(value *bool)
    GetZipOutput() *bool
    SetZipOutput(value *bool)
    GetImagesFolder() *string
    SetImagesFolder(value *string)
    GetImagesFolderAlias() *string
    SetImagesFolderAlias(value *string)
    GetReplaceBackslashWithYenSign() *bool
    SetReplaceBackslashWithYenSign(value *bool)
    GetSaveFormat() *string
    SetSaveFormat(value *string)
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
    ImlRenderingMode *string `json:"ImlRenderingMode,omitempty"`

    // Container class for xamlflow_pack save options.
    UpdateAmbiguousTextFont *bool `json:"UpdateAmbiguousTextFont,omitempty"`

    // Container class for xamlflow_pack save options.
    UpdateCreatedTimeProperty *bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for xamlflow_pack save options.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // Container class for xamlflow_pack save options.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for xamlflow_pack save options.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for xamlflow_pack save options.
    ZipOutput *bool `json:"ZipOutput,omitempty"`

    // Container class for xamlflow_pack save options.
    ImagesFolder *string `json:"ImagesFolder,omitempty"`

    // Container class for xamlflow_pack save options.
    ImagesFolderAlias *string `json:"ImagesFolderAlias,omitempty"`

    // Container class for xamlflow_pack save options.
    ReplaceBackslashWithYenSign *bool `json:"ReplaceBackslashWithYenSign,omitempty"`

    // Container class for xamlflow_pack save options.
    SaveFormat *string `json:"SaveFormat,omitempty"`
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

func (obj *XamlFlowPackSaveOptionsData) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["AllowEmbeddingPostScriptFonts"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.AllowEmbeddingPostScriptFonts = &parsedValue
        }

    } else if jsonValue, exists := json["allowEmbeddingPostScriptFonts"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.AllowEmbeddingPostScriptFonts = &parsedValue
        }

    }

    if jsonValue, exists := json["CustomTimeZoneInfoData"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ITimeZoneInfoData = new(TimeZoneInfoData)
            modelInstance.Deserialize(parsedValue)
            obj.CustomTimeZoneInfoData = modelInstance
        }

    } else if jsonValue, exists := json["customTimeZoneInfoData"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance ITimeZoneInfoData = new(TimeZoneInfoData)
            modelInstance.Deserialize(parsedValue)
            obj.CustomTimeZoneInfoData = modelInstance
        }

    }

    if jsonValue, exists := json["Dml3DEffectsRenderingMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Dml3DEffectsRenderingMode = &parsedValue
        }

    } else if jsonValue, exists := json["dml3DEffectsRenderingMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Dml3DEffectsRenderingMode = &parsedValue
        }

    }

    if jsonValue, exists := json["DmlEffectsRenderingMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DmlEffectsRenderingMode = &parsedValue
        }

    } else if jsonValue, exists := json["dmlEffectsRenderingMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DmlEffectsRenderingMode = &parsedValue
        }

    }

    if jsonValue, exists := json["DmlRenderingMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DmlRenderingMode = &parsedValue
        }

    } else if jsonValue, exists := json["dmlRenderingMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DmlRenderingMode = &parsedValue
        }

    }

    if jsonValue, exists := json["FileName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FileName = &parsedValue
        }

    } else if jsonValue, exists := json["fileName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FileName = &parsedValue
        }

    }

    if jsonValue, exists := json["ImlRenderingMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ImlRenderingMode = &parsedValue
        }

    } else if jsonValue, exists := json["imlRenderingMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ImlRenderingMode = &parsedValue
        }

    }

    if jsonValue, exists := json["UpdateAmbiguousTextFont"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UpdateAmbiguousTextFont = &parsedValue
        }

    } else if jsonValue, exists := json["updateAmbiguousTextFont"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UpdateAmbiguousTextFont = &parsedValue
        }

    }

    if jsonValue, exists := json["UpdateCreatedTimeProperty"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UpdateCreatedTimeProperty = &parsedValue
        }

    } else if jsonValue, exists := json["updateCreatedTimeProperty"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UpdateCreatedTimeProperty = &parsedValue
        }

    }

    if jsonValue, exists := json["UpdateFields"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UpdateFields = &parsedValue
        }

    } else if jsonValue, exists := json["updateFields"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UpdateFields = &parsedValue
        }

    }

    if jsonValue, exists := json["UpdateLastPrintedProperty"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UpdateLastPrintedProperty = &parsedValue
        }

    } else if jsonValue, exists := json["updateLastPrintedProperty"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UpdateLastPrintedProperty = &parsedValue
        }

    }

    if jsonValue, exists := json["UpdateLastSavedTimeProperty"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UpdateLastSavedTimeProperty = &parsedValue
        }

    } else if jsonValue, exists := json["updateLastSavedTimeProperty"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UpdateLastSavedTimeProperty = &parsedValue
        }

    }

    if jsonValue, exists := json["ZipOutput"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ZipOutput = &parsedValue
        }

    } else if jsonValue, exists := json["zipOutput"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ZipOutput = &parsedValue
        }

    }

    if jsonValue, exists := json["ImagesFolder"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ImagesFolder = &parsedValue
        }

    } else if jsonValue, exists := json["imagesFolder"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ImagesFolder = &parsedValue
        }

    }

    if jsonValue, exists := json["ImagesFolderAlias"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ImagesFolderAlias = &parsedValue
        }

    } else if jsonValue, exists := json["imagesFolderAlias"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ImagesFolderAlias = &parsedValue
        }

    }

    if jsonValue, exists := json["ReplaceBackslashWithYenSign"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ReplaceBackslashWithYenSign = &parsedValue
        }

    } else if jsonValue, exists := json["replaceBackslashWithYenSign"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ReplaceBackslashWithYenSign = &parsedValue
        }

    }
}

func (obj *XamlFlowPackSaveOptionsData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *XamlFlowPackSaveOptionsData) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.FileName == nil {
        return errors.New("Property FileName in XamlFlowPackSaveOptionsData is required.")
    }
    if obj.CustomTimeZoneInfoData != nil {
        if err := obj.CustomTimeZoneInfoData.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *XamlFlowPackSaveOptionsData) GetAllowEmbeddingPostScriptFonts() *bool {
    return obj.AllowEmbeddingPostScriptFonts
}

func (obj *XamlFlowPackSaveOptionsData) SetAllowEmbeddingPostScriptFonts(value *bool) {
    obj.AllowEmbeddingPostScriptFonts = value
}

func (obj *XamlFlowPackSaveOptionsData) GetCustomTimeZoneInfoData() ITimeZoneInfoData {
    return obj.CustomTimeZoneInfoData
}

func (obj *XamlFlowPackSaveOptionsData) SetCustomTimeZoneInfoData(value ITimeZoneInfoData) {
    obj.CustomTimeZoneInfoData = value
}

func (obj *XamlFlowPackSaveOptionsData) GetDml3DEffectsRenderingMode() *string {
    return obj.Dml3DEffectsRenderingMode
}

func (obj *XamlFlowPackSaveOptionsData) SetDml3DEffectsRenderingMode(value *string) {
    obj.Dml3DEffectsRenderingMode = value
}

func (obj *XamlFlowPackSaveOptionsData) GetDmlEffectsRenderingMode() *string {
    return obj.DmlEffectsRenderingMode
}

func (obj *XamlFlowPackSaveOptionsData) SetDmlEffectsRenderingMode(value *string) {
    obj.DmlEffectsRenderingMode = value
}

func (obj *XamlFlowPackSaveOptionsData) GetDmlRenderingMode() *string {
    return obj.DmlRenderingMode
}

func (obj *XamlFlowPackSaveOptionsData) SetDmlRenderingMode(value *string) {
    obj.DmlRenderingMode = value
}

func (obj *XamlFlowPackSaveOptionsData) GetFileName() *string {
    return obj.FileName
}

func (obj *XamlFlowPackSaveOptionsData) SetFileName(value *string) {
    obj.FileName = value
}

func (obj *XamlFlowPackSaveOptionsData) GetImlRenderingMode() *string {
    return obj.ImlRenderingMode
}

func (obj *XamlFlowPackSaveOptionsData) SetImlRenderingMode(value *string) {
    obj.ImlRenderingMode = value
}

func (obj *XamlFlowPackSaveOptionsData) GetUpdateAmbiguousTextFont() *bool {
    return obj.UpdateAmbiguousTextFont
}

func (obj *XamlFlowPackSaveOptionsData) SetUpdateAmbiguousTextFont(value *bool) {
    obj.UpdateAmbiguousTextFont = value
}

func (obj *XamlFlowPackSaveOptionsData) GetUpdateCreatedTimeProperty() *bool {
    return obj.UpdateCreatedTimeProperty
}

func (obj *XamlFlowPackSaveOptionsData) SetUpdateCreatedTimeProperty(value *bool) {
    obj.UpdateCreatedTimeProperty = value
}

func (obj *XamlFlowPackSaveOptionsData) GetUpdateFields() *bool {
    return obj.UpdateFields
}

func (obj *XamlFlowPackSaveOptionsData) SetUpdateFields(value *bool) {
    obj.UpdateFields = value
}

func (obj *XamlFlowPackSaveOptionsData) GetUpdateLastPrintedProperty() *bool {
    return obj.UpdateLastPrintedProperty
}

func (obj *XamlFlowPackSaveOptionsData) SetUpdateLastPrintedProperty(value *bool) {
    obj.UpdateLastPrintedProperty = value
}

func (obj *XamlFlowPackSaveOptionsData) GetUpdateLastSavedTimeProperty() *bool {
    return obj.UpdateLastSavedTimeProperty
}

func (obj *XamlFlowPackSaveOptionsData) SetUpdateLastSavedTimeProperty(value *bool) {
    obj.UpdateLastSavedTimeProperty = value
}

func (obj *XamlFlowPackSaveOptionsData) GetZipOutput() *bool {
    return obj.ZipOutput
}

func (obj *XamlFlowPackSaveOptionsData) SetZipOutput(value *bool) {
    obj.ZipOutput = value
}

func (obj *XamlFlowPackSaveOptionsData) GetImagesFolder() *string {
    return obj.ImagesFolder
}

func (obj *XamlFlowPackSaveOptionsData) SetImagesFolder(value *string) {
    obj.ImagesFolder = value
}

func (obj *XamlFlowPackSaveOptionsData) GetImagesFolderAlias() *string {
    return obj.ImagesFolderAlias
}

func (obj *XamlFlowPackSaveOptionsData) SetImagesFolderAlias(value *string) {
    obj.ImagesFolderAlias = value
}

func (obj *XamlFlowPackSaveOptionsData) GetReplaceBackslashWithYenSign() *bool {
    return obj.ReplaceBackslashWithYenSign
}

func (obj *XamlFlowPackSaveOptionsData) SetReplaceBackslashWithYenSign(value *bool) {
    obj.ReplaceBackslashWithYenSign = value
}

func (obj *XamlFlowPackSaveOptionsData) GetSaveFormat() *string {
    return obj.SaveFormat
}

func (obj *XamlFlowPackSaveOptionsData) SetSaveFormat(value *string) {
    obj.SaveFormat = value
}

