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
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2011/api/models"
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
        Name: ToStringPointer("FullName"),
        Enabled: ToBoolPointer(true),
        CalculateOnExit: ToBoolPointer(true),
        StatusText: ToStringPointer(""),
        TextInputType: "Regular",
        TextInputDefault: ToStringPointer("No name"),
    }

    options := map[string]interface{}{
        "nodePath": "sections/0",
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }
    actual, _, err := client.WordsApi.UpdateFormField(ctx, remoteFileName, requestFormField, int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.FormField, "Validate UpdateFormField response.");
    assert.Equal(t, "FullName", *actual.FormField.Name, "Validate UpdateFormField response.");
    assert.Equal(t, "", *actual.FormField.StatusText, "Validate UpdateFormField response.");
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
        Name: ToStringPointer("FullName"),
        Enabled: ToBoolPointer(true),
        CalculateOnExit: ToBoolPointer(true),
        StatusText: ToStringPointer(""),
        TextInputType: "Regular",
        TextInputDefault: ToStringPointer("No name"),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }
    actual, _, err := client.WordsApi.UpdateFormField(ctx, remoteFileName, requestFormField, int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.FormField, "Validate UpdateFormFieldWithoutNodePath response.");
    assert.Equal(t, "FullName", *actual.FormField.Name, "Validate UpdateFormFieldWithoutNodePath response.");
    assert.Equal(t, "", *actual.FormField.StatusText, "Validate UpdateFormFieldWithoutNodePath response.");
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
    actual, _, err := client.WordsApi.GetFormField(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.FormField, "Validate GetFormField response.");
    assert.Equal(t, "FullName", *actual.FormField.Name, "Validate GetFormField response.");
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
    actual, _, err := client.WordsApi.GetFormField(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.FormField, "Validate GetFormFieldWithoutNodePath response.");
    assert.Equal(t, "FullName", *actual.FormField.Name, "Validate GetFormFieldWithoutNodePath response.");
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
    actual, _, err := client.WordsApi.GetFormFields(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.FormFields, "Validate GetFormFields response.");
    assert.NotNil(t, actual.FormFields.List, "Validate GetFormFields response.");
    assert.Equal(t, 5, len(actual.FormFields.List), "Validate GetFormFields response.");
    assert.Equal(t, "FullName", *actual.FormFields.List[0].Name, "Validate GetFormFields response.");
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
    actual, _, err := client.WordsApi.GetFormFields(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.FormFields, "Validate GetFormFieldsWithoutNodePath response.");
    assert.NotNil(t, actual.FormFields.List, "Validate GetFormFieldsWithoutNodePath response.");
    assert.Equal(t, 5, len(actual.FormFields.List), "Validate GetFormFieldsWithoutNodePath response.");
    assert.Equal(t, "FullName", *actual.FormFields.List[0].Name, "Validate GetFormFieldsWithoutNodePath response.");
}

// Test for insert form field without node path.
func Test_FormField_InsertFormField(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/FormFields"
    remoteFileName := "TestInsertFormField.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile("Common/test_multi_pages.docx"), remoteDataFolder + "/" + remoteFileName)

    requestFormField := models.FormFieldTextInput{
        Name: ToStringPointer("FullName"),
        Enabled: ToBoolPointer(true),
        CalculateOnExit: ToBoolPointer(true),
        StatusText: ToStringPointer(""),
        TextInputType: "Regular",
        TextInputDefault: ToStringPointer("123"),
        TextInputFormat: ToStringPointer("UPPERCASE"),
    }

    options := map[string]interface{}{
        "nodePath": "sections/0/paragraphs/0",
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }
    actual, _, err := client.WordsApi.InsertFormField(ctx, remoteFileName, requestFormField, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.FormField, "Validate InsertFormField response.");
    assert.Equal(t, "FullName", *actual.FormField.Name, "Validate InsertFormField response.");
    assert.Equal(t, "", *actual.FormField.StatusText, "Validate InsertFormField response.");
}

// Test for insert form field without node path.
func Test_FormField_InsertFormFieldWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/FormFields"
    remoteFileName := "TestInsertFormFieldWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile("Common/test_multi_pages.docx"), remoteDataFolder + "/" + remoteFileName)

    requestFormField := models.FormFieldTextInput{
        Name: ToStringPointer("FullName"),
        Enabled: ToBoolPointer(true),
        CalculateOnExit: ToBoolPointer(true),
        StatusText: ToStringPointer(""),
        TextInputType: "Regular",
        TextInputDefault: ToStringPointer("123"),
        TextInputFormat: ToStringPointer("UPPERCASE"),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }
    actual, _, err := client.WordsApi.InsertFormField(ctx, remoteFileName, requestFormField, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.FormField, "Validate InsertFormFieldWithoutNodePath response.");
    assert.Equal(t, "FullName", *actual.FormField.Name, "Validate InsertFormFieldWithoutNodePath response.");
    assert.Equal(t, "", *actual.FormField.StatusText, "Validate InsertFormFieldWithoutNodePath response.");
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
