/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="save_result.go">
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

// Result of saving.
type SaveResultResult struct {
    // Result of saving.
    AdditionalItems []FileLinkResult `json:"AdditionalItems,omitempty"`

    // Result of saving.
    DestDocument FileLinkResult `json:"DestDocument,omitempty"`

    // Result of saving.
    SourceDocument FileLinkResult `json:"SourceDocument,omitempty"`
}

type SaveResult struct {
    // Result of saving.
    AdditionalItems []FileLink `json:"AdditionalItems,omitempty"`

    // Result of saving.
    DestDocument IFileLink `json:"DestDocument,omitempty"`

    // Result of saving.
    SourceDocument IFileLink `json:"SourceDocument,omitempty"`
}

type ISaveResult interface {
    IsSaveResult() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileContent) []FileContent
}

func (SaveResult) IsSaveResult() bool {
    return true
}


func (obj *SaveResult) Initialize() {
    if (obj.AdditionalItems != nil) {
        for _, element := range obj.AdditionalItems {
            element.Initialize()
        }
    }

    if (obj.DestDocument != nil) {
        obj.DestDocument.Initialize()
    }

    if (obj.SourceDocument != nil) {
        obj.SourceDocument.Initialize()
    }


}

func (obj *SaveResult) CollectFilesContent(resultFilesContent []FileContent) []FileContent {
    if (obj.AdditionalItems != nil) {
        for _, element := range obj.AdditionalItems {
            resultFilesContent = element.CollectFilesContent(resultFilesContent)
        }
    }

    if (obj.DestDocument != nil) {
        resultFilesContent = obj.DestDocument.CollectFilesContent(resultFilesContent)
    }

    if (obj.SourceDocument != nil) {
        resultFilesContent = obj.SourceDocument.CollectFilesContent(resultFilesContent)
    }

    return resultFilesContent
}


