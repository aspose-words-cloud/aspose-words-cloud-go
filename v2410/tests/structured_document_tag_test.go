/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="structured_document_tag_test.go">
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

// Example of how to use structured document tags.
package api_test

import (
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2410/api/models"
)

// Test for getting SDT objects from document.
func Test_StructuredDocumentTag_GetStructuredDocumentTags(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/StructuredDocumentTag"
    localFile := "DocumentElements/StructuredDocumentTag/StructuredDocumentTag.docx"
    remoteFileName := "TestGetStructuredDocumentTags.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "sections/0/body/paragraphs/0",
        "folder": remoteDataFolder,
    }

    request := &models.GetStructuredDocumentTagsRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetStructuredDocumentTags(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for getting SDT objects from document online.
func Test_StructuredDocumentTag_GetStructuredDocumentTagsOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/StructuredDocumentTag/StructuredDocumentTag.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
        "nodePath": "sections/0/body/paragraphs/0",
    }

    request := &models.GetStructuredDocumentTagsOnlineRequest{
        Document: requestDocument,
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetStructuredDocumentTagsOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for getting SDT object from document.
func Test_StructuredDocumentTag_GetStructuredDocumentTag(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/StructuredDocumentTag"
    localFile := "DocumentElements/StructuredDocumentTag/StructuredDocumentTag.docx"
    remoteFileName := "TestGetStructuredDocumentTag.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "sections/0/body/paragraphs/0",
        "folder": remoteDataFolder,
    }

    request := &models.GetStructuredDocumentTagRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetStructuredDocumentTag(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for getting SDT object from document online.
func Test_StructuredDocumentTag_GetStructuredDocumentTagOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/StructuredDocumentTag/StructuredDocumentTag.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
        "nodePath": "sections/0/body/paragraphs/0",
    }

    request := &models.GetStructuredDocumentTagOnlineRequest{
        Document: requestDocument,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetStructuredDocumentTagOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for adding SDT object.
func Test_StructuredDocumentTag_InsertStructuredDocumentTag(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/StructuredDocumentTag"
    localFile := "DocumentElements/StructuredDocumentTag/StructuredDocumentTag.docx"
    remoteFileName := "TestInsetStructuredDocumentTag.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestStructuredDocumentTag := models.StructuredDocumentTagInsert{
        SdtType: ToStringPointer("ComboBox"),
        Level: ToStringPointer("Inline"),
    }

    options := map[string]interface{}{
        "nodePath": "sections/0/body/paragraphs/0",
        "folder": remoteDataFolder,
    }

    request := &models.InsertStructuredDocumentTagRequest{
        Name: ToStringPointer(remoteFileName),
        StructuredDocumentTag: &requestStructuredDocumentTag,
        Optionals: options,
    }

    _, _, err := client.WordsApi.InsertStructuredDocumentTag(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for adding SDT object online.
func Test_StructuredDocumentTag_InsertStructuredDocumentTagOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/StructuredDocumentTag/StructuredDocumentTag.docx"

    requestDocument := OpenFile(t, localFile)
    requestStructuredDocumentTag := models.StructuredDocumentTagInsert{
        SdtType: ToStringPointer("ComboBox"),
        Level: ToStringPointer("Inline"),
    }

    options := map[string]interface{}{
        "nodePath": "sections/0/body/paragraphs/0",
    }

    request := &models.InsertStructuredDocumentTagOnlineRequest{
        Document: requestDocument,
        StructuredDocumentTag: &requestStructuredDocumentTag,
        Optionals: options,
    }

    _, _, err := client.WordsApi.InsertStructuredDocumentTagOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for deleting SDT object.
func Test_StructuredDocumentTag_DeleteStructuredDocumentTag(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/StructuredDocumentTag"
    localFile := "DocumentElements/StructuredDocumentTag/StructuredDocumentTag.docx"
    remoteFileName := "TestDeleteStructuredDocumentTag.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "nodePath": "sections/0/body/paragraphs/0",
        "folder": remoteDataFolder,
    }

    request := &models.DeleteStructuredDocumentTagRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, err := client.WordsApi.DeleteStructuredDocumentTag(ctx, request)

    if err != nil {
        t.Error(err)
    }

}

// Test for deleting SDT object online.
func Test_StructuredDocumentTag_DeleteStructuredDocumentTagOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/StructuredDocumentTag/StructuredDocumentTag.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
        "nodePath": "sections/0/body/paragraphs/0",
    }

    request := &models.DeleteStructuredDocumentTagOnlineRequest{
        Document: requestDocument,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.DeleteStructuredDocumentTagOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for updating SDT object.
func Test_StructuredDocumentTag_UpdateStructuredDocumentTag(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/StructuredDocumentTag"
    localFile := "DocumentElements/StructuredDocumentTag/StructuredDocumentTag.docx"
    remoteFileName := "TestUpdateStructuredDocumentTag.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestStructuredDocumentTagListItems0 := models.StructuredDocumentTagListItem{
        DisplayText: ToStringPointer("Aspose Words"),
        Value: ToStringPointer("1"),
    }
    requestStructuredDocumentTagListItems1 := models.StructuredDocumentTagListItem{
        DisplayText: ToStringPointer("Hello world"),
        Value: ToStringPointer("2"),
    }
    requestStructuredDocumentTagListItems := []models.IStructuredDocumentTagListItem{
        &requestStructuredDocumentTagListItems0,
        &requestStructuredDocumentTagListItems1,
    }
    requestStructuredDocumentTag := models.StructuredDocumentTagUpdate{
        ListItems: requestStructuredDocumentTagListItems,
    }

    options := map[string]interface{}{
        "nodePath": "sections/0/body/paragraphs/0",
        "folder": remoteDataFolder,
    }

    request := &models.UpdateStructuredDocumentTagRequest{
        Name: ToStringPointer(remoteFileName),
        Index: ToInt32Pointer(int32(0)),
        StructuredDocumentTag: &requestStructuredDocumentTag,
        Optionals: options,
    }

    _, _, err := client.WordsApi.UpdateStructuredDocumentTag(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for updating SDT object online.
func Test_StructuredDocumentTag_UpdateStructuredDocumentTagOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/StructuredDocumentTag/StructuredDocumentTag.docx"

    requestDocument := OpenFile(t, localFile)
    requestStructuredDocumentTagListItems0 := models.StructuredDocumentTagListItem{
        DisplayText: ToStringPointer("Aspose Words"),
        Value: ToStringPointer("1"),
    }
    requestStructuredDocumentTagListItems1 := models.StructuredDocumentTagListItem{
        DisplayText: ToStringPointer("Hello world"),
        Value: ToStringPointer("2"),
    }
    requestStructuredDocumentTagListItems := []models.IStructuredDocumentTagListItem{
        &requestStructuredDocumentTagListItems0,
        &requestStructuredDocumentTagListItems1,
    }
    requestStructuredDocumentTag := models.StructuredDocumentTagUpdate{
        ListItems: requestStructuredDocumentTagListItems,
    }

    options := map[string]interface{}{
        "nodePath": "sections/0/body/paragraphs/0",
    }

    request := &models.UpdateStructuredDocumentTagOnlineRequest{
        Document: requestDocument,
        StructuredDocumentTag: &requestStructuredDocumentTag,
        Index: ToInt32Pointer(int32(0)),
        Optionals: options,
    }

    _, _, err := client.WordsApi.UpdateStructuredDocumentTagOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}
