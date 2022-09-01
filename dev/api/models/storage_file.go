/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="storage_file.go">
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

// File or folder information.
type StorageFileResult struct {
    // File or folder information.
    IsFolder bool `json:"IsFolder,omitempty"`

    // File or folder information.
    ModifiedDate Time `json:"ModifiedDate,omitempty"`

    // File or folder information.
    Name string `json:"Name,omitempty"`

    // File or folder information.
    Path string `json:"Path,omitempty"`

    // File or folder information.
    Size int32 `json:"Size,omitempty"`
}

type StorageFile struct {
    // File or folder information.
    IsFolder *bool `json:"IsFolder,omitempty"`

    // File or folder information.
    ModifiedDate *Time `json:"ModifiedDate,omitempty"`

    // File or folder information.
    Name *string `json:"Name,omitempty"`

    // File or folder information.
    Path *string `json:"Path,omitempty"`

    // File or folder information.
    Size *int32 `json:"Size,omitempty"`
}

type IStorageFile interface {
    IsStorageFile() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileContent) []FileContent
}

func (StorageFile) IsStorageFile() bool {
    return true
}


func (obj *StorageFile) Initialize() {
}

func (obj *StorageFile) CollectFilesContent(resultFilesContent []FileContent) []FileContent {
    return resultFilesContent
}


