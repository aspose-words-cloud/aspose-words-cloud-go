/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="compress_document_test.go">
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

// Example of how to compress document for reduce document size.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2408/api/models"
)

// Test for compression document.
func Test_CompressDocument_CompressDocument(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteFolder := remoteBaseTestDataFolder + "/DocumentActions/CompressDocument"
    localFolder := "DocumentActions/CompressDocument"
    localName := "TestCompress.docx"
    remoteName := "TestCompress.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFolder + "/" + localName), remoteFolder + "/" + remoteName)

    requestCompressOptions := models.CompressOptions{
    }

    options := map[string]interface{}{
        "folder": remoteFolder,
    }

    request := &models.CompressDocumentRequest{
        Name: ToStringPointer(remoteName),
        CompressOptions: &requestCompressOptions,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.CompressDocument(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetDocument(), "Validate CompressDocument response.");
}

// Test for compression document online.
func Test_CompressDocument_CompressDocumentOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFolder := "DocumentActions/CompressDocument"
    localName := "TestCompress.docx"

    requestDocument := OpenFile(t, localFolder + "/" + localName)
    requestCompressOptions := models.CompressOptions{
    }

    options := map[string]interface{}{
    }

    request := &models.CompressDocumentOnlineRequest{
        Document: requestDocument,
        CompressOptions: &requestCompressOptions,
        Optionals: options,
    }

    _, _, err := client.WordsApi.CompressDocumentOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}
