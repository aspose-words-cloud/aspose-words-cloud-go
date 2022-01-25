/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="form_field.go">
 *   Copyright (c) 2022 Aspose.Words for Cloud
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
type FormFieldResult struct {
    // FromField.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // FromField.
    NodeId string `json:"NodeId,omitempty"`

    // FromField.
    CalculateOnExit bool `json:"CalculateOnExit,omitempty"`

    // FromField.
    Enabled bool `json:"Enabled,omitempty"`

    // FromField.
    EntryMacro string `json:"EntryMacro,omitempty"`

    // FromField.
    ExitMacro string `json:"ExitMacro,omitempty"`

    // FromField.
    HelpText string `json:"HelpText,omitempty"`

    // FromField.
    Name string `json:"Name,omitempty"`

    // FromField.
    OwnHelp bool `json:"OwnHelp,omitempty"`

    // FromField.
    OwnStatus bool `json:"OwnStatus,omitempty"`

    // FromField.
    StatusText string `json:"StatusText,omitempty"`
}

type FormField struct {
    // FromField.
    Link IWordsApiLink `json:"Link,omitempty"`

    // FromField.
    NodeId *string `json:"NodeId,omitempty"`

    // FromField.
    CalculateOnExit *bool `json:"CalculateOnExit,omitempty"`

    // FromField.
    Enabled *bool `json:"Enabled,omitempty"`

    // FromField.
    EntryMacro *string `json:"EntryMacro,omitempty"`

    // FromField.
    ExitMacro *string `json:"ExitMacro,omitempty"`

    // FromField.
    HelpText *string `json:"HelpText,omitempty"`

    // FromField.
    Name *string `json:"Name,omitempty"`

    // FromField.
    OwnHelp *bool `json:"OwnHelp,omitempty"`

    // FromField.
    OwnStatus *bool `json:"OwnStatus,omitempty"`

    // FromField.
    StatusText *string `json:"StatusText,omitempty"`
}

type IFormField interface {
    IsFormField() bool
    Initialize()
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
}


