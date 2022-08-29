/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="file_content.go">
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

import (
	"crypto/rand"
	"fmt"
	"io"
)

type FileContentResult struct {
    // File name in multipart.
    Id string `json:"Id"`

    // File name.
    Filename string `json:"Filename"`

    // File data.
    Content io.ReadCloser `json:"-"`
}

type FileContent struct {
    // File name in multipart.
    Id string `json:"Id"`

    // File name.
    Filename string `json:"Filename"`

    // File data.
    Content io.ReadCloser `json:"-"`
}

type IFileContent interface {
    IsFileContent() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileContent) []FileContent
}

func (FileContent) IsFileContent() bool {
    return true
}


func (obj *FileContent) Initialize() {
    obj.Id = createRandomFileContentId()
}

func (obj *FileContent) CollectFilesContent(resultFilesContent []FileContent) []FileContent {
    return append(resultFilesContent, *obj)
}

func createRandomFileContentId() string {
	b := make([]byte, 16)
	rand.Read(b)
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}
