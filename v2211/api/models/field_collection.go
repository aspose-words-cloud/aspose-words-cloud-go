/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="field_collection.go">
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

// DTO container with a collection of fields.
type FieldCollectionResult struct {
    // DTO container with a collection of fields.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // DTO container with a collection of fields.
    List []FieldResult `json:"List,omitempty"`
}

type FieldCollection struct {
    // DTO container with a collection of fields.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with a collection of fields.
    List []Field `json:"List,omitempty"`
}

type IFieldCollection interface {
    IsFieldCollection() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
}

func (FieldCollection) IsFieldCollection() bool {
    return true
}

func (FieldCollection) IsLinkElement() bool {
    return true
}

func (obj *FieldCollection) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.List != nil) {
        for _, objElementList := range obj.List {
            objElementList.Initialize()
        }
    }

}

func (obj *FieldCollection) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}


