/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="compare_data.go">
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

// Container class for compare documents.
type CompareDataResult struct {
    // Container class for compare documents.
    Author string `json:"Author,omitempty"`

    // Container class for compare documents.
    CompareOptions CompareOptionsResult `json:"CompareOptions,omitempty"`

    // Container class for compare documents.
    ComparingWithDocument string `json:"ComparingWithDocument,omitempty"`

    // Container class for compare documents.
    DateTime Time `json:"DateTime,omitempty"`

    // Container class for compare documents.
    ResultDocumentFormat string `json:"ResultDocumentFormat,omitempty"`
}

type CompareData struct {
    // Container class for compare documents.
    Author *string `json:"Author,omitempty"`

    // Container class for compare documents.
    CompareOptions ICompareOptions `json:"CompareOptions,omitempty"`

    // Container class for compare documents.
    ComparingWithDocument *string `json:"ComparingWithDocument,omitempty"`

    // Container class for compare documents.
    DateTime *Time `json:"DateTime,omitempty"`

    // Container class for compare documents.
    ResultDocumentFormat *string `json:"ResultDocumentFormat,omitempty"`
}

type ICompareData interface {
    IsCompareData() bool
    Initialize()
}

func (CompareData) IsCompareData() bool {
    return true
}


func (obj *CompareData) Initialize() {
}


