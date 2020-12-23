/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="document.go">
 *   Copyright (c) 2020 Aspose.Words for Cloud
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

// Represents Words document DTO.
type DocumentResult struct {
    // Represents Words document DTO.
    DocumentProperties DocumentPropertiesResult `json:"DocumentProperties,omitempty"`

    // Represents Words document DTO.
    FileName string `json:"FileName,omitempty"`

    // Represents Words document DTO.
    IsEncrypted bool `json:"IsEncrypted,omitempty"`

    // Represents Words document DTO.
    IsSigned bool `json:"IsSigned,omitempty"`

    // Represents Words document DTO.
    Links []LinkResult `json:"Links,omitempty"`

    // Represents Words document DTO.
    SourceFormat string `json:"SourceFormat,omitempty"`
}

type Document struct {
    // Represents Words document DTO.
    DocumentProperties IDocumentProperties `json:"DocumentProperties,omitempty"`

    // Represents Words document DTO.
    FileName *string `json:"FileName,omitempty"`

    // Represents Words document DTO.
    IsEncrypted *bool `json:"IsEncrypted,omitempty"`

    // Represents Words document DTO.
    IsSigned *bool `json:"IsSigned,omitempty"`

    // Represents Words document DTO.
    Links []Link `json:"Links,omitempty"`

    // Represents Words document DTO.
    SourceFormat *string `json:"SourceFormat,omitempty"`
}

type IDocument interface {
    IsDocument() bool
}
func (Document) IsDocument() bool {
    return true
}


