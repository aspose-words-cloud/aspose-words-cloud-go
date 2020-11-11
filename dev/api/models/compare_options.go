/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="compare_options.go">
 *   Copyright (c) 2020 Aspose.Words for Cloud
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

// Container class for compare documents options.
type CompareOptions struct {
    // Container class for compare documents options.
    IgnoreCaseChanges *bool `json:"IgnoreCaseChanges,omitempty"`

    // Container class for compare documents options.
    IgnoreComments *bool `json:"IgnoreComments,omitempty"`

    // Container class for compare documents options.
    IgnoreFields *bool `json:"IgnoreFields,omitempty"`

    // Container class for compare documents options.
    IgnoreFootnotes *bool `json:"IgnoreFootnotes,omitempty"`

    // Container class for compare documents options.
    IgnoreFormatting *bool `json:"IgnoreFormatting,omitempty"`

    // Container class for compare documents options.
    IgnoreHeadersAndFooters *bool `json:"IgnoreHeadersAndFooters,omitempty"`

    // Container class for compare documents options.
    IgnoreTables *bool `json:"IgnoreTables,omitempty"`

    // Container class for compare documents options.
    IgnoreTextboxes *bool `json:"IgnoreTextboxes,omitempty"`

    // Container class for compare documents options.
    Target string `json:"Target,omitempty"`
}

type ICompareOptions interface {
    IsCompareOptions() bool
}
func (CompareOptions) IsCompareOptions() bool {
    return true
}

