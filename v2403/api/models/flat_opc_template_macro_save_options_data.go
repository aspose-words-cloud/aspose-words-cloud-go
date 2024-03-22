/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="flat_opc_template_macro_save_options_data.go">
 *   Copyright (c) 2024 Aspose.Words for Cloud
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

// Container class for fopc_template_macro save options.

type IFlatOpcTemplateMacroSaveOptionsData interface {
    IsFlatOpcTemplateMacroSaveOptionsData() bool
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

type FlatOpcTemplateMacroSaveOptionsData struct {
    // Container class for fopc_template_macro save options.
    AllowEmbeddingPostScriptFonts *bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for fopc_template_macro save options.
    CustomTimeZoneInfoData ITimeZoneInfoData `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for fopc_template_macro save options.
    Dml3DEffectsRenderingMode *string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for fopc_template_macro save options.
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for fopc_template_macro save options.
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // Container class for fopc_template_macro save options.
    FileName *string `json:"FileName,omitempty"`

    // Container class for fopc_template_macro save options.
    ImlRenderingMode *string `json:"ImlRenderingMode,omitempty"`

    // Container class for fopc_template_macro save options.
    UpdateCreatedTimeProperty *bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for fopc_template_macro save options.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // Container class for fopc_template_macro save options.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for fopc_template_macro save options.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for fopc_template_macro save options.
    ZipOutput *bool `json:"ZipOutput,omitempty"`

    // Container class for fopc_template_macro save options.
    Compliance *string `json:"Compliance,omitempty"`

    // Container class for fopc_template_macro save options.
    CompressionLevel *string `json:"CompressionLevel,omitempty"`

    // Container class for fopc_template_macro save options.
    Password *string `json:"Password,omitempty"`

    // Container class for fopc_template_macro save options.
    PrettyFormat *bool `json:"PrettyFormat,omitempty"`

    // Container class for fopc_template_macro save options.
    SaveFormat *string `json:"SaveFormat,omitempty"`
}

func (FlatOpcTemplateMacroSaveOptionsData) IsFlatOpcTemplateMacroSaveOptionsData() bool {
    return true
}

func (FlatOpcTemplateMacroSaveOptionsData) IsOoxmlSaveOptionsData() bool {
    return true
}

func (FlatOpcTemplateMacroSaveOptionsData) IsSaveOptionsData() bool {
    return true
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) Initialize() {
    var _SaveFormat = "fopc_template_macro"
    obj.SaveFormat = &_SaveFormat


    if (obj.CustomTimeZoneInfoData != nil) {
        obj.CustomTimeZoneInfoData.Initialize()
    }


}

func (obj *FlatOpcTemplateMacroSaveOptionsData) Deserialize(json map[string]interface{}) {
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

func (obj *FlatOpcTemplateMacroSaveOptionsData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.FileName == nil {
        return errors.New("Property FileName in FlatOpcTemplateMacroSaveOptionsData is required.")
    }
    if obj.CustomTimeZoneInfoData != nil {
        if err := obj.CustomTimeZoneInfoData.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) GetAllowEmbeddingPostScriptFonts() *bool {
    return obj.AllowEmbeddingPostScriptFonts
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) SetAllowEmbeddingPostScriptFonts(value *bool) {
    obj.AllowEmbeddingPostScriptFonts = value
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) GetCustomTimeZoneInfoData() ITimeZoneInfoData {
    return obj.CustomTimeZoneInfoData
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) SetCustomTimeZoneInfoData(value ITimeZoneInfoData) {
    obj.CustomTimeZoneInfoData = value
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) GetDml3DEffectsRenderingMode() *string {
    return obj.Dml3DEffectsRenderingMode
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) SetDml3DEffectsRenderingMode(value *string) {
    obj.Dml3DEffectsRenderingMode = value
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) GetDmlEffectsRenderingMode() *string {
    return obj.DmlEffectsRenderingMode
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) SetDmlEffectsRenderingMode(value *string) {
    obj.DmlEffectsRenderingMode = value
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) GetDmlRenderingMode() *string {
    return obj.DmlRenderingMode
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) SetDmlRenderingMode(value *string) {
    obj.DmlRenderingMode = value
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) GetFileName() *string {
    return obj.FileName
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) SetFileName(value *string) {
    obj.FileName = value
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) GetImlRenderingMode() *string {
    return obj.ImlRenderingMode
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) SetImlRenderingMode(value *string) {
    obj.ImlRenderingMode = value
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) GetUpdateCreatedTimeProperty() *bool {
    return obj.UpdateCreatedTimeProperty
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) SetUpdateCreatedTimeProperty(value *bool) {
    obj.UpdateCreatedTimeProperty = value
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) GetUpdateFields() *bool {
    return obj.UpdateFields
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) SetUpdateFields(value *bool) {
    obj.UpdateFields = value
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) GetUpdateLastPrintedProperty() *bool {
    return obj.UpdateLastPrintedProperty
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) SetUpdateLastPrintedProperty(value *bool) {
    obj.UpdateLastPrintedProperty = value
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) GetUpdateLastSavedTimeProperty() *bool {
    return obj.UpdateLastSavedTimeProperty
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) SetUpdateLastSavedTimeProperty(value *bool) {
    obj.UpdateLastSavedTimeProperty = value
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) GetZipOutput() *bool {
    return obj.ZipOutput
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) SetZipOutput(value *bool) {
    obj.ZipOutput = value
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) GetCompliance() *string {
    return obj.Compliance
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) SetCompliance(value *string) {
    obj.Compliance = value
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) GetCompressionLevel() *string {
    return obj.CompressionLevel
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) SetCompressionLevel(value *string) {
    obj.CompressionLevel = value
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) GetPassword() *string {
    return obj.Password
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) SetPassword(value *string) {
    obj.Password = value
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) GetPrettyFormat() *bool {
    return obj.PrettyFormat
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) SetPrettyFormat(value *bool) {
    obj.PrettyFormat = value
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) GetSaveFormat() *string {
    return obj.SaveFormat
}

func (obj *FlatOpcTemplateMacroSaveOptionsData) SetSaveFormat(value *string) {
    obj.SaveFormat = value
}

