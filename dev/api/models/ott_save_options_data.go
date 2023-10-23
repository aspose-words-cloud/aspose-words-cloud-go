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

import (
    "errors"
)

// Container class for ott save options.

type IOttSaveOptionsData interface {
    IsOttSaveOptionsData() bool
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
    GetIsStrictSchema11() *bool
    SetIsStrictSchema11(value *bool)
    GetMeasureUnit() *string
    SetMeasureUnit(value *string)
    GetPassword() *string
    SetPassword(value *string)
    GetPrettyFormat() *bool
    SetPrettyFormat(value *bool)
    GetSaveFormat() *string
    SetSaveFormat(value *string)
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

func (obj *OttSaveOptionsData) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["IsStrictSchema11"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsStrictSchema11 = &parsedValue
        }

    } else if jsonValue, exists := json["isStrictSchema11"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsStrictSchema11 = &parsedValue
        }

    }

    if jsonValue, exists := json["MeasureUnit"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.MeasureUnit = &parsedValue
        }

    } else if jsonValue, exists := json["measureUnit"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.MeasureUnit = &parsedValue
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

func (obj *OttSaveOptionsData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *OttSaveOptionsData) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.FileName == nil {
        return errors.New("Property FileName in OttSaveOptionsData is required.")
    }
    if obj.AllowEmbeddingPostScriptFonts == nil {
        return errors.New("Property AllowEmbeddingPostScriptFonts in OttSaveOptionsData is required.")
    }

    if obj.CustomTimeZoneInfoData == nil {
        return errors.New("Property CustomTimeZoneInfoData in OttSaveOptionsData is required.")
    }

    if obj.CustomTimeZoneInfoData != nil {
        if err := obj.CustomTimeZoneInfoData.Validate(); err != nil {
            return err
        }
    }

    if obj.Dml3DEffectsRenderingMode == nil {
        return errors.New("Property Dml3DEffectsRenderingMode in OttSaveOptionsData is required.")
    }

    if obj.DmlEffectsRenderingMode == nil {
        return errors.New("Property DmlEffectsRenderingMode in OttSaveOptionsData is required.")
    }

    if obj.DmlRenderingMode == nil {
        return errors.New("Property DmlRenderingMode in OttSaveOptionsData is required.")
    }

    if obj.FileName == nil {
        return errors.New("Property FileName in OttSaveOptionsData is required.")
    }

    if obj.ImlRenderingMode == nil {
        return errors.New("Property ImlRenderingMode in OttSaveOptionsData is required.")
    }

    if obj.UpdateCreatedTimeProperty == nil {
        return errors.New("Property UpdateCreatedTimeProperty in OttSaveOptionsData is required.")
    }

    if obj.UpdateFields == nil {
        return errors.New("Property UpdateFields in OttSaveOptionsData is required.")
    }

    if obj.UpdateLastPrintedProperty == nil {
        return errors.New("Property UpdateLastPrintedProperty in OttSaveOptionsData is required.")
    }

    if obj.UpdateLastSavedTimeProperty == nil {
        return errors.New("Property UpdateLastSavedTimeProperty in OttSaveOptionsData is required.")
    }

    if obj.ZipOutput == nil {
        return errors.New("Property ZipOutput in OttSaveOptionsData is required.")
    }

    if obj.IsStrictSchema11 == nil {
        return errors.New("Property IsStrictSchema11 in OttSaveOptionsData is required.")
    }

    if obj.MeasureUnit == nil {
        return errors.New("Property MeasureUnit in OttSaveOptionsData is required.")
    }

    if obj.Password == nil {
        return errors.New("Property Password in OttSaveOptionsData is required.")
    }

    if obj.PrettyFormat == nil {
        return errors.New("Property PrettyFormat in OttSaveOptionsData is required.")
    }

    if obj.SaveFormat == nil {
        return errors.New("Property SaveFormat in OttSaveOptionsData is required.")
    }

    return nil;
}

func (obj *OttSaveOptionsData) GetAllowEmbeddingPostScriptFonts() *bool {
    return obj.AllowEmbeddingPostScriptFonts
}

func (obj *OttSaveOptionsData) SetAllowEmbeddingPostScriptFonts(value *bool) {
    obj.AllowEmbeddingPostScriptFonts = value
}

