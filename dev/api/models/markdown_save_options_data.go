/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="markdown_save_options_data.go">
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

// Container class for markdown save options.

type IMarkdownSaveOptionsData interface {
    IsMarkdownSaveOptionsData() bool
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
    GetTableContentAlignment() *string
    SetTableContentAlignment(value *string)
    GetSaveFormat() *string
    SetSaveFormat(value *string)
}

type MarkdownSaveOptionsData struct {
    // Container class for markdown save options.
    AllowEmbeddingPostScriptFonts *bool `json:"AllowEmbeddingPostScriptFonts,omitempty"`

    // Container class for markdown save options.
    CustomTimeZoneInfoData ITimeZoneInfoData `json:"CustomTimeZoneInfoData,omitempty"`

    // Container class for markdown save options.
    Dml3DEffectsRenderingMode *string `json:"Dml3DEffectsRenderingMode,omitempty"`

    // Container class for markdown save options.
    DmlEffectsRenderingMode *string `json:"DmlEffectsRenderingMode,omitempty"`

    // Container class for markdown save options.
    DmlRenderingMode *string `json:"DmlRenderingMode,omitempty"`

    // Container class for markdown save options.
    FileName *string `json:"FileName,omitempty"`

    // Container class for markdown save options.
    ImlRenderingMode *string `json:"ImlRenderingMode,omitempty"`

    // Container class for markdown save options.
    UpdateCreatedTimeProperty *bool `json:"UpdateCreatedTimeProperty,omitempty"`

    // Container class for markdown save options.
    UpdateFields *bool `json:"UpdateFields,omitempty"`

    // Container class for markdown save options.
    UpdateLastPrintedProperty *bool `json:"UpdateLastPrintedProperty,omitempty"`

    // Container class for markdown save options.
    UpdateLastSavedTimeProperty *bool `json:"UpdateLastSavedTimeProperty,omitempty"`

    // Container class for markdown save options.
    ZipOutput *bool `json:"ZipOutput,omitempty"`

    // Container class for markdown save options.
    Encoding *string `json:"Encoding,omitempty"`

    // Container class for markdown save options.
    ExportHeadersFootersMode *string `json:"ExportHeadersFootersMode,omitempty"`

    // Container class for markdown save options.
    ForcePageBreaks *bool `json:"ForcePageBreaks,omitempty"`

    // Container class for markdown save options.
    ParagraphBreak *string `json:"ParagraphBreak,omitempty"`

    // Container class for markdown save options.
    TableContentAlignment *string `json:"TableContentAlignment,omitempty"`

    // Container class for markdown save options.
    SaveFormat *string `json:"SaveFormat,omitempty"`
}

func (MarkdownSaveOptionsData) IsMarkdownSaveOptionsData() bool {
    return true
}

func (MarkdownSaveOptionsData) IsTxtSaveOptionsBaseData() bool {
    return true
}

func (MarkdownSaveOptionsData) IsSaveOptionsData() bool {
    return true
}

func (obj *MarkdownSaveOptionsData) Initialize() {
    var _SaveFormat = "md"
    obj.SaveFormat = &_SaveFormat


    if (obj.CustomTimeZoneInfoData != nil) {
        obj.CustomTimeZoneInfoData.Initialize()
    }


}

func (obj *MarkdownSaveOptionsData) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["TableContentAlignment"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.TableContentAlignment = &parsedValue
        }

    } else if jsonValue, exists := json["tableContentAlignment"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.TableContentAlignment = &parsedValue
        }

    }
}

func (obj *MarkdownSaveOptionsData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *MarkdownSaveOptionsData) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.FileName == nil {
        return errors.New("Property FileName in MarkdownSaveOptionsData is required.")
    }
    if obj.CustomTimeZoneInfoData != nil {
        if err := obj.CustomTimeZoneInfoData.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *MarkdownSaveOptionsData) GetAllowEmbeddingPostScriptFonts() *bool {
    return obj.AllowEmbeddingPostScriptFonts
}

func (obj *MarkdownSaveOptionsData) SetAllowEmbeddingPostScriptFonts(value *bool) {
    obj.AllowEmbeddingPostScriptFonts = value
}

func (obj *MarkdownSaveOptionsData) GetCustomTimeZoneInfoData() ITimeZoneInfoData {
    return obj.CustomTimeZoneInfoData
}

func (obj *MarkdownSaveOptionsData) SetCustomTimeZoneInfoData(value ITimeZoneInfoData) {
    obj.CustomTimeZoneInfoData = value
}

