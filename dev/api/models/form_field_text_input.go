/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="form_field_text_input.go">
 *   Copyright (c) 2021 Aspose.Words for Cloud
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

// FormField text input element.
type FormFieldTextInputResult struct {
    // FormField text input element.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // FormField text input element.
    NodeId string `json:"NodeId,omitempty"`

    // FormField text input element.
    CalculateOnExit bool `json:"CalculateOnExit,omitempty"`

    // FormField text input element.
    Enabled bool `json:"Enabled,omitempty"`

    // FormField text input element.
    EntryMacro string `json:"EntryMacro,omitempty"`

    // FormField text input element.
    ExitMacro string `json:"ExitMacro,omitempty"`

    // FormField text input element.
    HelpText string `json:"HelpText,omitempty"`

    // FormField text input element.
    Name string `json:"Name,omitempty"`

    // FormField text input element.
    OwnHelp bool `json:"OwnHelp,omitempty"`

    // FormField text input element.
    OwnStatus bool `json:"OwnStatus,omitempty"`

    // FormField text input element.
    StatusText string `json:"StatusText,omitempty"`

    // FormField text input element.
    MaxLength int32 `json:"MaxLength,omitempty"`

    // FormField text input element.
    TextInputDefault string `json:"TextInputDefault,omitempty"`

    // FormField text input element.
    TextInputFormat string `json:"TextInputFormat,omitempty"`

    // FormField text input element.
    TextInputType string `json:"TextInputType,omitempty"`
}

type FormFieldTextInput struct {
    // FormField text input element.
    Link IWordsApiLink `json:"Link,omitempty"`

    // FormField text input element.
    NodeId *string `json:"NodeId,omitempty"`

    // FormField text input element.
    CalculateOnExit *bool `json:"CalculateOnExit,omitempty"`

    // FormField text input element.
    Enabled *bool `json:"Enabled,omitempty"`

    // FormField text input element.
    EntryMacro *string `json:"EntryMacro,omitempty"`

    // FormField text input element.
    ExitMacro *string `json:"ExitMacro,omitempty"`

    // FormField text input element.
    HelpText *string `json:"HelpText,omitempty"`

    // FormField text input element.
    Name *string `json:"Name,omitempty"`

    // FormField text input element.
    OwnHelp *bool `json:"OwnHelp,omitempty"`

    // FormField text input element.
    OwnStatus *bool `json:"OwnStatus,omitempty"`

    // FormField text input element.
    StatusText *string `json:"StatusText,omitempty"`

    // FormField text input element.
    MaxLength *int32 `json:"MaxLength,omitempty"`

    // FormField text input element.
    TextInputDefault *string `json:"TextInputDefault,omitempty"`

    // FormField text input element.
    TextInputFormat *string `json:"TextInputFormat,omitempty"`

    // FormField text input element.
    TextInputType *string `json:"TextInputType,omitempty"`
}

type IFormFieldTextInput interface {
    IsFormFieldTextInput() bool
}
func (FormFieldTextInput) IsFormFieldTextInput() bool {
    return true
}

func (FormFieldTextInput) IsFormField() bool {
    return true
}


