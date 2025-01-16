/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="form_field_text_input.go">
 *   Copyright (c) 2025 Aspose.Words for Cloud
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

// FormField text input element.

type IFormFieldTextInput interface {
    IsFormFieldTextInput() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
    Validate() error
    GetLink() IWordsApiLink
    SetLink(value IWordsApiLink)
    GetNodeId() *string
    SetNodeId(value *string)
    GetName() *string
    SetName(value *string)
    GetEnabled() *bool
    SetEnabled(value *bool)
    GetStatusText() *string
    SetStatusText(value *string)
    GetOwnStatus() *bool
    SetOwnStatus(value *bool)
    GetHelpText() *string
    SetHelpText(value *string)
    GetOwnHelp() *bool
    SetOwnHelp(value *bool)
    GetCalculateOnExit() *bool
    SetCalculateOnExit(value *bool)
    GetEntryMacro() *string
    SetEntryMacro(value *string)
    GetExitMacro() *string
    SetExitMacro(value *string)
    GetTextInputFormat() *string
    SetTextInputFormat(value *string)
    GetTextInputType() *string
    SetTextInputType(value *string)
    GetTextInputDefault() *string
    SetTextInputDefault(value *string)
    GetMaxLength() *int32
    SetMaxLength(value *int32)
}

type FormFieldTextInput struct {
    // FormField text input element.
    Link IWordsApiLink `json:"Link,omitempty"`

    // FormField text input element.
    NodeId *string `json:"NodeId,omitempty"`

    // FormField text input element.
    Name *string `json:"Name,omitempty"`

    // FormField text input element.
    Enabled *bool `json:"Enabled,omitempty"`

    // FormField text input element.
    StatusText *string `json:"StatusText,omitempty"`

    // FormField text input element.
    OwnStatus *bool `json:"OwnStatus,omitempty"`

    // FormField text input element.
    HelpText *string `json:"HelpText,omitempty"`

    // FormField text input element.
    OwnHelp *bool `json:"OwnHelp,omitempty"`

    // FormField text input element.
    CalculateOnExit *bool `json:"CalculateOnExit,omitempty"`

    // FormField text input element.
    EntryMacro *string `json:"EntryMacro,omitempty"`

    // FormField text input element.
    ExitMacro *string `json:"ExitMacro,omitempty"`

    // FormField text input element.
    TextInputFormat *string `json:"TextInputFormat,omitempty"`

    // FormField text input element.
    TextInputType *string `json:"TextInputType,omitempty"`

    // FormField text input element.
    TextInputDefault *string `json:"TextInputDefault,omitempty"`

    // FormField text input element.
    MaxLength *int32 `json:"MaxLength,omitempty"`
}

func (FormFieldTextInput) IsFormFieldTextInput() bool {
    return true
}

func (FormFieldTextInput) IsFormField() bool {
    return true
}

func (FormFieldTextInput) IsNodeLink() bool {
    return true
}

func (FormFieldTextInput) IsLinkElement() bool {
    return true
}

func (obj *FormFieldTextInput) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }


}

func (obj *FormFieldTextInput) Deserialize(json map[string]interface{}) {
    if jsonValue, exists := json["Link"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IWordsApiLink = new(WordsApiLink)
            modelInstance.Deserialize(parsedValue)
            obj.Link = modelInstance
        }

    } else if jsonValue, exists := json["link"]; exists {
        if parsedValue, valid := jsonValue.(map[string]interface{}); valid {
            var modelInstance IWordsApiLink = new(WordsApiLink)
            modelInstance.Deserialize(parsedValue)
            obj.Link = modelInstance
        }

    }

    if jsonValue, exists := json["NodeId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NodeId = &parsedValue
        }

    } else if jsonValue, exists := json["nodeId"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.NodeId = &parsedValue
        }

    }

    if jsonValue, exists := json["Name"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Name = &parsedValue
        }

    } else if jsonValue, exists := json["name"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.Name = &parsedValue
        }

    }

    if jsonValue, exists := json["Enabled"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Enabled = &parsedValue
        }

    } else if jsonValue, exists := json["enabled"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Enabled = &parsedValue
        }

    }

    if jsonValue, exists := json["StatusText"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.StatusText = &parsedValue
        }

    } else if jsonValue, exists := json["statusText"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.StatusText = &parsedValue
        }

    }

    if jsonValue, exists := json["OwnStatus"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.OwnStatus = &parsedValue
        }

    } else if jsonValue, exists := json["ownStatus"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.OwnStatus = &parsedValue
        }

    }

    if jsonValue, exists := json["HelpText"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.HelpText = &parsedValue
        }

    } else if jsonValue, exists := json["helpText"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.HelpText = &parsedValue
        }

    }

    if jsonValue, exists := json["OwnHelp"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.OwnHelp = &parsedValue
        }

    } else if jsonValue, exists := json["ownHelp"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.OwnHelp = &parsedValue
        }

    }

    if jsonValue, exists := json["CalculateOnExit"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.CalculateOnExit = &parsedValue
        }

    } else if jsonValue, exists := json["calculateOnExit"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.CalculateOnExit = &parsedValue
        }

    }

    if jsonValue, exists := json["EntryMacro"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.EntryMacro = &parsedValue
        }

    } else if jsonValue, exists := json["entryMacro"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.EntryMacro = &parsedValue
        }

    }

    if jsonValue, exists := json["ExitMacro"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ExitMacro = &parsedValue
        }

    } else if jsonValue, exists := json["exitMacro"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.ExitMacro = &parsedValue
        }

    }

    if jsonValue, exists := json["TextInputFormat"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.TextInputFormat = &parsedValue
        }

    } else if jsonValue, exists := json["textInputFormat"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.TextInputFormat = &parsedValue
        }

    }

    if jsonValue, exists := json["TextInputType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.TextInputType = &parsedValue
        }

    } else if jsonValue, exists := json["textInputType"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.TextInputType = &parsedValue
        }

    }

    if jsonValue, exists := json["TextInputDefault"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.TextInputDefault = &parsedValue
        }

    } else if jsonValue, exists := json["textInputDefault"]; exists {
        if parsedValue, valid := jsonValue.(string); valid {
            obj.TextInputDefault = &parsedValue
        }

    }

    if jsonValue, exists := json["MaxLength"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.MaxLength = new(int32)
            *obj.MaxLength = int32(parsedValue)
        }

    } else if jsonValue, exists := json["maxLength"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.MaxLength = new(int32)
            *obj.MaxLength = int32(parsedValue)
        }

    }
}

