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
	"github.com/aspose-words-cloud/aspose-words-cloud-go/v2007/api/models"
	"path"
    "path/filepath"
	"testing"
)

func TestGetStyles(t *testing.T) {
	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Styles"), "GetStyles.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, filepath.Join("DocumentElements", "Styles"))
	remoteName := "TestGetStyles.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetStyles(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetStyle(t *testing.T) {
	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Styles"), "GetStyles.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, filepath.Join("DocumentElements", "Styles"))
	remoteName := "TestGetStyle.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetStyle(ctx, remoteName, "Heading 1", options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateStyle(t *testing.T) {
	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Styles"), "GetStyles.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, filepath.Join("DocumentElements", "Styles"))
	remoteName := "TestUpdateStyle.docx"
	styleUpdate := models.StyleUpdate{Name:"My Style"}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateStyle(ctx, remoteName, styleUpdate, "Heading 1", options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertStyle(t *testing.T) {
	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Styles"), "GetStyles.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, filepath.Join("DocumentElements", "Styles"))
	remoteName := "TestInsertStyle.docx"
	styleInsert := models.StyleInsert{StyleName:"My Style", StyleType:"Paragraph"}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertStyle(ctx, remoteName, styleInsert, options)

	if err != nil {
		t.Error(err)
	}
}

func TestCopyStyle(t *testing.T) {
	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Styles"), "GetStyles.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, filepath.Join("DocumentElements", "Styles"))
	remoteName := "TestCopyStyle.docx"
	styleCopy := models.StyleCopy{StyleName:"Heading 1"}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.CopyStyle(ctx, remoteName, styleCopy, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetStyleFromDocumentElement(t *testing.T) {
	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Styles"), "GetStyles.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, filepath.Join("DocumentElements", "Styles"))
	remoteName := "TestGetStyleFromDocumentElement.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetStyleFromDocumentElement(ctx, remoteName, "paragraphs/1/paragraphFormat", options)

	if err != nil {
		t.Error(err)
	}
}

func TestApplyStyleToDocumentElement(t *testing.T) {
	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Styles"), "GetStyles.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, filepath.Join("DocumentElements", "Styles"))
	remoteName := "TestApplyStyleToDocumentElement.docx"
	styleApply := models.StyleApply{StyleName:"Heading 1"}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.ApplyStyleToDocumentElement(ctx, remoteName, styleApply, "paragraphs/1/paragraphFormat", options)

	if err != nil {
		t.Error(err)
	}
}
