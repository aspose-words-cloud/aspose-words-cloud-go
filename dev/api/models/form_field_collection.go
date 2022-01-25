/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="form_field_collection.go">
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

// DTO container with a collection of form fields.
type FormFieldCollectionResult struct {
    // DTO container with a collection of form fields.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // DTO container with a collection of form fields.
    List []FormFieldResult `json:"List,omitempty"`
}

type FormFieldCollection struct {
    // DTO container with a collection of form fields.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with a collection of form fields.
    List []FormField `json:"List,omitempty"`
}

type IFormFieldCollection interface {
    IsFormFieldCollection() bool
    Initialize()
}

func (FormFieldCollection) IsFormFieldCollection() bool {
    return true
}

func (FormFieldCollection) IsLinkElement() bool {
    return true
}

func (obj *FormFieldCollection) Initialize() {
}


