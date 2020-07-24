/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="form_field_test.go">
 *   Copyright (c) 2020 Aspose.Words for Cloud
 * </copyright>
 * <summary>
 *   Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 * 
 *  The above copyright notice and this permission notice shall be included in all
 *  copies or substantial portions of the Software.
 * 
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *  SOFTWARE.
 * </summary>
 * --------------------------------------------------------------------------------
 */

// Example of how to work with form field.
package api_test

import (
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/dev/api/models"
)

// Test for posting form field.
func Test_FormField_UpdateFormField(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/FormFields"
    fieldFolder := "DocumentElements/FormFields"
    remoteFileName := "TestUpdateFormField.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(fieldFolder + "/FormFilled.docx"), remoteDataFolder + "/" + remoteFileName)

    requestFormField := models.FormFieldTextInput{
        Name: "FullName",
        Enabled: true,
        CalculateOnExit: true,
        StatusText: "",
        TextInputType: "Regular",
        TextInputDefault: "No name",
    }

    options := map[string]interface{}{
        "nodePath": "sections/0",
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }
    _, _, err := client.WordsApi.UpdateFormField(ctx, remoteFileName, requestFormField, int32(0), options)

    if err != nil {
        t.Error(err)
    }
}

// Test for posting form field without node path.
func Test_FormField_UpdateFormFieldWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/FormFields"
    fieldFolder := "DocumentElements/FormFields"
    remoteFileName := "TestUpdateFormFieldWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(fieldFolder + "/FormFilled.docx"), remoteDataFolder + "/" + remoteFileName)

    requestFormField := models.FormFieldTextInput{
        Name: "FullName",
        Enabled: true,
        CalculateOnExit: true,
        StatusText: "",
        TextInputType: "Regular",
        TextInputDefault: "No name",
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }
    _, _, err := client.WordsApi.UpdateFormField(ctx, remoteFileName, requestFormField, int32(0), options)

    if err != nil {
        t.Error(err)
    }
}

// Test for getting form field.
func Test_FormField_GetFormField(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/FormFields"
    fieldFolder := "DocumentElements/FormFields"
    remoteFileName := "TestGetFormField.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(fieldFolder + "/FormFilled.docx"), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "sections/0",
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetFormField(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }
}

// Test for getting form field without node path.
func Test_FormField_GetFormFieldWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/FormFields"
    fieldFolder := "DocumentElements/FormFields"
    remoteFileName := "TestGetFormFieldWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(fieldFolder + "/FormFilled.docx"), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetFormField(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }
}

// Test for getting form fields.
func Test_FormField_GetFormFields(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/FormFields"
    fieldFolder := "DocumentElements/FormFields"
    remoteFileName := "TestGetFormFields.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(fieldFolder + "/FormFilled.docx"), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "sections/0",
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetFormFields(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }
}

// Test for getting form fields without node path.
func Test_FormField_GetFormFieldsWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/FormFields"
    fieldFolder := "DocumentElements/FormFields"
    remoteFileName := "TestGetFormFieldsWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(fieldFolder + "/FormFilled.docx"), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetFormFields(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }
}

// Test for insert form field without node path.
func Test_FormField_InsertFormField(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/FormFields"
    remoteFileName := "TestInsertFormField.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile("Common/test_multi_pages.docx"), remoteDataFolder + "/" + remoteFileName)

    requestFormField := models.FormFieldTextInput{
        Name: "FullName",
        Enabled: true,
        CalculateOnExit: true,
        StatusText: "",
        TextInputType: "Regular",
        TextInputDefault: "123",
        TextInputFormat: "UPPERCASE",
    }

    options := map[string]interface{}{
        "nodePath": "sections/0/paragraphs/0",
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }
    _, _, err := client.WordsApi.InsertFormField(ctx, remoteFileName, requestFormField, options)

    if err != nil {
        t.Error(err)
    }
}

// Test for insert form field without node path.
func Test_FormField_InsertFormFieldWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/FormFields"
    remoteFileName := "TestInsertFormFieldWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile("Common/test_multi_pages.docx"), remoteDataFolder + "/" + remoteFileName)

    requestFormField := models.FormFieldTextInput{
        Name: "FullName",
        Enabled: true,
        CalculateOnExit: true,
        StatusText: "",
        TextInputType: "Regular",
        TextInputDefault: "123",
        TextInputFormat: "UPPERCASE",
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }
    _, _, err := client.WordsApi.InsertFormField(ctx, remoteFileName, requestFormField, options)

    if err != nil {
        t.Error(err)
    }
}

// Test for deleting form field.
func Test_FormField_DeleteFormField(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/FormFields"
    fieldFolder := "DocumentElements/FormFields"
    remoteFileName := "TestDeleteFormField.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(fieldFolder + "/FormFilled.docx"), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "sections/0",
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }
    _, err := client.WordsApi.DeleteFormField(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }
}

// Test for deleting form field without node path.
func Test_FormField_DeleteFormFieldWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/FormFields"
    fieldFolder := "DocumentElements/FormFields"
    remoteFileName := "TestDeleteFormFieldWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(fieldFolder + "/FormFilled.docx"), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }
    _, err := client.WordsApi.DeleteFormField(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }
}
