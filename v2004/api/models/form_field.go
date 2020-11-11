/*
* MIT License

* Copyright (c) 2019 Aspose Pty Ltd

* Permission is hereby granted, free of charge, to any person obtaining a copy
* of this software and associated documentation files (the "Software"), to deal
* in the Software without restriction, including without limitation the rights
* to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
* copies of the Software, and to permit persons to whom the Software is
* furnished to do so, subject to the following conditions:

* The above copyright notice and this permission notice shall be included in all
* copies or substantial portions of the Software.

* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
* IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
* FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
* AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
* LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
* OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
* SOFTWARE.
*/


package models



// FromField.
type FormField struct {

	// Gets or sets node id.
	NodeId string `json:"NodeId,omitempty"`

	Link *WordsApiLink `json:"link,omitempty"`

	// Gets or sets true if references to the specified form field are automatically updated whenever the field is exited.
	CalculateOnExit bool `json:"CalculateOnExit,omitempty"`

	// Gets or sets true if a form field is enabled.
	Enabled bool `json:"Enabled,omitempty"`

	// Gets or sets returns or sets an entry macro name for the form field.
	EntryMacro string `json:"EntryMacro,omitempty"`

	// Gets or sets returns or sets an exit macro name for the form field.
	ExitMacro string `json:"ExitMacro,omitempty"`

	// Gets or sets returns or sets the text that's displayed in a message box when the form field has the focus and the user presses F1.
	HelpText string `json:"HelpText,omitempty"`

	// Gets or sets the form field name.
	Name string `json:"Name,omitempty"`

	// Gets or sets specifies the source of the text that's displayed in a message box when a form field has the focus and the user presses F1.
	OwnHelp bool `json:"OwnHelp,omitempty"`

	// Gets or sets specifies the source of the text that's displayed in the status bar when a form field has the focus.
	OwnStatus bool `json:"OwnStatus,omitempty"`

	// Gets or sets returns or sets the text that's displayed in the status bar when a form field has the focus.
	StatusText string `json:"StatusText,omitempty"`
}

type IFormField interface {
	IsFormField() bool
}
func (FormField) IsFormField() bool {
	return true;
}
func (FormField) IsNodeLink() bool {
	return true;
}
