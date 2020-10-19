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

	"github.com/aspose-words-cloud/aspose-words-cloud-go/v2003/api/models"
)

func TestDeleteFormField(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "FormFields"), "FormFilled.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "FormFields")
	remoteName := "TestDeleteFormField.docx"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteFormField(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteFormFieldWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "FormFields"), "FormFilled.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "FormFields")
	remoteName := "TestDeleteFormFieldWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, err := client.WordsApi.DeleteFormFieldWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetFormField(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "FormFields"), "FormFilled.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "FormFields")
	remoteName := "TestGetFormField.docx"
	nodePath := "sections/0"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetFormField(ctx, remoteName, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetFormFieldWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "FormFields"), "FormFilled.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "FormFields")
	remoteName := "TestGetFormFieldWithoutNodePath.docx"
	index := 0
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetFormFieldWithoutNodePath(ctx, remoteName, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetFormFields(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "FormFields"), "FormFilled.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "FormFields")
	remoteName := "TestGetFormFields.docx"
	nodePath := "sections/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetFormFields(ctx, remoteName, nodePath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestGetFormFieldsWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "FormFields"), "FormFilled.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "FormFields")
	remoteName := "TestGetFormFieldsWithoutNodePath.docx"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.GetFormFieldsWithoutNodePath(ctx, remoteName, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertFormField(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "FormFields")
	remoteName := "TestInsertFormField.docx"
	var formField models.IFormField = models.FormFieldTextInput{
		Name:             "FullName",
		Enabled:          true,
		CalculateOnExit:  true,
		StatusText:       "",
		TextInputType:    "Regular",
		TextInputDefault: "123",
		TextInputFormat:  "UPPERCASE",
	}
	nodePath := "sections/0/paragraphs/0"
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertFormField(ctx, remoteName, formField, nodePath, options)

	if err != nil {
		t.Error(err)
	}
}

func TestInsertFormFieldWithoutNodePath(t *testing.T) {

	localFilePath := commonTestFile
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "FormFields")
	remoteName := "TestInsertFormFieldWithoutNodePath.docx"
	formField := models.FormFieldTextInput{
		Name:             "FullName",
		Enabled:          true,
		CalculateOnExit:  true,
		StatusText:       "",
		TextInputType:    "Regular",
		TextInputDefault: "123",
		TextInputFormat:  "UPPERCASE",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.InsertFormFieldWithoutNodePath(ctx, remoteName, formField, options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateFormField(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "FormFields"), "FormFilled.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "FormFields")
	remoteName := "TestUpdateFormField.docx"
	nodePath := "sections/0"
	index := 0
	formField := models.FormFieldTextInput{
		Name:             "FullName",
		Enabled:          true,
		CalculateOnExit:  true,
		TextInputType:    "Regular",
		TextInputDefault: "No name",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateFormField(ctx, remoteName, formField, nodePath, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateFormFieldWithoutNodePath(t *testing.T) {

	localFilePath := GetLocalPath(filepath.Join("DocumentElements", "FormFields"), "FormFilled.docx")
	remoteFolder := path.Join(remoteBaseTestDataFolder, "DocumentElements", "FormFields")
	remoteName := "TestUpdateFormFieldWithoutNodePath.docx"
	index := 0
	formField := models.FormFieldTextInput{
		Name:             "FullName",
		Enabled:          true,
		CalculateOnExit:  true,
		TextInputType:    "Regular",
		TextInputDefault: "No name",
	}
	options := map[string]interface{}{
		"folder": remoteFolder,
	}

	client, ctx := UploadFileToStorage(t, localFilePath, path.Join(remoteFolder, remoteName))

	_, _, err := client.WordsApi.UpdateFormFieldWithoutNodePath(ctx, remoteName, formField, int32(index), options)

	if err != nil {
		t.Error(err)
	}
}
