/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="pdf_encryption_details_data.go">
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

// Container class for details of encryption.
type PdfEncryptionDetailsDataResult struct {
    // Container class for details of encryption.
    OwnerPassword string `json:"OwnerPassword,omitempty"`

    // Container class for details of encryption.
    Permissions []string `json:"Permissions,omitempty"`

    // Container class for details of encryption.
    UserPassword string `json:"UserPassword,omitempty"`
}

type PdfEncryptionDetailsData struct {
    // Container class for details of encryption.
    OwnerPassword *string `json:"OwnerPassword,omitempty"`

    // Container class for details of encryption.
    Permissions []string `json:"Permissions,omitempty"`

    // Container class for details of encryption.
    UserPassword *string `json:"UserPassword,omitempty"`
}

type IPdfEncryptionDetailsData interface {
    IsPdfEncryptionDetailsData() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileContent) []FileContent
}

func (PdfEncryptionDetailsData) IsPdfEncryptionDetailsData() bool {
    return true
}


func (obj *PdfEncryptionDetailsData) Initialize() {
}

func (obj *PdfEncryptionDetailsData) CollectFilesContent(resultFilesContent []FileContent) []FileContent {
    return resultFilesContent
}


