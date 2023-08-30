/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="form_field.go">
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

// FromField.

type IFormField interface {
    IsFormField() bool
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
}

type FormField struct {
    // FromField.
    Link IWordsApiLink `json:"Link,omitempty"`

    // FromField.
    NodeId *string `json:"NodeId,omitempty"`

    // FromField.
    Name *string `json:"Name,omitempty"`

    // FromField.
    Enabled *bool `json:"Enabled,omitempty"`

    // FromField.
    StatusText *string `json:"StatusText,omitempty"`

    // FromField.
    OwnStatus *bool `json:"OwnStatus,omitempty"`

    // FromField.
    HelpText *string `json:"HelpText,omitempty"`

    // FromField.
    OwnHelp *bool `json:"OwnHelp,omitempty"`

    // FromField.
    CalculateOnExit *bool `json:"CalculateOnExit,omitempty"`

    // FromField.
    EntryMacro *string `json:"EntryMacro,omitempty"`

    // FromField.
    ExitMacro *string `json:"ExitMacro,omitempty"`
}

func (FormField) IsFormField() bool {
    return true
}

func (FormField) IsNodeLink() bool {
    return true
}

func (FormField) IsLinkElement() bool {
    return true
}

func (obj *FormField) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }


}

func (obj *FormField) Deserialize(json map[string]interface{}) {
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
}

func (obj *FormField) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}

func (obj *FormField) GetLink() IWordsApiLink {
    return obj.Link
}

func (obj *FormField) SetLink(value IWordsApiLink) {
    obj.Link = value
}

func (obj *FormField) GetNodeId() *string {
    return obj.NodeId
}

func (obj *FormField) SetNodeId(value *string) {
    obj.NodeId = value
}

func (obj *FormField) GetName() *string {
    return obj.Name
}

func (obj *FormField) SetName(value *string) {
    obj.Name = value
}

func (obj *FormField) GetEnabled() *bool {
    return obj.Enabled
}

func (obj *FormField) SetEnabled(value *bool) {
    obj.Enabled = value
}

func (obj *FormField) GetStatusText() *string {
    return obj.StatusText
}

func (obj *FormField) SetStatusText(value *string) {
    obj.StatusText = value
}

func (obj *FormField) GetOwnStatus() *bool {
    return obj.OwnStatus
}

func (obj *FormField) SetOwnStatus(value *bool) {
    obj.OwnStatus = value
}

func (obj *FormField) GetHelpText() *string {
    return obj.HelpText
}

func (obj *FormField) SetHelpText(value *string) {
    obj.HelpText = value
}

func (obj *FormField) GetOwnHelp() *bool {
    return obj.OwnHelp
}

func (obj *FormField) SetOwnHelp(value *bool) {
    obj.OwnHelp = value
}

func (obj *FormField) GetCalculateOnExit() *bool {
    return obj.CalculateOnExit
}

func (obj *FormField) SetCalculateOnExit(value *bool) {
    obj.CalculateOnExit = value
}

func (obj *FormField) GetEntryMacro() *string {
    return obj.EntryMacro
}

func (obj *FormField) SetEntryMacro(value *string) {
    obj.EntryMacro = value
}

func (obj *FormField) GetExitMacro() *string {
    return obj.ExitMacro
}

func (obj *FormField) SetExitMacro(value *string) {
    obj.ExitMacro = value
}

