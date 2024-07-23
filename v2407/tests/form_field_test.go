/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="form_field_test.go">
 *   Copyright (c) 2024 Aspose.Words for Cloud
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
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2407/api/models"
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
        TextInputType: ToStringPointer("Regular"),
        TextInputDefault: ToStringPointer("No name"),
        TextInputFormat: ToStringPointer(""),
    }

    options := map[string]interface{}{
        "nodePath": "sections/0",
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }

    request := &models.UpdateFormFieldRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        FormField: &requestFormField,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.UpdateFormField(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetFormField(), "Validate UpdateFormField response.");
    assert.Equal(t, "FullName", DereferenceValue(actual.GetFormField().GetName()), "Validate UpdateFormField response.");
    assert.Equal(t, "", DereferenceValue(actual.GetFormField().GetStatusText()), "Validate UpdateFormField response.");
}

// Test for posting form field online.
func Test_FormField_UpdateFormFieldOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    fieldFolder := "DocumentElements/FormFields"

    requestDocument := OpenFile(t, fieldFolder + "/FormFilled.docx")
    requestFormField := models.FormFieldTextInput{
        Name: ToStringPointer("FullName"),
        Enabled: ToBoolPointer(true),
        CalculateOnExit: ToBoolPointer(true),
        StatusText: ToStringPointer(""),
        TextInputType: ToStringPointer("Regular"),
        TextInputDefault: ToStringPointer("No name"),
        TextInputFormat: ToStringPointer(""),
    }

    options := map[string]interface{}{
        "nodePath": "sections/0",
    }

    request := &models.UpdateFormFieldOnlineRequest{
        Document: requestDocument,
        FormField: &requestFormField,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.UpdateFormFieldOnline(ctx, request)
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
        Name: ToStringPointer("FullName"),
        Enabled: ToBoolPointer(true),
        CalculateOnExit: ToBoolPointer(true),
        StatusText: ToStringPointer(""),
        TextInputType: ToStringPointer("Regular"),
        TextInputDefault: ToStringPointer("No name"),
        TextInputFormat: ToStringPointer(""),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }

    request := &models.UpdateFormFieldRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        FormField: &requestFormField,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.UpdateFormField(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetFormField(), "Validate UpdateFormFieldWithoutNodePath response.");
    assert.Equal(t, "FullName", DereferenceValue(actual.GetFormField().GetName()), "Validate UpdateFormFieldWithoutNodePath response.");
    assert.Equal(t, "", DereferenceValue(actual.GetFormField().GetStatusText()), "Validate UpdateFormFieldWithoutNodePath response.");
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

    request := &models.GetFormFieldRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetFormField(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetFormField(), "Validate GetFormField response.");
    assert.Equal(t, "FullName", DereferenceValue(actual.GetFormField().GetName()), "Validate GetFormField response.");
}

// Test for getting form field online.
func Test_FormField_GetFormFieldOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    fieldFolder := "DocumentElements/FormFields"

    requestDocument := OpenFile(t, fieldFolder + "/FormFilled.docx")

    options := map[string]interface{}{
        "nodePath": "sections/0",
    }

    request := &models.GetFormFieldOnlineRequest{
        Document: requestDocument,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetFormFieldOnline(ctx, request)
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

    request := &models.GetFormFieldRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetFormField(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetFormField(), "Validate GetFormFieldWithoutNodePath response.");
    assert.Equal(t, "FullName", DereferenceValue(actual.GetFormField().GetName()), "Validate GetFormFieldWithoutNodePath response.");
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

    request := &models.GetFormFieldsRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetFormFields(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetFormFields(), "Validate GetFormFields response.");
    assert.NotNil(t, actual.GetFormFields().GetList(), "Validate GetFormFields response.");
    assert.Equal(t, 5, len(actual.GetFormFields().GetList()), "Validate GetFormFields response.");
    assert.Equal(t, "FullName", DereferenceValue(actual.GetFormFields().GetList()[0].GetName()), "Validate GetFormFields response.");
}

