/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="document_properties_test.go">
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

// Example of how to get document properties.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2404/api/models"
)

// Test for getting document properties.
func Test_DocumentProperties_GetDocumentProperties(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/DocumentProperties"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetDocumentProperties.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.GetDocumentPropertiesRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetDocumentProperties(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetDocumentProperties(), "Validate GetDocumentProperties response.");
    assert.NotNil(t, actual.GetDocumentProperties().GetList(), "Validate GetDocumentProperties response.");
    assert.Equal(t, 24, len(actual.GetDocumentProperties().GetList()), "Validate GetDocumentProperties response.");
    assert.NotNil(t, actual.GetDocumentProperties().GetList()[0], "Validate GetDocumentProperties response.");
    assert.Equal(t, "Author", DereferenceValue(actual.GetDocumentProperties().GetList()[0].GetName()), "Validate GetDocumentProperties response.");
    assert.Equal(t, "", DereferenceValue(actual.GetDocumentProperties().GetList()[0].GetValue()), "Validate GetDocumentProperties response.");
}

// Test for getting document properties online.
func Test_DocumentProperties_GetDocumentPropertiesOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.GetDocumentPropertiesOnlineRequest{
        Document: requestDocument,
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetDocumentPropertiesOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// A test for GetDocumentProperty.
func Test_DocumentProperties_GetDocumentProperty(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/DocumentProperties"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetDocumentProperty.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.GetDocumentPropertyRequest{
        Name: ToStringPointer(remoteFileName),
        PropertyName: ToStringPointer("Author"),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetDocumentProperty(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetDocumentProperty(), "Validate GetDocumentProperty response.");
    assert.Equal(t, "Author", DereferenceValue(actual.GetDocumentProperty().GetName()), "Validate GetDocumentProperty response.");
    assert.Equal(t, "", DereferenceValue(actual.GetDocumentProperty().GetValue()), "Validate GetDocumentProperty response.");
}

// A test for GetDocumentProperty online.
func Test_DocumentProperties_GetDocumentPropertyOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.GetDocumentPropertyOnlineRequest{
        Document: requestDocument,
        PropertyName: ToStringPointer("Author"),
        Optionals: options,
    }

    _, _, err := client.WordsApi.GetDocumentPropertyOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for deleting document property.
func Test_DocumentProperties_DeleteDocumentProperty(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/DocumentProperties"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestDeleteDocumentProperty.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }

    request := &models.DeleteDocumentPropertyRequest{
        Name: ToStringPointer(remoteFileName),
        PropertyName: ToStringPointer("testProp"),
        Optionals: options,
    }

    _, err := client.WordsApi.DeleteDocumentProperty(ctx, request)

    if err != nil {
        t.Error(err)
    }

}

// Test for deleting document property online.
func Test_DocumentProperties_DeleteDocumentPropertyOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.DeleteDocumentPropertyOnlineRequest{
        Document: requestDocument,
        PropertyName: ToStringPointer("testProp"),
        Optionals: options,
    }

    _, _, err := client.WordsApi.DeleteDocumentPropertyOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for updating document property.
func Test_DocumentProperties_UpdateDocumentProperty(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentElements/DocumentProperties"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestUpdateDocumentProperty.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestProperty := models.DocumentPropertyCreateOrUpdate{
        Value: ToStringPointer("Imran Anwar"),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }

    request := &models.CreateOrUpdateDocumentPropertyRequest{
        Name: ToStringPointer(remoteFileName),
        PropertyName: ToStringPointer("AsposeAuthor"),
        Property: &requestProperty,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.CreateOrUpdateDocumentProperty(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetDocumentProperty(), "Validate UpdateDocumentProperty response.");
    assert.Equal(t, "AsposeAuthor", DereferenceValue(actual.GetDocumentProperty().GetName()), "Validate UpdateDocumentProperty response.");
    assert.Equal(t, "Imran Anwar", DereferenceValue(actual.GetDocumentProperty().GetValue()), "Validate UpdateDocumentProperty response.");
}

// Test for updating document property online.
func Test_DocumentProperties_UpdateDocumentPropertyOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)
    requestProperty := models.DocumentPropertyCreateOrUpdate{
        Value: ToStringPointer("Imran Anwar"),
    }

    options := map[string]interface{}{
    }

    request := &models.CreateOrUpdateDocumentPropertyOnlineRequest{
        Document: requestDocument,
        PropertyName: ToStringPointer("AsposeAuthor"),
        Property: &requestProperty,
        Optionals: options,
    }

    _, _, err := client.WordsApi.CreateOrUpdateDocumentPropertyOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}
