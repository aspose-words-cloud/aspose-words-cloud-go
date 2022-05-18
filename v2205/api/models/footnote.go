/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="footnote.go">
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

// DTO container with a footnote.
type FootnoteResult struct {
    // DTO container with a footnote.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // DTO container with a footnote.
    NodeId string `json:"NodeId,omitempty"`

    // DTO container with a footnote.
    Content StoryChildNodesResult `json:"Content,omitempty"`

    // DTO container with a footnote.
    FootnoteType string `json:"FootnoteType,omitempty"`

    // DTO container with a footnote.
    Position DocumentPositionResult `json:"Position,omitempty"`

    // DTO container with a footnote.
    ReferenceMark string `json:"ReferenceMark,omitempty"`

    // DTO container with a footnote.
    Text string `json:"Text,omitempty"`
}

type Footnote struct {
    // DTO container with a footnote.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with a footnote.
    NodeId *string `json:"NodeId,omitempty"`

    // DTO container with a footnote.
    Content IStoryChildNodes `json:"Content,omitempty"`

    // DTO container with a footnote.
    FootnoteType *string `json:"FootnoteType,omitempty"`

    // DTO container with a footnote.
    Position IDocumentPosition `json:"Position,omitempty"`

    // DTO container with a footnote.
    ReferenceMark *string `json:"ReferenceMark,omitempty"`

    // DTO container with a footnote.
    Text *string `json:"Text,omitempty"`
}

type IFootnote interface {
    IsFootnote() bool
    Initialize()
}

func (Footnote) IsFootnote() bool {
    return true
}

func (Footnote) IsFootnoteLink() bool {
    return true
}

func (Footnote) IsNodeLink() bool {
    return true
}

func (Footnote) IsLinkElement() bool {
    return true
}

func (obj *Footnote) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }





    if (obj.Content != nil) {
        obj.Content.Initialize()
    }





    if (obj.Position != nil) {
        obj.Position.Initialize()
    }





}


