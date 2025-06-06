/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="footnote_link.go">
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

// Footnote link.
type FootnoteLinkResult struct {
    // Footnote link.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // Footnote link.
    NodeId string `json:"NodeId,omitempty"`
}

type FootnoteLink struct {
    // Footnote link.
    Link IWordsApiLink `json:"Link,omitempty"`

    // Footnote link.
    NodeId *string `json:"NodeId,omitempty"`
}

type IFootnoteLink interface {
    IsFootnoteLink() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
}

func (FootnoteLink) IsFootnoteLink() bool {
    return true
}

func (FootnoteLink) IsNodeLink() bool {
    return true
}

func (FootnoteLink) IsLinkElement() bool {
    return true
}

func (obj *FootnoteLink) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }


}

func (obj *FootnoteLink) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}


