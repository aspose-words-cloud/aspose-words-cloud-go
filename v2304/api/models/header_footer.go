/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="header_footer.go">
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

// DTO container with a section element.
type HeaderFooterResult struct {
    // DTO container with a section element.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // DTO container with a section element.
    Type string `json:"Type,omitempty"`

    // DTO container with a section element.
    ChildNodes []NodeLinkResult `json:"ChildNodes,omitempty"`

    // DTO container with a section element.
    DrawingObjects LinkElementResult `json:"DrawingObjects,omitempty"`

    // DTO container with a section element.
    Paragraphs LinkElementResult `json:"Paragraphs,omitempty"`
}

type HeaderFooter struct {
    // DTO container with a section element.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with a section element.
    Type *string `json:"Type,omitempty"`

    // DTO container with a section element.
    ChildNodes []NodeLink `json:"ChildNodes,omitempty"`

    // DTO container with a section element.
    DrawingObjects ILinkElement `json:"DrawingObjects,omitempty"`

    // DTO container with a section element.
    Paragraphs ILinkElement `json:"Paragraphs,omitempty"`
}

type IHeaderFooter interface {
    IsHeaderFooter() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
}

func (HeaderFooter) IsHeaderFooter() bool {
    return true
}

func (HeaderFooter) IsHeaderFooterLink() bool {
    return true
}

func (HeaderFooter) IsLinkElement() bool {
    return true
}

func (obj *HeaderFooter) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.ChildNodes != nil) {
        for _, objElementChildNodes := range obj.ChildNodes {
            objElementChildNodes.Initialize()
        }
    }
    if (obj.DrawingObjects != nil) {
        obj.DrawingObjects.Initialize()
    }

    if (obj.Paragraphs != nil) {
        obj.Paragraphs.Initialize()
    }


}

func (obj *HeaderFooter) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}


