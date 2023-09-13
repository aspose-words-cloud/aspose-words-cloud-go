/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="form_field_checkbox.go">
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

// FormField checkbox element.

type IFormFieldCheckbox interface {
    IsFormFieldCheckbox() bool
    Initialize()
    Deserialize(json map[string]interface{})
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
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
    GetIsCheckBoxExactSize() *bool
    SetIsCheckBoxExactSize(value *bool)
    GetCheckBoxSize() *float64
    SetCheckBoxSize(value *float64)
    GetChecked() *bool
    SetChecked(value *bool)
}

type FormFieldCheckbox struct {
    // FormField checkbox element.
    Link IWordsApiLink `json:"Link,omitempty"`

    // FormField checkbox element.
    NodeId *string `json:"NodeId,omitempty"`

    // FormField checkbox element.
    Name *string `json:"Name,omitempty"`

    // FormField checkbox element.
    Enabled *bool `json:"Enabled,omitempty"`

    // FormField checkbox element.
    StatusText *string `json:"StatusText,omitempty"`

    // FormField checkbox element.
    OwnStatus *bool `json:"OwnStatus,omitempty"`

    // FormField checkbox element.
    HelpText *string `json:"HelpText,omitempty"`

    // FormField checkbox element.
    OwnHelp *bool `json:"OwnHelp,omitempty"`

    // FormField checkbox element.
    CalculateOnExit *bool `json:"CalculateOnExit,omitempty"`

    // FormField checkbox element.
    EntryMacro *string `json:"EntryMacro,omitempty"`

    // FormField checkbox element.
    ExitMacro *string `json:"ExitMacro,omitempty"`

    // FormField checkbox element.
    IsCheckBoxExactSize *bool `json:"IsCheckBoxExactSize,omitempty"`

    // FormField checkbox element.
    CheckBoxSize *float64 `json:"CheckBoxSize,omitempty"`

    // FormField checkbox element.
    Checked *bool `json:"Checked,omitempty"`
}

func (FormFieldCheckbox) IsFormFieldCheckbox() bool {
    return true
}

func (FormFieldCheckbox) IsFormField() bool {
    return true
}

func (FormFieldCheckbox) IsNodeLink() bool {
    return true
}

func (FormFieldCheckbox) IsLinkElement() bool {
    return true
}

func (obj *FormFieldCheckbox) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }


}

func (obj *FormFieldCheckbox) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["IsCheckBoxExactSize"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsCheckBoxExactSize = &parsedValue
        }

    } else if jsonValue, exists := json["isCheckBoxExactSize"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.IsCheckBoxExactSize = &parsedValue
        }

    }

    if jsonValue, exists := json["CheckBoxSize"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.CheckBoxSize = &parsedValue
        }

    } else if jsonValue, exists := json["checkBoxSize"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.CheckBoxSize = &parsedValue
        }

    }

    if jsonValue, exists := json["Checked"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Checked = &parsedValue
        }

    } else if jsonValue, exists := json["checked"]; exists {
        if parsedValue, valid := jsonValue.(bool); valid {
            obj.Checked = &parsedValue
        }

    }
}

func (obj *FormFieldCheckbox) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *FormFieldCheckbox) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *FormFieldCheckbox) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *FormFieldCheckbox) GetNodeId() *string {
    return obj.NodeId
}

func (obj *FormFieldCheckbox) SetNodeId(value *string) {
    obj.NodeId = value
}

func (obj *FormFieldCheckbox) GetName() *string {
    return obj.Name
}

func (obj *FormFieldCheckbox) SetName(value *string) {
    obj.Name = value
}

func (obj *FormFieldCheckbox) GetEnabled() *bool {
    return obj.Enabled
}

func (obj *FormFieldCheckbox) SetEnabled(value *bool) {
    obj.Enabled = value
}

func (obj *FormFieldCheckbox) GetStatusText() *string {
    return obj.StatusText
}

func (obj *FormFieldCheckbox) SetStatusText(value *string) {
    obj.StatusText = value
}

func (obj *FormFieldCheckbox) GetOwnStatus() *bool {
    return obj.OwnStatus
}

func (obj *FormFieldCheckbox) SetOwnStatus(value *bool) {
    obj.OwnStatus = value
}

func (obj *FormFieldCheckbox) GetHelpText() *string {
    return obj.HelpText
}

func (obj *FormFieldCheckbox) SetHelpText(value *string) {
    obj.HelpText = value
}

func (obj *FormFieldCheckbox) GetOwnHelp() *bool {
    return obj.OwnHelp
}

func (obj *FormFieldCheckbox) SetOwnHelp(value *bool) {
    obj.OwnHelp = value
}

func (obj *FormFieldCheckbox) GetCalculateOnExit() *bool {
    return obj.CalculateOnExit
}

func (obj *FormFieldCheckbox) SetCalculateOnExit(value *bool) {
    obj.CalculateOnExit = value
}

func (obj *FormFieldCheckbox) GetEntryMacro() *string {
    return obj.EntryMacro
}

func (obj *FormFieldCheckbox) SetEntryMacro(value *string) {
    obj.EntryMacro = value
}

func (obj *FormFieldCheckbox) GetExitMacro() *string {
    return obj.ExitMacro
}

func (obj *FormFieldCheckbox) SetExitMacro(value *string) {
    obj.ExitMacro = value
}

func (obj *FormFieldCheckbox) GetIsCheckBoxExactSize() *bool {
    return obj.IsCheckBoxExactSize
}

func (obj *FormFieldCheckbox) SetIsCheckBoxExactSize(value *bool) {
    obj.IsCheckBoxExactSize = value
}

func (obj *FormFieldCheckbox) GetCheckBoxSize() *float64 {
    return obj.CheckBoxSize
}

func (obj *FormFieldCheckbox) SetCheckBoxSize(value *float64) {
    obj.CheckBoxSize = value
}

func (obj *FormFieldCheckbox) GetChecked() *bool {
    return obj.Checked
}

func (obj *FormFieldCheckbox) SetChecked(value *bool) {
    obj.Checked = value
}

