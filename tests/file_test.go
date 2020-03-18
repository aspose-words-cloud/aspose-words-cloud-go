//
// MIT License

// Copyright (c) 2019 Aspose Pty Ltd

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

package api_test

import (
	"fmt"
	"path"
	"testing"

	"github.com/google/uuid"
)

func TestCopyFile(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "Storage")
	remoteSrcName := "TestCopyFile.docx"
	remoteDstName := fmt.Sprintf("TestCopyFile%s.docx", uuid.New().String())
	remoteSrcPath := path.Join(remoteFolder, remoteSrcName)
	remoteDstPath := path.Join(remoteFolder, remoteDstName)

	client, ctx := UploadFileToStorage(t, localFilePath, remoteSrcPath)

	_, err := client.WordsApi.CopyFile(ctx, remoteDstPath, remoteSrcPath, nil)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteFile(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "Storage")
	remoteName := "TestDeleteFile.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteFile(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestDownloadFile(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "Storage")
	remoteName := "TestDownloadFile.docx"
	fullName := path.Join(remoteFolder, remoteName)
	options := map[string]interface{}{}

	client, ctx := UploadFileToStorage(t, localFilePath, fullName)

	output, err := client.WordsApi.DownloadFile(ctx, fullName, options)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestMoveFile(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "Storage")
	remoteSrcName := "TestMoveFile.docx"
	remoteDstName := fmt.Sprintf("TesMoveFile%s.docx", uuid.New().String())
	remoteSrcPath := path.Join(remoteFolder, remoteSrcName)
	remoteDstPath := path.Join(remoteFolder, remoteDstName)

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteSrcName))

	_, err := client.WordsApi.MoveFile(ctx, remoteDstPath, remoteSrcPath, nil)

	if err != nil {
		t.Error(err)
	}
}
