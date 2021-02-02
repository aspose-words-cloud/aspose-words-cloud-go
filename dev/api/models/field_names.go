/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="field_names.go">
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

// Represents a collection of merge fields within a document.
type FieldNamesResult struct {
    // Represents a collection of merge fields within a document.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // Represents a collection of merge fields within a document.
    Names []string `json:"Names,omitempty"`
}

type FieldNames struct {
    // Represents a collection of merge fields within a document.
    Link IWordsApiLink `json:"Link,omitempty"`

    // Represents a collection of merge fields within a document.
    Names []string `json:"Names,omitempty"`
}

type IFieldNames interface {
    IsFieldNames() bool
}
func (FieldNames) IsFieldNames() bool {
    return true
}

func (FieldNames) IsLinkElement() bool {
    return true
}


