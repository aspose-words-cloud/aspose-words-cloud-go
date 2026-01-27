/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="rtf_save_options_data.go">
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

// Container class for rtf save options.

type IRtfSaveOptionsData interface {
    IsRtfSaveOptionsData() bool
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
    GetExportCompactSize() *bool
    SetExportCompactSize(value *bool)
    GetExportImagesForOldReaders() *bool
    SetExportImagesForOldReaders(value *bool)
    GetPrettyFormat() *bool
    SetPrettyFormat(value *bool)
    GetSaveImagesAsWmf() *bool
    SetSaveImagesAsWmf(value *bool)
    GetSaveFormat() *string
    SetSaveFormat(value *string)
}

type RtfSaveOptionsData struct {
    // Container class for rtf save options.
    AllowEmbeddingPostScriptFonts *bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for rtf save options.
    CustomTimeZoneInfoData ITimeZoneInfoData `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for rtf save options.
    Dml3DEffectsRenderingMode *string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for rtf save options.
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for rtf save options.
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // Container class for rtf save options.
    FileName *string `json:"FileName,omitempty"`

    // Container class for rtf save options.
    ImlRenderingMode *string `json:"ImlRenderingMode,omitempty"`

    // Container class for rtf save options.
    UpdateAmbiguousTextFont *bool `json:"UpdateAmbiguousTextFont,omitempty"`

    // Container class for rtf save options.
    UpdateCreatedTimeProperty *bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for rtf save options.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // Container class for rtf save options.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for rtf save options.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for rtf save options.
    ZipOutput *bool `json:"ZipOutput,omitempty"`

    // Container class for rtf save options.
    ExportCompactSize *bool `json:"ExportCompactSize,omitempty"`

    // Container class for rtf save options.
    ExportImagesForOldReaders *bool `json:"ExportImagesForOldReaders,omitempty"`

    // Container class for rtf save options.
    PrettyFormat *bool `json:"PrettyFormat,omitempty"`

    // Container class for rtf save options.
    SaveImagesAsWmf *bool `json:"SaveImagesAsWmf,omitempty"`

    // Container class for rtf save options.
    SaveFormat *string `json:"SaveFormat,omitempty"`
}

func (RtfSaveOptionsData) IsRtfSaveOptionsData() bool {
    return true
}

func (RtfSaveOptionsData) IsSaveOptionsData() bool {
    return true
}

func (obj *RtfSaveOptionsData) Initialize() {
    var _SaveFormat = "rtf"
    obj.SaveFormat = &_SaveFormat


    if (obj.CustomTimeZoneInfoData != nil) {
        obj.CustomTimeZoneInfoData.Initialize()
    }


}

func (obj *RtfSaveOptionsData) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["ExportCompactSize"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportCompactSize = &parsedValue
        }

    } else if jsonValue, exists := json["exportCompactSize"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportCompactSize = &parsedValue
        }

    }

    if jsonValue, exists := json["ExportImagesForOldReaders"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportImagesForOldReaders = &parsedValue
        }

    } else if jsonValue, exists := json["exportImagesForOldReaders"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ExportImagesForOldReaders = &parsedValue
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

    if jsonValue, exists := json["SaveImagesAsWmf"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.SaveImagesAsWmf = &parsedValue
        }

    } else if jsonValue, exists := json["saveImagesAsWmf"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.SaveImagesAsWmf = &parsedValue
        }

    }
}

func (obj *RtfSaveOptionsData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *RtfSaveOptionsData) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.FileName == nil {
        return errors.New("Property FileName in RtfSaveOptionsData is required.")
    }
    if obj.CustomTimeZoneInfoData != nil {
        if err := obj.CustomTimeZoneInfoData.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *RtfSaveOptionsData) GetAllowEmbeddingPostScriptFonts() *bool {
    return obj.AllowEmbeddingPostScriptFonts
}

func (obj *RtfSaveOptionsData) SetAllowEmbeddingPostScriptFonts(value *bool) {
    obj.AllowEmbeddingPostScriptFonts = value
}

func (obj *RtfSaveOptionsData) GetCustomTimeZoneInfoData() ITimeZoneInfoData {
    return obj.CustomTimeZoneInfoData
}

func (obj *RtfSaveOptionsData) SetCustomTimeZoneInfoData(value ITimeZoneInfoData) {
    obj.CustomTimeZoneInfoData = value
}

func (obj *RtfSaveOptionsData) GetDml3DEffectsRenderingMode() *string {
    return obj.Dml3DEffectsRenderingMode
}

func (obj *RtfSaveOptionsData) SetDml3DEffectsRenderingMode(value *string) {
    obj.Dml3DEffectsRenderingMode = value
}

func (obj *RtfSaveOptionsData) GetDmlEffectsRenderingMode() *string {
    return obj.DmlEffectsRenderingMode
}

func (obj *RtfSaveOptionsData) SetDmlEffectsRenderingMode(value *string) {
    obj.DmlEffectsRenderingMode = value
}

func (obj *RtfSaveOptionsData) GetDmlRenderingMode() *string {
    return obj.DmlRenderingMode
}

func (obj *RtfSaveOptionsData) SetDmlRenderingMode(value *string) {
    obj.DmlRenderingMode = value
}

func (obj *RtfSaveOptionsData) GetFileName() *string {
    return obj.FileName
}

func (obj *RtfSaveOptionsData) SetFileName(value *string) {
    obj.FileName = value
}

func (obj *RtfSaveOptionsData) GetImlRenderingMode() *string {
    return obj.ImlRenderingMode
}

func (obj *RtfSaveOptionsData) SetImlRenderingMode(value *string) {
    obj.ImlRenderingMode = value
}

func (obj *RtfSaveOptionsData) GetUpdateAmbiguousTextFont() *bool {
    return obj.UpdateAmbiguousTextFont
}

func (obj *RtfSaveOptionsData) SetUpdateAmbiguousTextFont(value *bool) {
    obj.UpdateAmbiguousTextFont = value
}

func (obj *RtfSaveOptionsData) GetUpdateCreatedTimeProperty() *bool {
    return obj.UpdateCreatedTimeProperty
}

func (obj *RtfSaveOptionsData) SetUpdateCreatedTimeProperty(value *bool) {
    obj.UpdateCreatedTimeProperty = value
}

func (obj *RtfSaveOptionsData) GetUpdateFields() *bool {
    return obj.UpdateFields
}

func (obj *RtfSaveOptionsData) SetUpdateFields(value *bool) {
    obj.UpdateFields = value
}

func (obj *RtfSaveOptionsData) GetUpdateLastPrintedProperty() *bool {
    return obj.UpdateLastPrintedProperty
}

func (obj *RtfSaveOptionsData) SetUpdateLastPrintedProperty(value *bool) {
    obj.UpdateLastPrintedProperty = value
}

func (obj *RtfSaveOptionsData) GetUpdateLastSavedTimeProperty() *bool {
    return obj.UpdateLastSavedTimeProperty
}

func (obj *RtfSaveOptionsData) SetUpdateLastSavedTimeProperty(value *bool) {
    obj.UpdateLastSavedTimeProperty = value
}

func (obj *RtfSaveOptionsData) GetZipOutput() *bool {
    return obj.ZipOutput
}

func (obj *RtfSaveOptionsData) SetZipOutput(value *bool) {
    obj.ZipOutput = value
}

func (obj *RtfSaveOptionsData) GetExportCompactSize() *bool {
    return obj.ExportCompactSize
}

func (obj *RtfSaveOptionsData) SetExportCompactSize(value *bool) {
    obj.ExportCompactSize = value
}

func (obj *RtfSaveOptionsData) GetExportImagesForOldReaders() *bool {
    return obj.ExportImagesForOldReaders
}

func (obj *RtfSaveOptionsData) SetExportImagesForOldReaders(value *bool) {
    obj.ExportImagesForOldReaders = value
}

func (obj *RtfSaveOptionsData) GetPrettyFormat() *bool {
    return obj.PrettyFormat
}

func (obj *RtfSaveOptionsData) SetPrettyFormat(value *bool) {
    obj.PrettyFormat = value
}

func (obj *RtfSaveOptionsData) GetSaveImagesAsWmf() *bool {
    return obj.SaveImagesAsWmf
}

func (obj *RtfSaveOptionsData) SetSaveImagesAsWmf(value *bool) {
    obj.SaveImagesAsWmf = value
}

func (obj *RtfSaveOptionsData) GetSaveFormat() *string {
    return obj.SaveFormat
}

func (obj *RtfSaveOptionsData) SetSaveFormat(value *string) {
    obj.SaveFormat = value
}

