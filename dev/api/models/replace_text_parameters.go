/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="replace_text_parameters.go">
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

// Class for document replace text request building.
type ReplaceTextParametersResult struct {
    // Class for document replace text request building.
    IsMatchCase bool `json:"IsMatchCase,omitempty"`

    // Class for document replace text request building.
    IsMatchWholeWord bool `json:"IsMatchWholeWord,omitempty"`

    // Class for document replace text request building.
    IsOldValueRegex bool `json:"IsOldValueRegex,omitempty"`

    // Class for document replace text request building.
    NewValue string `json:"NewValue,omitempty"`

    // Class for document replace text request building.
    OldValue string `json:"OldValue,omitempty"`
}

type ReplaceTextParameters struct {
    // Class for document replace text request building.
    IsMatchCase *bool `json:"IsMatchCase,omitempty"`

    // Class for document replace text request building.
    IsMatchWholeWord *bool `json:"IsMatchWholeWord,omitempty"`

    // Class for document replace text request building.
    IsOldValueRegex *bool `json:"IsOldValueRegex,omitempty"`

    // Class for document replace text request building.
    NewValue *string `json:"NewValue,omitempty"`

    // Class for document replace text request building.
    OldValue *string `json:"OldValue,omitempty"`
}

type IReplaceTextParameters interface {
    IsReplaceTextParameters() bool
}
func (ReplaceTextParameters) IsReplaceTextParameters() bool {
    return true
}


