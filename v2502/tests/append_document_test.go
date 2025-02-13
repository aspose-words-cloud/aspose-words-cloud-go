/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="append_document_test.go">
 *   Copyright (c) 2025 Aspose.Words for Cloud
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

// Example of how to append document.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2502/api/models"
)

// Test for appending document.
func Test_AppendDocument_AppendDocument(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentActions/AppendDocument"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestAppendDocument.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestDocumentListDocumentEntries0FileReference := models.CreateRemoteFileReference(remoteDataFolder + "/" + remoteFileName)
    requestDocumentListDocumentEntries0 := models.DocumentEntry{
        FileReference: &requestDocumentListDocumentEntries0FileReference,
        ImportFormatMode: ToStringPointer("KeepSourceFormatting"),
    }
    requestDocumentListDocumentEntries := []models.IDocumentEntry{
        &requestDocumentListDocumentEntries0,
    }
    requestDocumentList := models.DocumentEntryList{
        DocumentEntries: requestDocumentListDocumentEntries,
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }

    request := &models.AppendDocumentRequest{
        Name: ToStringPointer(remoteFileName),
        DocumentList: &requestDocumentList,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.AppendDocument(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetDocument(), "Validate AppendDocument response.");
    assert.Equal(t, "TestAppendDocument.docx", DereferenceValue(actual.GetDocument().GetFileName()), "Validate AppendDocument response.");
}

// Test for appending document online.
func Test_AppendDocument_AppendDocumentOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)
    requestDocumentListDocumentEntries0FileReferenceStream := OpenFile(t, localFile)
    requestDocumentListDocumentEntries0FileReference := models.CreateLocalFileReference(requestDocumentListDocumentEntries0FileReferenceStream)
    requestDocumentListDocumentEntries0 := models.DocumentEntry{
        FileReference: &requestDocumentListDocumentEntries0FileReference,
        ImportFormatMode: ToStringPointer("KeepSourceFormatting"),
    }
    requestDocumentListDocumentEntries := []models.IDocumentEntry{
        &requestDocumentListDocumentEntries0,
    }
    requestDocumentList := models.DocumentEntryList{
        DocumentEntries: requestDocumentListDocumentEntries,
    }

    options := map[string]interface{}{
    }

    request := &models.AppendDocumentOnlineRequest{
        Document: requestDocument,
        DocumentList: &requestDocumentList,
        Optionals: options,
    }

    _, _, err := client.WordsApi.AppendDocumentOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}
