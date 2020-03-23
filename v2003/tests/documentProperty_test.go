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
	"testing"

	"github.com/aspose-words-cloud/aspose-words-cloud-go/v2003/api/models"
)

func TestCreateOrUpdateDocumentProperty(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := "TestOut"
	remoteName := "TestCreateOrUpdateDocumentProperty.docx"
	propertyName := "AsposeAuthor"
	property := models.DocumentProperty{
		Name:  "Author",
		Value: "Imran Anwar",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.CreateOrUpdateDocumentProperty(ctx, remoteName, propertyName, property, options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteDocumentProperty(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DocumentProperties")
	remoteName := "TestDeleteDocumentProperty.docx"
	propertyName := "testProp"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteDocumentProperty(ctx, remoteName, propertyName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentProperties(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DocumentProperties")
	remoteName := "TestGetDocumentProperties.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetDocumentProperties(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentProperty(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "DocumentProperties")
	remoteName := "TestGetDocumentProperty.docx"
	propertyName := "Author"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetDocumentProperty(ctx, remoteName, propertyName, options)

	if err != nil {
		t.Error(err)
	}
}
