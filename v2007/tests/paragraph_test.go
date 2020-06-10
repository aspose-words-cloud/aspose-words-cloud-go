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

	"github.com/aspose-words-cloud/aspose-words-cloud-go/v2006/api/models"
)

func TestDeleteParagraph(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestDeleteParagraph.docx"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteParagraph(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteParagraphWithoutNodePath(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestDeleteParagraphWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteParagraphWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetParagraph(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestGetParagraph.docx"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetParagraph(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetParagraphFormat(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestGetParagraphFormat.docx"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetParagraphFormat(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetParagraphFormatWithoutNodePath(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestGetParagraphFormatWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetParagraphFormatWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetParagraphWithoutNodePath(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestGetParagraphWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetParagraphWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetParagraphs(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestGetParagraphs.docx"
	nodePath := "sections/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetParagraphs(ctx, remoteName, nodePath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetParagraphsWithoutNodePath(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestGetParagraphsWithoutNodePath.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetParagraphsWithoutNodePath(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertParagraph(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestInsertParagraph.docx"
	nodePath := "sections/0"
	paragraph := models.ParagraphInsert{
		Text: "This is a new paragraph for your document",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertParagraph(ctx, remoteName, paragraph, nodePath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertParagraphWithoutNodePath(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestInsertParagraph.docx"
	paragraph := models.ParagraphInsert{
		Text: "This is a new paragraph for your document",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertParagraphWithoutNodePath(ctx, remoteName, paragraph, options)

	if err != nil {
		t.Error(err)
	}
}

func TestRenderParagraph(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestRenderParagraph.docx"
	nodePath := "sections/0"
	format := "png"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	output, err := client.WordsApi.RenderParagraph(ctx, remoteName, format, nodePath, int32(index), options)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestRenderParagraphWithoutNodePath(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestRenderParagraphWithoutNodePath.docx"
	format := "png"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	output, err := client.WordsApi.RenderParagraphWithoutNodePath(ctx, remoteName, format, int32(index), options)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateParagraphFormat(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestUpdateParagraphFormat.docx"
	nodePath := "sections/0"
	index := 0
	dto := models.ParagraphFormat{
		Alignment: "Right",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateParagraphFormat(ctx, remoteName, dto, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateParagraphFormatWithoutNodePath(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Paragraphs")
	remoteName := "TestUpdateParagraphFormat.docx"
	index := 0
	dto := models.ParagraphFormat{
		Alignment: "Right",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateParagraphFormatWithoutNodePath(ctx, remoteName, dto, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetParagraphListFormat(t *testing.T) {
	filename := "ParagraphGetListFormat.doc"
	baseDirPath := path.Join("DocumentElements", "ParagraphListFormat")
	localFilePath := GetLocalPath(baseDirPath, filename)
	remoteFolder := path.Join(remoteBaseTestDataFolder, baseDirPath)
	nodePath := ""
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, filename))

	_, _, err := client.WordsApi.GetParagraphListFormat(ctx, filename, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetParagraphListFormatWithoutNodePath(t *testing.T) {
	filename := "ParagraphGetListFormat.doc"
	baseDirPath := path.Join("DocumentElements", "ParagraphListFormat")
	localFilePath := GetLocalPath(baseDirPath, filename)
	remoteFolder := path.Join(remoteBaseTestDataFolder, baseDirPath)
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, filename))

	_, _, err := client.WordsApi.GetParagraphListFormatWithoutNodePath(ctx, filename, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateParagraphListFormat(t *testing.T) {
	filename := "ParagraphUpdateListFormat.doc"
	baseDirPath := path.Join("DocumentElements", "ParagraphListFormat")
	localFilePath := GetLocalPath(baseDirPath, filename)
	remoteFolder := path.Join(remoteBaseTestDataFolder, baseDirPath)
	nodePath := ""
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}
	dto := models.ListFormatUpdate{ListId: 2}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, filename))

	_, _, err := client.WordsApi.UpdateParagraphListFormat(ctx, filename, dto, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateParagraphListFormatWithoutNodePath(t *testing.T) {
	filename := "ParagraphUpdateListFormat.doc"
	baseDirPath := path.Join("DocumentElements", "ParagraphListFormat")
	localFilePath := GetLocalPath(baseDirPath, filename)
	remoteFolder := path.Join(remoteBaseTestDataFolder, baseDirPath)
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}
	dto := models.ListFormatUpdate{ListId: 2}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, filename))

	_, _, err := client.WordsApi.UpdateParagraphListFormatWithoutNodePath(ctx, filename, dto, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteParagraphListFormat(t *testing.T) {
	filename := "ParagraphDeleteListFormat.doc"
	baseDirPath := path.Join("DocumentElements", "ParagraphListFormat")
	localFilePath := GetLocalPath(baseDirPath, filename)
	remoteFolder := path.Join(remoteBaseTestDataFolder, baseDirPath)
	nodePath := ""
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, filename))

	_, _, err := client.WordsApi.DeleteParagraphListFormat(ctx, filename, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteParagraphListFormatWithoutNodePath(t *testing.T) {
	filename := "ParagraphDeleteListFormat.doc"
	baseDirPath := path.Join("DocumentElements", "ParagraphListFormat")
	localFilePath := GetLocalPath(baseDirPath, filename)
	remoteFolder := path.Join(remoteBaseTestDataFolder, baseDirPath)
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, filename))

	_, _, err := client.WordsApi.DeleteParagraphListFormatWithoutNodePath(ctx, filename, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetParagraphTabStops(t *testing.T) {
	filename := "ParagraphTabStops.docx"
	baseDirPath := path.Join("DocumentElements", "Paragraphs")
	localFilePath := GetLocalPath(baseDirPath, filename)
	remoteFolder := path.Join(remoteBaseTestDataFolder, baseDirPath)
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, filename))

	_, _, err := client.WordsApi.GetParagraphTabStops(ctx, filename, "", int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetParagraphTabStopsWithoutNodePath(t *testing.T) {
	filename := "ParagraphTabStops.docx"
	baseDirPath := path.Join("DocumentElements", "Paragraphs")
	localFilePath := GetLocalPath(baseDirPath, filename)
	remoteFolder := path.Join(remoteBaseTestDataFolder, baseDirPath)
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, filename))

	_, _, err := client.WordsApi.GetParagraphTabStopsWithoutNodePath(ctx, filename, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertParagraphTabStop(t *testing.T) {
	filename := "ParagraphTabStops.docx"
	baseDirPath := path.Join("DocumentElements", "Paragraphs")
	localFilePath := GetLocalPath(baseDirPath, filename)
	remoteFolder := path.Join(remoteBaseTestDataFolder, baseDirPath)
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, filename))

	dto := models.TabStopInsert{Alignment: "Left", Leader: "None", Position: 72}

	_, _, err := client.WordsApi.InsertOrUpdateParagraphTabStop(ctx, filename, dto, "", int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertParagraphTabStopWithoutNodePath(t *testing.T) {
	filename := "ParagraphTabStops.docx"
	baseDirPath := path.Join("DocumentElements", "Paragraphs")
	localFilePath := GetLocalPath(baseDirPath, filename)
	remoteFolder := path.Join(remoteBaseTestDataFolder, baseDirPath)
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, filename))

	dto := models.TabStopInsert{Alignment: "Left", Leader: "None", Position: 72}

	_, _, err := client.WordsApi.InsertOrUpdateParagraphTabStopWithoutNodePath(ctx, filename, dto, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteAllParagraphTabStops(t *testing.T) {
	filename := "ParagraphTabStops.docx"
	baseDirPath := path.Join("DocumentElements", "Paragraphs")
	localFilePath := GetLocalPath(baseDirPath, filename)
	remoteFolder := path.Join(remoteBaseTestDataFolder, baseDirPath)
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, filename))

	_, _, err := client.WordsApi.DeleteAllParagraphTabStops(ctx, filename, "", int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteAllParagraphTabStopsWithoutNodePath(t *testing.T) {
	filename := "ParagraphTabStops.docx"
	baseDirPath := path.Join("DocumentElements", "Paragraphs")
	localFilePath := GetLocalPath(baseDirPath, filename)
	remoteFolder := path.Join(remoteBaseTestDataFolder, baseDirPath)
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, filename))

	_, _, err := client.WordsApi.DeleteAllParagraphTabStopsWithoutNodePath(ctx, filename, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteParagraphTabStop(t *testing.T) {
	filename := "ParagraphTabStops.docx"
	baseDirPath := path.Join("DocumentElements", "Paragraphs")
	localFilePath := GetLocalPath(baseDirPath, filename)
	remoteFolder := path.Join(remoteBaseTestDataFolder, baseDirPath)
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, filename))

	_, _, err := client.WordsApi.DeleteParagraphTabStop(ctx, filename, 72, "", int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteParagraphTabStopWithoutNodePath(t *testing.T) {
	filename := "ParagraphTabStops.docx"
	baseDirPath := path.Join("DocumentElements", "Paragraphs")
	localFilePath := GetLocalPath(baseDirPath, filename)
	remoteFolder := path.Join(remoteBaseTestDataFolder, baseDirPath)
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, filename))

	_, _, err := client.WordsApi.DeleteParagraphTabStopWithoutNodePath(ctx, filename, 72, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}
