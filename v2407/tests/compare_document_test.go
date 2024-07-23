/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="compare_document_test.go">
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

// Example of document comparison.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2407/api/models"
)

// Test for document comparison.
func Test_CompareDocument_CompareDocument(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteFolder := remoteBaseTestDataFolder + "/DocumentActions/CompareDocument"
    localFolder := "DocumentActions/CompareDocument"
    localName1 := "compareTestDoc1.doc"
    localName2 := "compareTestDoc2.doc"
    remoteName1 := "TestCompareDocument1.doc"
    remoteName2 := "TestCompareDocument2.doc"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFolder + "/" + localName1), remoteFolder + "/" + remoteName1)
    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFolder + "/" + localName2), remoteFolder + "/" + remoteName2)

    requestCompareDataFileReference := models.CreateRemoteFileReference(remoteFolder + "/" + remoteName2)
    requestCompareData := models.CompareData{
        Author: ToStringPointer("author"),
        DateTime: ToTimePointer(CreateTime(2015, 10, 26, 0, 0, 0)),
        FileReference: &requestCompareDataFileReference,
    }

    options := map[string]interface{}{
        "folder": remoteFolder,
        "destFileName": baseTestOutPath + "/TestCompareDocumentOut.doc",
    }

    request := &models.CompareDocumentRequest{
        Name: ToStringPointer(remoteName1),
        CompareData: &requestCompareData,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.CompareDocument(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetDocument(), "Validate CompareDocument response.");
    assert.Equal(t, "TestCompareDocumentOut.doc", DereferenceValue(actual.GetDocument().GetFileName()), "Validate CompareDocument response.");
}

// Test for document comparison online.
func Test_CompareDocument_CompareDocumentOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteFolder := remoteBaseTestDataFolder + "/DocumentActions/CompareDocument"
    localFolder := "DocumentActions/CompareDocument"
    localName1 := "compareTestDoc1.doc"
    localName2 := "compareTestDoc2.doc"
    remoteName2 := "TestCompareDocument2.doc"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFolder + "/" + localName2), remoteFolder + "/" + remoteName2)

    requestDocument := OpenFile(t, localFolder + "/" + localName1)
    requestCompareDataFileReference := models.CreateRemoteFileReference(remoteFolder + "/" + remoteName2)
    requestCompareData := models.CompareData{
        Author: ToStringPointer("author"),
        DateTime: ToTimePointer(CreateTime(2015, 10, 26, 0, 0, 0)),
        FileReference: &requestCompareDataFileReference,
    }

    options := map[string]interface{}{
        "destFileName": baseTestOutPath + "/TestCompareDocumentOut.doc",
    }

    request := &models.CompareDocumentOnlineRequest{
        Document: requestDocument,
        CompareData: &requestCompareData,
        Optionals: options,
    }

    _, _, err := client.WordsApi.CompareDocumentOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for document comparison online.
func Test_CompareDocument_CompareTwoDocumentOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteFolder := remoteBaseTestDataFolder + "/DocumentActions/CompareDocument"
    localFolder := "DocumentActions/CompareDocument"
    localName1 := "compareTestDoc1.doc"
    localName2 := "compareTestDoc2.doc"
    remoteName2 := "TestCompareDocument2.doc"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFolder + "/" + localName2), remoteFolder + "/" + remoteName2)

    requestDocument := OpenFile(t, localFolder + "/" + localName1)
    requestCompareDataFileReferenceStream := OpenFile(t, localFolder + "/" + localName2)
    requestCompareDataFileReference := models.CreateLocalFileReference(requestCompareDataFileReferenceStream)
    requestCompareData := models.CompareData{
        Author: ToStringPointer("author"),
        DateTime: ToTimePointer(CreateTime(2015, 10, 26, 0, 0, 0)),
        FileReference: &requestCompareDataFileReference,
    }

    options := map[string]interface{}{
        "destFileName": baseTestOutPath + "/TestCompareDocumentOut.doc",
    }

    request := &models.CompareDocumentOnlineRequest{
        Document: requestDocument,
        CompareData: &requestCompareData,
        Optionals: options,
    }

    _, _, err := client.WordsApi.CompareDocumentOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for document comparison with password protection.
func Test_CompareDocument_CompareDocumentWithPassword(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteFolder := remoteBaseTestDataFolder + "/DocumentActions/CompareDocument"
    localName := "DocWithPassword.docx"
    remoteName1 := "TestCompareDocument1.docx"
    remoteName2 := "TestCompareDocument2.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile("Common/" + localName), remoteFolder + "/" + remoteName1)
    UploadNextFileToStorage(t, ctx, client, GetLocalFile("Common/" + localName), remoteFolder + "/" + remoteName2)

    requestCompareDataFileReference := models.CreateRemoteFileReferenceWithPassword(remoteFolder + "/" + remoteName2, "12345")
    requestCompareData := models.CompareData{
        Author: ToStringPointer("author"),
        DateTime: ToTimePointer(CreateTime(2015, 10, 26, 0, 0, 0)),
        FileReference: &requestCompareDataFileReference,
    }

    options := map[string]interface{}{
        "folder": remoteFolder,
        "password": "12345",
        "destFileName": baseTestOutPath + "/TestCompareDocumentOut.docx",
    }

    request := &models.CompareDocumentRequest{
        Name: ToStringPointer(remoteName1),
        CompareData: &requestCompareData,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.CompareDocument(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetDocument(), "Validate CompareDocumentWithPassword response.");
    assert.Equal(t, "TestCompareDocumentOut.docx", DereferenceValue(actual.GetDocument().GetFileName()), "Validate CompareDocumentWithPassword response.");
}