func (obj *OttSaveOptionsData) GetCustomTimeZoneInfoData() ITimeZoneInfoData {
    return obj.CustomTimeZoneInfoData
}

func (obj *OttSaveOptionsData) SetCustomTimeZoneInfoData(value ITimeZoneInfoData) {
    obj.CustomTimeZoneInfoData = value
}

func (obj *OttSaveOptionsData) GetDml3DEffectsRenderingMode() *string {
    return obj.Dml3DEffectsRenderingMode
}

func (obj *OttSaveOptionsData) SetDml3DEffectsRenderingMode(value *string) {
    obj.Dml3DEffectsRenderingMode = value
}

func (obj *OttSaveOptionsData) GetDmlEffectsRenderingMode() *string {
    return obj.DmlEffectsRenderingMode
}

func (obj *OttSaveOptionsData) SetDmlEffectsRenderingMode(value *string) {
    obj.DmlEffectsRenderingMode = value
}

func (obj *OttSaveOptionsData) GetDmlRenderingMode() *string {
    return obj.DmlRenderingMode
}

func (obj *OttSaveOptionsData) SetDmlRenderingMode(value *string) {
    obj.DmlRenderingMode = value
}

func (obj *OttSaveOptionsData) GetFileName() *string {
    return obj.FileName
}

func (obj *OttSaveOptionsData) SetFileName(value *string) {
    obj.FileName = value
}

func (obj *OttSaveOptionsData) GetImlRenderingMode() *string {
    return obj.ImlRenderingMode
}

func (obj *OttSaveOptionsData) SetImlRenderingMode(value *string) {
    obj.ImlRenderingMode = value
}

func (obj *OttSaveOptionsData) GetUpdateCreatedTimeProperty() *bool {
    return obj.UpdateCreatedTimeProperty
}

func (obj *OttSaveOptionsData) SetUpdateCreatedTimeProperty(value *bool) {
    obj.UpdateCreatedTimeProperty = value
}

func (obj *OttSaveOptionsData) GetUpdateFields() *bool {
    return obj.UpdateFields
}

func (obj *OttSaveOptionsData) SetUpdateFields(value *bool) {
    obj.UpdateFields = value
}

func (obj *OttSaveOptionsData) GetUpdateLastPrintedProperty() *bool {
    return obj.UpdateLastPrintedProperty
}

func (obj *OttSaveOptionsData) SetUpdateLastPrintedProperty(value *bool) {
    obj.UpdateLastPrintedProperty = value
}

func (obj *OttSaveOptionsData) GetUpdateLastSavedTimeProperty() *bool {
    return obj.UpdateLastSavedTimeProperty
}

func (obj *OttSaveOptionsData) SetUpdateLastSavedTimeProperty(value *bool) {
    obj.UpdateLastSavedTimeProperty = value
}

func (obj *OttSaveOptionsData) GetZipOutput() *bool {
    return obj.ZipOutput
}

func (obj *OttSaveOptionsData) SetZipOutput(value *bool) {
    obj.ZipOutput = value
}

func (obj *OttSaveOptionsData) GetIsStrictSchema11() *bool {
    return obj.IsStrictSchema11
}

func (obj *OttSaveOptionsData) SetIsStrictSchema11(value *bool) {
    obj.IsStrictSchema11 = value
}

func (obj *OttSaveOptionsData) GetMeasureUnit() *string {
    return obj.MeasureUnit
}

func (obj *OttSaveOptionsData) SetMeasureUnit(value *string) {
    obj.MeasureUnit = value
}

func (obj *OttSaveOptionsData) GetPassword() *string {
    return obj.Password
}

func (obj *OttSaveOptionsData) SetPassword(value *string) {
    obj.Password = value
}

func (obj *OttSaveOptionsData) GetPrettyFormat() *bool {
    return obj.PrettyFormat
}

func (obj *OttSaveOptionsData) SetPrettyFormat(value *bool) {
    obj.PrettyFormat = value
}

func (obj *OttSaveOptionsData) GetSaveFormat() *string {
    return obj.SaveFormat
}

func (obj *OttSaveOptionsData) SetSaveFormat(value *string) {
    obj.SaveFormat = value
}

