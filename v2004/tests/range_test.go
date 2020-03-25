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

	"github.com/aspose-words-cloud/aspose-words-cloud-go/v2004/api/models"
)

func TestGetRangeText(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Range"), "RangeGet.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Range")
	remoteName := "TestGetRangeText.docx"
	rangeStartIdentifier := "id0.0.0"
	rangeEndIdentifier := "id0.0.1"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetRangeText(ctx, remoteName, rangeStartIdentifier, rangeEndIdentifier, options)

	if err != nil {
		t.Error(err)
	}
}

func TestRemoveRange(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Range"), "RangeGet.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Range")
	remoteName := "TestRemoveRange.docx"
	rangeStartIdentifier := "id0.0.0"
	rangeEndIdentifier := "id0.0.1"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.RemoveRange(ctx, remoteName, rangeStartIdentifier, rangeEndIdentifier, options)

	if err != nil {
		t.Error(err)
	}
}

func TestSaveAsRange(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Range"), "RangeGet.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Range")
	remoteName := "TestSaveAsRange.docx"
	rangeStartIdentifier := "id0.0.0"
	rangeEndIdentifier := "id0.0.1"
	newDocName := "NewDoc.docx"
	documentParameters := models.RangeDocument{
		DocumentName: path.Join(remoteFolder, newDocName),
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.SaveAsRange(ctx, remoteName, rangeStartIdentifier, documentParameters, rangeEndIdentifier, options)

	if err != nil {
		t.Error(err)
	}
}

func TestReplaceWithText(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Range"), "RangeGet.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Range")
	remoteName := "TestReplaceWithText.docx"
	rangeStartIdentifier := "id0.0.0"
	rangeEndIdentifier := "id0.0.1"
	rangeText := models.ReplaceRange{
		Text: "Replaced header",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.ReplaceWithText(ctx, remoteName, rangeStartIdentifier, rangeText, rangeEndIdentifier, options)

	if err != nil {
		t.Error(err)
	}
}
