/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="file_reference.go">
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

import (
	"crypto/rand"
	"fmt"
	"io"
)

type FileReferenceResult struct {
    // File source.
    Source string `json:"Source"`

    // File reference.
    Reference string `json:"Reference"`

    // File local data.
    Content io.ReadCloser `json:"-"`
}

type FileReference struct {
    // File source.
    Source string `json:"Source"`

    // File reference.
    Reference string `json:"Reference"`

    // File local data.
    Content io.ReadCloser `json:"-"`
}

type IFileReference interface {
    IsFileReference() bool
    Initialize()
    CollectFilesContent(resultFilesContent []FileReference) []FileReference
}

func (FileReference) IsFileReference() bool {
    return true
}


func (obj *FileReference) Initialize() {

}

func (obj *FileReference) CollectFilesContent(resultFilesContent []FileReference) []FileReference {
    if obj.Source == "Request" {
        return append(resultFilesContent, *obj)
    } else {
        return resultFilesContent;
    }
}

func createRandomFileReferenceId() string {
	b := make([]byte, 16)
	rand.Read(b)
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

func CreateRemoteFileReference(remoteFilePath string) FileReference {
    return FileReference {
        Source: "Storage",
        Reference: remoteFilePath,
        Content: nil,
    }
}

func CreateLocalFileReference(localFileContent io.ReadCloser) FileReference {
    return FileReference {
        Source: "Request",
        Reference: createRandomFileReferenceId(),
        Content: localFileContent,
    }
}
