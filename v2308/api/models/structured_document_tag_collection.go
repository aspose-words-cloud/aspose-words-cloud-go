/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="structured_document_tag_collection.go">
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

// DTO container with a collection of StructuredDocumentTags links.
type StructuredDocumentTagCollectionResult struct {
    // DTO container with a collection of StructuredDocumentTags links.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // DTO container with a collection of StructuredDocumentTags links.
    List []StructuredDocumentTagResult `json:"List,omitempty"`
}

type StructuredDocumentTagCollection struct {
    // DTO container with a collection of StructuredDocumentTags links.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with a collection of StructuredDocumentTags links.
    List []StructuredDocumentTag `json:"List,omitempty"`
}

type IStructuredDocumentTagCollection interface {
    IsStructuredDocumentTagCollection() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
}

func (StructuredDocumentTagCollection) IsStructuredDocumentTagCollection() bool {
    return true
}

func (StructuredDocumentTagCollection) IsLinkElement() bool {
    return true
}

func (obj *StructuredDocumentTagCollection) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.List != nil) {
        for _, objElementList := range obj.List {
            objElementList.Initialize()
        }
    }

}

func (obj *StructuredDocumentTagCollection) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}


