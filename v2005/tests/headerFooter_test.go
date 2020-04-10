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

	"github.com/aspose-words-cloud/aspose-words-cloud-go/v2005/api/models"
)

func TestDeleteHeaderFooter(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "HeaderFooters"), "HeadersFooters.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "HeaderFooters")
	remoteName := "TestDeleteHeaderFooter.docx"
	sectionPath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteHeaderFooter(ctx, remoteName, sectionPath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteHeadersFooters(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "HeaderFooters"), "HeadersFooters.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "HeaderFooters")
	remoteName := "TestDeleteHeadersFooters.docx"
	sectionPath := "sections/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteHeadersFooters(ctx, remoteName, sectionPath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetHeaderFooter(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "HeaderFooters"), "HeadersFooters.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "HeaderFooters")
	remoteName := "TestGetHeaderFooter.docx"
	headerFooterIndex := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetHeaderFooter(ctx, remoteName, int32(headerFooterIndex), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetHeaderFooters(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "HeaderFooters"), "HeadersFooters.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "HeaderFooters")
	remoteName := "TestGetHeaderFooters.docx"
	sectionPath := "sections/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetHeaderFooters(ctx, remoteName, sectionPath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertHeaderFooter(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "HeaderFooters"), "HeadersFooters.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "HeaderFooters")
	remoteName := "TestInsertHeaderFooter.docx"
	headerFooterType := "FooterEven"
	sectionPath := "sections/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertHeaderFooter(ctx, remoteName, headerFooterType, sectionPath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertPageNumbers(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Fields")
	remoteName := "TestInsertPageNumbers.docx"
	pageNumber := models.PageNumber{
		Alignment: "center",
		Format:    "{PAGE} of {NUMPAGES}",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertPageNumbers(ctx, remoteName, pageNumber, options)

	if err != nil {
		t.Error(err)
	}
}
