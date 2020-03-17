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

// Class for document replace text request building.
type ReplaceTextParameters struct {

	// Gets or sets old text value (or regex pattern ) to replace.
	OldValue string `json:"OldValue,omitempty"`

	// Gets or sets new text value to replace by.
	NewValue string `json:"NewValue,omitempty"`

	// Gets or sets a value indicating whether flag, true means the search is case-sensitive; false means the search is not case-sensitive.
	IsMatchCase bool `json:"IsMatchCase"`

	// Gets or sets a value indicating whether flag, means that only whole word matched are replaced.
	IsMatchWholeWord bool `json:"IsMatchWholeWord"`

	// Gets or sets a value indicating whether flag, means that  contains regex expression.
	IsOldValueRegex bool `json:"IsOldValueRegex"`
}
type IReplaceTextParameters interface {
	IsReplaceTextParameters() bool
}
func (ReplaceTextParameters) IsReplaceTextParameters() bool {
	return true;
}

