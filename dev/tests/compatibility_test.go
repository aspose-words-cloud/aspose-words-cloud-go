/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="compatibility_test.go">
 *   Copyright (c) 2026 Aspose.Words for Cloud
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

// Example of how to work with compatibility options.
package api_test

import (
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/dev/api/models"
)

// Test for optimize document to specific MS Word version.
func Test_Compatibility_OptimizeDocument(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/Compatibility"
    localFile := "Common/test_multi_pages.docx"
    remoteFileName := "TestOptimizeDocument.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)

    requestOptions := models.OptimizationOptions{
        MsWordVersion: ToStringPointer("Word2002"),
    }

    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.OptimizeDocumentRequest{
        Name: ToStringPointer(remoteFileName),
        Options: &requestOptions,
        Optionals: options,
    }

    _, err := client.WordsApi.OptimizeDocument(ctx, request)

    if err != nil {
        t.Error(err)
    }

}

// Test for optimize document to specific MS Word version.
func Test_Compatibility_OptimizeDocumentOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "Common/test_multi_pages.docx"

    requestDocument := OpenFile(t, localFile)
    requestOptions := models.OptimizationOptions{
        MsWordVersion: ToStringPointer("Word2002"),
    }

    options := map[string]interface{}{
    }

    request := &models.OptimizeDocumentOnlineRequest{
        Document: requestDocument,
        Options: &requestOptions,
        Optionals: options,
    }

    _, _, err := client.WordsApi.OptimizeDocumentOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

}