func (obj *FormFieldTextInput) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *FormFieldTextInput) Validate() error {
    if obj == nil {
        return errors.New("Invalid object.")
    }

    if obj.Name == nil {
        return errors.New("Property Name in FormFieldTextInput is required.")
    }
    if obj.TextInputFormat == nil {
        return errors.New("Property TextInputFormat in FormFieldTextInput is required.")
    }
    if obj.TextInputDefault == nil {
        return errors.New("Property TextInputDefault in FormFieldTextInput is required.")
    }
    if obj.Link != nil {
        if err := obj.Link.Validate(); err != nil {
            return err
        }
    }

    return nil;
}

func (obj *FormFieldTextInput) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *FormFieldTextInput) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *FormFieldTextInput) GetNodeId() *string {
    return obj.NodeId
}

func (obj *FormFieldTextInput) SetNodeId(value *string) {
    obj.NodeId = value
}

func (obj *FormFieldTextInput) GetName() *string {
    return obj.Name
}

func (obj *FormFieldTextInput) SetName(value *string) {
    obj.Name = value
}

func (obj *FormFieldTextInput) GetEnabled() *bool {
    return obj.Enabled
}

func (obj *FormFieldTextInput) SetEnabled(value *bool) {
    obj.Enabled = value
}

func (obj *FormFieldTextInput) GetStatusText() *string {
    return obj.StatusText
}

func (obj *FormFieldTextInput) SetStatusText(value *string) {
    obj.StatusText = value
}

func (obj *FormFieldTextInput) GetOwnStatus() *bool {
    return obj.OwnStatus
}

func (obj *FormFieldTextInput) SetOwnStatus(value *bool) {
    obj.OwnStatus = value
}

func (obj *FormFieldTextInput) GetHelpText() *string {
    return obj.HelpText
}

func (obj *FormFieldTextInput) SetHelpText(value *string) {
    obj.HelpText = value
}

func (obj *FormFieldTextInput) GetOwnHelp() *bool {
    return obj.OwnHelp
}

func (obj *FormFieldTextInput) SetOwnHelp(value *bool) {
    obj.OwnHelp = value
}

func (obj *FormFieldTextInput) GetCalculateOnExit() *bool {
    return obj.CalculateOnExit
}

func (obj *FormFieldTextInput) SetCalculateOnExit(value *bool) {
    obj.CalculateOnExit = value
}

func (obj *FormFieldTextInput) GetEntryMacro() *string {
    return obj.EntryMacro
}

func (obj *FormFieldTextInput) SetEntryMacro(value *string) {
    obj.EntryMacro = value
}

func (obj *FormFieldTextInput) GetExitMacro() *string {
    return obj.ExitMacro
}

func (obj *FormFieldTextInput) SetExitMacro(value *string) {
    obj.ExitMacro = value
}

func (obj *FormFieldTextInput) GetTextInputFormat() *string {
    return obj.TextInputFormat
}

func (obj *FormFieldTextInput) SetTextInputFormat(value *string) {
    obj.TextInputFormat = value
}

func (obj *FormFieldTextInput) GetTextInputType() *string {
    return obj.TextInputType
}

func (obj *FormFieldTextInput) SetTextInputType(value *string) {
    obj.TextInputType = value
}

func (obj *FormFieldTextInput) GetTextInputDefault() *string {
    return obj.TextInputDefault
}

func (obj *FormFieldTextInput) SetTextInputDefault(value *string) {
    obj.TextInputDefault = value
}

func (obj *FormFieldTextInput) GetMaxLength() *int32 {
    return obj.MaxLength
}

func (obj *FormFieldTextInput) SetMaxLength(value *int32) {
    obj.MaxLength = value
}

