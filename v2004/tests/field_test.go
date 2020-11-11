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

	"github.com/aspose-words-cloud/aspose-words-cloud-go/v2004/api/models"
)

func TestDeleteField(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Fields"), "GetField.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Fields")
	remoteName := "TestDeleteField.docx"
	nodePath := "sections/0/paragraphs/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteField(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteFieldWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Fields"), "GetField.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Fields")
	remoteName := "TestDeleteFieldWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteFieldWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteFields(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Fields")
	remoteName := "TestDeleteFields.docx"
	nodePath := "sections/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteFields(ctx, remoteName, nodePath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteFieldsWithoutNodePath(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Fields")
	remoteName := "TestDeleteFieldsWithoutNodePath.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteFieldsWithoutNodePath(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetField(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Fields"), "GetField.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Fields")
	remoteName := "TestGetField.docx"
	nodePath := "sections/0/paragraphs/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetField(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetFieldWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Fields"), "GetField.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Fields")
	remoteName := "TestGetFieldWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetFieldWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetFields(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Fields"), "GetField.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Fields")
	remoteName := "TestGetFields.docx"
	nodePath := "sections/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetFields(ctx, remoteName, nodePath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetFieldsWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Fields"), "GetField.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Fields")
	remoteName := "TestGetFieldsWithoutNodePath.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetFieldsWithoutNodePath(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertField(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Text"), "SampleWordDocument.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Fields")
	remoteName := "TestInsertField.docx"
	nodePath := "sections/0/paragraphs/0"
	field := models.FieldInsert{
		FieldCode: "{ NUMPAGES }",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertField(ctx, remoteName, field, nodePath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertFieldWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Text"), "SampleWordDocument.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Fields")
	remoteName := "TestInsertFieldWithoutNodePath.docx"
	field := models.FieldInsert{
		FieldCode: "{ NUMPAGES }",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertFieldWithoutNodePath(ctx, remoteName, field, options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateField(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "Fields"), "GetField.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Fields")
	remoteName := "TestUpdateField.docx"
	nodePath := "sections/0/paragraphs/0"
	index := 0
	field := models.FieldUpdate{
		FieldCode: "{ NUMPAGES }",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateField(ctx, remoteName, field, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateFields(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "Fields")
	remoteName := "TestUpdateFields.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateFields(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentFieldNames(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentActions", "MailMerge")
	remoteName := "TestGetDocumentFieldNames.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetDocumentFieldNames(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocumentFieldNamesOnline(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentActions", "MailMerge"), "SampleExecuteTemplate.docx")
	template, fileErr := os.Open(localFilePath)
	if fileErr != nil {
		t.Error(fileErr)
	}

	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	_, _, err := client.WordsApi.GetDocumentFieldNamesOnline(ctx, template, nil)

	if err != nil {
		t.Error(err)
	}
}