// Test for getting form fields online.
func Test_FormField_GetFormFieldsOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    fieldFolder := "DocumentElements/FormFields"

    requestDocument := OpenFile(t, fieldFolder + "/FormFilled.docx")

    options := map[string]interface{}{
        "nodePath": "sections/0",
    }

    request := &models.GetFormFieldsOnlineRequest{
        Document: requestDocument,
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetFormFieldsOnline(ctx, request)
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

    request := &models.GetFormFieldsRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetFormFields(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetFormFields(), "Validate GetFormFieldsWithoutNodePath response.");
    assert.NotNil(t, actual.GetFormFields().GetList(), "Validate GetFormFieldsWithoutNodePath response.");
    assert.Equal(t, 5, len(actual.GetFormFields().GetList()), "Validate GetFormFieldsWithoutNodePath response.");
    assert.Equal(t, "FullName", DereferenceValue(actual.GetFormFields().GetList()[0].GetName()), "Validate GetFormFieldsWithoutNodePath response.");
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
        TextInputType: ToStringPointer("Regular"),
        TextInputDefault: ToStringPointer("123"),
        TextInputFormat: ToStringPointer("UPPERCASE"),
    }

    options := map[string]interface{}{
        "nodePath": "sections/0/paragraphs/0",
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }

    request := &models.InsertFormFieldRequest{
        Name: ToStringPointer(remoteFileName),
        FormField: &requestFormField,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.InsertFormField(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetFormField(), "Validate InsertFormField response.");
    assert.Equal(t, "FullName", DereferenceValue(actual.GetFormField().GetName()), "Validate InsertFormField response.");
    assert.Equal(t, "", DereferenceValue(actual.GetFormField().GetStatusText()), "Validate InsertFormField response.");
}

// Test for insert form field without node path online.
func Test_FormField_InsertFormFieldOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    fieldFolder := "DocumentElements/FormFields"

    requestDocument := OpenFile(t, fieldFolder + "/FormFilled.docx")
    requestFormField := models.FormFieldTextInput{
        Name: ToStringPointer("FullName"),
        Enabled: ToBoolPointer(true),
        CalculateOnExit: ToBoolPointer(true),
        StatusText: ToStringPointer(""),
        TextInputType: ToStringPointer("Regular"),
        TextInputDefault: ToStringPointer("123"),
        TextInputFormat: ToStringPointer("UPPERCASE"),
    }

    options := map[string]interface{}{
        "nodePath": "sections/0/paragraphs/0",
    }

    request := &models.InsertFormFieldOnlineRequest{
        Document: requestDocument,
        FormField: &requestFormField,
        Optionals: options,
    }

    _, _, err := client.WordsApi.InsertFormFieldOnline(ctx, request)
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
        Name: ToStringPointer("FullName"),
        Enabled: ToBoolPointer(true),
        CalculateOnExit: ToBoolPointer(true),
        StatusText: ToStringPointer(""),
        TextInputType: ToStringPointer("Regular"),
        TextInputDefault: ToStringPointer("123"),
        TextInputFormat: ToStringPointer("UPPERCASE"),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }

    request := &models.InsertFormFieldRequest{
        Name: ToStringPointer(remoteFileName),
        FormField: &requestFormField,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.InsertFormField(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetFormField(), "Validate InsertFormFieldWithoutNodePath response.");
    assert.Equal(t, "FullName", DereferenceValue(actual.GetFormField().GetName()), "Validate InsertFormFieldWithoutNodePath response.");
    assert.Equal(t, "", DereferenceValue(actual.GetFormField().GetStatusText()), "Validate InsertFormFieldWithoutNodePath response.");
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

    request := &models.DeleteFormFieldRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, err := client.WordsApi.DeleteFormField(ctx, request)

    if err != nil {
        t.Error(err)
    }

}

// Test for deleting form field online.
func Test_FormField_DeleteFormFieldOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    fieldFolder := "DocumentElements/FormFields"

    requestDocument := OpenFile(t, fieldFolder + "/FormFilled.docx")

    options := map[string]interface{}{
        "nodePath": "sections/0",
    }

    request := &models.DeleteFormFieldOnlineRequest{
        Document: requestDocument,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.DeleteFormFieldOnline(ctx, request)
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

    request := &models.DeleteFormFieldRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, err := client.WordsApi.DeleteFormField(ctx, request)

    if err != nil {
        t.Error(err)
    }

}
