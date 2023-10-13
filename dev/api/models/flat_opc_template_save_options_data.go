/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="flat_opc_template_save_options_data.go">
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

// Container class for fopc_template save options.

type IFlatOpcTemplateSaveOptionsData interface {
    IsFlatOpcTemplateSaveOptionsData() bool
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
    GetCompliance() *string
    SetCompliance(value *string)
    GetCompressionLevel() *string
    SetCompressionLevel(value *string)
    GetPassword() *string
    SetPassword(value *string)
    GetPrettyFormat() *bool
    SetPrettyFormat(value *bool)
    GetSaveFormat() *string
    SetSaveFormat(value *string)
}

type FlatOpcTemplateSaveOptionsData struct {
    // Container class for fopc_template save options.
    AllowEmbeddingPostScriptFonts *bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for fopc_template save options.
    CustomTimeZoneInfoData ITimeZoneInfoData `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for fopc_template save options.
    Dml3DEffectsRenderingMode *string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for fopc_template save options.
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for fopc_template save options.
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // Container class for fopc_template save options.
    FileName *string `json:"FileName,omitempty"`

    // Container class for fopc_template save options.
    ImlRenderingMode *string `json:"ImlRenderingMode,omitempty"`

    // Container class for fopc_template save options.
    UpdateCreatedTimeProperty *bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for fopc_template save options.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // Container class for fopc_template save options.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for fopc_template save options.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for fopc_template save options.
    ZipOutput *bool `json:"ZipOutput,omitempty"`

    // Container class for fopc_template save options.
    Compliance *string `json:"Compliance,omitempty"`

    // Container class for fopc_template save options.
    CompressionLevel *string `json:"CompressionLevel,omitempty"`

    // Container class for fopc_template save options.
    Password *string `json:"Password,omitempty"`

    // Container class for fopc_template save options.
    PrettyFormat *bool `json:"PrettyFormat,omitempty"`

    // Container class for fopc_template save options.
    SaveFormat *string `json:"SaveFormat,omitempty"`
}

func (FlatOpcTemplateSaveOptionsData) IsFlatOpcTemplateSaveOptionsData() bool {
    return true
}

func (FlatOpcTemplateSaveOptionsData) IsOoxmlSaveOptionsData() bool {
    return true
}

func (FlatOpcTemplateSaveOptionsData) IsSaveOptionsData() bool {
    return true
}

func (obj *FlatOpcTemplateSaveOptionsData) Initialize() {
    var _SaveFormat = "fopc_template"
    obj.SaveFormat = &_SaveFormat


    if (obj.CustomTimeZoneInfoData != nil) {
        obj.CustomTimeZoneInfoData.Initialize()
    }


}

func (obj *FlatOpcTemplateSaveOptionsData) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["Compliance"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Compliance = &parsedValue
        }

    } else if jsonValue, exists := json["compliance"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Compliance = &parsedValue
        }

    }

    if jsonValue, exists := json["CompressionLevel"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.CompressionLevel = &parsedValue
        }

    } else if jsonValue, exists := json["compressionLevel"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.CompressionLevel = &parsedValue
        }

    }

    if jsonValue, exists := json["Password"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Password = &parsedValue
        }

    } else if jsonValue, exists := json["password"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Password = &parsedValue
        }

    }

    if jsonValue, exists := json["PrettyFormat"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.PrettyFormat = &parsedValue
        }

    } else if jsonValue, exists := json["prettyFormat"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.PrettyFormat = &parsedValue
        }

    }
}

func (obj *FlatOpcTemplateSaveOptionsData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *FlatOpcTemplateSaveOptionsData) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.FileName == nil {
        return errors.New("Property FileName in FlatOpcTemplateSaveOptionsData is required.")
    }

    return nil;
}

func (obj *FlatOpcTemplateSaveOptionsData) GetAllowEmbeddingPostScriptFonts() *bool {
    return obj.AllowEmbeddingPostScriptFonts
}

func (obj *FlatOpcTemplateSaveOptionsData) SetAllowEmbeddingPostScriptFonts(value *bool) {
    obj.AllowEmbeddingPostScriptFonts = value
}

func (obj *FlatOpcTemplateSaveOptionsData) GetCustomTimeZoneInfoData() ITimeZoneInfoData {
    return obj.CustomTimeZoneInfoData
}

func (obj *FlatOpcTemplateSaveOptionsData) SetCustomTimeZoneInfoData(value ITimeZoneInfoData) {
    obj.CustomTimeZoneInfoData = value
}

func (obj *FlatOpcTemplateSaveOptionsData) GetDml3DEffectsRenderingMode() *string {
    return obj.Dml3DEffectsRenderingMode
}

func (obj *FlatOpcTemplateSaveOptionsData) SetDml3DEffectsRenderingMode(value *string) {
    obj.Dml3DEffectsRenderingMode = value
}

func (obj *FlatOpcTemplateSaveOptionsData) GetDmlEffectsRenderingMode() *string {
    return obj.DmlEffectsRenderingMode
}

func (obj *FlatOpcTemplateSaveOptionsData) SetDmlEffectsRenderingMode(value *string) {
    obj.DmlEffectsRenderingMode = value
}

func (obj *FlatOpcTemplateSaveOptionsData) GetDmlRenderingMode() *string {
    return obj.DmlRenderingMode
}

func (obj *FlatOpcTemplateSaveOptionsData) SetDmlRenderingMode(value *string) {
    obj.DmlRenderingMode = value
}

func (obj *FlatOpcTemplateSaveOptionsData) GetFileName() *string {
    return obj.FileName
}

func (obj *FlatOpcTemplateSaveOptionsData) SetFileName(value *string) {
    obj.FileName = value
}

func (obj *FlatOpcTemplateSaveOptionsData) GetImlRenderingMode() *string {
    return obj.ImlRenderingMode
}

func (obj *FlatOpcTemplateSaveOptionsData) SetImlRenderingMode(value *string) {
    obj.ImlRenderingMode = value
}

func (obj *FlatOpcTemplateSaveOptionsData) GetUpdateCreatedTimeProperty() *bool {
    return obj.UpdateCreatedTimeProperty
}

func (obj *FlatOpcTemplateSaveOptionsData) SetUpdateCreatedTimeProperty(value *bool) {
    obj.UpdateCreatedTimeProperty = value
}

func (obj *FlatOpcTemplateSaveOptionsData) GetUpdateFields() *bool {
    return obj.UpdateFields
}

func (obj *FlatOpcTemplateSaveOptionsData) SetUpdateFields(value *bool) {
    obj.UpdateFields = value
}

func (obj *FlatOpcTemplateSaveOptionsData) GetUpdateLastPrintedProperty() *bool {
    return obj.UpdateLastPrintedProperty
}

func (obj *FlatOpcTemplateSaveOptionsData) SetUpdateLastPrintedProperty(value *bool) {
    obj.UpdateLastPrintedProperty = value
}

func (obj *FlatOpcTemplateSaveOptionsData) GetUpdateLastSavedTimeProperty() *bool {
    return obj.UpdateLastSavedTimeProperty
}

func (obj *FlatOpcTemplateSaveOptionsData) SetUpdateLastSavedTimeProperty(value *bool) {
    obj.UpdateLastSavedTimeProperty = value
}

func (obj *FlatOpcTemplateSaveOptionsData) GetZipOutput() *bool {
    return obj.ZipOutput
}

func (obj *FlatOpcTemplateSaveOptionsData) SetZipOutput(value *bool) {
    obj.ZipOutput = value
}

func (obj *FlatOpcTemplateSaveOptionsData) GetCompliance() *string {
    return obj.Compliance
}

func (obj *FlatOpcTemplateSaveOptionsData) SetCompliance(value *string) {
    obj.Compliance = value
}

func (obj *FlatOpcTemplateSaveOptionsData) GetCompressionLevel() *string {
    return obj.CompressionLevel
}

func (obj *FlatOpcTemplateSaveOptionsData) SetCompressionLevel(value *string) {
    obj.CompressionLevel = value
}

func (obj *FlatOpcTemplateSaveOptionsData) GetPassword() *string {
    return obj.Password
}

func (obj *FlatOpcTemplateSaveOptionsData) SetPassword(value *string) {
    obj.Password = value
}

func (obj *FlatOpcTemplateSaveOptionsData) GetPrettyFormat() *bool {
    return obj.PrettyFormat
}

func (obj *FlatOpcTemplateSaveOptionsData) SetPrettyFormat(value *bool) {
    obj.PrettyFormat = value
}

func (obj *FlatOpcTemplateSaveOptionsData) GetSaveFormat() *string {
    return obj.SaveFormat
}

func (obj *FlatOpcTemplateSaveOptionsData) SetSaveFormat(value *string) {
    obj.SaveFormat = value
}

