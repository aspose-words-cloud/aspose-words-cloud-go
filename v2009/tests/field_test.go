/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="field_test.go">
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

// Example of how to work with field.
package api_test

import (
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2009/api/models"
)

// Test for getting fields.
func Test_Field_GetFields(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Fields"
    fieldFolder := "DocumentElements/Fields"
    localFileName := "GetField.docx"
    remoteFileName := "TestGetFields.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(fieldFolder + "/" + localFileName), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "sections/0",
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetFields(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

}

// Test for getting fields without node path.
func Test_Field_GetFieldsWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Fields"
    fieldFolder := "DocumentElements/Fields"
    localFileName := "GetField.docx"
    remoteFileName := "TestGetFieldsWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(fieldFolder + "/" + localFileName), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetFields(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

}

// Test for getting field by index.
func Test_Field_GetField(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Fields"
    fieldFolder := "DocumentElements/Fields"
    localFileName := "GetField.docx"
    remoteFileName := "TestGetField.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(fieldFolder + "/" + localFileName), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "sections/0/paragraphs/0",
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetField(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for getting field by index without node path.
func Test_Field_GetFieldWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Fields"
    fieldFolder := "DocumentElements/Fields"
    localFileName := "GetField.docx"
    remoteFileName := "TestGetFieldWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(fieldFolder + "/" + localFileName), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.GetField(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for putting field.
func Test_Field_InsertField(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Fields"
    textFolder := "DocumentElements/Text"
    localFileName := "SampleWordDocument.docx"
    remoteFileName := "TestInsertField.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(textFolder + "/" + localFileName), remoteDataFolder + "/" + remoteFileName)

    requestField := models.FieldInsert{
        FieldCode: "{ NUMPAGES }",
    }

    options := map[string]interface{}{
        "nodePath": "sections/0/paragraphs/0",
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.InsertField(ctx, remoteFileName, requestField, options)

    if err != nil {
        t.Error(err)
    }

}

// Test for putting field without node path.
func Test_Field_InsertFieldWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Fields"
    textFolder := "DocumentElements/Text"
    localFileName := "SampleWordDocument.docx"
    remoteFileName := "TestInsertFieldWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(textFolder + "/" + localFileName), remoteDataFolder + "/" + remoteFileName)

    requestField := models.FieldInsert{
        FieldCode: "{ NUMPAGES }",
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.InsertField(ctx, remoteFileName, requestField, options)

    if err != nil {
        t.Error(err)
    }

}

// Test for posting field.
func Test_Field_UpdateField(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Fields"
    fieldFolder := "DocumentElements/Fields"
    localFileName := "GetField.docx"
    remoteFileName := "TestUpdateField.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(fieldFolder + "/" + localFileName), remoteDataFolder + "/" + remoteFileName)

    requestField := models.FieldUpdate{
        FieldCode: "{ NUMPAGES }",
    }

    options := map[string]interface{}{
        "nodePath": "sections/0/paragraphs/0",
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.UpdateField(ctx, remoteFileName, requestField, int32(0), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for inserting page numbers field.
func Test_Field_InsertPageNumbers(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Fields"
    localFileName := "test_multi_pages.docx"
    remoteFileName := "TestInsertPageNumbers.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile("Common/" + localFileName), remoteDataFolder + "/" + remoteFileName)

    requestPageNumber := models.PageNumber{
        Alignment: "center",
        Format: "{PAGE} of {NUMPAGES}",
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }
    _, _, err := client.WordsApi.InsertPageNumbers(ctx, remoteFileName, requestPageNumber, options)

    if err != nil {
        t.Error(err)
    }

}

// Test for deleting field.
func Test_Field_DeleteField(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Fields"
    fieldFolder := "DocumentElements/Fields"
    localFileName := "GetField.docx"
    remoteFileName := "TestDeleteField.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(fieldFolder + "/" + localFileName), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "sections/0/paragraphs/0",
        "folder": remoteDataFolder,
    }
    _, err := client.WordsApi.DeleteField(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for deleting field without node path.
func Test_Field_DeleteFieldWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Fields"
    fieldFolder := "DocumentElements/Fields"
    localFileName := "GetField.docx"
    remoteFileName := "TestDeleteFieldWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(fieldFolder + "/" + localFileName), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, err := client.WordsApi.DeleteField(ctx, remoteFileName, int32(0), options)

    if err != nil {
        t.Error(err)
    }

}

// Test for deleting paragraph fields.
func Test_Field_DeleteParagraphFields(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Fields"
    localFileName := "test_multi_pages.docx"
    remoteFileName := "TestDeleteParagraphFields.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile("Common/" + localFileName), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "paragraphs/0",
        "folder": remoteDataFolder,
    }
    _, err := client.WordsApi.DeleteFields(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

}

// Test for deleting paragraph fields without node path.
func Test_Field_DeleteParagraphFieldsWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Fields"
    localFileName := "test_multi_pages.docx"
    remoteFileName := "TestDeleteParagraphFieldsWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile("Common/" + localFileName), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, err := client.WordsApi.DeleteFields(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

}

// Test for deleting section fields.
func Test_Field_DeleteSectionFields(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Fields"
    localFileName := "test_multi_pages.docx"
    remoteFileName := "TestDeleteSectionFields.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile("Common/" + localFileName), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "sections/0",
        "folder": remoteDataFolder,
    }
    _, err := client.WordsApi.DeleteFields(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

}

// Test for deleting section fields without node path.
func Test_Field_DeleteSectionFieldsWithoutNodePath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Fields"
    localFileName := "test_multi_pages.docx"
    remoteFileName := "TestDeleteSectionFieldsWithoutNodePath.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile("Common/" + localFileName), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, err := client.WordsApi.DeleteFields(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

}

// Test for deleting paragraph fields in section.
func Test_Field_DeleteSectionParagraphFields(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Fields"
    localFileName := "test_multi_pages.docx"
    remoteFileName := "TestDeleteSectionParagraphFields.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile("Common/" + localFileName), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "sections/0/paragraphs/0",
        "folder": remoteDataFolder,
    }
    _, err := client.WordsApi.DeleteFields(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

}

// Test for deleting fields.
func Test_Field_DeleteDocumentFields(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Fields"
    localFileName := "test_multi_pages.docx"
    remoteFileName := "TestDeleteSectionParagraphFields.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile("Common/" + localFileName), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "",
        "folder": remoteDataFolder,
    }
    _, err := client.WordsApi.DeleteFields(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

}

// Test for posting updated fields.
func Test_Field_UpdateDocumentFields(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/Fields"
    localFileName := "test_multi_pages.docx"
    remoteFileName := "TestUpdateDocumentFields.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile("Common/" + localFileName), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }
    _, _, err := client.WordsApi.UpdateFields(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

}
