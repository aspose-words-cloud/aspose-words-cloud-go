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
	"os"
	"path"
	"path/filepath"
	"testing"
	"time"

	"github.com/aspose-words-cloud/aspose-words-cloud-go/v2006/api/models"
)

func TestAppendDocument(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "AppendDocument")
	remoteName := "TestAppendDocument.docx"
	remoteFilePath := path.Join(remoteFolder, remoteName)
	documentList := models.DocumentEntryList{
		DocumentEntries: []models.DocumentEntry{
			{
				Href:             remoteFilePath,
				ImportFormatMode: "KeepSourceFormatting",
			},
		},
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, remoteFilePath)

	_, _, err := client.WordsApi.AppendDocument(ctx, remoteName, documentList, options)

	if err != nil {
		t.Error(err)
	}
}

func TestClassify(t *testing.T) {

	text := "Try text classification"
	options := map[string]interface{}{
		"bestClassesCount": "3",
	}

	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	_, _, err := client.WordsApi.Classify(ctx, text, options)

	if err != nil {
		t.Error(err)
	}
}

func TestClassifyDocument(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("Common"), "test_multi_pages.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "Common")
	remoteName := "TestClassifyDocument.docx"
	options := map[string]interface{}{
		"folder":           remoteFolder,
		"bestClassesCount": "3",
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.ClassifyDocument(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestCompareDocument(t *testing.T) {

	localFilePath1 := GetLocalPath(filepath.Join("DocumentActions", "CompareDocument"), "compareTestDoc1.doc")
	localFilePath2 := GetLocalPath(filepath.Join("DocumentActions", "CompareDocument"), "compareTestDoc2.doc")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "CompareDocument")
	remoteName1 := "TestCompareDocument1.docx"
	remoteName2 := "TestCompareDocument2.docx"
	remotePath1 := path.Join(remoteFolder, remoteName1)
	remotePath2 := path.Join(remoteFolder, remoteName2)
	compareData := models.CompareData{
		ComparingWithDocument: remotePath2,
		Author:                "author",
		DateTime:              time.Date(2015, time.October, 26, 0, 0, 0, 0, time.UTC),
	}
	options := map[string]interface{}{
		"folder":       remoteFolder,
		"destFileName": path.Join("TestOut", "TestCompareDocumentOut.doc"),
	}

	client, ctx := UploadFileToStorage(t, localFilePath1, remotePath1)
	UploadNextFileToStorage(t, ctx, client, localFilePath2, remotePath2)

	_, _, err := client.WordsApi.CompareDocument(ctx, remoteName1, compareData, options)

	if err != nil {
		t.Error(err)
	}
}

func TestConvertDocument(t *testing.T) {

	format := "pdf"
	localFilePath := GetLocalPath(filepath.Join("DocumentActions", "ConvertDocument"), "test_uploadfile.docx")
	document, fileErr := os.Open(localFilePath)
	if fileErr != nil {
		t.Error(fileErr)
	}

	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	output, err := client.WordsApi.ConvertDocument(ctx, document, format, nil)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestCreateDocument(t *testing.T) {

	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "Document")
	remoteName := "TestCreateDocument.docx"
	options := map[string]interface{}{
		"fileName": remoteName,
		"folder":   remoteFolder,
	}

	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	_, _, err := client.WordsApi.CreateDocument(ctx, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocument(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "Document")
	remoteName := "TestGetDocument.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetDocument(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestLoadWebDocument(t *testing.T) {

	data := models.LoadWebDocumentData{
		LoadingDocumentUrl: "http://google.com",
		SaveOptions: &models.SaveOptionsData{
			FileName:                "google.doc",
			SaveFormat:              "doc",
			DmlEffectsRenderingMode: "1",
			DmlRenderingMode:        "1",
			UpdateSdtContent:        false,
			ZipOutput:               false,
		},
	}

	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	_, _, err := client.WordsApi.LoadWebDocument(ctx, data, nil)

	if err != nil {
		t.Error(err)
	}
}

func TestSplitDocument(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "SplitDocument")
	remoteName := "TestSplitDocument.docx"
	format := "text"
	options := map[string]interface{}{
		"folder": remoteFolder,
		"from":   int32(1),
		"to":     int32(2),
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.SplitDocument(ctx, remoteName, format, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentStatistics(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "Statistics")
	remoteName := "TestGetDocumentStatistics.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetDocumentStatistics(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentWithFormat(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "DocumentWithFormat")
	remoteName := "TestGetDocumentWithFormat.docx"
	format := "text"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	output, err := client.WordsApi.GetDocumentWithFormat(ctx, remoteName, format, options)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}

func TestSaveAs(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "ConvertDocument")
	remoteName := "TestSaveAs.docx"
	remoteDstName := "TestSaveAs.pdf"
	saveOptionsData := models.SaveOptionsData{
		SaveFormat: "pdf",
		FileName:   path.Join(remoteFolder, remoteDstName),
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.SaveAs(ctx, remoteName, saveOptionsData, options)

	if err != nil {
		t.Error(err)
	}
}

func TestSaveAsTiff(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "ConvertDocument")
	remoteName := "TestSaveAsTiff.docx"
	destFileName := "TestSaveAsTiff.tiff"
	saveOptions := models.TiffSaveOptionsData{
		FileName:   path.Join(remoteFolder, destFileName),
		SaveFormat: "tiff",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.SaveAsTiff(ctx, remoteName, saveOptions, options)

	if err != nil {
		t.Error(err)
	}
}
