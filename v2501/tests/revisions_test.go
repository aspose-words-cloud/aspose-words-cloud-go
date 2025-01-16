/*
 * --------------------------------------------------------------------------------
 * <copyright company="Aspose" file="revisions_test.go">
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

// Example of how to accept all revisions in document.
package api_test

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/v2501/api/models"
)

// Test for accepting revisions in document.
func Test_Revisions_AcceptAllRevisions(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentActions/Revisions"
    localFile := "DocumentElements/Revisions/TestRevisions.doc"
    remoteFileName := "TestAcceptAllRevisions.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }

    request := &models.AcceptAllRevisionsRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.AcceptAllRevisions(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetResult(), "Validate AcceptAllRevisions response.");
    assert.NotNil(t, actual.GetResult().GetDest(), "Validate AcceptAllRevisions response.");
}

// Test for accepting revisions in document online.
func Test_Revisions_AcceptAllRevisionsOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Revisions/TestRevisions.doc"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.AcceptAllRevisionsOnlineRequest{
        Document: requestDocument,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.AcceptAllRevisionsOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetDocument(), "Validate AcceptAllRevisionsOnline response.");
    assert.NotNil(t, actual.GetModel(), "Validate AcceptAllRevisionsOnline response.");
    assert.NotNil(t, actual.GetModel().GetResult(), "Validate AcceptAllRevisionsOnline response.");
    assert.NotNil(t, actual.GetModel().GetResult().GetDest(), "Validate AcceptAllRevisionsOnline response.");
}

// Test for rejecting revisions in document.
func Test_Revisions_RejectAllRevisions(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentActions/Revisions"
    localFile := "DocumentElements/Revisions/TestRevisions.doc"
    remoteFileName := "TestRejectAllRevisions.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
        "destFileName": baseTestOutPath + "/" + remoteFileName,
    }

    request := &models.RejectAllRevisionsRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.RejectAllRevisions(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetResult(), "Validate RejectAllRevisions response.");
    assert.NotNil(t, actual.GetResult().GetDest(), "Validate RejectAllRevisions response.");
}

// Test for rejecting revisions in document online.
func Test_Revisions_RejectAllRevisionsOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Revisions/TestRevisions.doc"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.RejectAllRevisionsOnlineRequest{
        Document: requestDocument,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.RejectAllRevisionsOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetDocument(), "Validate RejectAllRevisionsOnline response.");
    assert.NotNil(t, actual.GetModel(), "Validate RejectAllRevisionsOnline response.");
    assert.NotNil(t, actual.GetModel().GetResult(), "Validate RejectAllRevisionsOnline response.");
    assert.NotNil(t, actual.GetModel().GetResult().GetDest(), "Validate RejectAllRevisionsOnline response.");
}

// Test for getting revisions from document.
func Test_Revisions_GetAllRevisions(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    remoteDataFolder := remoteBaseTestDataFolder + "/DocumentActions/Revisions"
    localFile := "DocumentElements/Revisions/TestRevisions.doc"
    remoteFileName := "TestAcceptAllRevisions.docx"

    UploadNextFileToStorage(t, ctx, client, GetLocalFile(localFile), remoteDataFolder + "/" + remoteFileName)


    options := map[string]interface{}{
        "folder": remoteDataFolder,
    }

    request := &models.GetAllRevisionsRequest{
        Name: ToStringPointer(remoteFileName),
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetAllRevisions(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetRevisions(), "Validate GetAllRevisions response.");
    assert.Equal(t, 6, len(actual.GetRevisions().GetRevisions()), "Validate GetAllRevisions response.");
}

// Test for getting revisions online from document.
func Test_Revisions_GetAllRevisionsOnline(t *testing.T) {
    config := ReadConfiguration(t)
    client, ctx := PrepareTest(t, config)
    localFile := "DocumentElements/Revisions/TestRevisions.doc"

    requestDocument := OpenFile(t, localFile)

    options := map[string]interface{}{
    }

    request := &models.GetAllRevisionsOnlineRequest{
        Document: requestDocument,
        Optionals: options,
    }

    actual, _, err := client.WordsApi.GetAllRevisionsOnline(ctx, request)
    if err != nil {
        t.Error(err)
    }

    assert.NotNil(t, actual.GetRevisions(), "Validate GetAllRevisionsOnline response.");
    assert.Equal(t, 6, len(actual.GetRevisions().GetRevisions()), "Validate GetAllRevisionsOnline response.");
}
