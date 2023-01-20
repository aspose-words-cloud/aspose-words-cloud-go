/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="compare_options.go">
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

// DTO container with compare documents options.
type CompareOptionsResult struct {
    // DTO container with compare documents options.
    AcceptAllRevisionsBeforeComparison bool `json:"AcceptAllRevisionsBeforeComparison,omitempty"`

    // DTO container with compare documents options.
    IgnoreCaseChanges bool `json:"IgnoreCaseChanges,omitempty"`

    // DTO container with compare documents options.
    IgnoreComments bool `json:"IgnoreComments,omitempty"`

    // DTO container with compare documents options.
    IgnoreFields bool `json:"IgnoreFields,omitempty"`

    // DTO container with compare documents options.
    IgnoreFootnotes bool `json:"IgnoreFootnotes,omitempty"`

    // DTO container with compare documents options.
    IgnoreFormatting bool `json:"IgnoreFormatting,omitempty"`

    // DTO container with compare documents options.
    IgnoreHeadersAndFooters bool `json:"IgnoreHeadersAndFooters,omitempty"`

    // DTO container with compare documents options.
    IgnoreTables bool `json:"IgnoreTables,omitempty"`

    // DTO container with compare documents options.
    IgnoreTextboxes bool `json:"IgnoreTextboxes,omitempty"`

    // DTO container with compare documents options.
    Target string `json:"Target,omitempty"`
}

type CompareOptions struct {
    // DTO container with compare documents options.
    AcceptAllRevisionsBeforeComparison *bool `json:"AcceptAllRevisionsBeforeComparison,omitempty"`

    // DTO container with compare documents options.
    IgnoreCaseChanges *bool `json:"IgnoreCaseChanges,omitempty"`

    // DTO container with compare documents options.
    IgnoreComments *bool `json:"IgnoreComments,omitempty"`

    // DTO container with compare documents options.
    IgnoreFields *bool `json:"IgnoreFields,omitempty"`

    // DTO container with compare documents options.
    IgnoreFootnotes *bool `json:"IgnoreFootnotes,omitempty"`

    // DTO container with compare documents options.
    IgnoreFormatting *bool `json:"IgnoreFormatting,omitempty"`

    // DTO container with compare documents options.
    IgnoreHeadersAndFooters *bool `json:"IgnoreHeadersAndFooters,omitempty"`

    // DTO container with compare documents options.
    IgnoreTables *bool `json:"IgnoreTables,omitempty"`

    // DTO container with compare documents options.
    IgnoreTextboxes *bool `json:"IgnoreTextboxes,omitempty"`

    // DTO container with compare documents options.
    Target *string `json:"Target,omitempty"`
}

type ICompareOptions interface {
    IsCompareOptions() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
}

func (CompareOptions) IsCompareOptions() bool {
    return true
}


func (obj *CompareOptions) Initialize() {
}

func (obj *CompareOptions) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}


