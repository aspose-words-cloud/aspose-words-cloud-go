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
type FormFieldCheckboxResult struct {
    // FormField checkbox element.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // FormField checkbox element.
    NodeId string `json:"NodeId,omitempty"`

    // FormField checkbox element.
    CalculateOnExit bool `json:"CalculateOnExit,omitempty"`

    // FormField checkbox element.
    Enabled bool `json:"Enabled,omitempty"`

    // FormField checkbox element.
    EntryMacro string `json:"EntryMacro,omitempty"`

    // FormField checkbox element.
    ExitMacro string `json:"ExitMacro,omitempty"`

    // FormField checkbox element.
    HelpText string `json:"HelpText,omitempty"`

    // FormField checkbox element.
    Name string `json:"Name,omitempty"`

    // FormField checkbox element.
    OwnHelp bool `json:"OwnHelp,omitempty"`

    // FormField checkbox element.
    OwnStatus bool `json:"OwnStatus,omitempty"`

    // FormField checkbox element.
    StatusText string `json:"StatusText,omitempty"`

    // FormField checkbox element.
    CheckBoxSize float64 `json:"CheckBoxSize,omitempty"`

    // FormField checkbox element.
    Checked bool `json:"Checked,omitempty"`

    // FormField checkbox element.
    IsCheckBoxExactSize bool `json:"IsCheckBoxExactSize,omitempty"`
}

type FormFieldCheckbox struct {
    // FormField checkbox element.
    Link IWordsApiLink `json:"Link,omitempty"`

    // FormField checkbox element.
    NodeId *string `json:"NodeId,omitempty"`

    // FormField checkbox element.
    CalculateOnExit *bool `json:"CalculateOnExit,omitempty"`

    // FormField checkbox element.
    Enabled *bool `json:"Enabled,omitempty"`

    // FormField checkbox element.
    EntryMacro *string `json:"EntryMacro,omitempty"`

    // FormField checkbox element.
    ExitMacro *string `json:"ExitMacro,omitempty"`

    // FormField checkbox element.
    HelpText *string `json:"HelpText,omitempty"`

    // FormField checkbox element.
    Name *string `json:"Name,omitempty"`

    // FormField checkbox element.
    OwnHelp *bool `json:"OwnHelp,omitempty"`

    // FormField checkbox element.
    OwnStatus *bool `json:"OwnStatus,omitempty"`

    // FormField checkbox element.
    StatusText *string `json:"StatusText,omitempty"`

    // FormField checkbox element.
    CheckBoxSize *float64 `json:"CheckBoxSize,omitempty"`

    // FormField checkbox element.
    Checked *bool `json:"Checked,omitempty"`

    // FormField checkbox element.
    IsCheckBoxExactSize *bool `json:"IsCheckBoxExactSize,omitempty"`
}

type IFormFieldCheckbox interface {
    IsFormFieldCheckbox() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
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

func (obj *FormFieldCheckbox) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}


