/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="field_options.go">
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

// DTO for field options.

type IFieldOptions interface {
    IsFieldOptions() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    GetBuiltInTemplatesPaths() []string
    SetBuiltInTemplatesPaths(value []string)
    GetCurrentUser() IUserInformation
    SetCurrentUser(value IUserInformation)
    GetCustomTocStyleSeparator() *string
    SetCustomTocStyleSeparator(value *string)
    GetDefaultDocumentAuthor() *string
    SetDefaultDocumentAuthor(value *string)
    GetFieldIndexFormat() *string
    SetFieldIndexFormat(value *string)
    GetFieldUpdateCultureName() *string
    SetFieldUpdateCultureName(value *string)
    GetFieldUpdateCultureSource() *string
    SetFieldUpdateCultureSource(value *string)
    GetFileName() *string
    SetFileName(value *string)
    GetIsBidiTextSupportedOnUpdate() *bool
    SetIsBidiTextSupportedOnUpdate(value *bool)
    GetLegacyNumberFormat() *bool
    SetLegacyNumberFormat(value *bool)
    GetPreProcessCultureName() *string
    SetPreProcessCultureName(value *string)
    GetTemplateName() *string
    SetTemplateName(value *string)
    GetUseInvariantCultureNumberFormat() *bool
    SetUseInvariantCultureNumberFormat(value *bool)
}

type FieldOptions struct {
    // DTO for field options.
    BuiltInTemplatesPaths []string `json:"BuiltInTemplatesPaths,omitempty"`

    // DTO for field options.
    CurrentUser IUserInformation `json:"CurrentUser,omitempty"`

    // DTO for field options.
    CustomTocStyleSeparator *string `json:"CustomTocStyleSeparator,omitempty"`

    // DTO for field options.
    DefaultDocumentAuthor *string `json:"DefaultDocumentAuthor,omitempty"`

    // DTO for field options.
    FieldIndexFormat *string `json:"FieldIndexFormat,omitempty"`

    // DTO for field options.
    FieldUpdateCultureName *string `json:"FieldUpdateCultureName,omitempty"`

    // DTO for field options.
    FieldUpdateCultureSource *string `json:"FieldUpdateCultureSource,omitempty"`

    // DTO for field options.
    FileName *string `json:"FileName,omitempty"`

    // DTO for field options.
    IsBidiTextSupportedOnUpdate *bool `json:"IsBidiTextSupportedOnUpdate,omitempty"`

    // DTO for field options.
    LegacyNumberFormat *bool `json:"LegacyNumberFormat,omitempty"`

    // DTO for field options.
    PreProcessCultureName *string `json:"PreProcessCultureName,omitempty"`

    // DTO for field options.
    TemplateName *string `json:"TemplateName,omitempty"`

    // DTO for field options.
    UseInvariantCultureNumberFormat *bool `json:"UseInvariantCultureNumberFormat,omitempty"`
}

func (FieldOptions) IsFieldOptions() bool {
    return true
}


func (obj *FieldOptions) Initialize() {
    if (obj.CurrentUser != nil) {
        obj.CurrentUser.Initialize()
    }


}

func (obj *FieldOptions) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["BuiltInTemplatesPaths"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.BuiltInTemplatesPaths = make([]string, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(string); valid {
                    obj.BuiltInTemplatesPaths = append(obj.BuiltInTemplatesPaths, elementValue)
                }

            }
        }

    } else if jsonValue, exists := json["builtInTemplatesPaths"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.BuiltInTemplatesPaths = make([]string, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(string); valid {
                    obj.BuiltInTemplatesPaths = append(obj.BuiltInTemplatesPaths, elementValue)
                }

            }
        }

    }

    if jsonValue, exists := json["CurrentUser"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IUserInformation = new(UserInformation)
            modelInstance.Deserialize(parsedValue)
            obj.CurrentUser = modelInstance
        }

    } else if jsonValue, exists := json["currentUser"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IUserInformation = new(UserInformation)
            modelInstance.Deserialize(parsedValue)
            obj.CurrentUser = modelInstance
        }

    }

    if jsonValue, exists := json["CustomTocStyleSeparator"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.CustomTocStyleSeparator = &parsedValue
        }

    } else if jsonValue, exists := json["customTocStyleSeparator"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.CustomTocStyleSeparator = &parsedValue
        }

    }

    if jsonValue, exists := json["DefaultDocumentAuthor"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DefaultDocumentAuthor = &parsedValue
        }

    } else if jsonValue, exists := json["defaultDocumentAuthor"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.DefaultDocumentAuthor = &parsedValue
        }

    }

    if jsonValue, exists := json["FieldIndexFormat"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FieldIndexFormat = &parsedValue
        }

    } else if jsonValue, exists := json["fieldIndexFormat"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FieldIndexFormat = &parsedValue
        }

    }

    if jsonValue, exists := json["FieldUpdateCultureName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FieldUpdateCultureName = &parsedValue
        }

    } else if jsonValue, exists := json["fieldUpdateCultureName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FieldUpdateCultureName = &parsedValue
        }

    }

    if jsonValue, exists := json["FieldUpdateCultureSource"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FieldUpdateCultureSource = &parsedValue
        }

    } else if jsonValue, exists := json["fieldUpdateCultureSource"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.FieldUpdateCultureSource = &parsedValue
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

    if jsonValue, exists := json["IsBidiTextSupportedOnUpdate"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsBidiTextSupportedOnUpdate = &parsedValue
        }

    } else if jsonValue, exists := json["isBidiTextSupportedOnUpdate"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsBidiTextSupportedOnUpdate = &parsedValue
        }

    }

    if jsonValue, exists := json["LegacyNumberFormat"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.LegacyNumberFormat = &parsedValue
        }

    } else if jsonValue, exists := json["legacyNumberFormat"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.LegacyNumberFormat = &parsedValue
        }

    }

    if jsonValue, exists := json["PreProcessCultureName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.PreProcessCultureName = &parsedValue
        }

    } else if jsonValue, exists := json["preProcessCultureName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.PreProcessCultureName = &parsedValue
        }

    }

    if jsonValue, exists := json["TemplateName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.TemplateName = &parsedValue
        }

    } else if jsonValue, exists := json["templateName"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.TemplateName = &parsedValue
        }

    }

    if jsonValue, exists := json["UseInvariantCultureNumberFormat"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UseInvariantCultureNumberFormat = &parsedValue
        }

    } else if jsonValue, exists := json["useInvariantCultureNumberFormat"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.UseInvariantCultureNumberFormat = &parsedValue
        }

    }
}

