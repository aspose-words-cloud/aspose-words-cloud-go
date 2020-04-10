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

func TestDeleteRun(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Runs"), "Run.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Runs")
	remoteName := "TestDeleteRun.docx"
	paragraphPath := "paragraphs/1"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteRun(ctx, remoteName, paragraphPath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetRun(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestGetRun.docx"
	paragraphPath := "paragraphs/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetRun(ctx, remoteName, paragraphPath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetRunFont(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestGetRunFont.docx"
	paragraphPath := "paragraphs/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetRunFont(ctx, remoteName, paragraphPath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetRuns(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestGetRuns.docx"
	paragraphPath := "paragraphs/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetRuns(ctx, remoteName, paragraphPath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertRun(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Runs"), "Run.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Runs")
	remoteName := "TestInsertRun.docx"
	paragraphPath := "paragraphs/1"
	run := models.RunInsert{
		Text: "run with text",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertRun(ctx, remoteName, paragraphPath, run, options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateRun(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Runs"), "Run.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Runs")
	remoteName := "TestUpdateRun.docx"
	paragraphPath := "paragraphs/1"
	index := 0
	run := models.RunUpdate{
		Text: "run with text",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateRun(ctx, remoteName, run, paragraphPath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateRunFont(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestUpdateRunFont.docx"
	paragraphPath := "paragraphs/0"
	index := 0
	fontDto := models.Font{
		Bold: true,
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateRunFont(ctx, remoteName, fontDto, paragraphPath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}
