/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="document_with_format_test.go">
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

// Example of how to get document with different format.
package api_test

import (
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2507/api/models"
)

// Test for getting document with specified format.
func Test_DocumentWithFormat_GetDocumentWithFormat(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentActions/DocumentWithFormat"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetDocumentWithFormat.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.GetDocumentWithFormatRequest{
        Name: ToStringPointer(remoteFileName),
        Format: ToStringPointer("text"),
        Optionals: options,
    }

    _, err := client.WordsApi.GetDocumentWithFormat(ctx, request)
    if err != nil {
        t.Error(err)
    }

}

// Test for getting document with specified format.
func Test_DocumentWithFormat_GetDocumentWithFormatAndOutPath(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentActions/DocumentWithFormat"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestGetDocumentWithFormat.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "outPath": baseTestOutPath + "/TestGetDocumentWithFormatAndOutPath.text",
    }

    request := &models.GetDocumentWithFormatRequest{
        Name: ToStringPointer(remoteFileName),
        Format: ToStringPointer("text"),
        Optionals: options,
    }

    _, err := client.WordsApi.GetDocumentWithFormat(ctx, request)
    if err != nil {
        t.Error(err)
    }

}
