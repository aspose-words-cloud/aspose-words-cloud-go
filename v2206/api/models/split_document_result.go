/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="split_document_result.go">
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

// Result of splitting document.
type SplitDocumentResultResult struct {
    // Result of splitting document.
    Pages []FileLinkResult `json:"Pages,omitempty"`

    // Result of splitting document.
    SourceDocument FileLinkResult `json:"SourceDocument,omitempty"`

    // Result of splitting document.
    ZippedPages FileLinkResult `json:"ZippedPages,omitempty"`
}

type SplitDocumentResult struct {
    // Result of splitting document.
    Pages []FileLink `json:"Pages,omitempty"`

    // Result of splitting document.
    SourceDocument IFileLink `json:"SourceDocument,omitempty"`

    // Result of splitting document.
    ZippedPages IFileLink `json:"ZippedPages,omitempty"`
}

type ISplitDocumentResult interface {
    IsSplitDocumentResult() bool
    Initialize()
}

func (SplitDocumentResult) IsSplitDocumentResult() bool {
    return true
}


func (obj *SplitDocumentResult) Initialize() {
    if (obj.SourceDocument != nil) {
        obj.SourceDocument.Initialize()
    }



    if (obj.ZippedPages != nil) {
        obj.ZippedPages.Initialize()
    }

}


