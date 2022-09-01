/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="footnote_collection.go">
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

// DTO container with a collection of footnotes.
type FootnoteCollectionResult struct {
    // DTO container with a collection of footnotes.
    Link WordsApiLinkResult `json:"Link,omitempty"`

    // DTO container with a collection of footnotes.
    List []FootnoteResult `json:"List,omitempty"`
}

type FootnoteCollection struct {
    // DTO container with a collection of footnotes.
    Link IWordsApiLink `json:"Link,omitempty"`

    // DTO container with a collection of footnotes.
    List []Footnote `json:"List,omitempty"`
}

type IFootnoteCollection interface {
    IsFootnoteCollection() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileContent) []FileContent
}

func (FootnoteCollection) IsFootnoteCollection() bool {
    return true
}

func (FootnoteCollection) IsLinkElement() bool {
    return true
}

func (obj *FootnoteCollection) Initialize() {
    if (obj.Link != nil) {
        obj.Link.Initialize()
    }

    if (obj.List != nil) {
        for _, element := range obj.List {
            element.Initialize()
        }
    }


}

func (obj *FootnoteCollection) CollectFilesContent(resultFilesContent []FileContent) []FileContent {
    if (obj.Link != nil) {
        resultFilesContent = obj.Link.CollectFilesContent(resultFilesContent)
    }

    if (obj.List != nil) {
        for _, element := range obj.List {
            resultFilesContent = element.CollectFilesContent(resultFilesContent)
        }
    }

    return resultFilesContent
}


