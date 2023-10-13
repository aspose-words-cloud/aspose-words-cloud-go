/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="xaml_flow_save_options_data.go">
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

import (
    "errors"
)

// Container class for xaml flow save options.

type IXamlFlowSaveOptionsData interface {
    IsXamlFlowSaveOptionsData() bool
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
    GetSaveFormat() *string
    SetSaveFormat(value *string)
}

type XamlFlowSaveOptionsData struct {
    // Container class for xaml flow save options.
    AllowEmbeddingPostScriptFonts *bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for xaml flow save options.
    CustomTimeZoneInfoData ITimeZoneInfoData `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for xaml flow save options.
    Dml3DEffectsRenderingMode *string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for xaml flow save options.
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for xaml flow save options.
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // Container class for xaml flow save options.
    FileName *string `json:"FileName,omitempty"`

    // Container class for xaml flow save options.
    ImlRenderingMode *string `json:"ImlRenderingMode,omitempty"`

    // Container class for xaml flow save options.
    UpdateCreatedTimeProperty *bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for xaml flow save options.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // Container class for xaml flow save options.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for xaml flow save options.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for xaml flow save options.
    ZipOutput *bool `json:"ZipOutput,omitempty"`

    // Container class for xaml flow save options.
    ImagesFolder *string `json:"ImagesFolder,omitempty"`

    // Container class for xaml flow save options.
    ImagesFolderAlias *string `json:"ImagesFolderAlias,omitempty"`

    // Container class for xaml flow save options.
    SaveFormat *string `json:"SaveFormat,omitempty"`
}

func (XamlFlowSaveOptionsData) IsXamlFlowSaveOptionsData() bool {
    return true
}

func (XamlFlowSaveOptionsData) IsSaveOptionsData() bool {
    return true
}

func (obj *XamlFlowSaveOptionsData) Initialize() {
    var _SaveFormat = "xamlflow"
    obj.SaveFormat = &_SaveFormat


    if (obj.CustomTimeZoneInfoData != nil) {
        obj.CustomTimeZoneInfoData.Initialize()
    }


}

func (obj *XamlFlowSaveOptionsData) Deserialize(json map[string]interface{}) {
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
}

func (obj *XamlFlowSaveOptionsData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *XamlFlowSaveOptionsData) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.FileName == nil {
        return errors.New("Property FileName in XamlFlowSaveOptionsData is required.")
    }

    return nil;
}

func (obj *XamlFlowSaveOptionsData) GetAllowEmbeddingPostScriptFonts() *bool {
    return obj.AllowEmbeddingPostScriptFonts
}

func (obj *XamlFlowSaveOptionsData) SetAllowEmbeddingPostScriptFonts(value *bool) {
    obj.AllowEmbeddingPostScriptFonts = value
}

func (obj *XamlFlowSaveOptionsData) GetCustomTimeZoneInfoData() ITimeZoneInfoData {
    return obj.CustomTimeZoneInfoData
}

func (obj *XamlFlowSaveOptionsData) SetCustomTimeZoneInfoData(value ITimeZoneInfoData) {
    obj.CustomTimeZoneInfoData = value
}

func (obj *XamlFlowSaveOptionsData) GetDml3DEffectsRenderingMode() *string {
    return obj.Dml3DEffectsRenderingMode
}

func (obj *XamlFlowSaveOptionsData) SetDml3DEffectsRenderingMode(value *string) {
    obj.Dml3DEffectsRenderingMode = value
}

func (obj *XamlFlowSaveOptionsData) GetDmlEffectsRenderingMode() *string {
    return obj.DmlEffectsRenderingMode
}

func (obj *XamlFlowSaveOptionsData) SetDmlEffectsRenderingMode(value *string) {
    obj.DmlEffectsRenderingMode = value
}

func (obj *XamlFlowSaveOptionsData) GetDmlRenderingMode() *string {
    return obj.DmlRenderingMode
}

func (obj *XamlFlowSaveOptionsData) SetDmlRenderingMode(value *string) {
    obj.DmlRenderingMode = value
}

func (obj *XamlFlowSaveOptionsData) GetFileName() *string {
    return obj.FileName
}

func (obj *XamlFlowSaveOptionsData) SetFileName(value *string) {
    obj.FileName = value
}

func (obj *XamlFlowSaveOptionsData) GetImlRenderingMode() *string {
    return obj.ImlRenderingMode
}

func (obj *XamlFlowSaveOptionsData) SetImlRenderingMode(value *string) {
    obj.ImlRenderingMode = value
}

func (obj *XamlFlowSaveOptionsData) GetUpdateCreatedTimeProperty() *bool {
    return obj.UpdateCreatedTimeProperty
}

func (obj *XamlFlowSaveOptionsData) SetUpdateCreatedTimeProperty(value *bool) {
    obj.UpdateCreatedTimeProperty = value
}

func (obj *XamlFlowSaveOptionsData) GetUpdateFields() *bool {
    return obj.UpdateFields
}

func (obj *XamlFlowSaveOptionsData) SetUpdateFields(value *bool) {
    obj.UpdateFields = value
}

func (obj *XamlFlowSaveOptionsData) GetUpdateLastPrintedProperty() *bool {
    return obj.UpdateLastPrintedProperty
}

func (obj *XamlFlowSaveOptionsData) SetUpdateLastPrintedProperty(value *bool) {
    obj.UpdateLastPrintedProperty = value
}

func (obj *XamlFlowSaveOptionsData) GetUpdateLastSavedTimeProperty() *bool {
    return obj.UpdateLastSavedTimeProperty
}

func (obj *XamlFlowSaveOptionsData) SetUpdateLastSavedTimeProperty(value *bool) {
    obj.UpdateLastSavedTimeProperty = value
}

func (obj *XamlFlowSaveOptionsData) GetZipOutput() *bool {
    return obj.ZipOutput
}

func (obj *XamlFlowSaveOptionsData) SetZipOutput(value *bool) {
    obj.ZipOutput = value
}

func (obj *XamlFlowSaveOptionsData) GetImagesFolder() *string {
    return obj.ImagesFolder
}

func (obj *XamlFlowSaveOptionsData) SetImagesFolder(value *string) {
    obj.ImagesFolder = value
}

func (obj *XamlFlowSaveOptionsData) GetImagesFolderAlias() *string {
    return obj.ImagesFolderAlias
}

func (obj *XamlFlowSaveOptionsData) SetImagesFolderAlias(value *string) {
    obj.ImagesFolderAlias = value
}

func (obj *XamlFlowSaveOptionsData) GetSaveFormat() *string {
    return obj.SaveFormat
}

func (obj *XamlFlowSaveOptionsData) SetSaveFormat(value *string) {
    obj.SaveFormat = value
}

