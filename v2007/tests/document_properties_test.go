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
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2007/api/models"
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
    _, _, err := client.WordsApi.GetDocumentProperties(ctx, remoteFileName, options)

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
    _, _, err := client.WordsApi.GetDocumentProperty(ctx, remoteFileName, "Author", options)

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
        Value: "Imran Anwar",
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }
    _, _, err := client.WordsApi.CreateOrUpdateDocumentProperty(ctx, remoteFileName, "AsposeAuthor", requestProperty, options)

    if err != nil {
        t.Error(err)
    }
}
