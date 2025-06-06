/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="document_stat_data.go">
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

// Container for the document's statistical data.
type DocumentStatDataResult struct {
    // Container for the document's statistical data.
    FootnotesStatData FootnotesStatDataResult `json:"FootnotesStatData,omitempty"`

    // Container for the document's statistical data.
    PageCount int32 `json:"PageCount,omitempty"`

    // Container for the document's statistical data.
    PageStatData []PageStatDataResult `json:"PageStatData,omitempty"`

    // Container for the document's statistical data.
    ParagraphCount int32 `json:"ParagraphCount,omitempty"`

    // Container for the document's statistical data.
    WordCount int32 `json:"WordCount,omitempty"`
}

type DocumentStatData struct {
    // Container for the document's statistical data.
    FootnotesStatData IFootnotesStatData `json:"FootnotesStatData,omitempty"`

    // Container for the document's statistical data.
    PageCount *int32 `json:"PageCount,omitempty"`

    // Container for the document's statistical data.
    PageStatData []PageStatData `json:"PageStatData,omitempty"`

    // Container for the document's statistical data.
    ParagraphCount *int32 `json:"ParagraphCount,omitempty"`

    // Container for the document's statistical data.
    WordCount *int32 `json:"WordCount,omitempty"`
}

type IDocumentStatData interface {
    IsDocumentStatData() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
}

func (DocumentStatData) IsDocumentStatData() bool {
    return true
}


func (obj *DocumentStatData) Initialize() {
    if (obj.FootnotesStatData != nil) {
        obj.FootnotesStatData.Initialize()
    }

    if (obj.PageStatData != nil) {
        for _, objElementPageStatData := range obj.PageStatData {
            objElementPageStatData.Initialize()
        }
    }

}

func (obj *DocumentStatData) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    return resultFilesContent
}


