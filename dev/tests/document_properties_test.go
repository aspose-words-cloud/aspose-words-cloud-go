/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="document_properties_test.go">
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

// Example of how to get document properties.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/dev/api/models"
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
    actual, _, err := client.WordsApi.GetDocumentProperties(ctx, remoteFileName, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.DocumentProperties, "Validate GetDocumentProperties response.");
    assert.NotNil(t, actual.DocumentProperties.List, "Validate GetDocumentProperties response.");
    assert.Equal(t, 24, len(actual.DocumentProperties.List), "Validate GetDocumentProperties response.");
    assert.NotNil(t, actual.DocumentProperties.List[0], "Validate GetDocumentProperties response.");
    assert.Equal(t, "Author", *actual.DocumentProperties.List[0].Name, "Validate GetDocumentProperties response.");
    assert.Equal(t, "", *actual.DocumentProperties.List[0].Value, "Validate GetDocumentProperties response.");
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
    actual, _, err := client.WordsApi.GetDocumentProperty(ctx, remoteFileName, "Author", options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.DocumentProperty, "Validate GetDocumentProperty response.");
    assert.Equal(t, "Author", *actual.DocumentProperty.Name, "Validate GetDocumentProperty response.");
    assert.Equal(t, "", *actual.DocumentProperty.Value, "Validate GetDocumentProperty response.");
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
    _, err := client.WordsApi.DeleteDocumentProperty(ctx, remoteFileName, "testProp", options)

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
    actual, _, err := client.WordsApi.CreateOrUpdateDocumentProperty(ctx, remoteFileName, "AsposeAuthor", requestProperty, options)

    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.DocumentProperty, "Validate UpdateDocumentProperty response.");
    assert.Equal(t, "AsposeAuthor", *actual.DocumentProperty.Name, "Validate UpdateDocumentProperty response.");
    assert.Equal(t, "Imran Anwar", *actual.DocumentProperty.Value, "Validate UpdateDocumentProperty response.");
}
