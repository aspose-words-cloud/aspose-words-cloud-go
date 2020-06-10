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

	"github.com/aspose-words-cloud/aspose-words-cloud-go/v2006/api/models"
)

func TestDeleteSection(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Section")
	remoteName := "TestDeleteSection.docx"
	sectionIndex := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteSection(ctx, remoteName, int32(sectionIndex), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetHeaderFooterOfSection(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "HeaderFooters"), "HeadersFooters.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "HeaderFooters")
	remoteName := "TestGetHeaderFooterOfSection.docx"
	sectionIndex := 0
	headerFooterIndex := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetHeaderFooterOfSection(ctx, remoteName, int32(headerFooterIndex), int32(sectionIndex), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetSection(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Section")
	remoteName := "TestGetSection.docx"
	sectionIndex := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetSection(ctx, remoteName, int32(sectionIndex), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetSectionPageSetup(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "PageSetup")
	remoteName := "TestGetSectionPageSetup.docx"
	sectionIndex := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetSectionPageSetup(ctx, remoteName, int32(sectionIndex), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetSections(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Section")
	remoteName := "TestGetSections.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetSections(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateSectionPageSetup(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "PageSetup")
	remoteName := "TestUpdateSectionPageSetup.docx"
	sectionIndex := 0
	pageSetup := models.PageSetup{
		RtlGutter:   true,
		LeftMargin:  10.0,
		Orientation: "Landscape",
		PaperSize:   "A5",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateSectionPageSetup(ctx, remoteName, int32(sectionIndex), pageSetup, options)

	if err != nil {
		t.Error(err)
	}
}
