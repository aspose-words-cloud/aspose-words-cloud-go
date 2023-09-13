/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="form_field_drop_down.go">
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

// FormField dropdownlist element.

type IFormFieldDropDown interface {
    IsFormFieldDropDown() bool
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
    GetDropDownItems() []string
    SetDropDownItems(value []string)
    GetDropDownSelectedIndex() *int32
    SetDropDownSelectedIndex(value *int32)
}

type FormFieldDropDown struct {
    // FormField dropdownlist element.
    Link IWordsApiLink `json:"Link,omitempty"`

    // FormField dropdownlist element.
    NodeId *string `json:"NodeId,omitempty"`

    // FormField dropdownlist element.
    Name *string `json:"Name,omitempty"`

    // FormField dropdownlist element.
    Enabled *bool `json:"Enabled,omitempty"`

    // FormField dropdownlist element.
    StatusText *string `json:"StatusText,omitempty"`

    // FormField dropdownlist element.
    OwnStatus *bool `json:"OwnStatus,omitempty"`

    // FormField dropdownlist element.
    HelpText *string `json:"HelpText,omitempty"`

    // FormField dropdownlist element.
    OwnHelp *bool `json:"OwnHelp,omitempty"`

    // FormField dropdownlist element.
    CalculateOnExit *bool `json:"CalculateOnExit,omitempty"`

    // FormField dropdownlist element.
    EntryMacro *string `json:"EntryMacro,omitempty"`

    // FormField dropdownlist element.
    ExitMacro *string `json:"ExitMacro,omitempty"`

    // FormField dropdownlist element.
    DropDownItems []string `json:"DropDownItems,omitempty"`

    // FormField dropdownlist element.
    DropDownSelectedIndex *int32 `json:"DropDownSelectedIndex,omitempty"`
}

func (FormFieldDropDown) IsFormFieldDropDown() bool {
    return true
}

func (FormFieldDropDown) IsFormField() bool {
    return true
}

func (FormFieldDropDown) IsNodeLink() bool {
    return true
}

func (FormFieldDropDown) IsLinkElement() bool {
    return true
}

func (obj *FormFieldDropDown) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }


}

func (obj *FormFieldDropDown) Deserialize(json map[string]interface{}) {
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

    if jsonValue, exists := json["DropDownItems"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.DropDownItems = make([]string, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(string); valid {
                    obj.DropDownItems = append(obj.DropDownItems, elementValue)
                }

            }
        }

    } else if jsonValue, exists := json["dropDownItems"]; exists {
        if parsedValue, valid := jsonValue.([]interface{}); valid {
            obj.DropDownItems = make([]string, 0)
            for _, parsedElement := range parsedValue {
                if elementValue, valid := parsedElement.(string); valid {
                    obj.DropDownItems = append(obj.DropDownItems, elementValue)
                }

            }
        }

    }

    if jsonValue, exists := json["DropDownSelectedIndex"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.DropDownSelectedIndex = new(int32)
            *obj.DropDownSelectedIndex = int32(parsedValue)
        }

    } else if jsonValue, exists := json["dropDownSelectedIndex"]; exists {
        if parsedValue, valid := jsonValue.(float64); valid {
            obj.DropDownSelectedIndex = new(int32)
            *obj.DropDownSelectedIndex = int32(parsedValue)
        }

    }
}

func (obj *FormFieldDropDown) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *FormFieldDropDown) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *FormFieldDropDown) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *FormFieldDropDown) GetNodeId() *string {
    return obj.NodeId
}

func (obj *FormFieldDropDown) SetNodeId(value *string) {
    obj.NodeId = value
}

func (obj *FormFieldDropDown) GetName() *string {
    return obj.Name
}

func (obj *FormFieldDropDown) SetName(value *string) {
    obj.Name = value
}

func (obj *FormFieldDropDown) GetEnabled() *bool {
    return obj.Enabled
}

func (obj *FormFieldDropDown) SetEnabled(value *bool) {
    obj.Enabled = value
}

func (obj *FormFieldDropDown) GetStatusText() *string {
    return obj.StatusText
}

func (obj *FormFieldDropDown) SetStatusText(value *string) {
    obj.StatusText = value
}

func (obj *FormFieldDropDown) GetOwnStatus() *bool {
    return obj.OwnStatus
}

func (obj *FormFieldDropDown) SetOwnStatus(value *bool) {
    obj.OwnStatus = value
}

func (obj *FormFieldDropDown) GetHelpText() *string {
    return obj.HelpText
}

func (obj *FormFieldDropDown) SetHelpText(value *string) {
    obj.HelpText = value
}

func (obj *FormFieldDropDown) GetOwnHelp() *bool {
    return obj.OwnHelp
}

func (obj *FormFieldDropDown) SetOwnHelp(value *bool) {
    obj.OwnHelp = value
}

func (obj *FormFieldDropDown) GetCalculateOnExit() *bool {
    return obj.CalculateOnExit
}

func (obj *FormFieldDropDown) SetCalculateOnExit(value *bool) {
    obj.CalculateOnExit = value
}

func (obj *FormFieldDropDown) GetEntryMacro() *string {
    return obj.EntryMacro
}

func (obj *FormFieldDropDown) SetEntryMacro(value *string) {
    obj.EntryMacro = value
}

func (obj *FormFieldDropDown) GetExitMacro() *string {
    return obj.ExitMacro
}

func (obj *FormFieldDropDown) SetExitMacro(value *string) {
    obj.ExitMacro = value
}

func (obj *FormFieldDropDown) GetDropDownItems() []string {
    return obj.DropDownItems
}

func (obj *FormFieldDropDown) SetDropDownItems(value []string) {
    obj.DropDownItems = value
}

func (obj *FormFieldDropDown) GetDropDownSelectedIndex() *int32 {
    return obj.DropDownSelectedIndex
}

func (obj *FormFieldDropDown) SetDropDownSelectedIndex(value *int32) {
    obj.DropDownSelectedIndex = value
}

