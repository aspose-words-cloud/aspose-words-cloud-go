/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="text_save_options_data.go">
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

// Container class for text save options.

type ITextSaveOptionsData interface {
    IsTextSaveOptionsData() bool
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
    GetEncoding() *string
    SetEncoding(value *string)
    GetExportHeadersFootersMode() *string
    SetExportHeadersFootersMode(value *string)
    GetForcePageBreaks() *bool
    SetForcePageBreaks(value *bool)
    GetParagraphBreak() *string
    SetParagraphBreak(value *string)
    GetAddBidiMarks() *bool
    SetAddBidiMarks(value *bool)
    GetMaxCharactersPerLine() *int32
    SetMaxCharactersPerLine(value *int32)
    GetPreserveTableLayout() *bool
    SetPreserveTableLayout(value *bool)
    GetSimplifyListLabels() *bool
    SetSimplifyListLabels(value *bool)
    GetSaveFormat() *string
    SetSaveFormat(value *string)
}

type TextSaveOptionsData struct {
    // Container class for text save options.
    AllowEmbeddingPostScriptFonts *bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for text save options.
    CustomTimeZoneInfoData ITimeZoneInfoData `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for text save options.
    Dml3DEffectsRenderingMode *string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for text save options.
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for text save options.
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // Container class for text save options.
    FileName *string `json:"FileName,omitempty"`

    // Container class for text save options.
    ImlRenderingMode *string `json:"ImlRenderingMode,omitempty"`

    // Container class for text save options.
    UpdateCreatedTimeProperty *bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for text save options.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // Container class for text save options.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for text save options.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for text save options.
    ZipOutput *bool `json:"ZipOutput,omitempty"`

    // Container class for text save options.
    Encoding *string `json:"Encoding,omitempty"`

    // Container class for text save options.
    ExportHeadersFootersMode *string `json:"ExportHeadersFootersMode,omitempty"`

    // Container class for text save options.
    ForcePageBreaks *bool `json:"ForcePageBreaks,omitempty"`

    // Container class for text save options.
    ParagraphBreak *string `json:"ParagraphBreak,omitempty"`

    // Container class for text save options.
    AddBidiMarks *bool `json:"AddBidiMarks,omitempty"`

    // Container class for text save options.
    MaxCharactersPerLine *int32 `json:"MaxCharactersPerLine,omitempty"`

    // Container class for text save options.
    PreserveTableLayout *bool `json:"PreserveTableLayout,omitempty"`

    // Container class for text save options.
    SimplifyListLabels *bool `json:"SimplifyListLabels,omitempty"`

    // Container class for text save options.
    SaveFormat *string `json:"SaveFormat,omitempty"`
}

func (TextSaveOptionsData) IsTextSaveOptionsData() bool {
    return true
}

func (TextSaveOptionsData) IsTxtSaveOptionsBaseData() bool {
    return true
}

func (TextSaveOptionsData) IsSaveOptionsData() bool {
    return true
}

func (obj *TextSaveOptionsData) Initialize() {
    var _SaveFormat = "txt"
    obj.SaveFormat = &_SaveFormat


    if (obj.CustomTimeZoneInfoData != nil) {
        obj.CustomTimeZoneInfoData.Initialize()
    }


}

func (obj *TextSaveOptionsData) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["Encoding"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Encoding = &parsedValue
        }

    } else if jsonValue, exists := json["encoding"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Encoding = &parsedValue
        }

    }

    if jsonValue, exists := json["ExportHeadersFootersMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ExportHeadersFootersMode = &parsedValue
        }

    } else if jsonValue, exists := json["exportHeadersFootersMode"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ExportHeadersFootersMode = &parsedValue
        }

    }

    if jsonValue, exists := json["ForcePageBreaks"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ForcePageBreaks = &parsedValue
        }

    } else if jsonValue, exists := json["forcePageBreaks"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.ForcePageBreaks = &parsedValue
        }

    }

    if jsonValue, exists := json["ParagraphBreak"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ParagraphBreak = &parsedValue
        }

    } else if jsonValue, exists := json["paragraphBreak"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ParagraphBreak = &parsedValue
        }

    }

    if jsonValue, exists := json["AddBidiMarks"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.AddBidiMarks = &parsedValue
        }

    } else if jsonValue, exists := json["addBidiMarks"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.AddBidiMarks = &parsedValue
        }

    }

    if jsonValue, exists := json["MaxCharactersPerLine"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.MaxCharactersPerLine = new(int32)
            *obj.MaxCharactersPerLine = int32(parsedValue)
        }

    } else if jsonValue, exists := json["maxCharactersPerLine"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.MaxCharactersPerLine = new(int32)
            *obj.MaxCharactersPerLine = int32(parsedValue)
        }

    }

    if jsonValue, exists := json["PreserveTableLayout"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.PreserveTableLayout = &parsedValue
        }

    } else if jsonValue, exists := json["preserveTableLayout"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.PreserveTableLayout = &parsedValue
        }

    }

    if jsonValue, exists := json["SimplifyListLabels"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.SimplifyListLabels = &parsedValue
        }

    } else if jsonValue, exists := json["simplifyListLabels"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.SimplifyListLabels = &parsedValue
        }

    }
}

func (obj *TextSaveOptionsData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *TextSaveOptionsData) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.FileName == nil {
        return errors.New("Property FileName in TextSaveOptionsData is required.")
    }

    if obj.MaxCharactersPerLine == nil {
        return errors.New("Property MaxCharactersPerLine in TextSaveOptionsData is required.")
    }

    return nil;
}

func (obj *TextSaveOptionsData) GetAllowEmbeddingPostScriptFonts() *bool {
    return obj.AllowEmbeddingPostScriptFonts
}

func (obj *TextSaveOptionsData) SetAllowEmbeddingPostScriptFonts(value *bool) {
    obj.AllowEmbeddingPostScriptFonts = value
}

func (obj *TextSaveOptionsData) GetCustomTimeZoneInfoData() ITimeZoneInfoData {
    return obj.CustomTimeZoneInfoData
}

func (obj *TextSaveOptionsData) SetCustomTimeZoneInfoData(value ITimeZoneInfoData) {
    obj.CustomTimeZoneInfoData = value
}

func (obj *TextSaveOptionsData) GetDml3DEffectsRenderingMode() *string {
    return obj.Dml3DEffectsRenderingMode
}

func (obj *TextSaveOptionsData) SetDml3DEffectsRenderingMode(value *string) {
    obj.Dml3DEffectsRenderingMode = value
}

func (obj *TextSaveOptionsData) GetDmlEffectsRenderingMode() *string {
    return obj.DmlEffectsRenderingMode
}

func (obj *TextSaveOptionsData) SetDmlEffectsRenderingMode(value *string) {
    obj.DmlEffectsRenderingMode = value
}

func (obj *TextSaveOptionsData) GetDmlRenderingMode() *string {
    return obj.DmlRenderingMode
}

func (obj *TextSaveOptionsData) SetDmlRenderingMode(value *string) {
    obj.DmlRenderingMode = value
}

func (obj *TextSaveOptionsData) GetFileName() *string {
    return obj.FileName
}

func (obj *TextSaveOptionsData) SetFileName(value *string) {
    obj.FileName = value
}

func (obj *TextSaveOptionsData) GetImlRenderingMode() *string {
    return obj.ImlRenderingMode
}

func (obj *TextSaveOptionsData) SetImlRenderingMode(value *string) {
    obj.ImlRenderingMode = value
}

func (obj *TextSaveOptionsData) GetUpdateCreatedTimeProperty() *bool {
    return obj.UpdateCreatedTimeProperty
}

func (obj *TextSaveOptionsData) SetUpdateCreatedTimeProperty(value *bool) {
    obj.UpdateCreatedTimeProperty = value
}

func (obj *TextSaveOptionsData) GetUpdateFields() *bool {
    return obj.UpdateFields
}

func (obj *TextSaveOptionsData) SetUpdateFields(value *bool) {
    obj.UpdateFields = value
}

func (obj *TextSaveOptionsData) GetUpdateLastPrintedProperty() *bool {
    return obj.UpdateLastPrintedProperty
}

func (obj *TextSaveOptionsData) SetUpdateLastPrintedProperty(value *bool) {
    obj.UpdateLastPrintedProperty = value
}

func (obj *TextSaveOptionsData) GetUpdateLastSavedTimeProperty() *bool {
    return obj.UpdateLastSavedTimeProperty
}

func (obj *TextSaveOptionsData) SetUpdateLastSavedTimeProperty(value *bool) {
    obj.UpdateLastSavedTimeProperty = value
}

func (obj *TextSaveOptionsData) GetZipOutput() *bool {
    return obj.ZipOutput
}

func (obj *TextSaveOptionsData) SetZipOutput(value *bool) {
    obj.ZipOutput = value
}

func (obj *TextSaveOptionsData) GetEncoding() *string {
    return obj.Encoding
}

func (obj *TextSaveOptionsData) SetEncoding(value *string) {
    obj.Encoding = value
}

func (obj *TextSaveOptionsData) GetExportHeadersFootersMode() *string {
    return obj.ExportHeadersFootersMode
}

func (obj *TextSaveOptionsData) SetExportHeadersFootersMode(value *string) {
    obj.ExportHeadersFootersMode = value
}

func (obj *TextSaveOptionsData) GetForcePageBreaks() *bool {
    return obj.ForcePageBreaks
}

func (obj *TextSaveOptionsData) SetForcePageBreaks(value *bool) {
    obj.ForcePageBreaks = value
}

func (obj *TextSaveOptionsData) GetParagraphBreak() *string {
    return obj.ParagraphBreak
}

func (obj *TextSaveOptionsData) SetParagraphBreak(value *string) {
    obj.ParagraphBreak = value
}

func (obj *TextSaveOptionsData) GetAddBidiMarks() *bool {
    return obj.AddBidiMarks
}

func (obj *TextSaveOptionsData) SetAddBidiMarks(value *bool) {
    obj.AddBidiMarks = value
}

func (obj *TextSaveOptionsData) GetMaxCharactersPerLine() *int32 {
    return obj.MaxCharactersPerLine
}

func (obj *TextSaveOptionsData) SetMaxCharactersPerLine(value *int32) {
    obj.MaxCharactersPerLine = value
}

func (obj *TextSaveOptionsData) GetPreserveTableLayout() *bool {
    return obj.PreserveTableLayout
}

func (obj *TextSaveOptionsData) SetPreserveTableLayout(value *bool) {
    obj.PreserveTableLayout = value
}

func (obj *TextSaveOptionsData) GetSimplifyListLabels() *bool {
    return obj.SimplifyListLabels
}

func (obj *TextSaveOptionsData) SetSimplifyListLabels(value *bool) {
    obj.SimplifyListLabels = value
}

func (obj *TextSaveOptionsData) GetSaveFormat() *string {
    return obj.SaveFormat
}

func (obj *TextSaveOptionsData) SetSaveFormat(value *string) {
    obj.SaveFormat = value
}

