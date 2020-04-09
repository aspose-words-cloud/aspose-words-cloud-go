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

func TestDeleteFootnote(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Footnotes"), "Footnote.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Footnotes")
	remoteName := "TestDeleteFootnote.docx"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteFootnote(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteFootnoteWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Footnotes"), "Footnote.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Footnotes")
	remoteName := "TestDeleteFootnoteWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteFootnoteWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetFootnote(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Footnotes"), "Footnote.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Footnotes")
	remoteName := "TestGetFootnote.docx"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetFootnote(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetFootnoteWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Footnotes"), "Footnote.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Footnotes")
	remoteName := "TestGetFootnoteWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetFootnoteWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetFootnotes(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Footnotes"), "Footnote.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Footnotes")
	nodePath := "sections/0"
	remoteName := "TestGetFootnotes.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetFootnotes(ctx, remoteName, nodePath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetFootnotesWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Footnotes"), "Footnote.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Footnotes")
	remoteName := "TestGetFootnotesWithoutNodePath.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetFootnotesWithoutNodePath(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertFootnote(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Footnotes"), "Footnote.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Footnotes")
	remoteName := "TestInsertFootnote.docx"
	footnoteDto := models.FootnoteInsert{
		FootnoteType: "Endnote",
		Text:         "test endnote",
	}
	nodePath := "sections/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertFootnote(ctx, remoteName, footnoteDto, nodePath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertFootnoteWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Footnotes"), "Footnote.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Footnotes")
	remoteName := "TestInsertFootnoteWithoutNodePath.docx"
	footnoteDto := models.FootnoteInsert{
		FootnoteType: "Endnote",
		Text:         "test endnote",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertFootnoteWithoutNodePath(ctx, remoteName, footnoteDto, options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateFootnote(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Footnotes"), "Footnote.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Footnotes")
	remoteName := "TestUpdateFootnote.docx"
	nodePath := "sections/0"
	index := 0
	footnoteDto := models.FootnoteUpdate{
		Text: "new text is here",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateFootnote(ctx, remoteName, footnoteDto, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateFootnoteWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Footnotes"), "Footnote.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Footnotes")
	remoteName := "TestUpdateFootnoteWithoutNodePath.docx"
	index := 0
	footnoteDto := models.FootnoteUpdate{
		Text: "new text is here",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateFootnoteWithoutNodePath(ctx, remoteName, footnoteDto, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}
