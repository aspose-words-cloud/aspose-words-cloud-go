/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="file_link.go">
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

// Provides information for the file link.
type FileLinkResult struct {
    // Provides information for the file link.
    Href string `json:"Href,omitempty"`

    // Provides information for the file link.
    Rel string `json:"Rel,omitempty"`

    // Provides information for the file link.
    Title string `json:"Title,omitempty"`

    // Provides information for the file link.
    Type string `json:"Type,omitempty"`
}

type FileLink struct {
    // Provides information for the file link.
    Href *string `json:"Href,omitempty"`

    // Provides information for the file link.
    Rel *string `json:"Rel,omitempty"`

    // Provides information for the file link.
    Title *string `json:"Title,omitempty"`

    // Provides information for the file link.
    Type *string `json:"Type,omitempty"`
}

type IFileLink interface {
    IsFileLink() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileContent) []FileContent
}

func (FileLink) IsFileLink() bool {
    return true
}

func (FileLink) IsLink() bool {
    return true
}

func (obj *FileLink) Initialize() {
}

func (obj *FileLink) CollectFilesContent(resultFilesContent []FileContent) []FileContent {
    return resultFilesContent
}