func (obj *FieldOptions) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *FieldOptions) GetBuiltInTemplatesPaths() []string {
    return obj.BuiltInTemplatesPaths
}

func (obj *FieldOptions) SetBuiltInTemplatesPaths(value []string) {
    obj.BuiltInTemplatesPaths = value
}

func (obj *FieldOptions) GetCurrentUser() IUserInformation {
    return obj.CurrentUser
}

func (obj *FieldOptions) SetCurrentUser(value IUserInformation) {
    obj.CurrentUser = value
}

func (obj *FieldOptions) GetCustomTocStyleSeparator() *string {
    return obj.CustomTocStyleSeparator
}

func (obj *FieldOptions) SetCustomTocStyleSeparator(value *string) {
    obj.CustomTocStyleSeparator = value
}

func (obj *FieldOptions) GetDefaultDocumentAuthor() *string {
    return obj.DefaultDocumentAuthor
}

func (obj *FieldOptions) SetDefaultDocumentAuthor(value *string) {
    obj.DefaultDocumentAuthor = value
}

func (obj *FieldOptions) GetFieldIndexFormat() *string {
    return obj.FieldIndexFormat
}

func (obj *FieldOptions) SetFieldIndexFormat(value *string) {
    obj.FieldIndexFormat = value
}

func (obj *FieldOptions) GetFieldUpdateCultureName() *string {
    return obj.FieldUpdateCultureName
}

func (obj *FieldOptions) SetFieldUpdateCultureName(value *string) {
    obj.FieldUpdateCultureName = value
}

func (obj *FieldOptions) GetFieldUpdateCultureSource() *string {
    return obj.FieldUpdateCultureSource
}

func (obj *FieldOptions) SetFieldUpdateCultureSource(value *string) {
    obj.FieldUpdateCultureSource = value
}

func (obj *FieldOptions) GetFileName() *string {
    return obj.FileName
}

func (obj *FieldOptions) SetFileName(value *string) {
    obj.FileName = value
}

func (obj *FieldOptions) GetIsBidiTextSupportedOnUpdate() *bool {
    return obj.IsBidiTextSupportedOnUpdate
}

func (obj *FieldOptions) SetIsBidiTextSupportedOnUpdate(value *bool) {
    obj.IsBidiTextSupportedOnUpdate = value
}

func (obj *FieldOptions) GetLegacyNumberFormat() *bool {
    return obj.LegacyNumberFormat
}

func (obj *FieldOptions) SetLegacyNumberFormat(value *bool) {
    obj.LegacyNumberFormat = value
}

func (obj *FieldOptions) GetPreProcessCultureName() *string {
    return obj.PreProcessCultureName
}

func (obj *FieldOptions) SetPreProcessCultureName(value *string) {
    obj.PreProcessCultureName = value
}

func (obj *FieldOptions) GetTemplateName() *string {
    return obj.TemplateName
}

func (obj *FieldOptions) SetTemplateName(value *string) {
    obj.TemplateName = value
}

func (obj *FieldOptions) GetUseInvariantCultureNumberFormat() *bool {
    return obj.UseInvariantCultureNumberFormat
}

func (obj *FieldOptions) SetUseInvariantCultureNumberFormat(value *bool) {
    obj.UseInvariantCultureNumberFormat = value
}

