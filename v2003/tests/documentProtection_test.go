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
	"path"
	"path/filepath"
	"testing"

	"github.com/aspose-words-cloud/aspose-words-cloud-go/v2003/api/models"
)

func TestProtectDocument(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "DocumentProtection")
	remoteName := "TestProtectDocument.docx"
	protectionRequest := models.ProtectionRequest{
		NewPassword: "123",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.ProtectDocument(ctx, remoteName, protectionRequest, options)

	if err != nil {
		t.Error(err)
	}
}

func TestUnprotectDocument(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentActions", "DocumentProtection"), "SampleProtectedBlankWordDocument.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "DocumentProtection")
	remoteName := "TestUnprotectDocument.docx"
	protectionRequest := models.ProtectionRequest{
		Password: "aspose",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UnprotectDocument(ctx, remoteName, protectionRequest, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentProtection(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "DocumentProtection")
	remoteName := "TestGetDocumentProtection.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetDocumentProtection(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}
