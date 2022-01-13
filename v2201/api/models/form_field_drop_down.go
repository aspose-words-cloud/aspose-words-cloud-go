/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="form_field_drop_down.go">
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

// FormField dropdownlist element.
type FormFieldDropDownResult struct {
    // FormField dropdownlist element.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // FormField dropdownlist element.
    NodeId string `json:"NodeId,omitempty"`

    // FormField dropdownlist element.
    CalculateOnExit bool `json:"CalculateOnExit,omitempty"`

    // FormField dropdownlist element.
    Enabled bool `json:"Enabled,omitempty"`

    // FormField dropdownlist element.
    EntryMacro string `json:"EntryMacro,omitempty"`

    // FormField dropdownlist element.
    ExitMacro string `json:"ExitMacro,omitempty"`

    // FormField dropdownlist element.
    HelpText string `json:"HelpText,omitempty"`

    // FormField dropdownlist element.
    Name string `json:"Name,omitempty"`

    // FormField dropdownlist element.
    OwnHelp bool `json:"OwnHelp,omitempty"`

    // FormField dropdownlist element.
    OwnStatus bool `json:"OwnStatus,omitempty"`

    // FormField dropdownlist element.
    StatusText string `json:"StatusText,omitempty"`

    // FormField dropdownlist element.
    DropDownItems []string `json:"DropDownItems,omitempty"`

    // FormField dropdownlist element.
    DropDownSelectedIndex int32 `json:"DropDownSelectedIndex,omitempty"`
}

type FormFieldDropDown struct {
    // FormField dropdownlist element.
    Link IWordsApiLink `json:"Link,omitempty"`

    // FormField dropdownlist element.
    NodeId *string `json:"NodeId,omitempty"`

    // FormField dropdownlist element.
    CalculateOnExit *bool `json:"CalculateOnExit,omitempty"`

    // FormField dropdownlist element.
    Enabled *bool `json:"Enabled,omitempty"`

    // FormField dropdownlist element.
    EntryMacro *string `json:"EntryMacro,omitempty"`

    // FormField dropdownlist element.
    ExitMacro *string `json:"ExitMacro,omitempty"`

    // FormField dropdownlist element.
    HelpText *string `json:"HelpText,omitempty"`

    // FormField dropdownlist element.
    Name *string `json:"Name,omitempty"`

    // FormField dropdownlist element.
    OwnHelp *bool `json:"OwnHelp,omitempty"`

    // FormField dropdownlist element.
    OwnStatus *bool `json:"OwnStatus,omitempty"`

    // FormField dropdownlist element.
    StatusText *string `json:"StatusText,omitempty"`

    // FormField dropdownlist element.
    DropDownItems []string `json:"DropDownItems,omitempty"`

    // FormField dropdownlist element.
    DropDownSelectedIndex *int32 `json:"DropDownSelectedIndex,omitempty"`
}

type IFormFieldDropDown interface {
    IsFormFieldDropDown() bool
}
func (FormFieldDropDown) IsFormFieldDropDown() bool {
    return true
}

func (FormFieldDropDown) IsFormField() bool {
    return true
}