func (obj *MarkdownSaveOptionsData) GetDml3DEffectsRenderingMode() *string {
    return obj.Dml3DEffectsRenderingMode
}

func (obj *MarkdownSaveOptionsData) SetDml3DEffectsRenderingMode(value *string) {
    obj.Dml3DEffectsRenderingMode = value
}

func (obj *MarkdownSaveOptionsData) GetDmlEffectsRenderingMode() *string {
    return obj.DmlEffectsRenderingMode
}

func (obj *MarkdownSaveOptionsData) SetDmlEffectsRenderingMode(value *string) {
    obj.DmlEffectsRenderingMode = value
}

func (obj *MarkdownSaveOptionsData) GetDmlRenderingMode() *string {
    return obj.DmlRenderingMode
}

func (obj *MarkdownSaveOptionsData) SetDmlRenderingMode(value *string) {
    obj.DmlRenderingMode = value
}

func (obj *MarkdownSaveOptionsData) GetFileName() *string {
    return obj.FileName
}

func (obj *MarkdownSaveOptionsData) SetFileName(value *string) {
    obj.FileName = value
}

func (obj *MarkdownSaveOptionsData) GetImlRenderingMode() *string {
    return obj.ImlRenderingMode
}

func (obj *MarkdownSaveOptionsData) SetImlRenderingMode(value *string) {
    obj.ImlRenderingMode = value
}

func (obj *MarkdownSaveOptionsData) GetUpdateCreatedTimeProperty() *bool {
    return obj.UpdateCreatedTimeProperty
}

func (obj *MarkdownSaveOptionsData) SetUpdateCreatedTimeProperty(value *bool) {
    obj.UpdateCreatedTimeProperty = value
}

func (obj *MarkdownSaveOptionsData) GetUpdateFields() *bool {
    return obj.UpdateFields
}

func (obj *MarkdownSaveOptionsData) SetUpdateFields(value *bool) {
    obj.UpdateFields = value
}

func (obj *MarkdownSaveOptionsData) GetUpdateLastPrintedProperty() *bool {
    return obj.UpdateLastPrintedProperty
}

func (obj *MarkdownSaveOptionsData) SetUpdateLastPrintedProperty(value *bool) {
    obj.UpdateLastPrintedProperty = value
}

func (obj *MarkdownSaveOptionsData) GetUpdateLastSavedTimeProperty() *bool {
    return obj.UpdateLastSavedTimeProperty
}

func (obj *MarkdownSaveOptionsData) SetUpdateLastSavedTimeProperty(value *bool) {
    obj.UpdateLastSavedTimeProperty = value
}

func (obj *MarkdownSaveOptionsData) GetZipOutput() *bool {
    return obj.ZipOutput
}

func (obj *MarkdownSaveOptionsData) SetZipOutput(value *bool) {
    obj.ZipOutput = value
}

func (obj *MarkdownSaveOptionsData) GetEncoding() *string {
    return obj.Encoding
}

func (obj *MarkdownSaveOptionsData) SetEncoding(value *string) {
    obj.Encoding = value
}

func (obj *MarkdownSaveOptionsData) GetExportHeadersFootersMode() *string {
    return obj.ExportHeadersFootersMode
}

func (obj *MarkdownSaveOptionsData) SetExportHeadersFootersMode(value *string) {
    obj.ExportHeadersFootersMode = value
}

func (obj *MarkdownSaveOptionsData) GetForcePageBreaks() *bool {
    return obj.ForcePageBreaks
}

func (obj *MarkdownSaveOptionsData) SetForcePageBreaks(value *bool) {
    obj.ForcePageBreaks = value
}

func (obj *MarkdownSaveOptionsData) GetParagraphBreak() *string {
    return obj.ParagraphBreak
}

func (obj *MarkdownSaveOptionsData) SetParagraphBreak(value *string) {
    obj.ParagraphBreak = value
}

func (obj *MarkdownSaveOptionsData) GetTableContentAlignment() *string {
    return obj.TableContentAlignment
}

func (obj *MarkdownSaveOptionsData) SetTableContentAlignment(value *string) {
    obj.TableContentAlignment = value
}

func (obj *MarkdownSaveOptionsData) GetSaveFormat() *string {
    return obj.SaveFormat
}

func (obj *MarkdownSaveOptionsData) SetSaveFormat(value *string) {
    obj.SaveFormat = value
}

