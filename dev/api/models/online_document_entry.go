/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="online_document_entry.go">
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

// Represents a document which will be appended to the original resource document.
type OnlineDocumentEntryResult struct {
    // Represents a document which will be appended to the original resource document.
    EncryptedPassword string `json:"EncryptedPassword,omitempty"`

    // Represents a document which will be appended to the original resource document.
    ImportFormatMode string `json:"ImportFormatMode,omitempty"`

    // Represents a document which will be appended to the original resource document.
    File FileContentResult `json:"File,omitempty"`
}

type OnlineDocumentEntry struct {
    // Represents a document which will be appended to the original resource document.
    EncryptedPassword *string `json:"EncryptedPassword,omitempty"`

    // Represents a document which will be appended to the original resource document.
    ImportFormatMode *string `json:"ImportFormatMode,omitempty"`

    // Represents a document which will be appended to the original resource document.
    File IFileContent `json:"File,omitempty"`
}

type IOnlineDocumentEntry interface {
    IsOnlineDocumentEntry() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileContent) []FileContent
}

func (OnlineDocumentEntry) IsOnlineDocumentEntry() bool {
    return true
}

func (OnlineDocumentEntry) IsBaseDocumentEntry() bool {
    return true
}

func (obj *OnlineDocumentEntry) Initialize() {
    if (obj.File != nil) {
        obj.File.Initialize()
    }


}

func (obj *OnlineDocumentEntry) CollectFilesContent(resultFilesContent []FileContent) []FileContent {
    if (obj.File != nil) {
        resultFilesContent = obj.File.CollectFilesContent(resultFilesContent)
    }

    return resultFilesContent
}


