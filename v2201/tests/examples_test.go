/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="examplestest.go">
 *   Copyright (c) 2022 Aspose.Words for Cloud
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
package api_test

import (
    "os"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2201/api/models"
)

func Test_Examples_AcceptAllRevisions(t *testing.T) {
    documentsDir := GetExamplesDir()
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    wordsApi := client.WordsApi
    fileName:= "test_doc.docx"

    // Upload original document to cloud storage.
    myVar1, _ := os.Open(documentsDir + "/" + fileName)
    myVar2 := fileName
    uploadFileRequestOptions := map[string]interface{}{}
    uploadFileRequest := &models.UploadFileRequest{
        FileContent: myVar1,
        Path: ToStringPointer(myVar2),
        Optionals: uploadFileRequestOptions,
    }
    _, _, _ = wordsApi.UploadFile(ctx, uploadFileRequest)

    // Calls AcceptAllRevisions method for document in cloud.
    myVar3 := fileName
    requestOptions := map[string]interface{}{}
    request := &models.AcceptAllRevisionsRequest{
        Name: ToStringPointer(myVar3),
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.AcceptAllRevisions(ctx, request)
}

func Test_Examples_AcceptAllRevisionsOnline(t *testing.T) {
    documentsDir := GetExamplesDir()
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    UploadNextFileToStorage(t, ctx, client, documentsDir + "/test_doc.docx", "test_doc.docx")
    wordsApi := client.WordsApi
    fileName:= "test_doc.docx"

    // Calls AcceptAllRevisionsOnline method for document in cloud.
    requestDocument, _ := os.Open(documentsDir + "/" + fileName)
    requestOptions := map[string]interface{}{}
    request := &models.AcceptAllRevisionsOnlineRequest{
        Document: requestDocument,
        Optionals: requestOptions,
    }
    _, _, _ = wordsApi.AcceptAllRevisionsOnline(ctx, request)
}