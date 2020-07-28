/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="append_document_test.go">
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

// Example of how to append document.
package api_test

import (
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/dev/api/models"
)

// Test for appending document.
func Test_AppendDocument_AppendDocument(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentActions/AppendDocument"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestAppendDocument.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestDocumentListDocumentEntries0 := models.DocumentEntry{
        Href: remoteDataFolder + "/" + remoteFileName,
        ImportFormatMode: "KeepSourceFormatting",
    }
    requestDocumentListDocumentEntries := []models.DocumentEntry{
        requestDocumentListDocumentEntries0,
    }
    requestDocumentList := models.DocumentEntryList{
        DocumentEntries: requestDocumentListDocumentEntries,
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }
    _, _, err := client.WordsApi.AppendDocument(ctx, remoteFileName, requestDocumentList, options)

    if err != nil {
        t.Error(err)
    }
}

// Test for appending document online.
func Test_AppendDocument_AppendDocumentOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentActions/AppendDocument"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestAppendDocument.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestDocumentListDocumentEntries0 := models.DocumentEntry{
        Href: remoteDataFolder + "/" + remoteFileName,
        ImportFormatMode: "KeepSourceFormatting",
    }
    requestDocumentListDocumentEntries := []models.DocumentEntry{
        requestDocumentListDocumentEntries0,
    }
    requestDocumentList := models.DocumentEntryList{
        DocumentEntries: requestDocumentListDocumentEntries,
    }

    _, err := client.WordsApi.AppendDocumentOnline(ctx, OpenFile(t, localFile), requestDocumentList)

    if err != nil {
        t.Error(err)
    }
}
